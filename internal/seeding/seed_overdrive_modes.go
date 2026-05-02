package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type OverdriveMode struct {
	ID             int32
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Effect         string          `json:"effect"`
	Type           string          `json:"type"`
	FillRate       *float32        `json:"fill_rate"`
	ActionsToLearn []ActionToLearn `json:"actions_to_learn"`
}

func (o OverdriveMode) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", o),
		o.Name,
		o.Description,
		o.Effect,
		o.Type,
		h.DerefOrNil(o.FillRate),
	}
}

func (o OverdriveMode) GetID() int32 {
	return o.ID
}

func (o OverdriveMode) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   o.ID,
		Name: o.Name,
	}
}

func (o OverdriveMode) Error() string {
	return fmt.Sprintf("overdrive mode %s", o.Name)
}

type ActionToLearn struct {
	ID     int32
	UserID int32
	User   string `json:"user"`
	Amount int32  `json:"amount"`
}

func (a ActionToLearn) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.UserID,
		a.Amount,
	}
}

func (a ActionToLearn) GetID() int32 {
	return a.ID
}

func (a *ActionToLearn) SetID(id int32) {
	a.ID = id
}

func (a ActionToLearn) GetName() string {
	return a.User
}

func (a ActionToLearn) GetVersion() *int32 {
	return nil
}

func (a ActionToLearn) GetVal() int32 {
	return a.Amount
}

func (a ActionToLearn) Error() string {
	return fmt.Sprintf("action to learn with user: %s, amount: %d", a.User, a.Amount)
}

func (l *Lookup) seedOverdriveModes(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/overdrive_modes.json"

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
				Type:        database.OverdriveModeType(mode.Type),
				FillRate:    h.GetNullFloat64(mode.FillRate),
			})
			if err != nil {
				return h.NewErr(mode.Error(), err, "couldn't create overdrive mode")
			}

			mode.ID = dbOverdriveMode.ID
			l.OverdriveModes[mode.Name] = mode
			l.OverdriveModesID[mode.ID] = mode
		}
		return nil
	})
}

func (l *Lookup) seedOverdriveModesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/overdrive_modes.json"

	var overdriveModes []OverdriveMode
	err := loadJSONFile(string(srcPath), &overdriveModes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonMode := range overdriveModes {
			mode, err := GetResource(jsonMode.Name, l.OverdriveModes)
			if err != nil {
				return err
			}

			for _, action := range mode.ActionsToLearn {
				junction, err := createJunctionSeed(qtx, mode, action, l.seedODModeAction)
				if err != nil {
					return h.NewErr(mode.Error(), err)
				}

				err = qtx.CreateOverdriveModesActionsToLearnJunction(context.Background(), database.CreateOverdriveModesActionsToLearnJunctionParams{
					DataHash:        generateDataHash(junction),
					OverdriveModeID: junction.ParentID,
					ActionID:        junction.ChildID,
				})
				if err != nil {
					subjects := h.JoinErrSubjects(mode.Error(), action.Error())
					return h.NewErr(subjects, err, "couldn't junction overdrive mode action")
				}
			}
		}
		return nil
	})
}

func (l *Lookup) seedODModeAction(qtx *database.Queries, action ActionToLearn) (ActionToLearn, error) {
	var err error

	action.UserID, err = assignFK(action.User, l.Characters)
	if err != nil {
		return ActionToLearn{}, h.NewErr(action.Error(), err)
	}

	dbAction, err := qtx.CreateODModeAction(context.Background(), database.CreateODModeActionParams{
		DataHash: generateDataHash(action),
		UserID:   action.UserID,
		Amount:   action.Amount,
	})
	if err != nil {
		return ActionToLearn{}, h.NewErr(action.Error(), err, "couldn't create overdrive mode action")
	}

	action.ID = dbAction.ID

	return action, nil
}


func (l *Lookup) loop1SeedOverdriveModes(qtx *database.Queries, ctx context.Context) error {
	modes := dedupeRows(l.json.overdriveModes, l.Hashes)

	params := database.CreateOverdriveModeBulkParams{
		DataHash:    make([]string, len(modes)),
		Name:        make([]string, len(modes)),
		Description: make([]string, len(modes)),
		Effect:      make([]string, len(modes)),
		Type:        make([]database.OverdriveModeType, len(modes)),
		FillRate:    make([]sql.NullFloat64, len(modes)),
	}

	for i, m := range modes {
		params.DataHash[i] = generateDataHash(m)
		params.Name[i] = m.Name
		params.Description[i] = m.Description
		params.Effect[i] = m.Effect
		params.Type[i] = database.OverdriveModeType(m.Type)
		params.FillRate[i] = h.GetNullFloat64(m.FillRate)
	}

	dbRows, err := qtx.CreateOverdriveModeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create overdrive modes: %v", err)
	}

	for i, row := range dbRows {
		modes[i].ID = row.ID
		l.json.overdriveModes[i].ID = row.ID
		l.OverdriveModes[modes[i].Name] = modes[i]
		l.OverdriveModesID[row.ID] = modes[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) completeOverdriveModes() error {
	for i := range l.json.overdriveModes {
		mode := &l.json.overdriveModes[i]

		err := assignIDs(l, mode.ActionsToLearn)
		if err != nil {
			return err
		}

		l.OverdriveModes[mode.Name] = *mode
		l.OverdriveModesID[mode.ID] = *mode
	}

	return nil
}


func (l *Lookup) loop5SeedOdModeActions(qtx *database.Queries, ctx context.Context) error {
	actions, err := l.extractOdModeActions()
	if err != nil {
		return err
	}

	params := database.CreateODModeActionBulkParams{
		DataHash:   make([]string, len(actions)),
		UserID: 	make([]int32, len(actions)),
		Amount: 	make([]int32, len(actions)),
	}

	for i, a := range actions {
		params.DataHash[i] = generateDataHash(a)
		params.UserID[i] = a.UserID
		params.Amount[i] = a.Amount
	}

	dbRows, err := qtx.CreateODModeActionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create overdrive mode actions: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractOdModeActions() ([]ActionToLearn, error) {
	actions := []ActionToLearn{}
	var err error

	for i := range l.json.overdriveModes {
		mode := &l.json.overdriveModes[i]

		for j := range mode.ActionsToLearn {
			action := &mode.ActionsToLearn[j]

			action.UserID, err = assignFK(action.User, l.Characters)
			if err != nil {
				return nil, err
			}

			actions = append(actions, *action)
		}
	}

	return dedupeRows(actions, l.Hashes), nil
}