package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Mix struct {
	ID						int32
	OverdriveID          	int32
	Name                 	string           `json:"name"`
	Category             	string           `json:"category"`
	BestCombinations     	[]MixCombination `json:"best_combinations"`
	PossibleCombinations 	[]MixCombination `json:"possible_combinations"`
}

func (m Mix) ToHashFields() []any {
	return []any{
		m.OverdriveID,
		m.Category,
	}
}

func (m Mix) GetID() int32 {
	return m.ID
}


type MixCombination struct {
	ID				int32
	FirstItem    	string `json:"first_item"`
	SecondItem   	string `json:"second_item"`
	FirstItemID  	int32
	SecondItemID 	int32
	IsBestCombo  	bool
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

func (m MixCombination) GetID() int32 {
	return m.ID
}

type MixComboJunction struct {
	Junction
	IsBestCombo bool
}

func (m MixComboJunction) ToHashFields() []any {
	return []any{
		m.ParentID,
		m.ChildID,
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

			mix.ID = dbMix.ID
			l.mixes[mix.Name] = mix
		}

		return nil
	})
}


func (l *lookup) createMixesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/mixes.json"

	var mixes []Mix
	err := loadJSONFile(string(srcPath), &mixes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonMix := range mixes {
			mix, err := l.getMix(jsonMix.Name)
			if err != nil {
				return err
			}

			err = l.seedMixCombinations(qtx, mix)
			if err != nil {
				return err
			}
		}

		return nil
	})
}


func (l *lookup) seedMixCombinations(qtx *database.Queries, mix Mix) error {
	bestComboMap := getBestComboMap(mix)

	for _, combo := range mix.PossibleCombinations {
		var err error
		mixJunction := MixComboJunction{}

		key := createLookupKey(combo)
		if _, exists := bestComboMap[key]; exists {
			mixJunction.IsBestCombo = true
		}

		mixJunction.Junction, err = createJunctionSeed(qtx, mix, combo, l.seedMixCombination)
		if err != nil {
			return err
		}

		err = qtx.CreateMixesCombinationsJunction(context.Background(), database.CreateMixesCombinationsJunctionParams{
			DataHash:    generateDataHash(mixJunction),
			MixID:       mixJunction.ParentID,
			ComboID:     mixJunction.ChildID,
			IsBestCombo: mixJunction.IsBestCombo,
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



func (l *lookup) seedMixCombination(qtx *database.Queries, combo MixCombination) (MixCombination, error) {
	var err error
	combo.FirstItemID, err = assignFK(combo.FirstItem, l.getItem)
	if err != nil {
		return MixCombination{}, err
	}

	combo.SecondItemID, err = assignFK(combo.SecondItem, l.getItem)
	if err != nil {
		return MixCombination{}, err
	}

	dbCombo, err := qtx.CreateMixCombination(context.Background(), database.CreateMixCombinationParams{
		DataHash:     generateDataHash(combo),
		FirstItemID:  combo.FirstItemID,
		SecondItemID: combo.SecondItemID,
	})
	if err != nil {
		return MixCombination{}, fmt.Errorf("couldn't create Mix Combination: %s + %s: %v", combo.FirstItem, combo.SecondItem, err)
	}

	combo.ID = dbCombo.ID

	return combo, nil
}
