package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop5SeedBaseStats(qtx *database.Queries, ctx context.Context) error {
	stats, err := extractBaseStats()
	if err != nil {
		return err
	}

	params := database.CreateBaseStatBulkParams{
		DataHash: make([]string, len(stats)),
		StatID:   make([]int32, len(stats)),
		Value:    make([]int32, len(stats)),
	}

	for i, s := range stats {
		params.DataHash[i] = seedhelpers.GenerateDataHash(s)
		params.StatID[i] = s.StatID
		params.Value[i] = s.Value
	}

	dbRows, err := qtx.CreateBaseStatBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create base stats: %v", err)
	}

	for _, row := range dbRows {
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractBaseStats() ([]seedtypes.BaseStat, error) {
	stats := []seedtypes.BaseStat{}

	for i := range seedhelpers.RawJSON.Characters {
		char := &seedhelpers.RawJSON.Characters[i]

		statsNew, err := prepareBaseStats(char.BaseStats)
		if err != nil {
			return nil, err
		}

		stats = append(stats, statsNew...)
	}

	for i := range seedhelpers.RawJSON.Monsters {
		mon := &seedhelpers.RawJSON.Monsters[i]

		statsNew, err := prepareBaseStats(mon.BaseStats)
		if err != nil {
			return nil, err
		}

		stats = append(stats, statsNew...)

		for j := range mon.AlteredStates {
			stateBs, err := extractAltStateBaseStats(&mon.AlteredStates[j])
			if err != nil {
				return nil, err
			}
			stats = append(stats, stateBs...)
		}
	}

	for i := range seedhelpers.RawJSON.AeonStats {
		aeon := &seedhelpers.RawJSON.AeonStats[i]

		aVals, err := prepareBaseStats(aeon.AVals)
		if err != nil {
			return nil, err
		}
		stats = append(stats, aVals...)

		bVals, err := prepareBaseStats(aeon.BVals)
		if err != nil {
			return nil, err
		}
		stats = append(stats, bVals...)

		for j := range aeon.XVals {
			xVal := &aeon.XVals[j]

			xVals, err := prepareBaseStats(xVal.BaseStats)
			if err != nil {
				return nil, err
			}
			stats = append(stats, xVals...)
		}
	}

	return dedupeRows(stats, seedhelpers.Hashes), nil
}

func extractAltStateBaseStats(state *seedtypes.AlteredState) ([]seedtypes.BaseStat, error) {
	stats := []seedtypes.BaseStat{}
	var err error

	for i := range state.Changes {
		change := &state.Changes[i]

		if change.BaseStats == nil {
			continue
		}

		for j := range change.BaseStats {
			bs := &change.BaseStats[j]

			bs.StatID, err = seedhelpers.AssignFK(bs.StatName, lookup.L.Stats)
			if err != nil {
				return nil, err
			}

			stats = append(stats, *bs)
		}
	}

	return stats, nil
}

func prepareBaseStats(stats []seedtypes.BaseStat) ([]seedtypes.BaseStat, error) {
	statsNew := []seedtypes.BaseStat{}
	var err error

	for i := range stats {
		bs := &stats[i]

		bs.StatID, err = seedhelpers.AssignFK(bs.StatName, lookup.L.Stats)
		if err != nil {
			return nil, err
		}

		statsNew = append(statsNew, *bs)
	}

	return statsNew, nil
}
