package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Element struct {
	//id 			int32
	//dataHash		string
	Name string `json:"name"`
}

func (e Element) ToHashFields() []any {
	return []any{
		e.Name,
	}
}

func (l *lookup) seedElements(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/elements.json"

	var elements []Element
	err := loadJSONFile(string(srcPath), &elements)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, element := range elements {
			err = qtx.CreateElement(context.Background(), database.CreateElementParams{
				DataHash: generateDataHash(element),
				Name:     element.Name,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Element: %s: %v", element.Name, err)
			}
		}
		return nil
	})
}
