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
	Version				int32
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
		t.Version,
		t.TreasureType,
		t.LootType,
		t.IsPostAirship,
		t.IsAnimaTreasure,
		derefOrNil(t.Notes),
		derefOrNil(t.GilAmount),
	}
}


type TreasureCounter struct {
    Chest  int32
    Gift   int32
    Object int32
}


func (c *TreasureCounter) IncreaseCounter(treasureType string) int32 {
    switch treasureType {
    case "chest":
        c.Chest++
        return c.Chest
    case "gift":
        c.Gift++
        return c.Gift
    case "object":
        c.Object++
        return c.Object
    default:
        return 0
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

			counter := &TreasureCounter{}

			for j, treasure := range list.Treasures {
				treasure.TreasureListID = listID
				treasure.Version = counter.IncreaseCounter(treasure.TreasureType)
				
				err = qtx.CreateTreasure(context.Background(), database.CreateTreasureParams{
					DataHash: 			generateDataHash(treasure),
					TreasureListID: 	treasure.TreasureListID,
					Version:			treasure.Version,
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