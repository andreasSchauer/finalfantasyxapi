package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop5SeedMixes(qtx *database.Queries, ctx context.Context) error {
	mixes, err := extractMixes()
	if err != nil {
		return err
	}

	params := database.CreateMixBulkParams{
		DataHash:    make([]string, len(mixes)),
		OverdriveID: make([]int32, len(mixes)),
		Category:    make([]database.MixCategory, len(mixes)),
	}

	for i, m := range mixes {
		params.DataHash[i] = generateDataHash(m)
		params.OverdriveID[i] = m.OverdriveID
		params.Category[i] = database.MixCategory(m.Category)
	}

	dbRows, err := qtx.CreateMixBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create mixes: %v", err)
	}

	for i, row := range dbRows {
		mixes[i].ID = row.ID
		rawJSON.mixes[i].ID = row.ID
		lookup.L.Mixes[seedtypes.Key(mixes[i])] = mixes[i]
		lookup.L.MixesID[row.ID] = mixes[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMixes() ([]seedtypes.Mix, error) {
	mixes := []seedtypes.Mix{}
	var err error

	for i := range rawJSON.mixes {
		mix := &rawJSON.mixes[i]

		obj := seedtypes.LookupObject{
			Name: mix.Name,
		}

		mix.OverdriveID, err = assignFK(seedtypes.Key(obj), lookup.L.Overdrives)
		if err != nil {
			return nil, err
		}

		mixes = append(mixes, *mix)
	}

	return dedupeRows(mixes, hashes), nil
}

func completeMixes() error {
	for i := range rawJSON.mixes {
		mix := &rawJSON.mixes[i]

		err := assignIDs(mix.PossibleCombinations)
		if err != nil {
			return err
		}

		lookup.L.Mixes[seedtypes.Key(mix)] = *mix
		lookup.L.MixesID[mix.ID] = *mix
	}

	return nil
}

func loop6SeedMixCombinations(qtx *database.Queries, ctx context.Context) error {
	combos, err := extractMixCombinations()
	if err != nil {
		return err
	}

	params := database.CreateMixCombinationBulkParams{
		DataHash:     make([]string, len(combos)),
		MixID:        make([]int32, len(combos)),
		FirstItemID:  make([]int32, len(combos)),
		SecondItemID: make([]int32, len(combos)),
		IsBestCombo:  make([]bool, len(combos)),
	}

	for i, c := range combos {
		params.DataHash[i] = generateDataHash(c)
		params.MixID[i] = c.MixID
		params.FirstItemID[i] = c.FirstItemID
		params.SecondItemID[i] = c.SecondItemID
		params.IsBestCombo[i] = c.IsBestCombo
	}

	dbRows, err := qtx.CreateMixCombinationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create mix combinations: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMixCombinations() ([]seedtypes.MixCombination, error) {
	combos := []seedtypes.MixCombination{}
	var err error

	for i := range rawJSON.mixes {
		mix := &rawJSON.mixes[i]
		bestComboMap := getBestComboMap(*mix)

		for j := range mix.PossibleCombinations {
			combo := &mix.PossibleCombinations[j]
			combo.MixID = mix.ID

			key := seedtypes.Key(combo)
			combo.IsBestCombo = bestComboMap[key]

			combo.FirstItemID, err = assignFK(combo.FirstItem, lookup.L.Items)
			if err != nil {
				return nil, err
			}

			combo.SecondItemID, err = assignFK(combo.SecondItem, lookup.L.Items)
			if err != nil {
				return nil, err
			}

			if combo.FirstItemID > combo.SecondItemID {
				temp := combo.FirstItemID
				combo.FirstItemID = combo.SecondItemID
				combo.SecondItemID = temp
			}

			combos = append(combos, *combo)
		}
	}

	return dedupeRows(combos, hashes), nil
}

func getBestComboMap(mix seedtypes.Mix) map[string]bool {
	bestComboMap := make(map[string]bool)

	for _, combo := range mix.BestCombinations {
		key := seedtypes.Key(combo)
		bestComboMap[key] = true
	}

	return bestComboMap
}
