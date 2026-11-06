package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func createItemMonster(itemName string, mon apiresource.NamedAPIResource) apitypes.MonItemAmts {
	monster, _ := lookup.GetResourceByID(mon.ID, lookup.L.MonstersID)

	return apitypes.MonItemAmts{
		Monster:             mon,
		StealCommon:         getAmountMonItem(itemName, monster.Items.StealCommon),
		StealRare:           getAmountMonItem(itemName, monster.Items.StealRare),
		DropCommon:          getAmountMonItem(itemName, monster.Items.DropCommon),
		DropRare:            getAmountMonItem(itemName, monster.Items.DropRare),
		SecondaryDropCommon: getAmountMonItem(itemName, monster.Items.SecondaryDropCommon),
		SecondaryDropRare:   getAmountMonItem(itemName, monster.Items.SecondaryDropRare),
		Bribe:               getAmountMonItem(itemName, monster.Items.Bribe),
		Other:               getAmountOther(itemName, monster.Items.OtherItems),
	}
}

func getAmountMonItem(wantedItem string, ia *seedtypes.ItemAmount) int32 {
	if ia == nil || ia.ItemName != wantedItem {
		return 0
	}

	return ia.Amount
}

func getAmountOther(wantedItem string, otherItems []seedtypes.PossibleItem) int32 {
	for _, otherItem := range otherItems {
		if otherItem.ItemAmount.ItemName == wantedItem {
			return otherItem.ItemAmount.Amount
		}
	}

	return 0
}
