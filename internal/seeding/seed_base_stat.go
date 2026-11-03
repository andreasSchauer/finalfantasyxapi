package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop5SeedBaseStats(qtx *database.Queries, ctx context.Context) error {
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
		params.DataHash[i] = generateDataHash(s)
		params.StatID[i] = s.StatID
		params.Value[i] = s.Value
	}

	dbRows, err := qtx.CreateBaseStatBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create base stats: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractBaseStats() ([]seedtypes.BaseStat, error) {
	stats := []seedtypes.BaseStat{}

	for i := range rawJSON.characters {
		char := &rawJSON.characters[i]

		statsNew, err := prepareBaseStats(char.BaseStats)
		if err != nil {
			return nil, err
		}

		stats = append(stats, statsNew...)
	}

	for i := range rawJSON.monsters {
		mon := &rawJSON.monsters[i]

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

	for i := range rawJSON.aeonStats {
		aeon := &rawJSON.aeonStats[i]

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

	return dedupeRows(stats, hashes), nil
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

			bs.StatID, err = assignFK(bs.StatName, L.Stats)
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

		bs.StatID, err = assignFK(bs.StatName, L.Stats)
		if err != nil {
			return nil, err
		}

		statsNew = append(statsNew, *bs)
	}

	return statsNew, nil
}
