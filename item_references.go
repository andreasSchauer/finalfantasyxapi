package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type ItemAmount struct {
	Item   NamedAPIResource `json:"item"`
	Amount int32            `json:"amount"`
}

func (ia ItemAmount) IsZero() bool {
	return ia.Item.Name == ""
}

func (ia ItemAmount) getAPIResource() IsAPIResource {
	return ia.Item
}

func (cfg *Config) newItemAmount(itemType database.ItemType, itemName string, itemID, amount int32) ItemAmount {
	if itemName == "" {
		return ItemAmount{}
	}
	endpoint := "items"

	if itemType == database.ItemTypeKeyItem {
		endpoint = "key-items"
	}

	itemResource := cfg.newNamedAPIResourceSimple(endpoint, itemID, itemName)

	return ItemAmount{
		Item:   itemResource,
		Amount: amount,
	}
}

type PossibleItem struct {
	ItemAmount
	Chance int32 `json:"chance"`
}

func (ps PossibleItem) getAPIResource() IsAPIResource {
	return ps.Item.getAPIResource()
}

func (cfg *Config) newPossibleItem(itemType database.ItemType, itemName string, itemID, amount, chance int32) PossibleItem {
	itemAmount := cfg.newItemAmount(itemType, itemName, itemID, amount)

	return PossibleItem{
		ItemAmount: itemAmount,
		Chance:     chance,
	}
}
