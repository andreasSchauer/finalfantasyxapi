package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop4SeedFMVs(qtx *database.Queries, ctx context.Context) error {
	fmvs, err := l.extractFMVs()
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
		l.json.fmvs[i].ID = row.ID
		l.FMVs[fmvs[i].Name] = fmvs[i]
		l.FMVsID[row.ID] = fmvs[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractFMVs() ([]FMV, error) {
	fmvs := []FMV{}
	var err error

	for i := range l.json.fmvs {
		fmv := &l.json.fmvs[i]

		fmv.SongID, err = assignFKPtr(fmv.SongName, l.Songs)
		if err != nil {
			return nil, err
		}

		fmv.AreaID, err = assignFK(fmv.LocationArea, l.Areas)
		if err != nil {
			return nil, err
		}

		fmvs = append(fmvs, *fmv)
	}

	return dedupeRows(fmvs, l.Hashes), nil
}
