package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Property struct {
	//id 		int32
	//dataHash	string
	Name			string 		`json:"name"`
	Effect			string 		`json:"effect"`
	NullifyArmored 	*string		`json:"nullify_armored"`
}

func(p Property) ToHashFields() []any {
	return []any{
		p.Name,
		p.Effect,
		derefOrNil(p.NullifyArmored),
	}
}


func seedProperties(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/properties.json"

	var properties []Property
	err := loadJSONFile(string(srcPath), &properties)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, property := range properties {
			err = qtx.CreateProperty(context.Background(), database.CreatePropertyParams{
				DataHash: 		generateDataHash(property),
				Name: 			property.Name,
				Effect: 		property.Effect,
				NullifyArmored: nullNullifyArmored(property.NullifyArmored),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Property: %s: %v", property.Name, err)
			}
		}
		return nil
	})
}