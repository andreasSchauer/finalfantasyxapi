package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop3SeedShopItems(qtx *database.Queries, ctx context.Context) error {
	items, err := l.extractShopItems()
	if err != nil {
		return err
	}

	params := database.CreateShopItemBulkParams{
		DataHash: make([]string, len(items)),
		ItemID:   make([]int32, len(items)),
		Price:    make([]int32, len(items)),
	}

	for i, si := range items {
		params.DataHash[i] = generateDataHash(si)
		params.ItemID[i] = si.ItemID
		params.Price[i] = si.Price
	}

	dbRows, err := qtx.CreateShopItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create shop items: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractShopItems() ([]ShopItem, error) {
	items := []ShopItem{}
	var err error

	for i := range l.json.shops {
		shop := &l.json.shops[i]

		if shop.PreAirship != nil {
			for j := range shop.PreAirship.Items {
				item := &shop.PreAirship.Items[j]

				item.ItemID, err = assignFK(item.Name, l.Items)
				if err != nil {
					return nil, err
				}

				items = append(items, *item)
			}
		}

		if shop.PostAirship != nil {
			for j := range shop.PostAirship.Items {
				item := &shop.PostAirship.Items[j]

				item.ItemID, err = assignFK(item.Name, l.Items)
				if err != nil {
					return nil, err
				}

				items = append(items, *item)
			}
		}
	}

	return dedupeRows(items, l.Hashes), nil
}
