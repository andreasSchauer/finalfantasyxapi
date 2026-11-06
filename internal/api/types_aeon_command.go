package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertPossibleAbilityList(pa seedtypes.PossibleAbilityList) apitypes.PossibleAbilityList {
	return apitypes.PossibleAbilityList{
		User:      nameToNamedAPIResource(E.CharacterClasses, pa.User, nil),
		Abilities: keysToTypedAPIResources(E.Abilities, pa.Abilities),
	}
}
