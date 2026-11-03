package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop1SeedTopmenus(qtx *database.Queries, ctx context.Context) error {
	topmenus := dedupeRows(rawJSON.topmenus, hashes)

	params := database.CreateTopmenuBulkParams{
		DataHash: make([]string, len(topmenus)),
		Name:     make([]string, len(topmenus)),
	}

	for i, c := range topmenus {
		params.DataHash[i] = generateDataHash(c)
		params.Name[i] = c.Name
	}

	dbRows, err := qtx.CreateTopmenuBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create topmenu: %v", err)
	}

	for i, row := range dbRows {
		topmenus[i].ID = row.ID
		rawJSON.topmenus[i].ID = row.ID
		L.Topmenus[seedtypes.Key(topmenus[i])] = topmenus[i]
		L.TopmenusID[row.ID] = topmenus[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}
