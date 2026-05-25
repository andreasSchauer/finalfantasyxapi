package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop1SeedLocations(qtx *database.Queries, ctx context.Context) error {
	locations := dedupeRows(l.json.locations, l.Hashes)

	params := database.CreateLocationBulkParams{
		DataHash: 		make([]string, len(locations)),
		Name:     		make([]string, len(locations)),
		Availability: 	make([]database.AvailabilityType, len(locations)),
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
		l.json.locations[i].ID = row.ID
		l.Locations[Key(locations[i])] = locations[i]
		l.LocationsID[row.ID] = locations[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) completeLocations() error {
	for i := range l.json.locations {
		location := &l.json.locations[i]

		err := l.completeSublocations(location.Sublocations)
		if err != nil {
			return err
		}

		l.Locations[Key(location)] = *location
		l.LocationsID[location.ID] = *location
	}

	return nil
}
