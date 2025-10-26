package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Property struct {
	ID                      int32
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

func (p Property) GetID() int32 {
	return p.ID
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

			property.ID = dbProperty.ID
			l.properties[property.Name] = property
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

			relationShipFunctions := []func(*database.Queries, Property) error{
				l.createPropertyRelatedStats,
				l.createPropertyRemovedConditions,
				l.createPropertyStatChanges,
				l.createPropertyModifierChanges,
			}

			for _, function := range relationShipFunctions {
				err := function(qtx, property)
				if err != nil {
					return fmt.Errorf("property: %s: %v", property.Name, err)
				}
			}
		}

		return nil
	})

}

func (l *lookup) createPropertyRelatedStats(qtx *database.Queries, property Property) error {
	for _, jsonStat := range property.RelatedStats {
		junction, err := createJunction(property, jsonStat, l.getStat)
		if err != nil {
			return err
		}

		err = qtx.CreatePropertiesRelatedStatsJunction(context.Background(), database.CreatePropertiesRelatedStatsJunctionParams{
			DataHash:   generateDataHash(junction),
			PropertyID: junction.ParentID,
			StatID:     junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create related stats: %v", err)
		}
	}

	return nil
}

func (l *lookup) createPropertyRemovedConditions(qtx *database.Queries, property Property) error {
	for _, jsonCondition := range property.RemovedStatusConditions {
		junction, err := createJunction(property, jsonCondition, l.getStatusCondition)
		if err != nil {
			return err
		}

		err = qtx.CreatePropertiesRemovedStatusConditionsJunction(context.Background(), database.CreatePropertiesRemovedStatusConditionsJunctionParams{
			DataHash:          generateDataHash(junction),
			PropertyID:        junction.ParentID,
			StatusConditionID: junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create removed conditions: %v", err)
		}
	}

	return nil
}

func (l *lookup) createPropertyStatChanges(qtx *database.Queries, property Property) error {
	for _, statChange := range property.StatChanges {
		junction, err := createJunctionSeed(qtx, property, statChange, l.seedStatChange)
		if err != nil {
			return err
		}

		err = qtx.CreatePropertiesStatChangesJunction(context.Background(), database.CreatePropertiesStatChangesJunctionParams{
			DataHash:     generateDataHash(junction),
			PropertyID:   junction.ParentID,
			StatChangeID: junction.ChildID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *lookup) createPropertyModifierChanges(qtx *database.Queries, property Property) error {
	for _, modifierChange := range property.ModifierChanges {
		junction, err := createJunctionSeed(qtx, property, modifierChange, l.seedModifierChange)
		if err != nil {
			return err
		}

		err = qtx.CreatePropertiesModifierChangesJunction(context.Background(), database.CreatePropertiesModifierChangesJunctionParams{
			DataHash:         generateDataHash(junction),
			PropertyID:       junction.ParentID,
			ModifierChangeID: junction.ChildID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
