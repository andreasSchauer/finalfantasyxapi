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

func (tl TreasureList) Error() string {
	return fmt.Sprintf("treasures at %s", tl.LocationArea)
}

type Treasure struct {
	ID				int32
	Version         int32
	AreaID          int32
	TreasureType    string  		`json:"treasure_type"`
	LootType        string  		`json:"loot_type"`
	IsPostAirship   bool    		`json:"is_post_airship"`
	IsAnimaTreasure bool    		`json:"is_anima_treasure"`
	Notes           *string 		`json:"notes"`
	GilAmount       *int32  		`json:"gil_amount"`
	Items			[]ItemAmount	`json:"items"`
	Equipment		*FoundEquipment	`json:"equipment"`
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
		ObjPtrToHashID(t.Equipment),
	}
}


func (t Treasure) ToKeyFields() []any {
	return []any{
		t.AreaID,
		t.Version,
	}
}


func (t Treasure) GetID() int32 {
	return t.ID
}

func (t Treasure) Error() string {
	return fmt.Sprintf("treasure number: %d", t.Version)
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
			var err error

			locationArea := list.LocationArea
			list.LocationArea.ID, err = assignFK(locationArea, l.getArea)
			if err != nil {
				return getErr(list.Error(), err)
			}

			for j, treasure := range list.Treasures {
				treasure.AreaID = list.LocationArea.ID
				treasure.Version = int32(j + 1)

				dbTreasure, err := qtx.CreateTreasure(context.Background(), database.CreateTreasureParams{
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
					return getErr(treasure.Error(), err, "couldn't create treasure")
				}

				treasure.ID = dbTreasure.ID
				key := createLookupKey(treasure)
				l.treasures[key] = treasure
			}
		}

		return nil
	})
}


func (l *lookup) seedTreasuresRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/treasures.json"

	var treasureLists []TreasureList
	err := loadJSONFile(string(srcPath), &treasureLists)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, list := range treasureLists {
			list.LocationArea.ID, err = assignFK(list.LocationArea, l.getArea)
			if err != nil {
				return getErr(list.Error(), err)
			}

			for j, jsonTreasure := range list.Treasures {
				jsonTreasure.AreaID = list.LocationArea.ID
				jsonTreasure.Version = int32(j + 1)
				key := createLookupKey(jsonTreasure)

				treasure, err := l.getTreasure(key)
				if err != nil {
					return err
				}

				err = l.seedTreasureItemAmounts(qtx, treasure)
				if err != nil {
					return getErr(treasure.Error(), err)
				}

				treasure.Equipment, err = seedObjPtrAssignFK(qtx, treasure.Equipment, l.seedFoundEquipment)
				if err != nil {
					return getErr(treasure.Error(), err)
				}

				err = qtx.UpdateTreasure(context.Background(), database.UpdateTreasureParams{
					DataHash:        	generateDataHash(treasure),
					FoundEquipmentID: 	ObjPtrToNullInt32ID(treasure.Equipment),
					ID: 				treasure.ID,
				})
				if err != nil {
					return getErr(treasure.Error(), err, "couldn't update treasure")
				}
			}
		}
		return nil
	})
}


func (l *lookup) seedTreasureItemAmounts(qtx *database.Queries, treasure Treasure) error {
	for _, itemAmount := range treasure.Items {
		junction, err := createJunctionSeed(qtx, treasure, itemAmount, l.seedItemAmount)
		if err != nil {
			return err
		}

		err = qtx.CreateTreasuresItemsJunction(context.Background(), database.CreateTreasuresItemsJunctionParams{
			DataHash: generateDataHash(junction),
			TreasureID: junction.ParentID,
			ItemAmountID: junction.ChildID,
		})
		if err != nil {
			return getErr(itemAmount.Error(), err, "couldn't junction item amount")
		}
	}

	return nil
}