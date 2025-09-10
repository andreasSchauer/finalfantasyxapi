package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type OverdriveCommand struct {
	Name 			string 		`json:"name"`
	Description		string			`json:"description"`
	Rank			int32			`json:"rank"`
	Overdrives		[]Overdrive		`json:"overdrives"`
}


func(oc OverdriveCommand) ToHashFields() []any {
	return []any{
		oc.Name,
		oc.Description,
		oc.Rank,
	}
}


type Overdrive struct {	
	odCommandID			*int32
	Name				string		`json:""`
	Version				*int32		`json:""`
	Description			string		`json:""`
	Effect				string		`json:""`
	Rank				int32		`json:""`
	AppearsInHelpBar	bool		`json:""`
	CanCopycat			bool		`json:""`
	UnlockCondition		*string		`json:""`
	CountdownInSec		*int32		`json:""`
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
		})
		if err != nil {
			return fmt.Errorf("couldn't create Overdrive: %s: %v", overdrive.Name, err)
		}
	}

	return nil
}


