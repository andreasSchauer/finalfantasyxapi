package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop4SeedEncounterAreas(qtx *database.Queries, ctx context.Context) error {
	areas, err := extractEncounterAreas()
	if err != nil {
		return err
	}

	params := database.CreateEncounterAreaBulkParams{
		DataHash:      make([]string, len(areas)),
		AreaID:        make([]int32, len(areas)),
		Specification: make([]sql.NullString, len(areas)),
		Availability:  make([]database.AvailabilityType, len(areas)),
	}

	for i, a := range areas {
		params.DataHash[i] = generateDataHash(a)
		params.AreaID[i] = a.AreaID
		params.Specification[i] = h.GetNullString(a.Specification)
		params.Availability[i] = database.AvailabilityType(a.Availability)
	}

	dbRows, err := qtx.CreateEncounterAreaBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create encounter areas: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractEncounterAreas() ([]seedtypes.EncounterArea, error) {
	areas := []seedtypes.EncounterArea{}
	var err error

	for i := range rawJSON.monsterFormations {
		mf := &rawJSON.monsterFormations[i]

		for j := range mf.EncounterAreas {
			area := &mf.EncounterAreas[j]

			area.AreaID, err = assignFK(area.LocationArea, lookup.L.Areas)
			if err != nil {
				return nil, err
			}

			areas = append(areas, *area)
		}
	}

	return dedupeRows(areas, hashes), nil
}
