package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertMixCombination(mc seedtypes.MixCombination) apitypes.MixCombination {
	return apitypes.MixCombination{
		FirstItem:  nameToNamedAPIResource(E.Items, mc.FirstItem, nil),
		SecondItem: nameToNamedAPIResource(E.Items, mc.SecondItem, nil),
	}
}
