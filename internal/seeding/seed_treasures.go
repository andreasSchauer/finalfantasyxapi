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


type FoundEquipment struct {
	ID					int32
	EquipmentNameID		int32
	Name				string		`json:"name"`
	Abilities			[]string	`json:"abilities"`
	EmptySlotsAmount	int32		`json:"empty_slots_amount"`
}


func (f FoundEquipment) ToHashFields() []any {
	return []any{
		f.EquipmentNameID,
		f.EmptySlotsAmount,
	}
}


func (f FoundEquipment) GetID() int32 {
	return f.ID
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
				return fmt.Errorf("monster formations: %v", err)
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
					return fmt.Errorf("couldn't create Treasure: %s - treasure version: %d: %v", createLookupKey(locationArea), treasure.Version, err)
				}

				treasure.ID = dbTreasure.ID
				key := createLookupKey(treasure)
				l.treasures[key] = treasure
			}
		}

		return nil
	})
}


func (l *lookup) createTreasuresRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/treasures.json"

	var treasureLists []TreasureList
	err := loadJSONFile(string(srcPath), &treasureLists)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, list := range treasureLists {
			list.LocationArea.ID, err = assignFK(list.LocationArea, l.getArea)

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
					return fmt.Errorf("treasure: %s: %v", createLookupKey(treasure), err)
				}

				treasure.Equipment, err = seedObjPtrAssignFK(qtx, treasure.Equipment, l.seedFoundEquipment)
				if err != nil {
					return fmt.Errorf("treasure: %s: %v", createLookupKey(treasure), err)
				}

				err = qtx.UpdateTreasure(context.Background(), database.UpdateTreasureParams{
					DataHash:        	generateDataHash(treasure),
					AreaID:          	treasure.AreaID,
					Version:         	treasure.Version,
					TreasureType:    	database.TreasureType(treasure.TreasureType),
					LootType:        	database.LootType(treasure.LootType),
					IsPostAirship:   	treasure.IsPostAirship,
					IsAnimaTreasure: 	treasure.IsAnimaTreasure,
					Notes:           	getNullString(treasure.Notes),
					GilAmount:       	getNullInt32(treasure.GilAmount),
					FoundEquipmentID: 	ObjPtrToNullInt32ID(treasure.Equipment),
					ID: 				treasure.ID,
				})
				if err != nil {
					return fmt.Errorf("couldn't update Treasure: %s - treasure version: %d: %v", createLookupKey(list.LocationArea), treasure.Version, err)
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

		err = qtx.CreateTreasureItemAmountJunction(context.Background(), database.CreateTreasureItemAmountJunctionParams{
			DataHash: generateDataHash(junction),
			TreasureID: junction.ParentID,
			ItemAmountID: junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create junction for item Amount: %s: %v", createLookupKey(itemAmount), err)
		}
	}

	return nil
}


func (l *lookup) seedFoundEquipment(qtx *database.Queries, foundEquipment FoundEquipment) (FoundEquipment, error) {
	var err error

	foundEquipment.EquipmentNameID, err = assignFK(foundEquipment.Name, l.getEquipmentName)
	if err != nil {
		return FoundEquipment{}, err
	}

	dbFoundEquipment, err := qtx.CreateFoundEquipmentPiece(context.Background(), database.CreateFoundEquipmentPieceParams{
		DataHash: generateDataHash(foundEquipment),
		EquipmentNameID: foundEquipment.EquipmentNameID,
		EmptySlotsAmount: foundEquipment.EmptySlotsAmount,
	})
	if err != nil {
		return FoundEquipment{}, fmt.Errorf("couldn't create found equipment %s: %v", foundEquipment.Name, err)
	}

	foundEquipment.ID = dbFoundEquipment.ID

	err = l.seedFoundEquipmentAbilities(qtx, foundEquipment)
	if err != nil {
		return FoundEquipment{}, fmt.Errorf("found equipment: %s: %v", foundEquipment.Name, err)
	}

	return foundEquipment, nil
}


func (l *lookup) seedFoundEquipmentAbilities(qtx *database.Queries, foundEquipment FoundEquipment) error {
	for _, autoAbility := range foundEquipment.Abilities {
		junction, err := createJunction(foundEquipment, autoAbility, l.getAutoAbility)
		if err != nil {
			return fmt.Errorf("couldn't create junction with auto ability %s: %v", autoAbility, err)
		}

		err = qtx.CreateFoundEquipmentAutoAbilityJunction(context.Background(), database.CreateFoundEquipmentAutoAbilityJunctionParams{
			DataHash: generateDataHash(junction),
			FoundEquipmentID: junction.ParentID,
			AutoAbilityID: junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't seed junction with auto ability %s: %v", autoAbility, err)
		}
	}

	return nil
}