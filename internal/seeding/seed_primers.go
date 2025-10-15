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
	Name          string `json:"name"`
	AlBhedLetter  string `json:"al_bhed_letter"`
	EnglishLetter string `json:"english_letter"`
	KeyItemID     int32
}

func (p Primer) ToHashFields() []any {
	return []any{
		p.AlBhedLetter,
		p.EnglishLetter,
		p.KeyItemID,
	}
}


func (l *lookup) seedPrimers(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/primers.json"

	var primers []Primer
	err := loadJSONFile(string(srcPath), &primers)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, primer := range primers {
			var err error
			
			primer.KeyItemID, err = assignFK(primer.Name, l.getKeyItem)
			if err != nil {
				return err
			}

			err = qtx.CreatePrimer(context.Background(), database.CreatePrimerParams{
				DataHash:      generateDataHash(primer),
				KeyItemID:     primer.KeyItemID,
				AlBhedLetter:  primer.AlBhedLetter,
				EnglishLetter: primer.EnglishLetter,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Primer: %s: %v", primer.Name, err)
			}
		}
		return nil
	})
}
