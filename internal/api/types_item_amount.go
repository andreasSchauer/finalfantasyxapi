package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ItemAmount struct {
	Item   TypedAPIResource `json:"item"`
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

func newItemAmount(res TypedAPIResource, amount int32) ItemAmount {
	return ItemAmount{
		Item:   res,
		Amount: amount,
	}
}

func convertItemAmount(cfg *Config, input seeding.ItemAmount) ItemAmount {
	return keyToTypedResourceAmount(cfg, cfg.e.masterItems, input, newItemAmount)
}


type PossibleItem struct {
	ItemAmount
	Chance int32 `json:"chance"`
}

func (ps PossibleItem) GetAPIResource() APIResource {
	return ps.Item.GetAPIResource()
}

func newPossibleItem(cfg *Config, item seeding.ItemAmount, chance int32) PossibleItem {
	return PossibleItem{
		ItemAmount: convertItemAmount(cfg, item),
		Chance:     chance,
	}
}

func convertPossibleItem(cfg *Config, item seeding.PossibleItem) PossibleItem {
	return newPossibleItem(cfg, item.ItemAmount, item.Chance)
}
