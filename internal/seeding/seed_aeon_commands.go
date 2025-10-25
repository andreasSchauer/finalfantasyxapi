package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type AeonCommand struct {
	ID			int32
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Effect      string  `json:"effect"`
	Topmenu		string	`json:"topmenu"`
	Cursor      *string `json:"cursor"`
}

func (c AeonCommand) ToHashFields() []any {
	return []any{
		c.Name,
		c.Description,
		c.Effect,
		c.Topmenu,
		derefOrNil(c.Cursor),
	}
}

func (c AeonCommand) GetID() int32 {
	return c.ID
}

func (l *lookup) seedAeonCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/aeon_commands.json"

	var aeon_commands []AeonCommand
	err := loadJSONFile(string(srcPath), &aeon_commands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, command := range aeon_commands {
			dbAeonCommand, err := qtx.CreateAeonCommand(context.Background(), database.CreateAeonCommandParams{
				DataHash:    	generateDataHash(command),
				Name:        	command.Name,
				Description: 	command.Description,
				Effect:      	command.Effect,
				Topmenu: 		database.TopmenuType(command.Topmenu),
				Cursor:      	nullTargetType(command.Cursor),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Aeon Command: %s: %v", command.Name, err)
			}

			command.ID = dbAeonCommand.ID
			l.aeonCommands[command.Name] = command
		}
		return nil
	})
}
