package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type TreasureList struct {
	Treasures []Treasure	`json:"treasures"`
}


type Treasure struct {
	//id 				int32
	//dataHash			string
	TreasureListID		int32
	TreasureType		string					`json:"treasure_type"`
	LootType			string					`json:"loot_type"`
	IsPostAirship		bool					`json:"is_post_airship"`
	IsAnimaTreasure		bool					`json:"is_anima_treasure"`
	Notes				*string					`json:"notes"`
	GilAmount			*int32					`json:"gil_amount"`
}

func(t Treasure) ToHashFields() []any {
	return []any{
		t.TreasureListID,
		t.TreasureType,
		t.LootType,
		t.IsPostAirship,
		t.IsAnimaTreasure,
		derefOrNil(t.Notes),
		derefOrNil(t.GilAmount),
	}
}


func seedTreasures(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/treasures.json"

	var treasureLists []TreasureList
	err := loadJSONFile(string(srcPath), &treasureLists)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for i, list := range treasureLists {
			listID, err := qtx.CreateTreasureList(context.Background())
			if err != nil {
				return fmt.Errorf("couldn't create Treasure List: %d: %v", i, err)
			}

			for j, treasure := range list.Treasures {
				treasure.TreasureListID = listID
				
				err = qtx.CreateTreasure(context.Background(), database.CreateTreasureParams{
					DataHash: 			generateDataHash(treasure),
					TreasureListID: 	treasure.TreasureListID,
					TreasureType: 		database.TreasureType(treasure.TreasureType),
					LootType: 			database.LootType(treasure.LootType),
					IsPostAirship: 		treasure.IsPostAirship,
					IsAnimaTreasure: 	treasure.IsAnimaTreasure,
					Notes: 				getNullString(treasure.Notes),
					GilAmount: 			getNullInt32(treasure.GilAmount),
				})
				if err != nil {
					return fmt.Errorf("couldn't create Treasure: %d - %d: %v", i, j, err)
				}
			}
		}
		return nil
	})
}