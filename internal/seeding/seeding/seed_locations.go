package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop1SeedLocations(qtx *database.Queries, ctx context.Context) error {
	locations := dedupeRows(rawJSON.locations, hashes)

	params := database.CreateLocationBulkParams{
		DataHash:     make([]string, len(locations)),
		Name:         make([]string, len(locations)),
		Availability: make([]database.AvailabilityType, len(locations)),
	}

	for i, l := range locations {
		params.DataHash[i] = generateDataHash(l)
		params.Name[i] = l.Name
		params.Availability[i] = database.AvailabilityType(l.Availability)
	}

	dbRows, err := qtx.CreateLocationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create locations: %v", err)
	}

	for i, row := range dbRows {
		locations[i].ID = row.ID
		rawJSON.locations[i].ID = row.ID
		lookup.L.Locations[seedtypes.Key(locations[i])] = locations[i]
		lookup.L.LocationsID[row.ID] = locations[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func completeLocations() error {
	for i := range rawJSON.locations {
		location := &rawJSON.locations[i]

		err := completeSublocations(location.Sublocations)
		if err != nil {
			return err
		}

		lookup.L.Locations[seedtypes.Key(location)] = *location
		lookup.L.LocationsID[location.ID] = *location
	}

	return nil
}
