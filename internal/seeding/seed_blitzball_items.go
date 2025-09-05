package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type BlitzballItemList struct {
	//id 		int32
	//dataHash	string
	Category	string		`json:"category"`
	Slot		string		`json:"slot"`
}

func(b BlitzballItemList) ToHashFields() []any {
	return []any{
		b.Category,
		b.Slot,
	}
}


func seedBlitzballItems(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/blitzball_items.json"

	var blitzballItemLists []BlitzballItemList
	err := loadJSONFile(string(srcPath), &blitzballItemLists)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, list := range blitzballItemLists {
			err = qtx.CreateBlitzballItemList(context.Background(), database.CreateBlitzballItemListParams{
				DataHash: 		generateDataHash(list),
				Category: 		database.BlitzballTournamentCategory((list.Category)),
				Slot: 			database.BlitzballItemSlot(list.Slot),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Blitzball Item List: %s - %s: %v", list.Category, list.Slot, err)
			}
		}
		return nil
	})
}