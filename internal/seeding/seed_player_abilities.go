package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type PlayerAbilityJSON struct {
	Name				string		`json:"name"`
	Version				*int32		`json:"version"`
	Specification		*string		`json:"specification"`
	Description			*string		`json:"description"`
	Effect				string		`json:"effect"`
	Submenu				*string		`json:"submenu"`
	CanUseOutsideBattle	bool		`json:"can_use_outside_battle"`
	MPCost				*int32		`json:"mp_cost"`
	Rank				*int32		`json:"rank"`
	AppearsInHelpBar	bool		`json:"appears_in_help_bar"`
	CanCopycat			bool		`json:"can_copycat"`
	Cursor				*string		`json:"cursor"`
	OpenMenu			*string		`json:"open_menu"`
}


type PlayerAbility struct {
	Ability		Ability
	Info		PlayerAbilityInfo
}



type PlayerAbilityInfo struct {
	AbilityID			int32
	Description			*string
	Effect				string
	Submenu				*string
	CanUseOutsideBattle	bool
	MPCost				*int32
	Rank				*int32
	AppearsInHelpBar	bool
	CanCopycat			bool
	Cursor				*string
	OpenMenu			*string
}

func(a PlayerAbilityInfo) ToHashFields() []any {
	return []any{
		a.AbilityID,
		derefOrNil(a.Description),
		a.Effect,
		derefOrNil(a.Submenu),
		a.CanUseOutsideBattle,
		derefOrNil(a.MPCost),
		derefOrNil(a.Rank),
		a.AppearsInHelpBar,
		a.CanCopycat,
		derefOrNil(a.Cursor),
		derefOrNil(a.OpenMenu),
	}
}


func seedPlayerAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/player_abilities.json"

	var json_data []PlayerAbilityJSON

	err := loadJSONFile(string(srcPath), &json_data)
	if err != nil {
		return err
	}

	playerAbilities := jsonToPlayerAbilities(json_data)

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, playerAbility := range playerAbilities {
			ability := playerAbility.Ability
			
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


			info := playerAbility.Info
			info.AbilityID = dbAbility.ID

			err = qtx.CreatePlayerAbility(context.Background(), database.CreatePlayerAbilityParams{
				DataHash: 				generateDataHash(info),
				AbilityID: 				info.AbilityID,
				Description: 			getNullString(info.Description),
				Effect: 				info.Effect,
				Submenu: 				nullSubmenuType(info.Submenu),
				CanUseOutsideBattle: 	info.CanUseOutsideBattle,
				MpCost: 				getNullInt32(info.MPCost),
				Rank: 					getNullInt32(info.Rank),
				AppearsInHelpBar: 		info.AppearsInHelpBar,
				CanCopycat: 			info.CanCopycat,
				Cursor: 				nullTargetType(info.Cursor),
				OpenMenu: 				nullSubmenuType(info.OpenMenu),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Player Ability: %s: %v", ability.Name, err)
			}
		}
		return nil
	})
}


func jsonToPlayerAbilities(json_data []PlayerAbilityJSON) []PlayerAbility {
	var playerAbilities []PlayerAbility

	for _, item := range json_data {
		ability := Ability{
            Name:          item.Name,
            Version:       item.Version,
            Specification: item.Specification,
            Type:          database.AbilityTypePlayerAbility,
        }

		info := PlayerAbilityInfo{
            Description:         item.Description,
            Effect:              item.Effect,
            CanUseOutsideBattle: item.CanUseOutsideBattle,
            MPCost:              item.MPCost,
            Rank:                item.Rank,
            AppearsInHelpBar:    item.AppearsInHelpBar,
            CanCopycat:          item.CanCopycat,
			Cursor: 			 item.Cursor,
			OpenMenu: 			 item.OpenMenu,
        }

		playerAbilities = append(playerAbilities, PlayerAbility{
			Ability: ability,
			Info: info,
		})
	}

	return playerAbilities
}