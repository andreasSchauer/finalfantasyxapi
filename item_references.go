package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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

func (ia ItemAmount) GetVal() int32 {
	return ia.Amount
}

func newItemAmount(res NamedAPIResource, amount int32) ItemAmount {
	return ItemAmount{
		Item: res,
		Amount: amount,
	}
}

func (cfg *Config) createItemAmount(name string, amount int32) (ItemAmount, error) {
	iItems := cfg.e.items
	iKeyItems := cfg.e.keyItems
	var ia ItemAmount

	_, ok := iItems.objLookup[name]
	if ok {
		ia = nameToResourceAmount(cfg, iItems, name, nil, amount, newItemAmount)
		return ia, nil
	}

	_, ok = iKeyItems.objLookup[name]
	if ok {
		ia = nameToResourceAmount(cfg, iKeyItems, name, nil, amount, newItemAmount)
		return ia, nil
	}

	return ItemAmount{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("item '%s' does not exist.", name), nil)
}


func (cfg *Config) newItemAmountOld(itemType database.ItemType, itemName string, itemID, amount int32) ItemAmount {
	if itemName == "" {
		return ItemAmount{}
	}
	var endpoint string

	switch itemType {
	case database.ItemTypeItem:
		endpoint = cfg.e.items.endpoint
	case database.ItemTypeKeyItem:
		endpoint = cfg.e.keyItems.endpoint
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

func (ps PossibleItem) GetAPIResource() APIResource {
	return ps.Item.GetAPIResource()
}

func (cfg *Config) newPossibleItem(itemType database.ItemType, itemName string, itemID, amount, chance int32) PossibleItem {
	itemAmount := cfg.newItemAmountOld(itemType, itemName, itemID, amount)

	return PossibleItem{
		ItemAmount: itemAmount,
		Chance:     chance,
	}
}
