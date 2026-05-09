package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

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

			item.ItemAmount.ID, err = l.GetHashID(item.ItemAmount)
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

			item.ItemAmount.ID, err = l.GetHashID(item.ItemAmount)
			if err != nil {
				return nil, err
			}

			items = append(items, *item)
		}
	}

	return dedupeRows(items, l.Hashes), nil
}
