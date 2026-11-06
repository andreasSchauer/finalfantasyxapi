package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertMonsterItems(items seedtypes.MonsterItems) apitypes.MonsterItems {
	monItems := apitypes.MonsterItems{
		DropChance:          items.DropChance,
		DropCondition:       items.DropCondition,
		OtherItemsCondition: items.OtherItemsCondition,
		OtherItems:          ah.ConvertObjSlice(items.OtherItems, convertPossibleItem),
	}

	if items.StealCommon != nil {
		stealCommon := E.AllItems.toResAmtFromName(*items.StealCommon)
		monItems.StealCommon = &stealCommon
	}

	if items.StealRare != nil {
		stealRare := E.AllItems.toResAmtFromName(*items.StealRare)
		monItems.StealRare = &stealRare
	}

	if items.DropCommon != nil {
		dropCommon := E.AllItems.toResAmtFromName(*items.DropCommon)
		monItems.DropCommon = &dropCommon
	}

	if items.DropRare != nil {
		dropRare := E.AllItems.toResAmtFromName(*items.DropRare)
		monItems.DropRare = &dropRare
	}

	if items.SecondaryDropCommon != nil {
		secDropCommon := E.AllItems.toResAmtFromName(*items.SecondaryDropCommon)
		monItems.SecondaryDropCommon = &secDropCommon
	}

	if items.SecondaryDropRare != nil {
		secDropRare := E.AllItems.toResAmtFromName(*items.SecondaryDropRare)
		monItems.SecondaryDropRare = &secDropRare
	}

	if items.Bribe != nil {
		bribe := E.AllItems.toResAmtFromName(*items.Bribe)
		monItems.Bribe = &bribe
	}

	return monItems
}
