package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type TreasureList struct {
	LocationArea LocationArea `json:"location_area"`
	Treasures    []Treasure   `json:"treasures"`
}

type Treasure struct {
	//id 				int32
	//dataHash			string
	Version         int32
	AreaID          int32
	TreasureType    string  `json:"treasure_type"`
	LootType        string  `json:"loot_type"`
	IsPostAirship   bool    `json:"is_post_airship"`
	IsAnimaTreasure bool    `json:"is_anima_treasure"`
	Notes           *string `json:"notes"`
	GilAmount       *int32  `json:"gil_amount"`
}

func (t Treasure) ToHashFields() []any {
	return []any{
		t.Version,
		t.AreaID,
		t.TreasureType,
		t.LootType,
		t.IsPostAirship,
		t.IsAnimaTreasure,
		derefOrNil(t.Notes),
		derefOrNil(t.GilAmount),
	}
}

func (l *lookup) seedTreasures(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/treasures.json"

	var treasureLists []TreasureList
	err := loadJSONFile(string(srcPath), &treasureLists)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, list := range treasureLists {
			locationArea := list.LocationArea
			locationAreaID, err := l.getAreaID(locationArea)
			if err != nil {
				return fmt.Errorf("treasures: %v", err)
			}

			for j, treasure := range list.Treasures {
				treasure.AreaID = locationAreaID
				treasure.Version = int32(j + 1)

				err = qtx.CreateTreasure(context.Background(), database.CreateTreasureParams{
					DataHash:        generateDataHash(treasure),
					AreaID:          treasure.AreaID,
					Version:         treasure.Version,
					TreasureType:    database.TreasureType(treasure.TreasureType),
					LootType:        database.LootType(treasure.LootType),
					IsPostAirship:   treasure.IsPostAirship,
					IsAnimaTreasure: treasure.IsAnimaTreasure,
					Notes:           getNullString(treasure.Notes),
					GilAmount:       getNullInt32(treasure.GilAmount),
				})
				if err != nil {
					return fmt.Errorf("couldn't create Treasure: %s - %s - %s - %d - %d: %v", locationArea.Location, locationArea.SubLocation, locationArea.Area, derefOrNil(locationArea.Version), treasure.Version, err)
				}
			}
		}
		
		return nil
	})
}
