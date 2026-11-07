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
		stealCommon := E.AllItems.ToResAmtFromName(*items.StealCommon)
		monItems.StealCommon = &stealCommon
	}

	if items.StealRare != nil {
		stealRare := E.AllItems.ToResAmtFromName(*items.StealRare)
		monItems.StealRare = &stealRare
	}

	if items.DropCommon != nil {
		dropCommon := E.AllItems.ToResAmtFromName(*items.DropCommon)
		monItems.DropCommon = &dropCommon
	}

	if items.DropRare != nil {
		dropRare := E.AllItems.ToResAmtFromName(*items.DropRare)
		monItems.DropRare = &dropRare
	}

	if items.SecondaryDropCommon != nil {
		secDropCommon := E.AllItems.ToResAmtFromName(*items.SecondaryDropCommon)
		monItems.SecondaryDropCommon = &secDropCommon
	}

	if items.SecondaryDropRare != nil {
		secDropRare := E.AllItems.ToResAmtFromName(*items.SecondaryDropRare)
		monItems.SecondaryDropRare = &secDropRare
	}

	if items.Bribe != nil {
		bribe := E.AllItems.ToResAmtFromName(*items.Bribe)
		monItems.Bribe = &bribe
	}

	return monItems
}
