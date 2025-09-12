package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type OverdriveCommand struct {
	Name 			string 			`json:"name"`
	Description		string			`json:"description"`
	Rank			int32			`json:"rank"`
	OpenMenu		*string			`json:"open_menu"`
	Overdrives		[]Overdrive		`json:"overdrives"`
}


func(oc OverdriveCommand) ToHashFields() []any {
	return []any{
		oc.Name,
		oc.Description,
		oc.Rank,
		derefOrNil(oc.OpenMenu),
	}
}


type Overdrive struct {	
	odCommandID			*int32
	Name				string		`json:"name"`
	Version				*int32		`json:"version"`
	Description			string		`json:"description"`
	Effect				string		`json:"effect"`
	Rank				int32		`json:"rank"`
	AppearsInHelpBar	bool		`json:"appears_in_help_bar"`
	CanCopycat			bool		`json:"can_copycat"`
	UnlockCondition		*string		`json:"unlock_condition"`
	CountdownInSec		*int32		`json:"countdown_in_sec"`
	Cursor				*string		`json:"cursor"`
}


func(o Overdrive) ToHashFields() []any {
	return []any{
		derefOrNil(o.odCommandID),
		o.Name,
		derefOrNil(o.Version),
		o.Description,
		o.Effect,
		o.Rank,
		o.AppearsInHelpBar,
		o.CanCopycat,
		derefOrNil(o.UnlockCondition),
		derefOrNil(o.CountdownInSec),
		derefOrNil(o.Cursor),
	}
}



func seedOverdriveCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrive_commands.json"

	var overdriveCommands []OverdriveCommand
	err := loadJSONFile(string(srcPath), &overdriveCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, command := range overdriveCommands {
			var overdriveCommandID *int32

			if command.Name != "" {
				dbODCommand, err := qtx.CreateOverdriveCommand(context.Background(), database.CreateOverdriveCommandParams{
					DataHash: 		generateDataHash(command),
					Name: 			command.Name,
					Description: 	command.Description,
					Rank: 			command.Rank,
					OpenMenu: 		nullSubmenuType(command.OpenMenu),
				})
				if err != nil {
					return fmt.Errorf("couldn't create Overdrive Command: %s: %v", command.Name, err)
				}

				overdriveCommandID = &dbODCommand.ID
			}

			err = seedOverdrives(qtx, command, overdriveCommandID)
			if err != nil {
				return err
			}
		}
		
		return nil
	})
}


func seedOverdrives(qtx *database.Queries, command OverdriveCommand, odCommandID *int32) error {
	for _, overdrive := range command.Overdrives {
		overdrive.odCommandID = odCommandID

		err := qtx.CreateOverdrive(context.Background(), database.CreateOverdriveParams{
			DataHash: 			generateDataHash(overdrive),
			OdCommandID: 		getNullInt32(overdrive.odCommandID),
			Name: 				overdrive.Name,
			Version: 			getNullInt32(overdrive.Version),
			Description: 		overdrive.Description,
			Effect: 			overdrive.Effect,
			Rank: 				overdrive.Rank,
			AppearsInHelpBar: 	overdrive.AppearsInHelpBar,
			CanCopycat: 		overdrive.CanCopycat,
			UnlockCondition: 	getNullString(overdrive.UnlockCondition),
			CountdownInSec: 	getNullInt32(overdrive.CountdownInSec),
			Cursor: 			nullTargetType(overdrive.Cursor),
		})
		if err != nil {
			return fmt.Errorf("couldn't create Overdrive: %s: %v", overdrive.Name, err)
		}
	}

	return nil
}


