package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop4SeedFMVs(qtx *database.Queries, ctx context.Context) error {
	fmvs, err := extractFMVs()
	if err != nil {
		return err
	}

	params := database.CreateFMVBulkParams{
		DataHash:            make([]string, len(fmvs)),
		Name:                make([]string, len(fmvs)),
		Translation:         make([]sql.NullString, len(fmvs)),
		CutsceneDescription: make([]string, len(fmvs)),
		SongID:              make([]sql.NullInt32, len(fmvs)),
		AreaID:              make([]int32, len(fmvs)),
	}

	for i, f := range fmvs {
		params.DataHash[i] = generateDataHash(f)
		params.Name[i] = f.Name
		params.Translation[i] = h.GetNullString(f.Translation)
		params.CutsceneDescription[i] = f.CutsceneDescription
		params.SongID[i] = h.GetNullInt32(f.SongID)
		params.AreaID[i] = f.AreaID
	}

	dbRows, err := qtx.CreateFMVBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create fmvs: %v", err)
	}

	for i, row := range dbRows {
		fmvs[i].ID = row.ID
		rawJSON.fmvs[i].ID = row.ID
		L.FMVs[seedtypes.Key(fmvs[i])] = fmvs[i]
		L.FMVsID[row.ID] = fmvs[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractFMVs() ([]seedtypes.FMV, error) {
	fmvs := []seedtypes.FMV{}
	var err error

	for i := range rawJSON.fmvs {
		fmv := &rawJSON.fmvs[i]

		fmv.SongID, err = assignFKPtr(fmv.SongName, L.Songs)
		if err != nil {
			return nil, err
		}

		fmv.AreaID, err = assignFK(fmv.LocationArea, L.Areas)
		if err != nil {
			return nil, err
		}

		fmvs = append(fmvs, *fmv)
	}

	return dedupeRows(fmvs, hashes), nil
}
