package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop2SeedKeyItems(qtx *database.Queries, ctx context.Context) error {
	keyItems, err := extractKeyItems()
	if err != nil {
		return err
	}

	params := database.CreateKeyItemBulkParams{
		DataHash:     make([]string, len(keyItems)),
		MasterItemID: make([]int32, len(keyItems)),
		Category:     make([]database.KeyItemCategory, len(keyItems)),
		Description:  make([]string, len(keyItems)),
		Effect:       make([]string, len(keyItems)),
	}

	for i, ki := range keyItems {
		params.DataHash[i] = generateDataHash(ki)
		params.MasterItemID[i] = ki.MasterItem.ID
		params.Category[i] = database.KeyItemCategory(ki.Category)
		params.Description[i] = ki.Description
		params.Effect[i] = ki.Effect
	}

	dbRows, err := qtx.CreateKeyItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create key items: %v", err)
	}

	for i, row := range dbRows {
		keyItems[i].ID = row.ID
		rawJSON.keyItems[i].ID = row.ID
		lookup.L.KeyItems[seedtypes.Key(keyItems[i])] = keyItems[i]
		lookup.L.KeyItemsID[row.ID] = keyItems[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractKeyItems() ([]seedtypes.KeyItem, error) {
	keyItems := []seedtypes.KeyItem{}
	var err error

	for i := range rawJSON.keyItems {
		keyItem := &rawJSON.keyItems[i]

		keyItem.MasterItem.ID, err = assignFK(keyItem.Name, lookup.L.MasterItems)
		if err != nil {
			return nil, err
		}
		keyItem.MasterItem.Type = database.ItemTypeKeyItem

		keyItems = append(keyItems, *keyItem)
	}

	return dedupeRows(keyItems, hashes), nil
}
