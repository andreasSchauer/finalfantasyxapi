package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1SeedMasterItems(qtx *database.Queries, ctx context.Context) error {
	items := extractMasterItems()

	params := database.CreateMasterItemBulkParams{
		DataHash: make([]string, len(items)),
		Name:     make([]string, len(items)),
		Type:     make([]database.ItemType, len(items)),
	}

	for i, mi := range items {
		params.DataHash[i] = sh.GenerateDataHash(mi)
		params.Name[i] = mi.Name
		params.Type[i] = mi.Type
	}

	dbRows, err := qtx.CreateMasterItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create master items: %v", err)
	}

	for i, row := range dbRows {
		items[i].ID = row.ID
		lookup.L.MasterItems[seedtypes.Key(items[i])] = items[i]
		lookup.L.MasterItemsID[row.ID] = items[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMasterItems() []seedtypes.MasterItem {
	masterItems := []seedtypes.MasterItem{}

	for _, i := range sh.RawJSON.Items {
		i.MasterItem.Type = database.ItemTypeItem
		masterItems = append(masterItems, i.MasterItem)
	}

	for _, i := range sh.RawJSON.KeyItems {
		i.MasterItem.Type = database.ItemTypeKeyItem
		masterItems = append(masterItems, i.MasterItem)
	}

	return dedupeRows(masterItems, sh.Hashes)
}
