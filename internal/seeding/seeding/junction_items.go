package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

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
