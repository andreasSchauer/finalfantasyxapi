package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

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
		l.OverdriveModes[Key(modes[i])] = modes[i]
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

		l.OverdriveModes[Key(mode)] = *mode
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
		DataHash: make([]string, len(actions)),
		UserID:   make([]int32, len(actions)),
		Amount:   make([]int32, len(actions)),
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

func (l *Lookup) getOverdriveModeActions(om OverdriveMode) ([]ActionToLearn, error) {
	return om.ActionsToLearn, nil
}

func (l *Lookup) seedJuncOverdriveModeActions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrive modes + actions"
	jParams, err := processJunctions(l, desc, l.json.overdriveModes, l.getOverdriveModeActions)
	if err != nil {
		return err
	}

	return qtx.CreateOverdriveModesActionsToLearnJunctionBulk(ctx, database.CreateOverdriveModesActionsToLearnJunctionBulkParams{
		DataHash:        jParams.DataHashes,
		OverdriveModeID: jParams.ParentIDs,
		ActionID:        jParams.ChildIDs,
	})
}
