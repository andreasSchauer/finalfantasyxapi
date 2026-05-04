package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop2SeedSublocations(qtx *database.Queries, ctx context.Context) error {
	sublocations := l.extractSublocations()

	params := database.CreateSublocationBulkParams{
		DataHash:   make([]string, len(sublocations)),
		LocationID: make([]int32, len(sublocations)),
		Name:       make([]string, len(sublocations)),
	}

	for i, s := range sublocations {
		params.DataHash[i] = generateDataHash(s)
		params.LocationID[i] = s.Location.ID
		params.Name[i] = s.Name
	}

	dbRows, err := qtx.CreateSublocationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create sublocations: %v", err)
	}

	for i, row := range dbRows {
		sublocations[i].ID = row.ID
		l.Sublocations[sublocations[i].Name] = sublocations[i]
		l.SublocationsID[row.ID] = sublocations[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractSublocations() []Sublocation {
	sublocations := []Sublocation{}

	for i := range l.json.locations {
		location := &l.json.locations[i]

		for j := range location.Sublocations {
			sublocation := &location.Sublocations[j]
			sublocation.Location = *location
			sublocations = append(sublocations, *sublocation)
		}
	}

	return dedupeRows(sublocations, l.Hashes)
}

func (l *Lookup) completeSublocations(sublocations []Sublocation) error {
	for i := range sublocations {
		sublocation := &sublocations[i]

		err := l.completeAreas(sublocation.Areas)
		if err != nil {
			return err
		}

		l.Sublocations[sublocation.Name] = *sublocation
		l.SublocationsID[sublocation.ID] = *sublocation
	}

	return nil
}
