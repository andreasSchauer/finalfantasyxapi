package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertPossibleAbilityList(pa seedtypes.PossibleAbilityList) apitypes.PossibleAbilityList {
	return apitypes.PossibleAbilityList{
		User:      handlerinput.NameToNamedAPIResource(E.CharacterClasses, pa.User, nil),
		Abilities: handlerinput.KeysToTypedAPIResources(E.Abilities, pa.Abilities),
	}
}
