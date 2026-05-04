package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop5SeedMixes(qtx *database.Queries, ctx context.Context) error {
	mixes, err := l.extractMixes()
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
		l.json.mixes[i].ID = row.ID
		l.Mixes[mixes[i].Name] = mixes[i]
		l.MixesID[row.ID] = mixes[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMixes() ([]Mix, error) {
	mixes := []Mix{}
	var err error

	for i := range l.json.mixes {
		mix := &l.json.mixes[i]

		obj := LookupObject{
			Name: mix.Name,
		}

		mix.OverdriveID, err = assignFK(Key(obj), l.Overdrives)
		if err != nil {
			return nil, err
		}

		mixes = append(mixes, *mix)
	}

	return dedupeRows(mixes, l.Hashes), nil
}

func (l *Lookup) completeMixes() error {
	for i := range l.json.mixes {
		mix := &l.json.mixes[i]

		err := assignIDs(l, mix.PossibleCombinations)
		if err != nil {
			return err
		}

		l.Mixes[mix.Name] = *mix
		l.MixesID[mix.ID] = *mix
	}

	return nil
}

func (l *Lookup) loop6SeedMixCombinations(qtx *database.Queries, ctx context.Context) error {
	combos, err := l.extractMixCombinations()
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMixCombinations() ([]MixCombination, error) {
	combos := []MixCombination{}
	var err error

	for i := range l.json.mixes {
		mix := &l.json.mixes[i]
		bestComboMap := getBestComboMap(*mix)

		for j := range mix.PossibleCombinations {
			combo := &mix.PossibleCombinations[j]
			combo.MixID = mix.ID

			key := Key(combo)
			combo.IsBestCombo = bestComboMap[key]

			combo.FirstItemID, err = assignFK(combo.FirstItem, l.Items)
			if err != nil {
				return nil, err
			}

			combo.SecondItemID, err = assignFK(combo.SecondItem, l.Items)
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

	return dedupeRows(combos, l.Hashes), nil
}

func getBestComboMap(mix Mix) map[string]bool {
	bestComboMap := make(map[string]bool)

	for _, combo := range mix.BestCombinations {
		key := Key(combo)
		bestComboMap[key] = true
	}

	return bestComboMap
}
