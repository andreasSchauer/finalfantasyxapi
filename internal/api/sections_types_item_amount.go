package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ItemAmountSimple struct {
	ia            seeding.ItemAmount `json:"-"`
	ItemAmountStr string             `json:"item"`
	ItemType      database.ItemType  `json:"item_type"`
}

func convertItemAmountSimple(cfg *Config, ia seeding.ItemAmount) ItemAmountSimple {
	itemLookup, _ := seeding.GetResource(ia.ItemName, cfg.l.MasterItems)
	itemStr := nameAmountString(ia.ItemName, nil, nil, ia.Amount)

	return ItemAmountSimple{
		ia:            ia,
		ItemAmountStr: itemStr,
		ItemType:      itemLookup.Type,
	}
}

func posItemToItemAmtSimple(cfg *Config, posItem seeding.PossibleItem) ItemAmountSimple {
	return convertItemAmountSimple(cfg, posItem.ItemAmount)
}

func sortSimpleItemAmountsByID(cfg *Config, s []ItemAmountSimple) []ItemAmountSimple {
	slices.SortStableFunc(s, func(a, b ItemAmountSimple) int {
		A := getMasterItemID(cfg, a)
		B := getMasterItemID(cfg, b)

		if A < B {
			return -1
		}

		if A > B {
			return 1
		}

		return 0
	})

	return s
}

func getMasterItemID(cfg *Config, ia ItemAmountSimple) int32 {
	if ia.ItemType == database.ItemTypeItem {
		itemLookup, _ := seeding.GetResource(ia.ia.ItemName, cfg.l.Items)
		return itemLookup.MasterItem.ID
	}

	itemLookup, _ := seeding.GetResource(ia.ia.ItemName, cfg.l.KeyItems)
	return itemLookup.MasterItem.ID
}
