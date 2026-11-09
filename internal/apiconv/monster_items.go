package apiconv

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func ConvertMonsterItems(items seedtypes.MonsterItems) apitypes.MonsterItems {
	monItems := apitypes.MonsterItems{
		DropChance:          items.DropChance,
		DropCondition:       items.DropCondition,
		OtherItemsCondition: items.OtherItemsCondition,
		OtherItems:          ah.ConvertObjSlice(items.OtherItems, ConvertPossibleItem),
	}

	if items.StealCommon != nil {
		stealCommon := apilookups.E.AllItems.ToResAmtFromName(*items.StealCommon)
		monItems.StealCommon = &stealCommon
	}

	if items.StealRare != nil {
		stealRare := apilookups.E.AllItems.ToResAmtFromName(*items.StealRare)
		monItems.StealRare = &stealRare
	}

	if items.DropCommon != nil {
		dropCommon := apilookups.E.AllItems.ToResAmtFromName(*items.DropCommon)
		monItems.DropCommon = &dropCommon
	}

	if items.DropRare != nil {
		dropRare := apilookups.E.AllItems.ToResAmtFromName(*items.DropRare)
		monItems.DropRare = &dropRare
	}

	if items.SecondaryDropCommon != nil {
		secDropCommon := apilookups.E.AllItems.ToResAmtFromName(*items.SecondaryDropCommon)
		monItems.SecondaryDropCommon = &secDropCommon
	}

	if items.SecondaryDropRare != nil {
		secDropRare := apilookups.E.AllItems.ToResAmtFromName(*items.SecondaryDropRare)
		monItems.SecondaryDropRare = &secDropRare
	}

	if items.Bribe != nil {
		bribe := apilookups.E.AllItems.ToResAmtFromName(*items.Bribe)
		monItems.Bribe = &bribe
	}

	return monItems
}
