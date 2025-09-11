package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type OverdriveAbilityJSON struct {
	Name				string		`json:"name"`
	Version				*int32		`json:"version"`
	Specification		*string		`json:"specification"`
}


type OverdriveAbility struct {
	Ability		Ability
	Info		OverdriveAbilityInfo
}



type OverdriveAbilityInfo struct {
	AbilityID			int32
}

func(a OverdriveAbilityInfo) ToHashFields() []any {
	return []any{
		a.AbilityID,
	}
}


func seedOverdriveAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrive_abilities.json"

	var json_data []OverdriveAbilityJSON

	err := loadJSONFile(string(srcPath), &json_data)
	if err != nil {
		return err
	}

	overdriveAbilities := jsonToOverdriveAbilities(json_data)

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, overdriveAbility := range overdriveAbilities {
			ability := overdriveAbility.Ability
			
			dbAbility, err := qtx.CreateAbility(context.Background(), database.CreateAbilityParams{
				DataHash: 		generateDataHash(ability),
				Name: 			ability.Name,
				Version: 		getNullInt32(ability.Version),
				Specification: 	getNullString(ability.Specification),
				Type: 			ability.Type,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Ability: %s-%d, type: %s: %v", ability.Name, *ability.Version, ability.Type, err)
			}

			
			info := overdriveAbility.Info
			info.AbilityID = dbAbility.ID

			err = qtx.CreateEnemyAbility(context.Background(), database.CreateEnemyAbilityParams{
				DataHash: 				generateDataHash(info),
				AbilityID: 				info.AbilityID,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Overdrive Ability: %s: %v", ability.Name, err)
			}
		}
		return nil
	})
}


func jsonToOverdriveAbilities(json_data []OverdriveAbilityJSON) []OverdriveAbility {
	var overdriveAbilities []OverdriveAbility

	for _, item := range json_data {
		ability := Ability{
            Name:          item.Name,
            Version:       item.Version,
            Specification: item.Specification,
            Type:          database.AbilityTypeOverdriveAbility,
        }

		info := OverdriveAbilityInfo{}

		overdriveAbilities = append(overdriveAbilities, OverdriveAbility{
			Ability: ability,
			Info: info,
		})
	}

	return overdriveAbilities
}