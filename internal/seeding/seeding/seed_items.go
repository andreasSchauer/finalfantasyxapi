package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop2SeedItems(qtx *database.Queries, ctx context.Context) error {
	items, err := extractItems()
	if err != nil {
		return err
	}

	params := database.CreateItemBulkParams{
		DataHash:     make([]string, len(items)),
		MasterItemID: make([]int32, len(items)),
		Description:  make([]string, len(items)),
		Effect:       make([]string, len(items)),
		Category:     make([]database.ItemCategory, len(items)),
		Usability:    make([]database.ItemUsability, len(items)),
		BasePrice:    make([]sql.NullInt32, len(items)),
		SellValue:    make([]int32, len(items)),
	}

	for i, item := range items {
		params.DataHash[i] = generateDataHash(item)
		params.MasterItemID[i] = item.MasterItem.ID
		params.Description[i] = item.Description
		params.Effect[i] = item.Effect
		params.Category[i] = database.ItemCategory(item.Category)
		params.Usability[i] = database.ItemUsability(item.Usability)
		params.BasePrice[i] = h.GetNullInt32(item.BasePrice)
		params.SellValue[i] = item.SellValue
	}

	dbRows, err := qtx.CreateItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create items: %v", err)
	}

	for i, row := range dbRows {
		items[i].ID = row.ID
		rawJSON.items[i].ID = row.ID
		lookup.L.Items[seedtypes.Key(items[i])] = items[i]
		lookup.L.ItemsID[row.ID] = items[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractItems() ([]seedtypes.Item, error) {
	items := []seedtypes.Item{}
	var err error

	for i := range rawJSON.items {
		item := &rawJSON.items[i]

		item.MasterItem.ID, err = assignFK(item.Name, lookup.L.MasterItems)
		if err != nil {
			return nil, err
		}
		item.MasterItem.Type = database.ItemTypeItem

		items = append(items, *item)
	}

	return dedupeRows(items, hashes), nil
}

func completeItems() error {
	for i := range rawJSON.items {
		item := &rawJSON.items[i]

		if len(item.BattleInteractions) > 0 {
			item.ItemAbility.ID = item.ID

			err := completeBattleInteractions(item.BattleInteractions)
			if err != nil {
				return err
			}

			lookup.L.ItemAbilities[seedtypes.Key(item.ItemAbility)] = item.ItemAbility
			lookup.L.ItemAbilitiesID[item.ID] = item.ItemAbility
		}

		lookup.L.Items[seedtypes.Key(item)] = *item
		lookup.L.ItemsID[item.ID] = *item
	}

	return nil
}

func getItemAvailableMenus(i seedtypes.Item) ([]seedtypes.Submenu, error) {
	return lookup.GetResources(i.AvailableMenus, lookup.L.Submenus)
}

func seedJuncItemsAvailableMenus(qtx *database.Queries, ctx context.Context) error {
	const desc string = "items + available menus"
	jParams, err := processJunctions(desc, rawJSON.items, getItemAvailableMenus)
	if err != nil {
		return err
	}

	return qtx.CreateItemsAvailableMenusJunctionBulk(ctx, database.CreateItemsAvailableMenusJunctionBulkParams{
		DataHash:  jParams.DataHashes,
		ItemID:    jParams.ParentIDs,
		SubmenuID: jParams.ChildIDs,
	})
}

func getItemRelatedStats(i seedtypes.Item) ([]seedtypes.Stat, error) {
	return lookup.GetResources(i.RelatedStats, lookup.L.Stats)
}

func seedJuncItemsRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "items + related stats"
	jParams, err := processJunctions(desc, rawJSON.items, getItemRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateItemsRelatedStatsJunctionBulk(ctx, database.CreateItemsRelatedStatsJunctionBulkParams{
		DataHash: jParams.DataHashes,
		ItemID:   jParams.ParentIDs,
		StatID:   jParams.ChildIDs,
	})
}
