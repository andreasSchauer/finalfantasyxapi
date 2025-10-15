package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type OverdriveMode struct {
	ID				int32
	Name           	string          `json:"name"`
	Description    	string          `json:"description"`
	Effect         	string          `json:"effect"`
	Type           	string          `json:"type"`
	FillRate       	*float32        `json:"fill_rate"`
	ActionsToLearn 	[]ActionToLearn `json:"actions_to_learn"`
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

func (o OverdriveMode) GetID() int32 {
	return o.ID
}


type ActionToLearn struct {
	ID		int32
	UserID 	int32
	User   	string `json:"user"`
	Amount 	int32  `json:"amount"`
}

func (a ActionToLearn) ToHashFields() []any {
	return []any{
		a.UserID,
		a.Amount,
	}
}

func (a ActionToLearn) GetID() int32 {
	return a.ID
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

			mode.ID = dbOverdriveMode.ID
			l.overdriveModes[mode.Name] = mode
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

			for _, action := range mode.ActionsToLearn {
				junction, err := createJunctionSeed(qtx, mode, action, l.seedODModeAction)
				if err != nil {
					return err
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
		}
		return nil
	})
}


func (l *lookup) seedODModeAction(qtx *database.Queries, action ActionToLearn) (ActionToLearn, error) {
	var err error

	action.UserID, err = assignFK(action.User, l.getCharacter)
	if err != nil {
		return ActionToLearn{}, err
	}

	dbAction, err := qtx.CreateODModeAction(context.Background(), database.CreateODModeActionParams{
		DataHash: generateDataHash(action),
		UserID:   action.UserID,
		Amount:   action.Amount,
	})
	if err != nil {
		return ActionToLearn{}, err
	}

	action.ID = dbAction.ID

	return action, nil
 }