package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type BlitzballPosition struct {
	ID       int32
	Category string          `json:"category"`
	Slot     string          `json:"slot"`
	Items    []BlitzballItem `json:"items"`
}

func (b BlitzballPosition) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", b),
		b.Category,
		b.Slot,
	}
}

func (b BlitzballPosition) GetID() int32 {
	return b.ID
}

func (b BlitzballPosition) ToKeyFields() []any {
	return []any{
		b.Category,
		b.Slot,
	}
}

func (b BlitzballPosition) Error() string {
	return fmt.Sprintf("blitzball position %s in %s", b.Slot, b.Category)
}

func (b BlitzballPosition) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   b.ID,
		Name: fmt.Sprintf("%s - %s", b.Category, b.Slot),
	}
}

func (b BlitzballPosition) GetItemAmounts() []ItemAmount {
	ias := []ItemAmount{}

	for _, item := range b.Items {
		ias = append(ias, item.ItemAmount)
	}

	return ias
}

type PossibleItem struct {
	ID         int32
	ItemAmount ItemAmount `json:"item"`
	Chance     int32      `json:"chance"`
}

func (i PossibleItem) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", i),
		i.ItemAmount.ID,
		i.Chance,
	}
}

func (i PossibleItem) GetID() int32 {
	return i.ID
}

func (i *PossibleItem) SetID(id int32) {
	i.ID = id
}

func (i PossibleItem) Error() string {
	return fmt.Sprintf("possible item %s, amount: %d, chance: %d", i.ItemAmount.ItemName, i.ItemAmount.Amount, i.Chance)
}

type BlitzballItem struct {
	ID				int32
	PositionID 		int32
	PossibleItem
}

func (b BlitzballItem) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", b),
		b.PositionID,
		b.PossibleItem.ID,
	}
}

func (b BlitzballItem) GetID() int32 {
	return b.ID
}

func (b *BlitzballItem) SetID(id int32) {
	b.ID = id
}

func (b BlitzballItem) Error() string {
	return fmt.Sprintf("blitzball item %s, chance %d, position id %d", b.ItemAmount.ItemName, b.Chance, b.PositionID)
}


func (l *Lookup) loop1SeedBlitzballPositions(qtx *database.Queries, ctx context.Context) error {
	positions := dedupeRows(l.json.blitzballPositions, l.Hashes)

	params := database.CreateBlitzballPositionBulkParams{
		DataHash: make([]string, len(positions)),
		Category: make([]database.BlitzballTournamentCategory, len(positions)),
		Slot:     make([]database.BlitzballPositionSlot, len(positions)),
	}

	for i, p := range positions {
		params.DataHash[i] = generateDataHash(p)
		params.Category[i] = database.BlitzballTournamentCategory(p.Category)
		params.Slot[i] = database.BlitzballPositionSlot(p.Slot)
	}

	dbRows, err := qtx.CreateBlitzballPositionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create blitzball position: %v", err)
	}

	for i, row := range dbRows {
		positions[i].ID = row.ID
		l.json.blitzballPositions[i].ID = row.ID
		key := Key(positions[i])
		l.Positions[key] = positions[i]
		l.PositionsID[row.ID] = positions[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) completeBlitzballPositions() error {
	for i := range l.json.blitzballPositions {
		pos := &l.json.blitzballPositions[i]
		err := assignIDs(l, pos.Items)
		if err != nil {
			return err
		}

		l.Positions[Key(*pos)] = *pos
		l.PositionsID[pos.ID] = *pos
	}

	return nil
}

func (l *Lookup) loop4SeedBlitzballItems(qtx *database.Queries, ctx context.Context) error {
	items, err := l.extractBlitzballItems()
	if err != nil {
		return err
	}

	params := database.CreateBlitzballItemBulkParams{
		DataHash:       make([]string, len(items)),
		PositionID:     make([]int32, len(items)),
		PossibleItemID: make([]int32, len(items)),
	}

	for i, bi := range items {
		params.DataHash[i] = generateDataHash(bi)
		params.PositionID[i] = bi.PositionID
		params.PossibleItemID[i] = bi.PossibleItem.ID
	}

	dbRows, err := qtx.CreateBlitzballItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create blitzball items: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractBlitzballItems() ([]BlitzballItem, error) {
	items := []BlitzballItem{}
	var err error

	for i := range l.json.blitzballPositions {
		pos := &l.json.blitzballPositions[i]

		for j := range pos.Items {
			item := &pos.Items[j]
			item.PositionID = pos.ID

			item.PossibleItem.ID, err = l.getHashID(item.PossibleItem)
			if err != nil {
				return nil, err
			}

			items = append(items, *item)
		}
	}

	return dedupeRows(items, l.Hashes), nil
}

func (l *Lookup) loop3SeedPossibleItems(qtx *database.Queries, ctx context.Context) error {
	items, err := l.extractPossibleItems()
	if err != nil {
		return err
	}

	params := database.CreatePossibleItemBulkParams{
		DataHash:     make([]string, len(items)),
		ItemAmountID: make([]int32, len(items)),
		Chance:       make([]int32, len(items)),
	}

	for i, pi := range items {
		params.DataHash[i] = generateDataHash(pi)
		params.ItemAmountID[i] = pi.ItemAmount.ID
		params.Chance[i] = pi.Chance
	}

	dbRows, err := qtx.CreatePossibleItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create possible items: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractPossibleItems() ([]PossibleItem, error) {
	items := []PossibleItem{}
	var err error

	for i := range l.json.blitzballPositions {
		pos := &l.json.blitzballPositions[i]

		for j := range pos.Items {
			item := &pos.Items[j]

			item.ItemAmount.ID, err = l.getHashID(item.ItemAmount)
			if err != nil {
				return nil, err
			}

			items = append(items, item.PossibleItem)
		}
	}

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		if mon.Items == nil {
			continue
		}

		for j := range mon.Items.OtherItems {
			item := &mon.Items.OtherItems[j]

			item.ItemAmount.ID, err = l.getHashID(item.ItemAmount)
			if err != nil {
				return nil, err
			}

			items = append(items, *item)
		}
	}

	return dedupeRows(items, l.Hashes), nil
}
