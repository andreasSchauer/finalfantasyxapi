package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop5SeedRonsoRages(qtx *database.Queries, ctx context.Context) error {
	rages := extractRonsoRages()

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
		lookup.L.RonsoRages[seedtypes.Key(rages[i])] = rages[i]
		lookup.L.RonsoRagesID[row.ID] = rages[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractRonsoRages() []seedtypes.RonsoRage {
	rages := []seedtypes.RonsoRage{}

	for _, overdrive := range rawJSON.overdrives {
		if overdrive.User != "kimahri" {
			continue
		}

		rage := seedtypes.RonsoRage{
			ID:        0,
			Overdrive: overdrive,
		}

		rages = append(rages, rage)
	}

	return dedupeRows(rages, hashes)
}
