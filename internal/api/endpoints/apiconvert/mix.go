package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func MixCombination(mc seedtypes.MixCombination) apitypes.MixCombination {
	return apitypes.MixCombination{
		FirstItem:  handlerinput.NameToNamedAPIResource(apilookups.E.Items, mc.FirstItem, nil),
		SecondItem: handlerinput.NameToNamedAPIResource(apilookups.E.Items, mc.SecondItem, nil),
	}
}
