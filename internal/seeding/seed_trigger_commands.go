package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type TriggerCommandJSON struct {
	Name				string		`json:"name"`
	Version				*int32		`json:"version"`
	Description			string		`json:"description"`
	Effect				string		`json:"effect"`
	Rank				int32		`json:"rank"`
	AppearsInHelpBar	bool		`json:"appears_in_help_bar"`
	CanCopycat			bool		`json:"can_copycat"`
}


type TriggerCommand struct {
	Ability		Ability
	Info		TriggerCommandInfo
}



type TriggerCommandInfo struct {
	AbilityID			int32
	Description			string
	Effect				string		`json:"effect"`
	Rank				int32		`json:"rank"`
	AppearsInHelpBar	bool		`json:"appears_in_help_bar"`
	CanCopycat			bool		`json:"can_copycat"`
}

func(t TriggerCommandInfo) ToHashFields() []any {
	return []any{
		t.AbilityID,
		t.Description,
		t.Effect,
		t.Rank,
		t.AppearsInHelpBar,
		t.CanCopycat,
	}
}


func seedTriggerCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/trigger_commands.json"

	var json_data []TriggerCommandJSON

	err := loadJSONFile(string(srcPath), &json_data)
	if err != nil {
		return err
	}

	triggerCommands := jsonToTriggerCommands(json_data)

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, command := range triggerCommands {
			ability := command.Ability
			
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

			
			info := command.Info
			info.AbilityID = dbAbility.ID

			err = qtx.CreateTriggerCommand(context.Background(), database.CreateTriggerCommandParams{
				DataHash: 				generateDataHash(info),
				AbilityID: 				info.AbilityID,
				Description: 			info.Description,
				Effect: 				info.Effect,
				Rank: 					info.Rank,
				AppearsInHelpBar: 		info.AppearsInHelpBar,
				CanCopycat: 			info.CanCopycat,
			})
			if err != nil {
				return fmt.Errorf("couldn't create TriggerCommand: %s: %v", ability.Name, err)
			}
		}
		return nil
	})
}


func jsonToTriggerCommands(json_data []TriggerCommandJSON) []TriggerCommand {
	var triggerCommands []TriggerCommand

	for _, item := range json_data {
		ability := Ability{
            Name:          item.Name,
            Version:       item.Version,
            Specification: nil,
            Type:          database.AbilityTypeTriggerCommand,
        }

		info := TriggerCommandInfo{
            Description: 		item.Description,
			Effect:             item.Effect,
            Rank:               item.Rank,
            AppearsInHelpBar:   item.AppearsInHelpBar,
            CanCopycat:         item.CanCopycat,
        }

		triggerCommands = append(triggerCommands, TriggerCommand{
			Ability: ability,
			Info: info,
		})
	}

	return triggerCommands
}