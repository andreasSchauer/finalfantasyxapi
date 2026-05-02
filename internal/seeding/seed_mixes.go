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
		fmt.Sprintf("%T", m),
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

func (m Mix) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   m.ID,
		Name: m.Name,
	}
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
		fmt.Sprintf("%T", m),
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

func (m *MixCombination) SetID(id int32) {
	m.ID = id
}

func (m MixCombination) Error() string {
	return fmt.Sprintf("mix combination with first item: %s, second item: %s", m.FirstItem, m.SecondItem)
}

func (l *Lookup) seedMixes(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/mixes.json"

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
				return h.NewErr(mix.Error(), err)
			}

			dbMix, err := qtx.CreateMix(context.Background(), database.CreateMixParams{
				DataHash:    generateDataHash(mix),
				OverdriveID: mix.OverdriveID,
				Category:    database.MixCategory(mix.Category),
			})
			if err != nil {
				return h.NewErr(mix.Error(), err, "couldn't create mix")
			}

			mix.ID = dbMix.ID
			l.Mixes[mix.Name] = mix
			l.MixesID[mix.ID] = mix
		}

		return nil
	})
}

func (l *Lookup) seedMixesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/mixes.json"

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
				return h.NewErr(mix.Error(), err)
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

		key := Key(combo)
		combo.IsBestCombo = bestComboMap[key]

		_, err = seedObjAssignID(qtx, combo, l.seedMixCombination)
		if err != nil {
			return err
		}
	}

	return nil
}

func getBestComboMap(mix Mix) map[string]bool {
	bestComboMap := make(map[string]bool)

	for _, combo := range mix.BestCombinations {
		key := Key(combo)
		bestComboMap[key] = true
	}

	return bestComboMap
}

func (l *Lookup) seedMixCombination(qtx *database.Queries, combo MixCombination) (MixCombination, error) {
	var err error

	combo.FirstItemID, err = assignFK(combo.FirstItem, l.Items)
	if err != nil {
		return MixCombination{}, h.NewErr(combo.Error(), err)
	}

	combo.SecondItemID, err = assignFK(combo.SecondItem, l.Items)
	if err != nil {
		return MixCombination{}, h.NewErr(combo.Error(), err)
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
		return MixCombination{}, h.NewErr(combo.Error(), err, "couldn't create mix combination")
	}

	combo.ID = dbCombo.ID

	return combo, nil
}

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

		mix.OverdriveID, err = assignFK(mix.Name, l.Overdrives)
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
