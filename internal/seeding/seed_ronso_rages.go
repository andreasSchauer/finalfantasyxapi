package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop5SeedRonsoRages(qtx *database.Queries, ctx context.Context) error {
	rages := l.extractRonsoRages()

	params := database.CreateRonsoRageBulkParams{
		DataHash:    make([]string, len(rages)),
		OverdriveID: make([]int32, len(rages)),
	}

	for i, r := range rages {
		params.DataHash[i] = generateDataHash(r)
		params.OverdriveID[i] = r.Overdrive.ID
	}

	dbRows, err := qtx.CreateRonsoRageBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ronso rages: %v", err)
	}

	for i, row := range dbRows {
		rages[i].ID = row.ID
		l.RonsoRages[Key(rages[i])] = rages[i]
		l.RonsoRagesID[row.ID] = rages[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractRonsoRages() []RonsoRage {
	rages := []RonsoRage{}

	for _, overdrive := range l.json.overdrives {
		if overdrive.User != "kimahri" {
			continue
		}

		rage := RonsoRage{
			ID:        0,
			Overdrive: overdrive,
		}

		rages = append(rages, rage)
	}

	return dedupeRows(rages, l.Hashes)
}
