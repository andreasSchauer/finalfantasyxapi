package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop2SeedAlteredStates(qtx *database.Queries, ctx context.Context) error {
	states := extractMonsterAlteredStates()

	params := database.CreateAlteredStateBulkParams{
		DataHash:    make([]string, len(states)),
		MonsterID:   make([]int32, len(states)),
		Condition:   make([]string, len(states)),
		IsTemporary: make([]bool, len(states)),
	}

	for i, s := range states {
		params.DataHash[i] = generateDataHash(s)
		params.MonsterID[i] = s.MonsterID
		params.Condition[i] = s.Condition
		params.IsTemporary[i] = s.IsTemporary
	}

	dbRows, err := qtx.CreateAlteredStateBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create altered states: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterAlteredStates() []seedtypes.AlteredState {
	states := []seedtypes.AlteredState{}

	for i := range rawJSON.monsters {
		mon := &rawJSON.monsters[i]

		for j := range mon.AlteredStates {
			state := &mon.AlteredStates[j]
			state.MonsterID = mon.ID
			states = append(states, *state)
		}
	}

	return dedupeRows(states, hashes)
}

func completeAlteredStates(states []seedtypes.AlteredState) error {
	for i := range states {
		state := &states[i]

		err := assignID(state)
		if err != nil {
			return err
		}

		err = completeAlts(state.Changes)
		if err != nil {
			return err
		}
	}

	return nil
}
