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

func (p Primer) Error() string {
	return fmt.Sprintf("primer %s", p.Name)
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
				return getErr(primer.Error(), err)
			}

			err = qtx.CreatePrimer(context.Background(), database.CreatePrimerParams{
				DataHash:      generateDataHash(primer),
				KeyItemID:     primer.KeyItemID,
				AlBhedLetter:  primer.AlBhedLetter,
				EnglishLetter: primer.EnglishLetter,
			})
			if err != nil {
				return getErr(primer.Error(), err, "couldn't create primer")
			}
		}
		return nil
	})
}
