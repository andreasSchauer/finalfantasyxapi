package apiconv

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func ConvertMonsterAbility(ability seedtypes.MonsterAbility) apitypes.MonsterAbility {
	return apitypes.MonsterAbility{
		Ability:  handlerinput.KeyToTypedAPIResource(apilookups.E.Abilities, ability.AbilityReference),
		IsForced: ability.IsForced,
		IsUnused: ability.IsUnused,
	}
}

func ConvertAlt(asc seedtypes.Alt) apitypes.Alt {
	return apitypes.Alt{
		Properties:       h.SliceOrNil(handlerinput.NamesToNamedAPIResources(apilookups.E.Properties, asc.Properties)),
		AutoAbilities:    h.SliceOrNil(handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, asc.AutoAbilities)),
		BaseStats:        h.SliceOrNil(apilookups.E.Stats.ToResAmtType(asc.BaseStats, th.NewBaseStat)),
		ElemResists:      h.SliceOrNil(th.NamesToElemResists(asc.ElemResists)),
		StatusImmunities: h.SliceOrNil(handlerinput.NamesToNamedAPIResources(apilookups.E.StatusConditions, asc.StatusImmunities)),
		AddedStatus:      ah.ConvertObjPtr(asc.AddedStatus, ConvertInflictedStatus),
	}
}
