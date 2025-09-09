package seeding

import (
	"context"
	"database/sql"
	"fmt"
	
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type AutoAbility struct {
	//id 			int32
	//dataHash		string
	Name         		string   	`json:"name"`
	Description			*string   	`json:"description"`
	Effect				string   	`json:"effect"`
	Type				string   	`json:"type"`
	Category			string   	`json:"category"`
	AbilityValue		*int32   	`json:"ability_value"`
	ActivationCondition	string   	`json:"activation_condition"`
	Counter				*string   	`json:"counter"`
	GradualRecovery		*string   	`json:"gradual_recovery"`
	OnHitElement		*string   	`json:"on_hit_element"`
	ConversionFrom		*string   	`json:"conversion_from"`
	ConversionTo		*string   	`json:"conversion_to"`
}

func (a AutoAbility) ToHashFields() []any {
	return []any{
		a.Name,
		derefOrNil(a.Description),
		a.Effect,
		a.Type,
		a.Category,
		derefOrNil(a.AbilityValue),
		a.ActivationCondition,
		derefOrNil(a.Counter),
		derefOrNil(a.GradualRecovery),
		derefOrNil(a.OnHitElement),
		derefOrNil(a.ConversionFrom),
		derefOrNil(a.ConversionTo),
	}
}


func seedAutoAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/auto_abilities.json"

	var autoAbilities []AutoAbility
	err := loadJSONFile(string(srcPath), &autoAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, ability := range autoAbilities {
			err = qtx.CreateAutoAbility(context.Background(), database.CreateAutoAbilityParams{
				DataHash:     			generateDataHash(ability),
				Name:         			ability.Name,
				Description: 			getNullString(ability.Description),
				Effect: 				ability.Effect,
				Type: 					database.EquipType(ability.Type),
				Category: 				database.AutoAbilityCategory(ability.Category),
				AbilityValue: 			getNullInt32(ability.AbilityValue),
				ActivationCondition: 	database.AaActivationCondition(ability.ActivationCondition),
				Counter: 				nullCounterType(ability.Counter),
				GradualRecovery: 		nullRecoveryType(ability.GradualRecovery),
				OnHitElement: 			nullElementType(ability.OnHitElement),
				ConversionFrom: 		nullParameter(ability.ConversionFrom),
				ConversionTo: 			nullParameter(ability.ConversionTo),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Auto-Ability: %s: %v", ability.Name, err)
			}
		}
		return nil
	})
}
