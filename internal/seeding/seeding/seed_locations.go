package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1SeedLocations(qtx *database.Queries, ctx context.Context) error {
	locations := dedupeRows(seedhelpers.RawJSON.Locations, seedhelpers.Hashes)

	params := database.CreateLocationBulkParams{
		DataHash:     make([]string, len(locations)),
		Name:         make([]string, len(locations)),
		Availability: make([]database.AvailabilityType, len(locations)),
	}

	for i, l := range locations {
		params.DataHash[i] = seedhelpers.GenerateDataHash(l)
		params.Name[i] = l.Name
		params.Availability[i] = database.AvailabilityType(l.Availability)
	}

	dbRows, err := qtx.CreateLocationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create locations: %v", err)
	}

	for i, row := range dbRows {
		locations[i].ID = row.ID
		seedhelpers.RawJSON.Locations[i].ID = row.ID
		lookup.L.Locations[seedtypes.Key(locations[i])] = locations[i]
		lookup.L.LocationsID[row.ID] = locations[i]
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func completeLocations() error {
	for i := range seedhelpers.RawJSON.Locations {
		location := &seedhelpers.RawJSON.Locations[i]

		err := completeSublocations(location.Sublocations)
		if err != nil {
			return err
		}

		lookup.L.Locations[seedtypes.Key(location)] = *location
		lookup.L.LocationsID[location.ID] = *location
	}

	return nil
}
