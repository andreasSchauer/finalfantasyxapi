package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop5SeedAltStateChanges(qtx *database.Queries, ctx context.Context) error {
	changes, err := l.extractAltStateChanges()
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
		params.DataHash[i] = generateDataHash(c)
		params.AlteredStateID[i] = c.AlteredStateID
		params.AlterationType[i] = database.AlterationType(c.AlterationType)
		params.Distance[i] = h.GetNullInt32(c.Distance)
		params.AddedStatusID[i] = h.ObjPtrToNullInt32ID(c.AddedStatus)
	}

	dbRows, err := qtx.CreateAltStateChangeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create alt state changes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAltStateChanges() ([]AltStateChange, error) {
	changes := []AltStateChange{}

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		changesNew, err := l.prepareAltStateChanges(mon.AlteredStates)
		if err != nil {
			return nil, err
		}

		changes = append(changes, changesNew...)
	}

	return dedupeRows(changes, l.Hashes), nil
}

func (l *Lookup) prepareAltStateChanges(states []AlteredState) ([]AltStateChange, error) {
	changes := []AltStateChange{}
	var err error

	for i := range states {
		state := &states[i]

		for j := range state.Changes {
			change := &state.Changes[j]

			change.AlteredStateID, err = l.getHashID(state)
			if err != nil {
				return nil, err
			}

			if change.AddedStatus != nil {
				change.AddedStatus.ID, err = l.getHashID(change.AddedStatus)
				if err != nil {
					return nil, err
				}
			}

			changes = append(changes, *change)
		}
	}

	return changes, nil
}

func (l *Lookup) completeAltStateChanges(changes []AltStateChange) error {
	for i := range changes {
		change := &changes[i]

		err := l.assignID(change)
		if err != nil {
			return err
		}

		err = assignIDs(l, change.BaseStats)
		if err != nil {
			return err
		}

		err = assignIDs(l, change.ElemResists)
		if err != nil {
			return err
		}
	}

	return nil
}
