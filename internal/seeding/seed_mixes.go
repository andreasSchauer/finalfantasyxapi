package seeding

import (
	"context"
	"database/sql"
	"fmt"
	
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Mix struct {
	//id 			int32
	//dataHash		string
	OverdriveID				int32
	Name         			string  		   `json:"name"`
	Category				string   		   `json:"category"`
	BestCombinations		[]MixCombination   `json:"best_combinations"`
	PossibleCombinations	[]MixCombination   `json:"possible_combinations"`
}

func (m Mix) ToHashFields() []any {
	return []any{
		m.OverdriveID,
		m.Category,
	}
}


type MixCombination struct {
	FirstItem		string   `json:"first_item"`
	SecondItem		string   `json:"second_item"`
	FirstItemID		int32
	SecondItemID	int32
	IsBestCombo		bool
}


func (m MixCombination) ToHashFields() []any {
	return []any{
		m.FirstItemID,
		m.SecondItemID,
	}
}


type MixComboJunction struct {
	MixID		int32
	ComboID		int32
	IsBestCombo	bool
}


func (m MixComboJunction) ToHashFields() []any {
	return []any{
		m.MixID,
		m.ComboID,
		m.IsBestCombo,
	}
}


func seedMixes(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/mixes.json"

	var mixes []Mix
	err := loadJSONFile(string(srcPath), &mixes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, mix := range mixes {
			dbOverdrive, err := qtx.GetOverdriveByName(context.Background(), mix.Name)
			if err != nil {
				return err
			}

			mix.OverdriveID = dbOverdrive.ID

			dbMix, err := qtx.CreateMix(context.Background(), database.CreateMixParams{
				DataHash:     	generateDataHash(mix),
				OverdriveID: 	mix.OverdriveID,
				Category: 		database.MixCategory(mix.Category),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Mix: %s: %v", mix.Name, err)
			}

			items, err := qtx.GetItems(context.Background())
			if err != nil {
				return err
			}

			itemNameToID := make(map[string]int32, len(items))
			for _, item := range items {
				itemNameToID[*convertNullString(item.Name)] = item.ItemID
			}

			err = seedMixCombinations(qtx, mix, dbMix.ID, itemNameToID)
			if err != nil {
				return err
			}
		}

		return nil
	})
}




func seedMixCombinations(qtx *database.Queries, mix Mix, dbMixID int32, lookup map[string]int32) error {
	bestComboMap := make(map[string]struct{})
	for _, combo := range mix.BestCombinations {
		key := combo.FirstItem + "|" + combo.SecondItem
		bestComboMap[key] = struct{}{}
	}

	for _, combo := range mix.PossibleCombinations {
		mixComboJunction := MixComboJunction{
			MixID: dbMixID,
		}
		key := combo.FirstItem + "|" + combo.SecondItem
		if _, exists := bestComboMap[key]; exists {
			mixComboJunction.IsBestCombo = true
		}

		dbCombo, err := seedMixCombination(qtx, combo, lookup)
		if err != nil {
			return err
		}

		mixComboJunction.ComboID = dbCombo.ID

		err = qtx.CreateMixComboJunction(context.Background(), database.CreateMixComboJunctionParams{
			DataHash: generateDataHash(mixComboJunction),
			MixID: mixComboJunction.MixID,
			ComboID: mixComboJunction.ComboID,
			IsBestCombo: mixComboJunction.IsBestCombo,
		})
		if err != nil {
			return fmt.Errorf("couldn't create Junction between Mix: %s and Combo %s-%s: %v", mix.Name, combo.FirstItem, combo.SecondItem, err)
		}
	}

	return nil
}



func seedMixCombination(qtx *database.Queries, combo MixCombination, lookup map[string]int32) (database.MixCombination, error) {
	firstItemID, found := lookup[combo.FirstItem]
	if !found {
		return database.MixCombination{}, fmt.Errorf("couldn't find Item %s", combo.FirstItem)
	}

	secondItemID, found := lookup[combo.SecondItem]
	if !found {
		return database.MixCombination{}, fmt.Errorf("couldn't find Item %s", combo.FirstItem)
	}

	combo.FirstItemID = firstItemID
	combo.SecondItemID = secondItemID

	dbCombo, err := qtx.CreateMixCombination(context.Background(), database.CreateMixCombinationParams{
		DataHash: generateDataHash(combo),
		FirstItemID: combo.FirstItemID,
		SecondItemID: combo.SecondItemID,
	})
	if err != nil {
		return database.MixCombination{}, fmt.Errorf("couldn't create Mix Combination: %s + %s: %v", combo.FirstItem, combo.SecondItem, err)
	}

	return dbCombo, nil
}