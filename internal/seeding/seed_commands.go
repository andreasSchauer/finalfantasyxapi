package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Command struct {
	//id 		int32
	//dataHash	string
	Name			string 		`json:"name"`
	Description		string		`json:"description"`
	Effect			string		`json:"effect"`
	Category		string		`json:"category"`
	Cursor			*string		`json:"cursor"`
}

func(c Command) ToHashFields() []any {
	return []any{
		c.Name,
		c.Description,
		c.Effect,
		c.Category,
		derefOrNil(c.Cursor),
	}
}


func seedCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/commands.json"

	var commands []Command
	err := loadJSONFile(string(srcPath), &commands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, command := range commands {
			err = qtx.CreateCommand(context.Background(), database.CreateCommandParams{
				DataHash: 		generateDataHash(command),
				Name: 			command.Name,
				Description: 	command.Description,
				Effect: 		command.Effect,
				Category: 		database.CommandCategory(command.Category),
				Cursor: 		nullTargetType(command.Cursor),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Command: %s: %v", command.Name, err)
			}
		}
		return nil
	})
}