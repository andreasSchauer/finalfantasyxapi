package typehelpers

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func NewPossibleItem(item seedtypes.ItemAmount, chance int32) apitypes.PossibleItem {
	return apitypes.PossibleItem{
		ResourceAmount: apilookups.E.AllItems.ToResAmtFromName(item),
		Chance:         chance,
	}
}
