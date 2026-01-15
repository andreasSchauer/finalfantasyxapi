package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Affinity struct {
	ID           int32
	Name         string   `json:"name"`
	DamageFactor *float32 `json:"damage_factor"`
}

func (a Affinity) ToHashFields() []any {
	return []any{
		a.Name,
		h.DerefOrNil(a.DamageFactor),
	}
}

func (a Affinity) GetID() int32 {
	return a.ID
}

func (a Affinity) Error() string {
	return fmt.Sprintf("elemental affinity %s", a.Name)
}

func (l *Lookup) seedAffinities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/elemental_affinities.json"

	var affinities []Affinity
	err := loadJSONFile(string(srcPath), &affinities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, affinity := range affinities {
			dbAffinity, err := qtx.CreateAffinity(context.Background(), database.CreateAffinityParams{
				DataHash:     generateDataHash(affinity),
				Name:         affinity.Name,
				DamageFactor: h.GetNullFloat64(affinity.DamageFactor),
			})
			if err != nil {
				return h.NewErr(affinity.Error(), err, "couldn't create elemental affinity")
			}

			affinity.ID = dbAffinity.ID
			l.Affinities[affinity.Name] = affinity
			l.AffinitiesID[affinity.ID] = affinity
		}
		return nil
	})
}
