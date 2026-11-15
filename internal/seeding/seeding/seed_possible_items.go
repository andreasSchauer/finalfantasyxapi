package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3SeedPossibleItems(qtx *database.Queries, ctx context.Context) error {
	items, err := extractPossibleItems()
	if err != nil {
		return err
	}

	params := database.CreatePossibleItemBulkParams{
		DataHash:     make([]string, len(items)),
		ItemAmountID: make([]int32, len(items)),
		Chance:       make([]int32, len(items)),
	}

	for i, pi := range items {
		params.DataHash[i] = seedhelpers.GenerateDataHash(pi)
		params.ItemAmountID[i] = pi.ItemAmount.ID
		params.Chance[i] = pi.Chance
	}

	dbRows, err := qtx.CreatePossibleItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create possible items: %v", err)
	}

	for _, row := range dbRows {
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractPossibleItems() ([]seedtypes.PossibleItem, error) {
	items := []seedtypes.PossibleItem{}
	var err error

	for i := range seedhelpers.RawJSON.BlitzballPositions {
		pos := &seedhelpers.RawJSON.BlitzballPositions[i]

		for j := range pos.Items {
			item := &pos.Items[j]

			item.ItemAmount.ID, err = seedhelpers.GetHashID(item.ItemAmount)
			if err != nil {
				return nil, err
			}

			items = append(items, item.PossibleItem)
		}
	}

	for i := range seedhelpers.RawJSON.Monsters {
		mon := &seedhelpers.RawJSON.Monsters[i]

		if mon.Items == nil {
			continue
		}

		for j := range mon.Items.OtherItems {
			item := &mon.Items.OtherItems[j]

			item.ItemAmount.ID, err = seedhelpers.GetHashID(item.ItemAmount)
			if err != nil {
				return nil, err
			}

			items = append(items, *item)
		}
	}

	return dedupeRows(items, seedhelpers.Hashes), nil
}
