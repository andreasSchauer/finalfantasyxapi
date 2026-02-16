package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type OverdriveCommand struct {
	ID          int32
	CharClassID *int32
	SubmenuID   *int32
	Name        string `json:"name"`
	Description string `json:"description"`
	User        string `json:"user"`
	Rank        int32  `json:"rank"`
	Topmenu     string `json:"topmenu"`
	OpenSubmenu string `json:"open_submenu"`
}

func (oc OverdriveCommand) ToHashFields() []any {
	return []any{
		oc.Name,
		oc.Description,
		oc.Rank,
		oc.Topmenu,
		oc.OpenSubmenu,
		h.DerefOrNil(oc.CharClassID),
		h.DerefOrNil(oc.SubmenuID),
	}
}

func (oc OverdriveCommand) GetID() int32 {
	return oc.ID
}

func (oc OverdriveCommand) Error() string {
	return fmt.Sprintf("overdrive command %s", oc.Name)
}

func (o OverdriveCommand) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   o.ID,
		Name: o.Name,
	}
}

func (l *Lookup) seedOverdriveCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/overdrive_commands.json"

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
				return h.NewErr(command.Error(), err, "couldn't create overdrive command")
			}

			command.ID = dbODCommand.ID
			l.OverdriveCommands[command.Name] = command
			l.OverdriveCommandsID[command.ID] = command
		}

		return nil
	})
}

func (l *Lookup) seedOverdriveCommandsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/overdrive_commands.json"

	var overdriveCommands []OverdriveCommand
	err := loadJSONFile(string(srcPath), &overdriveCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCommand := range overdriveCommands {
			command, err := GetResource(jsonCommand.Name, l.OverdriveCommands)
			if err != nil {
				return err
			}

			command.CharClassID, err = assignFKPtr(&command.User, l.CharClasses)
			if err != nil {
				return h.NewErr(command.Error(), err)
			}

			command.SubmenuID, err = assignFKPtr(&command.OpenSubmenu, l.Submenus)
			if err != nil {
				return h.NewErr(command.Error(), err)
			}

			err = qtx.UpdateOverdriveCommand(context.Background(), database.UpdateOverdriveCommandParams{
				DataHash:         generateDataHash(command),
				CharacterClassID: h.GetNullInt32(command.CharClassID),
				SubmenuID:        h.GetNullInt32(command.SubmenuID),
				ID:               command.ID,
			})
			if err != nil {
				return h.NewErr(command.Error(), err, "couldn't update overdrive command")
			}
		}

		return nil
	})
}
