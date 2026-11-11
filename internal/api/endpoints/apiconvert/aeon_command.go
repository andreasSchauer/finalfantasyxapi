package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func PossibleAbilityList(pa seedtypes.PossibleAbilityList) apitypes.PossibleAbilityList {
	return apitypes.PossibleAbilityList{
		User:      handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, pa.User, nil),
		Abilities: handlerinput.KeysToTypedAPIResources(apilookups.E.Abilities, pa.Abilities),
	}
}
