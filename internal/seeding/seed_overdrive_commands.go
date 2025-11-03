package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type OverdriveCommand struct {
	ID          int32
	CharClassID	*int32
	SubmenuID	*int32
	Name        string     	`json:"name"`
	Description string      `json:"description"`
	User		string		`json:"user"`
	Rank        int32       `json:"rank"`
	Topmenu     string     	`json:"topmenu"`
	OpenSubmenu string     	`json:"open_submenu"`
}

func (oc OverdriveCommand) ToHashFields() []any {
	return []any{
		oc.Name,
		oc.Description,
		oc.Rank,
		oc.Topmenu,
		oc.OpenSubmenu,
		derefOrNil(oc.CharClassID),
		derefOrNil(oc.SubmenuID),
	}
}

func (oc OverdriveCommand) GetID() int32 {
	return oc.ID
}

func (oc OverdriveCommand) Error() string {
	return fmt.Sprintf("overdrive command %s", oc.Name)
}



func (l *lookup) seedOverdriveCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrive_commands.json"

	var overdriveCommands []OverdriveCommand
	err := loadJSONFile(string(srcPath), &overdriveCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, command := range overdriveCommands {
			dbODCommand, err := qtx.CreateOverdriveCommand(context.Background(), database.CreateOverdriveCommandParams{
				DataHash:    generateDataHash(command),
				Name:        command.Name,
				Description: command.Description,
				Rank:        command.Rank,
				Topmenu:     database.TopmenuType(command.Topmenu),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Overdrive Command: %s: %v", command.Name, err)
			}

			command.ID = dbODCommand.ID
			l.overdriveCommands[command.Name] = command
		}

		return nil
	})
}


func (l *lookup) seedOverdriveCommandsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrive_commands.json"

	var overdriveCommands []OverdriveCommand
	err := loadJSONFile(string(srcPath), &overdriveCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCommand := range overdriveCommands {
			command, err := l.getOverdriveCommand(jsonCommand.Name)
			if err != nil {
				return err
			}

			command.CharClassID, err = assignFKPtr(&command.User, l.getCharacterClass)
			if err != nil {
				return err
			}

			command.SubmenuID, err = assignFKPtr(&command.OpenSubmenu, l.getSubmenu)
			if err != nil {
				return err
			}

			err = qtx.UpdateOverdriveCommand(context.Background(), database.UpdateOverdriveCommandParams{
				DataHash:    		generateDataHash(command),
				CharacterClassID: 	getNullInt32(command.CharClassID),
				SubmenuID: 			getNullInt32(command.SubmenuID),
				ID:					command.ID,
			})
			if err != nil {
				return fmt.Errorf("couldn't update Overdrive Command: %s: %v", command.Name, err)
			}
		}

		return nil
	})
}