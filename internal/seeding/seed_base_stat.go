package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop5SeedBaseStats(qtx *database.Queries, ctx context.Context) error {
	stats, err := l.extractBaseStats()
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractBaseStats() ([]BaseStat, error) {
	stats := []BaseStat{}

	for i := range l.json.characters {
		char := &l.json.characters[i]

		statsNew, err := l.prepareBaseStats(char.BaseStats)
		if err != nil {
			return nil, err
		}

		stats = append(stats, statsNew...)
	}

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		statsNew, err := l.prepareBaseStats(mon.BaseStats)
		if err != nil {
			return nil, err
		}

		stats = append(stats, statsNew...)

		for j := range mon.AlteredStates {
			stateBs, err := l.extractAltStateBaseStats(&mon.AlteredStates[j])
			if err != nil {
				return nil, err
			}
			stats = append(stats, stateBs...)
		}
	}

	for i := range l.json.aeonStats {
		aeon := &l.json.aeonStats[i]

		aVals, err := l.prepareBaseStats(aeon.AVals)
		if err != nil {
			return nil, err
		}
		stats = append(stats, aVals...)

		bVals, err := l.prepareBaseStats(aeon.BVals)
		if err != nil {
			return nil, err
		}
		stats = append(stats, bVals...)

		for j := range aeon.XVals {
			xVal := &aeon.XVals[j]

			xVals, err := l.prepareBaseStats(xVal.BaseStats)
			if err != nil {
				return nil, err
			}
			stats = append(stats, xVals...)
		}
	}

	return dedupeRows(stats, l.Hashes), nil
}

func (l *Lookup) extractAltStateBaseStats(state *AlteredState) ([]BaseStat, error) {
	stats := []BaseStat{}
	var err error

	for i := range state.Changes {
		change := &state.Changes[i]

		if change.BaseStats == nil {
			continue
		}

		for j := range change.BaseStats {
			bs := &change.BaseStats[j]

			bs.StatID, err = assignFK(bs.StatName, l.Stats)
			if err != nil {
				return nil, err
			}

			stats = append(stats, *bs)
		}
	}

	return stats, nil
}

func (l *Lookup) prepareBaseStats(stats []BaseStat) ([]BaseStat, error) {
	statsNew := []BaseStat{}
	var err error

	for i := range stats {
		bs := &stats[i]

		bs.StatID, err = assignFK(bs.StatName, l.Stats)
		if err != nil {
			return nil, err
		}

		statsNew = append(statsNew, *bs)
	}

	return statsNew, nil
}
