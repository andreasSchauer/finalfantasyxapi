package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type FMV struct {
	//id 			int32
	//dataHash		string
	Name                string       `json:"name"`
	Translation         *string      `json:"translation"`
	CutsceneDescription string       `json:"cutscene_description"`
	SongName            *string      `json:"music"`
	LocationArea        LocationArea `json:"location_area"`
	SongID              *int32
	AreaID              int32
}

func (f FMV) ToHashFields() []any {
	return []any{
		f.Name,
		h.DerefOrNil(f.Translation),
		f.CutsceneDescription,
		h.DerefOrNil(f.SongName),
		h.DerefOrNil(f.SongID),
		f.AreaID,
	}
}

func (f FMV) Error() string {
	return fmt.Sprintf("fmv %s", f.Name)
}

func (l *Lookup) seedFMVs(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/fmvs.json"

	var fmvs []FMV
	err := loadJSONFile(string(srcPath), &fmvs)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, fmv := range fmvs {
			var err error

			fmv.SongID, err = assignFKPtr(fmv.SongName, l.Songs)
			if err != nil {
				return h.GetErr(fmv.Error(), err)
			}

			fmv.AreaID, err = assignFK(fmv.LocationArea, l.Areas)
			if err != nil {
				return h.GetErr(fmv.Error(), err)
			}

			err = qtx.CreateFMV(context.Background(), database.CreateFMVParams{
				DataHash:            generateDataHash(fmv),
				Name:                fmv.Name,
				Translation:         h.GetNullString(fmv.Translation),
				CutsceneDescription: fmv.CutsceneDescription,
				SongID:              h.GetNullInt32(fmv.SongID),
				AreaID:              fmv.AreaID,
			})
			if err != nil {
				return h.GetErr(fmv.Error(), err, "couldn't create fmv")
			}
		}
		return nil
	})
}
