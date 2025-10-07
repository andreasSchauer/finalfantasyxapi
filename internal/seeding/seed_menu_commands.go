package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type MenuCommand struct {
	//id 		int32
	//dataHash	string
	Name        string `json:"name"`
	Description string `json:"description"`
	Effect      string `json:"effect"`
}

func (c MenuCommand) ToHashFields() []any {
	return []any{
		c.Name,
		c.Description,
		c.Effect,
	}
}

func (l *lookup) seedMenuCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/menu_commands.json"

	var menu_commands []MenuCommand
	err := loadJSONFile(string(srcPath), &menu_commands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, command := range menu_commands {
			err = qtx.CreateMenuCommand(context.Background(), database.CreateMenuCommandParams{
				DataHash:    generateDataHash(command),
				Name:        command.Name,
				Description: command.Description,
				Effect:      command.Effect,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Menu Command: %s: %v", command.Name, err)
			}
		}
		return nil
	})
}
