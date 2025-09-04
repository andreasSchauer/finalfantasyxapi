package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type DefaultAbilitiesEntry struct {
	//id 		int32
	//dataHash	string
	Name				string 		`json:"name"`
}

func(d DefaultAbilitiesEntry) ToHashFields() []any {
	return []any{
		d.Name,
	}
}


func seedDefaultAbilitiesEntries(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/default_abilities.json"

	var entries []DefaultAbilitiesEntry
	err := loadJSONFile(string(srcPath), &entries)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, entry := range entries {
			err = qtx.CreateDefaultAbilitesEntry(context.Background(), database.CreateDefaultAbilitesEntryParams{
				DataHash: 				generateDataHash(entry),
				Name: 					entry.Name,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Default Abilities Entry: %s: %v", entry.Name, err)
			}
		}
		return nil
	})
}