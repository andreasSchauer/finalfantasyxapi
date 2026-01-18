package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ItemAmount struct {
	Item   NamedAPIResource `json:"item"`
	Amount int32            `json:"amount"`
}

func (ia ItemAmount) IsZero() bool {
	return ia.Item.Name == ""
}

func (ia ItemAmount) GetAPIResource() APIResource {
	return ia.Item
}

func (ia ItemAmount) GetName() string {
	return ia.Item.Name
}

func (ia ItemAmount) GetVersion() *int32 {
	return nil
}

func (ia ItemAmount) GetVal() int32 {
	return ia.Amount
}

func newItemAmount(res NamedAPIResource, amount int32) ItemAmount {
	return ItemAmount{
		Item: res,
		Amount: amount,
	}
}

func (cfg *Config) createItemAmount(input seeding.ItemAmount) ItemAmount {
	iItems := cfg.e.items
	iKeyItems := cfg.e.keyItems
	var ia ItemAmount

	_, ok := iItems.objLookup[input.ItemName]
	if ok {
		ia = nameToResourceAmount(cfg, iItems, input, newItemAmount)
		return ia
	}

	_, ok = iKeyItems.objLookup[input.ItemName]
	if ok {
		ia = nameToResourceAmount(cfg, iKeyItems, input, newItemAmount)
		return ia
	}

	return ItemAmount{}
}


type PossibleItem struct {
	ItemAmount
	Chance int32 `json:"chance"`
}

func (ps PossibleItem) GetAPIResource() APIResource {
	return ps.Item.GetAPIResource()
}

func (cfg *Config) newPossibleItem(item seeding.ItemAmount, chance int32) PossibleItem {
	return PossibleItem{
		ItemAmount: cfg.createItemAmount(item),
		Chance:     chance,
	}
}
