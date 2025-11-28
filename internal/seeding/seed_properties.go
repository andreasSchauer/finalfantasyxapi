package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Property struct {
	ID                      int32
	Name                    string           `json:"name"`
	Effect                  string           `json:"effect"`
	RelatedStats            []string         `json:"related_stats"`
	RemovedStatusConditions []string         `json:"removed_status_conditions"`
	NullifyArmored          *string          `json:"h.Nullify_armored"`
	StatChanges             []StatChange     `json:"stat_changes"`
	ModifierChanges         []ModifierChange `json:"modifier_changes"`
}

func (p Property) ToHashFields() []any {
	return []any{
		p.Name,
		p.Effect,
		h.DerefOrNil(p.NullifyArmored),
	}
}

func (p Property) GetID() int32 {
	return p.ID
}

func (p Property) Error() string {
	return fmt.Sprintf("property %s", p.Name)
}

func (l *Lookup) seedProperties(db *database.Queries, dbConn *sql.DB) error {
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
				NullifyArmored: h.NullNullifyArmored(property.NullifyArmored),
			})
			if err != nil {
				return h.GetErr(property.Error(), err, "couldn't create property")
			}

			property.ID = dbProperty.ID
			l.Properties[property.Name] = property
		}
		return nil
	})
}

func (l *Lookup) seedPropertiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/properties.json"

	var properties []Property
	err := loadJSONFile(string(srcPath), &properties)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonProperty := range properties {
			property, err := GetResource(jsonProperty.Name, l.Properties)
			if err != nil {
				return err
			}

			relationShipFunctions := []func(*database.Queries, Property) error{
				l.seedPropertyRelatedStats,
				l.seedPropertyRemovedConditions,
				l.seedPropertyStatChanges,
				l.seedPropertyModifierChanges,
			}

			for _, function := range relationShipFunctions {
				err := function(qtx, property)
				if err != nil {
					return h.GetErr(property.Error(), err)
				}
			}
		}

		return nil
	})

}

func (l *Lookup) seedPropertyRelatedStats(qtx *database.Queries, property Property) error {
	for _, jsonStat := range property.RelatedStats {
		junction, err := createJunction(property, jsonStat, l.Stats)
		if err != nil {
			return err
		}

		err = qtx.CreatePropertiesRelatedStatsJunction(context.Background(), database.CreatePropertiesRelatedStatsJunctionParams{
			DataHash:   generateDataHash(junction),
			PropertyID: junction.ParentID,
			StatID:     junction.ChildID,
		})
		if err != nil {
			return h.GetErr(jsonStat, err, "couldn't junction related stat")
		}
	}

	return nil
}

func (l *Lookup) seedPropertyRemovedConditions(qtx *database.Queries, property Property) error {
	for _, jsonCondition := range property.RemovedStatusConditions {
		junction, err := createJunction(property, jsonCondition, l.StatusConditions)
		if err != nil {
			return err
		}

		err = qtx.CreatePropertiesRemovedStatusConditionsJunction(context.Background(), database.CreatePropertiesRemovedStatusConditionsJunctionParams{
			DataHash:          generateDataHash(junction),
			PropertyID:        junction.ParentID,
			StatusConditionID: junction.ChildID,
		})
		if err != nil {
			return h.GetErr(jsonCondition, err, "couldn't junction removed status condition")
		}
	}

	return nil
}

func (l *Lookup) seedPropertyStatChanges(qtx *database.Queries, property Property) error {
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
			return h.GetErr(statChange.Error(), err, "couldn't junction stat change")
		}
	}

	return nil
}

func (l *Lookup) seedPropertyModifierChanges(qtx *database.Queries, property Property) error {
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
			return h.GetErr(modifierChange.Error(), err, "couldn't junction modifier change")
		}
	}

	return nil
}
