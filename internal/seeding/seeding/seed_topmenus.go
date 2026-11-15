package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1SeedTopmenus(qtx *database.Queries, ctx context.Context) error {
	topmenus := dedupeRows(sh.RawJSON.Topmenus, sh.Hashes)

	params := database.CreateTopmenuBulkParams{
		DataHash: make([]string, len(topmenus)),
		Name:     make([]string, len(topmenus)),
	}

	for i, c := range topmenus {
		params.DataHash[i] = sh.GenerateDataHash(c)
		params.Name[i] = c.Name
	}

	dbRows, err := qtx.CreateTopmenuBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create topmenu: %v", err)
	}

	for i, row := range dbRows {
		topmenus[i].ID = row.ID
		sh.RawJSON.Topmenus[i].ID = row.ID
		lookup.L.Topmenus[seedtypes.Key(topmenus[i])] = topmenus[i]
		lookup.L.TopmenusID[row.ID] = topmenus[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}
