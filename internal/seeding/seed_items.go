package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop2SeedItems(qtx *database.Queries, ctx context.Context) error {
	items, err := l.extractItems()

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
		l.json.items[i].ID = row.ID
		l.Items[items[i].Name] = items[i]
		l.ItemsID[row.ID] = items[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractItems() ([]Item, error) {
	items := []Item{}
	var err error

	for i := range l.json.items {
		item := &l.json.items[i]

		item.MasterItem.ID, err = assignFK(item.Name, l.MasterItems)
		if err != nil {
			return nil, err
		}
		item.MasterItem.Type = database.ItemTypeItem

		items = append(items, *item)
	}

	return dedupeRows(items, l.Hashes), nil
}

func (l *Lookup) completeItems() error {
	for i := range l.json.items {
		item := &l.json.items[i]

		if len(item.BattleInteractions) > 0 {
			item.ItemAbility.ID = item.ID

			err := l.completeBattleInteractions(item.BattleInteractions)
			if err != nil {
				return err
			}

			l.ItemAbilities[Key(item.ItemAbility)] = item.ItemAbility
			l.ItemAbilitiesID[item.ID] = item.ItemAbility
		}

		l.Items[item.Name] = *item
		l.ItemsID[item.ID] = *item
	}

	return nil
}

func (l *Lookup) getItemAvailableMenus(i Item) ([]Submenu, error) {
	return getResources(i.AvailableMenus, l.Submenus)
}

func (l *Lookup) seedJuncItemsAvailableMenus(qtx *database.Queries, ctx context.Context) error {
	const desc string = "items + available menus"
	jParams, err := processJunctions(l, desc, l.json.items, l.getItemAvailableMenus)
	if err != nil {
		return err
	}

	return qtx.CreateItemsAvailableMenusJunctionBulk(ctx, database.CreateItemsAvailableMenusJunctionBulkParams{
		DataHash:  jParams.DataHashes,
		ItemID:    jParams.ParentIDs,
		SubmenuID: jParams.ChildIDs,
	})
}

func (l *Lookup) getItemRelatedStats(i Item) ([]Stat, error) {
	return getResources(i.RelatedStats, l.Stats)
}

func (l *Lookup) seedJuncItemsRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "items + related stats"
	jParams, err := processJunctions(l, desc, l.json.items, l.getItemRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateItemsRelatedStatsJunctionBulk(ctx, database.CreateItemsRelatedStatsJunctionBulkParams{
		DataHash: jParams.DataHashes,
		ItemID:   jParams.ParentIDs,
		StatID:   jParams.ChildIDs,
	})
}
