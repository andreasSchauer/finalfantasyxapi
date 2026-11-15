package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1SeedLocations(qtx *database.Queries, ctx context.Context) error {
	locations := sh.DedupeRows(sh.RawJSON.Locations, sh.Hashes)

	params := database.CreateLocationBulkParams{
		DataHash:     make([]string, len(locations)),
		Name:         make([]string, len(locations)),
		Availability: make([]database.AvailabilityType, len(locations)),
	}

	for i, l := range locations {
		params.DataHash[i] = sh.GenerateDataHash(l)
		params.Name[i] = l.Name
		params.Availability[i] = database.AvailabilityType(l.Availability)
	}

	dbRows, err := qtx.CreateLocationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create locations: %v", err)
	}

	for i, row := range dbRows {
		locations[i].ID = row.ID
		sh.RawJSON.Locations[i].ID = row.ID
		lookup.L.Locations[seedtypes.Key(locations[i])] = locations[i]
		lookup.L.LocationsID[row.ID] = locations[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func CompleteLocations() error {
	for i := range sh.RawJSON.Locations {
		location := &sh.RawJSON.Locations[i]

		err := completeSublocations(location.Sublocations)
		if err != nil {
			return err
		}

		lookup.L.Locations[seedtypes.Key(location)] = *location
		lookup.L.LocationsID[location.ID] = *location
	}

	return nil
}
