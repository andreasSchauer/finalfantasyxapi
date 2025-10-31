package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type TriggerCommand struct {
	ID 					int32
	Ability
	Description        	string              `json:"description"`
	Effect             	string              `json:"effect"`
	Topmenu            	string              `json:"topmenu"`
	RelatedStats		[]string			`json:"related_stats"`
	Cursor             	string              `json:"cursor"`
	BattleInteractions 	[]BattleInteraction `json:"battle_interactions"`
}

func (t TriggerCommand) ToHashFields() []any {
	return []any{
		t.Ability.ID,
		t.Description,
		t.Effect,
		t.Topmenu,
		t.Cursor,
	}
}

func (t TriggerCommand) GetID() int32 {
	return t.ID
}

func (t TriggerCommand) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        t.Name,
		Version:     t.Version,
		AbilityType: string(database.AbilityTypeTriggerCommand),
	}
}

func (l *lookup) seedTriggerCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/trigger_commands.json"

	var triggerCommands []TriggerCommand

	err := loadJSONFile(string(srcPath), &triggerCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, command := range triggerCommands {
			var err error
			command.Type = database.AbilityTypeTriggerCommand

			command.Ability, err = seedObjAssignID(qtx, command.Ability, l.seedAbility)
			if err != nil {
				return err
			}

			dbTriggerCommand, err := qtx.CreateTriggerCommand(context.Background(), database.CreateTriggerCommandParams{
				DataHash:    generateDataHash(command),
				AbilityID:   command.Ability.ID,
				Description: command.Description,
				Effect:      command.Effect,
				Topmenu:     database.TopmenuType(command.Topmenu),
				Cursor:      database.TargetType(command.Cursor),
			})
			if err != nil {
				return fmt.Errorf("couldn't create TriggerCommand: %s: %v", command.Name, err)
			}

			command.ID = dbTriggerCommand.ID
			key := createLookupKey(command.Ability)
			l.triggerCommands[key] = command
		}
		return nil
	})
}

func (l *lookup) createTriggerCommandsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/trigger_commands.json"

	var triggerCommands []TriggerCommand

	err := loadJSONFile(string(srcPath), &triggerCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCommand := range triggerCommands {
			abilityRef := jsonCommand.GetAbilityRef()

			command, err := l.getTriggerCommand(abilityRef)
			if err != nil {
				return err
			}

			err = l.seedTriggerCommandRelatedStats(qtx, command)
			if err != nil {
				return err
			}

			l.currentAbility = command.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, command.BattleInteractions)
			if err != nil {
				return err
			}
		}

		return nil
	})
}



func (l *lookup) seedTriggerCommandRelatedStats(qtx *database.Queries, command TriggerCommand) error {
	for _, jsonStat := range command.RelatedStats {
		junction, err := createJunction(command, jsonStat, l.getStat)
		if err != nil {
			return err
		}

		err = qtx.CreateTriggerCommandsRelatedStatsJunction(context.Background(), database.CreateTriggerCommandsRelatedStatsJunctionParams{
			DataHash: 			generateDataHash(junction),
			TriggerCommandID: 	junction.ParentID,
			StatID: 			junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("command %s: %v", createLookupKey(command.Ability), err)
		}
	}

	return nil
}