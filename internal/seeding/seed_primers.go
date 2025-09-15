package seeding

import (
	"context"
	"database/sql"
	"fmt"
	
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Primer struct {
	//id 			int32
	//dataHash		string
	Name         	string   	`json:"name"`
	AlBhedLetter	string 		`json:"al_bhed_letter"`
	EnglishLetter	string		`json:"english_letter"`
	KeyItemID		int32
}

func (p Primer) ToHashFields() []any {
	return []any{
		p.AlBhedLetter,
		p.EnglishLetter,
		p.KeyItemID,
	}
}


func seedPrimers(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/primers.json"

	var primers []Primer
	err := loadJSONFile(string(srcPath), &primers)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		keyItems, err := qtx.GetKeyItems(context.Background())
		if err != nil {
			return err
		}

		keyItemNameToID := make(map[string]int32, len(keyItems))
		for _, keyItem := range keyItems {
			keyItemNameToID[*convertNullString(keyItem.Name)] = keyItem.KeyItemID
		}

		for _, primer := range primers {
			keyItemID, found := keyItemNameToID[primer.Name]
			if !found {
				return fmt.Errorf("couldn't find Key Item %s", primer.Name)
			}

			primer.KeyItemID = keyItemID

			err = qtx.CreatePrimer(context.Background(), database.CreatePrimerParams{
				DataHash:     generateDataHash(primer),
				KeyItemID: 		primer.KeyItemID,
				AlBhedLetter: 	primer.AlBhedLetter,
				EnglishLetter: 	primer.EnglishLetter,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Primer: %s: %v", primer.Name, err)
			}
		}
		return nil
	})
}
