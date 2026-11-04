package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop2SeedSublocations(qtx *database.Queries, ctx context.Context) error {
	sublocations := extractSublocations()

	params := database.CreateSublocationBulkParams{
		DataHash:     make([]string, len(sublocations)),
		LocationID:   make([]int32, len(sublocations)),
		Name:         make([]string, len(sublocations)),
		Availability: make([]database.AvailabilityType, len(sublocations)),
	}

	for i, s := range sublocations {
		params.DataHash[i] = generateDataHash(s)
		params.LocationID[i] = s.LocationID
		params.Name[i] = s.Name
		params.Availability[i] = database.AvailabilityType(s.Availability)
	}

	dbRows, err := qtx.CreateSublocationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create sublocations: %v", err)
	}

	for i, row := range dbRows {
		sublocations[i].ID = row.ID
		lookup.L.Sublocations[seedtypes.Key(sublocations[i])] = sublocations[i]
		lookup.L.SublocationsID[row.ID] = sublocations[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractSublocations() []seedtypes.Sublocation {
	sublocations := []seedtypes.Sublocation{}

	for i := range rawJSON.locations {
		location := &rawJSON.locations[i]

		for j := range location.Sublocations {
			sublocation := &location.Sublocations[j]
			sublocation.Location = location.Name
			sublocation.LocationID = location.ID
			sublocations = append(sublocations, *sublocation)
		}
	}

	return dedupeRows(sublocations, hashes)
}

func completeSublocations(sublocations []seedtypes.Sublocation) error {
	for i := range sublocations {
		sublocation := &sublocations[i]

		err := completeAreas(sublocation.Areas)
		if err != nil {
			return err
		}

		lookup.L.Sublocations[seedtypes.Key(sublocation)] = *sublocation
		lookup.L.SublocationsID[sublocation.ID] = *sublocation
	}

	return nil
}
