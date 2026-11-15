package seedhandlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1OverdriveModes(qtx *database.Queries, ctx context.Context) error {
	modes := sh.DedupeRows(sh.RawJSON.OverdriveModes, sh.Hashes)

	params := database.CreateOverdriveModeBulkParams{
		DataHash:    make([]string, len(modes)),
		Name:        make([]string, len(modes)),
		Description: make([]string, len(modes)),
		Effect:      make([]string, len(modes)),
		Type:        make([]database.OverdriveModeType, len(modes)),
		FillRate:    make([]sql.NullFloat64, len(modes)),
	}

	for i, m := range modes {
		params.DataHash[i] = sh.GenerateDataHash(m)
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
		sh.RawJSON.OverdriveModes[i].ID = row.ID
		lookup.L.OverdriveModes[seedtypes.Key(modes[i])] = modes[i]
		lookup.L.OverdriveModesID[row.ID] = modes[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func CompleteOverdriveModes() error {
	for i := range sh.RawJSON.OverdriveModes {
		mode := &sh.RawJSON.OverdriveModes[i]

		err := sh.AssignIDs(mode.ActionsToLearn)
		if err != nil {
			return err
		}

		lookup.L.OverdriveModes[seedtypes.Key(mode)] = *mode
		lookup.L.OverdriveModesID[mode.ID] = *mode
	}

	return nil
}

func Loop5OdModeActions(qtx *database.Queries, ctx context.Context) error {
	actions, err := extractOdModeActions()
	if err != nil {
		return err
	}

	params := database.CreateODModeActionBulkParams{
		DataHash: make([]string, len(actions)),
		UserID:   make([]int32, len(actions)),
		Amount:   make([]int32, len(actions)),
	}

	for i, a := range actions {
		params.DataHash[i] = sh.GenerateDataHash(a)
		params.UserID[i] = a.UserID
		params.Amount[i] = a.Amount
	}

	dbRows, err := qtx.CreateODModeActionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create overdrive mode actions: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractOdModeActions() ([]seedtypes.ActionToLearn, error) {
	actions := []seedtypes.ActionToLearn{}
	var err error

	for i := range sh.RawJSON.OverdriveModes {
		mode := &sh.RawJSON.OverdriveModes[i]

		for j := range mode.ActionsToLearn {
			action := &mode.ActionsToLearn[j]

			action.UserID, err = sh.AssignFK(action.User, lookup.L.Characters)
			if err != nil {
				return nil, err
			}

			actions = append(actions, *action)
		}
	}

	return sh.DedupeRows(actions, sh.Hashes), nil
}
