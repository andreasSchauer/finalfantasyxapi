package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Stat struct {
	ID			int32
	Name    	string `json:"name"`
	Effect  	string `json:"effect"`
	MinVal  	int32  `json:"min_val"`
	MaxVal  	int32  `json:"max_val"`
	MaxVal2 	*int32 `json:"max_val_2"`
	SphereID 	*int32
	Sphere		string	`json:"sphere"`
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


func (s Stat) GetID() *int32 {
	return &s.ID
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
				return fmt.Errorf("couldn't create Stat: %s: %v", stat.Name, err)
			}

			stat.ID = dbStat.ID
			l.stats[stat.Name] = stat
		}
		return nil
	})
}



func (l *lookup) createStatsRelationships(db *database.Queries, dbConn *sql.DB) error {
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

			sphere, err := l.getItem(jsonStat.Sphere)
			if err != nil {
				return err
			}
			stat.SphereID = &sphere.ID

			err = qtx.UpdateStat(context.Background(), database.UpdateStatParams{
				DataHash: 	generateDataHash(stat),
				Name:   	stat.Name,
				Effect: 	stat.Effect,
				MinVal:   	stat.MinVal,
				MaxVal:  	stat.MaxVal,
				MaxVal2:  	getNullInt32(stat.MaxVal2),
				SphereID: 	getNullInt32(stat.SphereID),
				ID: 		stat.ID,
			})
			if err != nil {
				return fmt.Errorf("couldn't update Stat: %s: %v", stat.Name, err)
			}
			
			l.stats[stat.Name] = stat
		}
		return nil
	})
}