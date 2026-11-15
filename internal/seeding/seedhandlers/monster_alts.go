package seedhandlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop5Alts(qtx *database.Queries, ctx context.Context) error {
	changes, err := extractAlts()
	if err != nil {
		return err
	}

	params := database.CreateAltStateChangeBulkParams{
		DataHash:       make([]string, len(changes)),
		AlteredStateID: make([]int32, len(changes)),
		AlterationType: make([]database.AlterationType, len(changes)),
		Distance:       make([]sql.NullInt32, len(changes)),
		AddedStatusID:  make([]sql.NullInt32, len(changes)),
	}

	for i, c := range changes {
		params.DataHash[i] = sh.GenerateDataHash(c)
		params.AlteredStateID[i] = c.AlteredStateID
		params.AlterationType[i] = database.AlterationType(c.AlterationType)
		params.Distance[i] = h.GetNullInt32(c.Distance)
		params.AddedStatusID[i] = h.ObjPtrToNullInt32ID(c.AddedStatus)
	}

	dbRows, err := qtx.CreateAltStateChangeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create alts: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractAlts() ([]seedtypes.Alt, error) {
	changes := []seedtypes.Alt{}

	for i := range sh.RawJSON.Monsters {
		mon := &sh.RawJSON.Monsters[i]

		changesNew, err := prepareAlts(mon.AlteredStates)
		if err != nil {
			return nil, err
		}

		changes = append(changes, changesNew...)
	}

	return sh.DedupeRows(changes, sh.Hashes), nil
}

func prepareAlts(states []seedtypes.AlteredState) ([]seedtypes.Alt, error) {
	changes := []seedtypes.Alt{}
	var err error

	for i := range states {
		state := &states[i]

		for j := range state.Changes {
			change := &state.Changes[j]

			change.AlteredStateID, err = sh.GetHashID(state)
			if err != nil {
				return nil, err
			}

			if change.AddedStatus != nil {
				change.AddedStatus.ID, err = sh.GetHashID(change.AddedStatus)
				if err != nil {
					return nil, err
				}
			}

			changes = append(changes, *change)
		}
	}

	return changes, nil
}

func completeAlts(changes []seedtypes.Alt) error {
	for i := range changes {
		change := &changes[i]

		err := sh.AssignID(change)
		if err != nil {
			return err
		}

		err = sh.AssignIDs(change.BaseStats)
		if err != nil {
			return err
		}

		err = sh.AssignIDs(change.ElemResists)
		if err != nil {
			return err
		}
	}

	return nil
}
