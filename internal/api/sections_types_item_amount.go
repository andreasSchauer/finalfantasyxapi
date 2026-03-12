package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ItemAmountSimple struct {
	ia            seeding.ItemAmount `json:"-"`
	ItemAmountStr string             `json:"item"`
	ItemType      database.ItemType  `json:"item_type"`
}

func convertItemAmountSimple(cfg *Config, ia seeding.ItemAmount) string {
	return h.NameAmountString(ia.ItemName, nil, nil, ia.Amount)
}

func posItemToItemAmtSimple(cfg *Config, posItem seeding.PossibleItem) string {
	return convertItemAmountSimple(cfg, posItem.ItemAmount)
}

func sortItemAmountsByID(cfg *Config, s []seeding.ItemAmount) []seeding.ItemAmount {
	slices.SortStableFunc(s, func(a, b seeding.ItemAmount) int {
		A := getMasterItemID(cfg, a)
		B := getMasterItemID(cfg, b)

		if A < B {
			return -1
		}

		if A > B {
			return 1
		}

		if A == B {
			if a.Amount < b.Amount {
				return -1
			}

			if a.Amount > b.Amount {
				return 1
			}
		}

		return 0
	})

	return s
}

func getMasterItemID(cfg *Config, ia seeding.ItemAmount) int32 {
	lookup, _ := seeding.GetResource(ia.ItemName, cfg.l.MasterItems)

	if lookup.Type == database.ItemTypeItem {
		itemLookup, _ := seeding.GetResource(ia.ItemName, cfg.l.Items)
		return itemLookup.MasterItem.ID
	}

	itemLookup, _ := seeding.GetResource(ia.ItemName, cfg.l.KeyItems)
	return itemLookup.MasterItem.ID
}
