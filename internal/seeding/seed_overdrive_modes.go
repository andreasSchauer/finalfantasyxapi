package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type OverdriveMode struct {
	//id 		int32
	//dataHash	string
	Name		string 		`json:"name"`
	Description	string		`json:"description"`
	Effect		string 		`json:"effect"`
	Type		string		`json:"type"`
	FillRate	*float32	`json:"fill_rate"`
}

func(o OverdriveMode) ToHashFields() []any {
	return []any{
		o.Name,
		o.Description,
		o.Effect,
		o.Type,
		derefOrNil(o.FillRate),
	}
}


func seedOverdriveModes(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrive_modes.json"

	var overdriveModes []OverdriveMode
	err := loadJSONFile(string(srcPath), &overdriveModes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, mode := range overdriveModes {
			err = qtx.CreateOverdriveMode(context.Background(), database.CreateOverdriveModeParams{
				DataHash: 		generateDataHash(mode),
				Name: 			mode.Name,
				Description: 	mode.Description,
				Effect: 		mode.Effect,
				Type:			database.OverdriveType(mode.Type),
				FillRate: 		getNullFloat64(mode.FillRate),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Overdrive Mode: %s: %v", mode.Name, err)
			}
		}
		return nil
	})
}



