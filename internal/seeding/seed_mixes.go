package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Mix struct {
	ID                   int32
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

func (m Mix) GetID() int32 {
	return m.ID
}

func (m Mix) Error() string {
	return fmt.Sprintf("mix %s", m.Name)
}

type MixCombination struct {
	ID           int32
	MixID        int32
	FirstItem    string `json:"first_item"`
	SecondItem   string `json:"second_item"`
	FirstItemID  int32
	SecondItemID int32
	IsBestCombo  bool
}

func (m MixCombination) ToHashFields() []any {
	return []any{
		m.MixID,
		m.FirstItemID,
		m.SecondItemID,
		m.IsBestCombo,
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

func (m MixCombination) Error() string {
	return fmt.Sprintf("mix combination with first item: %s, second item: %s", m.FirstItem, m.SecondItem)
}

func (l *Lookup) seedMixes(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/mixes.json"

	var mixes []Mix
	err := loadJSONFile(string(srcPath), &mixes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, mix := range mixes {
			var err error

			key := LookupObject{
				Name: mix.Name,
			}

			mix.OverdriveID, err = assignFK(key, l.Overdrives)
			if err != nil {
				return h.GetErr(mix.Error(), err)
			}

			dbMix, err := qtx.CreateMix(context.Background(), database.CreateMixParams{
				DataHash:    generateDataHash(mix),
				OverdriveID: mix.OverdriveID,
				Category:    database.MixCategory(mix.Category),
			})
			if err != nil {
				return h.GetErr(mix.Error(), err, "couldn't create mix")
			}

			mix.ID = dbMix.ID
			l.Mixes[mix.Name] = mix
		}

		return nil
	})
}

func (l *Lookup) seedMixesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/mixes.json"

	var mixes []Mix
	err := loadJSONFile(string(srcPath), &mixes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonMix := range mixes {
			mix, err := GetResource(jsonMix.Name, l.Mixes)
			if err != nil {
				return err
			}

			err = l.seedMixCombinations(qtx, mix)
			if err != nil {
				return h.GetErr(mix.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedMixCombinations(qtx *database.Queries, mix Mix) error {
	bestComboMap := getBestComboMap(mix)

	for _, combo := range mix.PossibleCombinations {
		var err error
		combo.MixID = mix.ID

		key := CreateLookupKey(combo)
		if _, exists := bestComboMap[key]; exists {
			combo.IsBestCombo = true
		}

		_, err = seedObjAssignID(qtx, combo, l.seedMixCombination)
		if err != nil {
			return err
		}
	}

	return nil
}

func getBestComboMap(mix Mix) map[string]struct{} {
	bestComboMap := make(map[string]struct{})

	for _, combo := range mix.BestCombinations {
		key := CreateLookupKey(combo)
		bestComboMap[key] = struct{}{}
	}

	return bestComboMap
}

func (l *Lookup) seedMixCombination(qtx *database.Queries, combo MixCombination) (MixCombination, error) {
	var err error

	combo.FirstItemID, err = assignFK(combo.FirstItem, l.Items)
	if err != nil {
		return MixCombination{}, h.GetErr(combo.Error(), err)
	}

	combo.SecondItemID, err = assignFK(combo.SecondItem, l.Items)
	if err != nil {
		return MixCombination{}, h.GetErr(combo.Error(), err)
	}

	if combo.FirstItemID > combo.SecondItemID {
		temp := combo.FirstItemID
		combo.FirstItemID = combo.SecondItemID
		combo.SecondItemID = temp
	}

	dbCombo, err := qtx.CreateMixCombination(context.Background(), database.CreateMixCombinationParams{
		DataHash:     generateDataHash(combo),
		MixID:        combo.MixID,
		FirstItemID:  combo.FirstItemID,
		SecondItemID: combo.SecondItemID,
		IsBestCombo:  combo.IsBestCombo,
	})
	if err != nil {
		return MixCombination{}, h.GetErr(combo.Error(), err, "couldn't create mix combination")
	}

	combo.ID = dbCombo.ID

	return combo, nil
}
