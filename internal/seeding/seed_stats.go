package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Stat struct {
	//id 		int32
	//dataHash	string
	Name		string 	`json:"name"`
	Effect		string 	`json:"effect"`
	MinVal		int32	`json:"min_val"`
	MaxVal		int32	`json:"max_val"`
	MaxVal2		*int32	`json:"max_val_2"`
}

func(s Stat) ToHashFields() []any {
	return []any{
		s.Name,
		s.Effect,
		s.MinVal,
		s.MaxVal,
		derefOrNil(s.MaxVal2),
	}
}


func seedStats(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/stats.json"

	var stats []Stat
	err := loadJSONFile(string(srcPath), &stats)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, stat := range stats {
			err = qtx.CreateStat(context.Background(), database.CreateStatParams{
				DataHash: 	generateDataHash(stat.ToHashFields()),
				Name: 		stat.Name,
				Effect: 	stat.Effect,
				MinVal: 	stat.MinVal,
				MaxVal: 	stat.MaxVal,
				MaxVal2: 	getNullInt32(stat.MaxVal2),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Stat: %s: %v", stat.Name, err)
			}
		}
		return nil
	})
}



