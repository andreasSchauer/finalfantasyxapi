package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type OverdriveMode struct {
	//id 		int32
	//dataHash	string
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Effect         string          `json:"effect"`
	Type           string          `json:"type"`
	FillRate       *float32        `json:"fill_rate"`
	ActionsToLearn []ActionToLearn `json:"actions_to_learn"`
}

func (o OverdriveMode) ToHashFields() []any {
	return []any{
		o.Name,
		o.Description,
		o.Effect,
		o.Type,
		derefOrNil(o.FillRate),
	}
}

type OverdriveModeLookup struct {
	OverdriveMode
	ID int32
}

type ActionToLearn struct {
	UserID int32
	User   string `json:"user"`
	Amount int32  `json:"amount"`
}

func (a ActionToLearn) ToHashFields() []any {
	return []any{
		a.UserID,
		a.Amount,
	}
}

func (l *lookup) seedOverdriveModes(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrive_modes.json"

	var overdriveModes []OverdriveMode
	err := loadJSONFile(string(srcPath), &overdriveModes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, mode := range overdriveModes {
			dbOverdriveMode, err := qtx.CreateOverdriveMode(context.Background(), database.CreateOverdriveModeParams{
				DataHash:    generateDataHash(mode),
				Name:        mode.Name,
				Description: mode.Description,
				Effect:      mode.Effect,
				Type:        database.OverdriveType(mode.Type),
				FillRate:    getNullFloat64(mode.FillRate),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Overdrive Mode: %s: %v", mode.Name, err)
			}

			l.overdriveModes[mode.Name] = OverdriveModeLookup{
				OverdriveMode: mode,
				ID:            dbOverdriveMode.ID,
			}
		}
		return nil
	})
}

func (l *lookup) createOverdriveModesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrive_modes.json"

	var overdriveModes []OverdriveMode
	err := loadJSONFile(string(srcPath), &overdriveModes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonMode := range overdriveModes {
			mode, err := l.getOverdriveMode(jsonMode.Name)
			if err != nil {
				return err
			}

			for i, action := range mode.ActionsToLearn {
				user, err := l.getCharacter(action.User)
				if err != nil {
					return err
				}

				action.UserID = user.ID
				mode.ActionsToLearn[i] = action

				dbAction, err := qtx.CreateODModeAction(context.Background(), database.CreateODModeActionParams{
					DataHash: generateDataHash(action),
					UserID:   action.UserID,
					Amount:   action.Amount,
				})
				if err != nil {
					return err
				}

				junction := Junction{
					ParentID: 	mode.ID,
					ChildID:  	dbAction.ID,
				}

				err = qtx.CreateODModeActionJunction(context.Background(), database.CreateODModeActionJunctionParams{
					DataHash:        generateDataHash(junction),
					OverdriveModeID: junction.ParentID,
					ActionID:        junction.ChildID,
				})
				if err != nil {
					return err
				}
			}

			l.overdriveModes[mode.Name] = mode
		}
		return nil
	})
}
