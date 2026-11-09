package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertPossibleAbilityList(pa seedtypes.PossibleAbilityList) apitypes.PossibleAbilityList {
	return apitypes.PossibleAbilityList{
		User:      handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, pa.User, nil),
		Abilities: handlerinput.KeysToTypedAPIResources(apilookups.E.Abilities, pa.Abilities),
	}
}
