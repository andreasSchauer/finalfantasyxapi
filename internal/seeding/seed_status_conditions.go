package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type StatusCondition struct {
	//id 		int32
	//dataHash	string
	Name		string 		`json:"name"`
	Effect		string 		`json:"effect"`
}

func(s StatusCondition) ToHashFields() []any {
	return []any{
		s.Name,
		s.Effect,
	}
}


func seedStatusConditions(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/status_conditions.json"

	var statusConditions []StatusCondition
	err := loadJSONFile(string(srcPath), &statusConditions)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, sc := range statusConditions {
			err = qtx.CreateStatusCondition(context.Background(), database.CreateStatusConditionParams{
				DataHash: 	generateDataHash(sc.ToHashFields()),
				Name: 		sc.Name,
				Effect: 	sc.Effect,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Status Condition: %s: %v", sc.Name, err)
			}
		}
		return nil
	})
}