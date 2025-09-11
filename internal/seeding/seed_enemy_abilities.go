package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type EnemyAbilityJSON struct {
	Name				string		`json:"name"`
	Version				*int32		`json:"version"`
	Specification		*string		`json:"specification"`
	Effect				*string		`json:"effect"`
	Rank				int32		`json:"rank"`
	AppearsInHelpBar	bool		`json:"appears_in_help_bar"`
	CanCopycat			bool		`json:"can_copycat"`
}


type EnemyAbility struct {
	Ability		Ability
	Info		EnemyAbilityInfo
}


type EnemyAbilityInfo struct {
	AbilityID			int32
	Effect				*string
	Rank				int32
	AppearsInHelpBar	bool
	CanCopycat			bool
}

func(a EnemyAbilityInfo) ToHashFields() []any {
	return []any{
		a.AbilityID,
		derefOrNil(a.Effect),
		a.Rank,
		a.AppearsInHelpBar,
		a.CanCopycat,
	}
}


func seedEnemyAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/enemy_abilities.json"

	var json_data []EnemyAbilityJSON

	err := loadJSONFile(string(srcPath), &json_data)
	if err != nil {
		return err
	}

	enemyAbilities := jsonToEnemyAbilities(json_data)

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, enemyAbility := range enemyAbilities {
			ability := enemyAbility.Ability
			
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

			
			info := enemyAbility.Info
			info.AbilityID = dbAbility.ID

			err = qtx.CreateEnemyAbility(context.Background(), database.CreateEnemyAbilityParams{
				DataHash: 				generateDataHash(info),
				AbilityID: 				info.AbilityID,
				Effect: 				getNullString(info.Effect),
				Rank: 					info.Rank,
				AppearsInHelpBar: 		info.AppearsInHelpBar,
				CanCopycat: 			info.CanCopycat,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Enemy Ability: %s: %v", ability.Name, err)
			}
		}
		return nil
	})
}


func jsonToEnemyAbilities(json_data []EnemyAbilityJSON) []EnemyAbility {
	var enemyAbilities []EnemyAbility

	for _, item := range json_data {
		ability := Ability{
            Name:          item.Name,
            Version:       item.Version,
            Specification: item.Specification,
            Type:          database.AbilityTypeEnemyAbility,
        }

		info := EnemyAbilityInfo{
            Effect:              item.Effect,
            Rank:                item.Rank,
            AppearsInHelpBar:    item.AppearsInHelpBar,
            CanCopycat:          item.CanCopycat,
        }

		enemyAbilities = append(enemyAbilities, EnemyAbility{
			Ability: ability,
			Info: info,
		})
	}

	return enemyAbilities
}