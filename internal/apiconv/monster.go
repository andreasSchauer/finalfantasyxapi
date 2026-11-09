package apiconv

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/typehelpers"
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
