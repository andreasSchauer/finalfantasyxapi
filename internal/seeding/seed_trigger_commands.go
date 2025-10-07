package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type TriggerCommand struct {
	Ability
	AbilityAttributes
	AbilityID   int32
	Description string `json:"description"`
	Effect      string `json:"effect"`
	Cursor      string `json:"cursor"`
}

func (t TriggerCommand) ToHashFields() []any {
	return []any{
		t.AbilityID,
		t.Description,
		t.Effect,
		t.Cursor,
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
			ability := command.Ability
			attributes := command.AbilityAttributes
			ability.Type = database.AbilityTypeTriggerCommand

			dbAbility, err := l.seedAbility(qtx, attributes, ability)
			if err != nil {
				return err
			}

			command.AbilityID = dbAbility.ID

			err = qtx.CreateTriggerCommand(context.Background(), database.CreateTriggerCommandParams{
				DataHash:    generateDataHash(command),
				AbilityID:   command.AbilityID,
				Description: command.Description,
				Effect:      command.Effect,
				Cursor:      database.TargetType(command.Cursor),
			})
			if err != nil {
				return fmt.Errorf("couldn't create TriggerCommand: %s: %v", ability.Name, err)
			}
		}
		return nil
	})
}
