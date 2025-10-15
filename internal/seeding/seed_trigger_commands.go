package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type TriggerCommand struct {
	Ability
	Description string `json:"description"`
	Effect      string `json:"effect"`
	Topmenu     string `json:"topmenu"`
	Cursor      string `json:"cursor"`
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

			command.Ability, err = seedObjAssignFK(qtx, command.Ability, l.seedAbility)
			if err != nil {
				return err
			}

			err = qtx.CreateTriggerCommand(context.Background(), database.CreateTriggerCommandParams{
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
		}
		return nil
	})
}
