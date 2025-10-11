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
	Name           			string  			`json:"name"`
	Effect         			string  			`json:"effect"`
	RelatedStats			[]string			`json:"related_stats"`
	RemovedStatusConditions	[]string			`json:"removed_status_conditions"`
	NullifyArmored 			*string 			`json:"nullify_armored"`
	StatChanges				[]StatChange		`json:"stat_changes"`
	ModifierChanges			[]ModifierChange	`json:"modifier_changes"`
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
	ID			int32
}


type PropertyStatJunction struct {
	PropertyID	int32
	StatID		int32
}

func (p PropertyStatJunction) ToHashFields() []any {
	return []any{
		p.PropertyID,
		p.StatID,
	}
}


type PropertyStatusConditionJunction struct {
	PropertyID			int32
	StatusConditionID	int32
}

func (s PropertyStatusConditionJunction) ToHashFields() []any {
	return []any{
		s.PropertyID,
		s.StatusConditionID,
	}
}


type PropertyStatChangeJunction struct {
	PropertyID		int32
	StatChangeID	int32
}

func (p PropertyStatChangeJunction) ToHashFields() []any {
	return []any{
		p.PropertyID,
		p.StatChangeID,
	}
}


type PropertyModifierChangeJunction struct {
	PropertyID	int32
	ModifierID	int32
}

func (p PropertyModifierChangeJunction) ToHashFields() []any {
	return []any{
		p.PropertyID,
		p.ModifierID,
	}
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
				Property: 	property,
				ID: 		dbProperty.ID,
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

			relationshipFunctions := []func(*database.Queries, PropertyLookup) error{
				l.createPropertyRelatedStats,
				l.createPropertyRemovedConditions,
				l.createPropertyStatChanges,
				l.createPropertyModifierChanges,
			}

			for _, function := range relationshipFunctions {
				err := function(db, property)
				if err != nil {
					return err
				}
			}
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

		junction := PropertyStatJunction{
			PropertyID: property.ID,
			StatID: 	stat.ID,
		}

		err = qtx.CreatePropertyStatJunction(context.Background(), database.CreatePropertyStatJunctionParams{
			DataHash: 		generateDataHash(junction),
			PropertyID: 	junction.PropertyID,
			StatID: 		junction.StatID,
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

		junction := PropertyStatusConditionJunction{
			PropertyID: 		property.ID,
			StatusConditionID: 	condition.ID,
		}

		err = qtx.CreatePropertyStatusConditionJunction(context.Background(), database.CreatePropertyStatusConditionJunctionParams{
			DataHash: 			generateDataHash(junction),
			PropertyID: 		junction.PropertyID,
			StatusConditionID: 	junction.StatusConditionID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}



func (l *lookup) createPropertyStatChanges(qtx *database.Queries, property PropertyLookup) error {
	for _, statChange := range property.StatChanges {
		dbStatChange, err := l.seedStatChange(qtx, statChange)
		if err != nil {
			return err
		}

		junction := PropertyStatChangeJunction{
			PropertyID: 	property.ID,
			StatChangeID: 	dbStatChange.ID,
		}

		err = qtx.CreatePropertyStatChangeJunction(context.Background(), database.CreatePropertyStatChangeJunctionParams{
			DataHash: 		generateDataHash(junction),
			PropertyID: 	junction.PropertyID,
			StatChangeID: 	junction.StatChangeID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}



func (l *lookup) createPropertyModifierChanges(qtx *database.Queries, property PropertyLookup) error {
	for _, modifierChange := range property.ModifierChanges {
		dbModifierChange, err := l.seedModifierChange(qtx, modifierChange)
		if err != nil {
			return err
		}

		junction := PropertyModifierChangeJunction{
			PropertyID: 	property.ID,
			ModifierID: 	dbModifierChange.ID,
		}

		err = qtx.CreatePropertyModifierChangeJunction(context.Background(), database.CreatePropertyModifierChangeJunctionParams{
			DataHash: 			generateDataHash(junction),
			PropertyID: 		junction.PropertyID,
			ModifierChangeID: 	junction.ModifierID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}