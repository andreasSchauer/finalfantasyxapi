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
	OverdriveID          int32
	Name                 string           `json:"name"`
	Category             string           `json:"category"`
	BestCombinations     []MixCombination `json:"best_combinations"`
	PossibleCombinations []MixCombination `json:"possible_combinations"`
}

func (m Mix) ToHashFields() []any {
	return []any{
		m.OverdriveID,
		m.Category,
	}
}

type MixCombination struct {
	FirstItem    string `json:"first_item"`
	SecondItem   string `json:"second_item"`
	FirstItemID  int32
	SecondItemID int32
	IsBestCombo  bool
}

func (m MixCombination) ToHashFields() []any {
	return []any{
		m.FirstItemID,
		m.SecondItemID,
	}
}


func (m MixCombination) ToKeyFields() []any {
	return []any{
		m.FirstItem,
		m.SecondItem,
	}
}


type MixComboJunction struct {
	MixID       int32
	ComboID     int32
	IsBestCombo bool
}

func (m MixComboJunction) ToHashFields() []any {
	return []any{
		m.MixID,
		m.ComboID,
		m.IsBestCombo,
	}
}

func (l *lookup) seedMixes(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/mixes.json"

	var mixes []Mix
	err := loadJSONFile(string(srcPath), &mixes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, mix := range mixes {
			ability := Ability{
				Name: mix.Name,
			}

			overdrive, err := l.getOverdrive(ability)
			if err != nil {
				return err
			}

			mix.OverdriveID = overdrive.ID

			dbMix, err := qtx.CreateMix(context.Background(), database.CreateMixParams{
				DataHash:    generateDataHash(mix),
				OverdriveID: mix.OverdriveID,
				Category:    database.MixCategory(mix.Category),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Mix: %s: %v", mix.Name, err)
			}

			err = l.seedMixCombinations(qtx, mix, dbMix.ID)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (l *lookup) seedMixCombinations(qtx *database.Queries, mix Mix, dbMixID int32) error {
	bestComboMap := getBestComboMap(mix)

	for _, combo := range mix.PossibleCombinations {
		junction := MixComboJunction{
			MixID: dbMixID,
		}
		key := createLookupKey(combo)
		if _, exists := bestComboMap[key]; exists {
			junction.IsBestCombo = true
		}

		dbCombo, err := l.seedMixCombination(qtx, combo)
		if err != nil {
			return err
		}

		junction.ComboID = dbCombo.ID

		err = qtx.CreateMixComboJunction(context.Background(), database.CreateMixComboJunctionParams{
			DataHash:    generateDataHash(junction),
			MixID:       junction.MixID,
			ComboID:     junction.ComboID,
			IsBestCombo: junction.IsBestCombo,
		})
		if err != nil {
			return fmt.Errorf("couldn't create Junction between Mix: %s and Combo %s-%s: %v", mix.Name, combo.FirstItem, combo.SecondItem, err)
		}
	}

	return nil
}


func getBestComboMap(mix Mix) map[string]struct{} {
	bestComboMap := make(map[string]struct{})
	for _, combo := range mix.BestCombinations {
		key := createLookupKey(combo)
		bestComboMap[key] = struct{}{}
	}

	return bestComboMap
}


func (l *lookup) seedMixCombination(qtx *database.Queries, combo MixCombination) (database.MixCombination, error) {
	firstItem, err := l.getItem(combo.FirstItem)
	if err != nil {
		return database.MixCombination{}, err
	}

	secondItem, err := l.getItem(combo.SecondItem)
	if err != nil {
		return database.MixCombination{}, err
	}

	combo.FirstItemID = firstItem.ID
	combo.SecondItemID = secondItem.ID

	dbCombo, err := qtx.CreateMixCombination(context.Background(), database.CreateMixCombinationParams{
		DataHash:     generateDataHash(combo),
		FirstItemID:  combo.FirstItemID,
		SecondItemID: combo.SecondItemID,
	})
	if err != nil {
		return database.MixCombination{}, fmt.Errorf("couldn't create Mix Combination: %s + %s: %v", combo.FirstItem, combo.SecondItem, err)
	}

	return dbCombo, nil
}
