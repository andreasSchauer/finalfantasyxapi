package main

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ItemAmountSub struct {
	seeding.ItemAmount
	ItemType			database.ItemType	`json:"item_type"`
}

func createSubItemAmountPtr(cfg *Config, ia *seeding.ItemAmount) *ItemAmountSub {
	if ia == nil {
		return nil
	}

	itemAmountSub := createSubItemAmount(cfg, *ia)
	return &itemAmountSub
}

func createSubItemAmount(cfg *Config, ia seeding.ItemAmount) ItemAmountSub {
	itemLookup, _ := seeding.GetResource(ia.ItemName, cfg.l.MasterItems)

	return ItemAmountSub{
		ItemAmount: ia,
		ItemType: 	itemLookup.Type,
	}
}

func sortItemAmountSubsByID(cfg *Config, s []ItemAmountSub) []ItemAmountSub {
	slices.SortStableFunc(s, func (a, b ItemAmountSub) int{
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

func getMasterItemID(cfg *Config, ia ItemAmountSub) int32 {
	if ia.ItemType == database.ItemTypeItem {
		itemLookup, _ := seeding.GetResource(ia.ItemName, cfg.l.Items)
		return itemLookup.MasterItem.ID
	}

	itemLookup, _ := seeding.GetResource(ia.ItemName, cfg.l.KeyItems)
	return itemLookup.MasterItem.ID
}