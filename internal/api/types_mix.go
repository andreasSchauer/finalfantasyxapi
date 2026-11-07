package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertMixCombination(mc seedtypes.MixCombination) apitypes.MixCombination {
	return apitypes.MixCombination{
		FirstItem:  handlerinput.NameToNamedAPIResource(E.Items, mc.FirstItem, nil),
		SecondItem: handlerinput.NameToNamedAPIResource(E.Items, mc.SecondItem, nil),
	}
}
