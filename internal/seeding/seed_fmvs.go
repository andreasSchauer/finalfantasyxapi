package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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
		derefOrNil(f.Translation),
		f.CutsceneDescription,
		derefOrNil(f.SongName),
		derefOrNil(f.SongID),
		f.AreaID,
	}
}

func (f FMV) Error() string {
	return fmt.Sprintf("fmv %s", f.Name)
}


func (l *lookup) seedFMVs(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/fmvs.json"

	var fmvs []FMV
	err := loadJSONFile(string(srcPath), &fmvs)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, fmv := range fmvs {
			var err error 

			fmv.SongID, err = assignFKPtr(fmv.SongName, l.getSong)
			if err != nil {
				return err
			}


			fmv.AreaID, err = assignFK(fmv.LocationArea, l.getArea)
			if err != nil {
				return err
			}

			err = qtx.CreateFMV(context.Background(), database.CreateFMVParams{
				DataHash:            generateDataHash(fmv),
				Name:                fmv.Name,
				Translation:         getNullString(fmv.Translation),
				CutsceneDescription: fmv.CutsceneDescription,
				SongID:              getNullInt32(fmv.SongID),
				AreaID:              fmv.AreaID,
			})
			if err != nil {
				return fmt.Errorf("couldn't create FMV: %s: %v", fmv.Name, err)
			}
		}
		return nil
	})
}
