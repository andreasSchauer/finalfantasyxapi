package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertPossibleAbilityList(pa seedtypes.PossibleAbilityList) apitypes.PossibleAbilityList {
	return apitypes.PossibleAbilityList{
		User:      NameToNamedAPIResource(E.CharacterClasses, pa.User, nil),
		Abilities: KeysToTypedAPIResources(E.Abilities, pa.Abilities),
	}
}
