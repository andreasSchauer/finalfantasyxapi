package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Stat struct {
	ID       int32
	Name     string `json:"name"`
	Effect   string `json:"effect"`
	MinVal   int32  `json:"min_val"`
	MaxVal   int32  `json:"max_val"`
	MaxVal2  *int32 `json:"max_val_2"`
	SphereID *int32
	Sphere   string `json:"sphere"`
}

func (s Stat) ToHashFields() []any {
	return []any{
		s.Name,
		s.Effect,
		s.MinVal,
		s.MaxVal,
		derefOrNil(s.MaxVal2),
		derefOrNil(s.SphereID),
	}
}

func (s Stat) GetID() int32 {
	return s.ID
}

func (s Stat) Error() string {
	return fmt.Sprintf("stat %s", s.Name)
}

func (l *lookup) seedStats(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/stats.json"

	var stats []Stat
	err := loadJSONFile(string(srcPath), &stats)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, stat := range stats {
			dbStat, err := qtx.CreateStat(context.Background(), database.CreateStatParams{
				DataHash: generateDataHash(stat),
				Name:     stat.Name,
				Effect:   stat.Effect,
				MinVal:   stat.MinVal,
				MaxVal:   stat.MaxVal,
				MaxVal2:  getNullInt32(stat.MaxVal2),
			})
			if err != nil {
				return getErr(stat.Error(), err, "couldn't create stat")
			}

			stat.ID = dbStat.ID
			l.stats[stat.Name] = stat
		}
		return nil
	})
}

func (l *lookup) seedStatsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/stats.json"

	var stats []Stat
	err := loadJSONFile(string(srcPath), &stats)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonStat := range stats {
			stat, err := l.getStat(jsonStat.Name)
			if err != nil {
				return err
			}

			stat.SphereID, err = assignFKPtr(&jsonStat.Sphere, l.getItem)
			if err != nil {
				return getErr(stat.Error(), err)
			}

			err = qtx.UpdateStat(context.Background(), database.UpdateStatParams{
				DataHash: generateDataHash(stat),
				SphereID: getNullInt32(stat.SphereID),
				ID:       stat.ID,
			})
			if err != nil {
				return getErr(stat.Error(), err, "couldn't update stat")
			}

			l.stats[stat.Name] = stat
		}
		return nil
	})
}
