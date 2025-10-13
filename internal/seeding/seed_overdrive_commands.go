package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type OverdriveCommand struct {
	ID          *int32
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Rank        int32       `json:"rank"`
	Topmenu     *string     `json:"topmenu"`
	OpenMenu    *string     `json:"open_menu"`
	Overdrives  []Overdrive `json:"overdrives"`
}

func (oc OverdriveCommand) ToHashFields() []any {
	return []any{
		oc.Name,
		oc.Description,
		oc.Rank,
		derefOrNil(oc.Topmenu),
		derefOrNil(oc.OpenMenu),
	}
}

type Overdrive struct {
	ID          int32
	ODCommandID *int32
	Ability
	Description     string  `json:"description"`
	Effect          string  `json:"effect"`
	Topmenu         *string `json:"topmenu"`
	UnlockCondition *string `json:"unlock_condition"`
	CountdownInSec  *int32  `json:"countdown_in_sec"`
	Cursor          *string `json:"cursor"`
}

func (o Overdrive) ToHashFields() []any {
	return []any{
		derefOrNil(o.ODCommandID),
		o.Name,
		derefOrNil(o.Version),
		o.Description,
		o.Effect,
		derefOrNil(o.Topmenu),
		derefOrNil(o.Attributes.ID),
		derefOrNil(o.UnlockCondition),
		derefOrNil(o.CountdownInSec),
		derefOrNil(o.Cursor),
	}
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
			if command.Name != "" {
				dbODCommand, err := qtx.CreateOverdriveCommand(context.Background(), database.CreateOverdriveCommandParams{
					DataHash:    generateDataHash(command),
					Name:        command.Name,
					Description: command.Description,
					Rank:        command.Rank,
					Topmenu:     nullTopmenuType(command.Topmenu),
				})
				if err != nil {
					return fmt.Errorf("couldn't create Overdrive Command: %s: %v", command.Name, err)
				}

				command.ID = &dbODCommand.ID
			}

			err = l.seedOverdrives(qtx, command)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (l *lookup) seedOverdrives(qtx *database.Queries, command OverdriveCommand) error {
	for _, overdrive := range command.Overdrives {
		overdrive.ODCommandID = command.ID

		dbAttributes, err := l.seedAbilityAttributes(qtx, overdrive.Ability)
		if err != nil {
			return err
		}

		overdrive.Attributes.ID = &dbAttributes.ID

		dbOverdrive, err := qtx.CreateOverdrive(context.Background(), database.CreateOverdriveParams{
			DataHash:        generateDataHash(overdrive),
			OdCommandID:     getNullInt32(overdrive.ODCommandID),
			Name:            overdrive.Name,
			Version:         getNullInt32(overdrive.Version),
			Description:     overdrive.Description,
			Effect:          overdrive.Effect,
			Topmenu:         nullTopmenuType(overdrive.Topmenu),
			AttributesID:    *overdrive.Attributes.ID,
			UnlockCondition: getNullString(overdrive.UnlockCondition),
			CountdownInSec:  getNullInt32(overdrive.CountdownInSec),
			Cursor:          nullTargetType(overdrive.Cursor),
		})
		if err != nil {
			return fmt.Errorf("couldn't create Overdrive: %s: %v", overdrive.Name, err)
		}

		overdrive.ID = dbOverdrive.ID
		key := createLookupKey(overdrive.Ability)
		l.overdrives[key] = overdrive
	}

	return nil
}
