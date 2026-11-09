package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertMonsterAbility(ability seedtypes.MonsterAbility) apitypes.MonsterAbility {
	return apitypes.MonsterAbility{
		Ability:  handlerinput.KeyToTypedAPIResource(apilookups.E.Abilities, ability.AbilityReference),
		IsForced: ability.IsForced,
		IsUnused: ability.IsUnused,
	}
}
