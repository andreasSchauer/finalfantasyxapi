package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ItemAmountSub struct {
	seeding.ItemAmount
	ItemType			database.ItemType	`json:"item_type"`
}

func (cfg *Config) createSubItemAmountPtr(ia *seeding.ItemAmount) *ItemAmountSub {
	if ia == nil {
		return nil
	}

	itemAmountSub := cfg.createSubItemAmount(*ia)
	return &itemAmountSub
}

func (cfg *Config) createSubItemAmount(ia seeding.ItemAmount) ItemAmountSub {
	itemLookup, _ := seeding.GetResource(ia.ItemName, cfg.l.MasterItems)

	return ItemAmountSub{
		ItemAmount: ia,
		ItemType: 	itemLookup.Type,
	}
}