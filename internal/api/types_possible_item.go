package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func newPossibleItem(item seedtypes.ItemAmount, chance int32) apitypes.PossibleItem {
	return apitypes.PossibleItem{
		ResourceAmount: E.AllItems.ToResAmtFromName(item),
		Chance:         chance,
	}
}

func convertPossibleItem(item seedtypes.PossibleItem) apitypes.PossibleItem {
	return newPossibleItem(item.ItemAmount, item.Chance)
}
