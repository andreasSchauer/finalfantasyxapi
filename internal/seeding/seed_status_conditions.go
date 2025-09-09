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
	Name			string 		`json:"name"`
	Effect			string 		`json:"effect"`
	NullifyArmored 	*string		`json:"nullify_armored"`
}

func(s StatusCondition) ToHashFields() []any {
	return []any{
		s.Name,
		s.Effect,
		derefOrNil(s.NullifyArmored),
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
		for _, condition := range statusConditions {
			err = qtx.CreateStatusCondition(context.Background(), database.CreateStatusConditionParams{
				DataHash: 		generateDataHash(condition),
				Name: 			condition.Name,
				Effect: 		condition.Effect,
				NullifyArmored: nullNullifyArmored(condition.NullifyArmored),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Status Condition: %s: %v", condition.Name, err)
			}
		}
		return nil
	})
}