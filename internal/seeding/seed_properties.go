package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Property struct {
	//id 		int32
	//dataHash	string
	Name                    string           `json:"name"`
	Effect                  string           `json:"effect"`
	RelatedStats            []string         `json:"related_stats"`
	RemovedStatusConditions []string         `json:"removed_status_conditions"`
	NullifyArmored          *string          `json:"nullify_armored"`
	StatChanges             []StatChange     `json:"stat_changes"`
	ModifierChanges         []ModifierChange `json:"modifier_changes"`
}

func (p Property) ToHashFields() []any {
	return []any{
		p.Name,
		p.Effect,
		derefOrNil(p.NullifyArmored),
	}
}

type PropertyLookup struct {
	Property
	ID int32
}

func (l *lookup) seedProperties(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/properties.json"

	var properties []Property
	err := loadJSONFile(string(srcPath), &properties)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, property := range properties {
			dbProperty, err := qtx.CreateProperty(context.Background(), database.CreatePropertyParams{
				DataHash:       generateDataHash(property),
				Name:           property.Name,
				Effect:         property.Effect,
				NullifyArmored: nullNullifyArmored(property.NullifyArmored),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Property: %s: %v", property.Name, err)
			}

			l.properties[property.Name] = PropertyLookup{
				Property: property,
				ID:       dbProperty.ID,
			}
		}
		return nil
	})
}

func (l *lookup) createPropertiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/properties.json"

	var properties []Property
	err := loadJSONFile(string(srcPath), &properties)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonProperty := range properties {
			property, err := l.getProperty(jsonProperty.Name)
			if err != nil {
				return err
			}

			err = l.createPropertyRelatedStats(db, property)
			if err != nil {
				return err
			}

			err = l.createPropertyRemovedConditions(db, property)
			if err != nil {
				return err
			}

			statChangesNew, err := l.createPropertyStatChanges(db, property)
			if err != nil {
				return err
			}

			modifierChangesNew, err := l.createPropertyModifierChanges(db, property)
			if err != nil {
				return err
			}

			property.StatChanges = statChangesNew
			property.ModifierChanges = modifierChangesNew
			l.properties[property.Name] = property
		}

		return nil
	})

}

func (l *lookup) createPropertyRelatedStats(qtx *database.Queries, property PropertyLookup) error {
	for _, jsonStat := range property.RelatedStats {
		stat, err := l.getStat(jsonStat)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: 	property.ID,
			ChildID:  	stat.ID,
		}

		err = qtx.CreatePropertyStatJunction(context.Background(), database.CreatePropertyStatJunctionParams{
			DataHash:   generateDataHash(junction),
			PropertyID: junction.ParentID,
			StatID:     junction.ChildID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *lookup) createPropertyRemovedConditions(qtx *database.Queries, property PropertyLookup) error {
	for _, jsonCondition := range property.RemovedStatusConditions {
		condition, err := l.getStatusCondition(jsonCondition)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: 	property.ID,
			ChildID:  	condition.ID,
		}

		err = qtx.CreatePropertyStatusConditionJunction(context.Background(), database.CreatePropertyStatusConditionJunctionParams{
			DataHash:          generateDataHash(junction),
			PropertyID:        junction.ParentID,
			StatusConditionID: junction.ChildID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *lookup) createPropertyStatChanges(qtx *database.Queries, property PropertyLookup) ([]StatChange, error) {
	for i, statChange := range property.StatChanges {
		dbStatChange, err := l.seedStatChange(qtx, statChange)
		if err != nil {
			return []StatChange{}, err
		}
		statChange.StatID = dbStatChange.StatID
		property.StatChanges[i] = statChange

		junction := Junction{
			ParentID: 	property.ID,
			ChildID:  	dbStatChange.ID,
		}

		err = qtx.CreatePropertyStatChangeJunction(context.Background(), database.CreatePropertyStatChangeJunctionParams{
			DataHash:     generateDataHash(junction),
			PropertyID:   junction.ParentID,
			StatChangeID: junction.ChildID,
		})
		if err != nil {
			return []StatChange{}, err
		}
	}

	return property.StatChanges, nil
}

func (l *lookup) createPropertyModifierChanges(qtx *database.Queries, property PropertyLookup) ([]ModifierChange, error) {
	for i, modifierChange := range property.ModifierChanges {
		dbModifierChange, err := l.seedModifierChange(qtx, modifierChange)
		if err != nil {
			return []ModifierChange{}, err
		}
		modifierChange.ModifierID = dbModifierChange.ModifierID
		property.ModifierChanges[i] = modifierChange

		junction := Junction{
			ParentID: 	property.ID,
			ChildID:  	dbModifierChange.ID,
		}

		err = qtx.CreatePropertyModifierChangeJunction(context.Background(), database.CreatePropertyModifierChangeJunctionParams{
			DataHash:         generateDataHash(junction),
			PropertyID:       junction.ParentID,
			ModifierChangeID: junction.ChildID,
		})
		if err != nil {
			return []ModifierChange{}, err
		}
	}

	return property.ModifierChanges, nil
}
