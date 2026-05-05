package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop1SeedMasterItems(qtx *database.Queries, ctx context.Context) error {
	items := l.extractMasterItems()

	params := database.CreateMasterItemBulkParams{
		DataHash: make([]string, len(items)),
		Name:     make([]string, len(items)),
		Type:     make([]database.ItemType, len(items)),
	}

	for i, mi := range items {
		params.DataHash[i] = generateDataHash(mi)
		params.Name[i] = mi.Name
		params.Type[i] = mi.Type
	}

	dbRows, err := qtx.CreateMasterItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create master items: %v", err)
	}

	for i, row := range dbRows {
		items[i].ID = row.ID
		l.MasterItems[Key(items[i])] = items[i]
		l.MasterItemsID[row.ID] = items[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMasterItems() []MasterItem {
	masterItems := []MasterItem{}

	for _, i := range l.json.items {
		i.MasterItem.Type = database.ItemTypeItem
		masterItems = append(masterItems, i.MasterItem)
	}

	for _, i := range l.json.keyItems {
		i.MasterItem.Type = database.ItemTypeKeyItem
		masterItems = append(masterItems, i.MasterItem)
	}

	return dedupeRows(masterItems, l.Hashes)
}
