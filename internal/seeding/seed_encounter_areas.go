package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop4SeedEncounterAreas(qtx *database.Queries, ctx context.Context) error {
	areas, err := l.extractEncounterAreas()
	if err != nil {
		return err
	}

	params := database.CreateEncounterAreaBulkParams{
		DataHash:      make([]string, len(areas)),
		AreaID:        make([]int32, len(areas)),
		Specification: make([]sql.NullString, len(areas)),
	}

	for i, a := range areas {
		params.DataHash[i] = generateDataHash(a)
		params.AreaID[i] = a.AreaID
		params.Specification[i] = h.GetNullString(a.Specification)
	}

	dbRows, err := qtx.CreateEncounterAreaBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create encounter areas: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractEncounterAreas() ([]EncounterArea, error) {
	areas := []EncounterArea{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		for j := range mf.EncounterAreas {
			area := &mf.EncounterAreas[j]

			area.AreaID, err = assignFK(area.LocationArea, l.Areas)
			if err != nil {
				return nil, err
			}

			areas = append(areas, *area)
		}
	}

	return dedupeRows(areas, l.Hashes), nil
}
