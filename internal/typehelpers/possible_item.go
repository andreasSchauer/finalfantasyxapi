package typehelpers

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func NewPossibleItem(item seedtypes.ItemAmount, chance int32) apitypes.PossibleItem {
	return apitypes.PossibleItem{
		ResourceAmount: apilookups.E.AllItems.ToResAmtFromName(item),
		Chance:         chance,
	}
}
