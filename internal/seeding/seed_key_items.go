package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop2SeedKeyItems(qtx *database.Queries, ctx context.Context) error {
	keyItems, err := l.extractKeyItems()

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
		l.json.keyItems[i].ID = row.ID
		l.KeyItems[keyItems[i].Name] = keyItems[i]
		l.KeyItemsID[row.ID] = keyItems[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractKeyItems() ([]KeyItem, error) {
	keyItems := []KeyItem{}
	var err error

	for i := range l.json.keyItems {
		keyItem := &l.json.keyItems[i]

		keyItem.MasterItem.ID, err = assignFK(keyItem.Name, l.MasterItems)
		if err != nil {
			return nil, err
		}
		keyItem.MasterItem.Type = database.ItemTypeKeyItem

		keyItems = append(keyItems, *keyItem)
	}

	return dedupeRows(keyItems, l.Hashes), nil
}
