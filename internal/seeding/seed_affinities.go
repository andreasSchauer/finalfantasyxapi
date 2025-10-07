package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Affinity struct {
	//id 			int32
	//dataHash		string
	Name         string   `json:"name"`
	DamageFactor *float32 `json:"damage_factor"`
}

func (a Affinity) ToHashFields() []any {
	return []any{
		a.Name,
		derefOrNil(a.DamageFactor),
	}
}

func (l *lookup) seedAffinities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/elemental_affinities.json"

	var affinities []Affinity
	err := loadJSONFile(string(srcPath), &affinities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, affinity := range affinities {
			err = qtx.CreateAffinity(context.Background(), database.CreateAffinityParams{
				DataHash:     generateDataHash(affinity),
				Name:         affinity.Name,
				DamageFactor: getNullFloat64(affinity.DamageFactor),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Elemental Affinity: %s: %v", affinity.Name, err)
			}
		}
		return nil
	})
}
