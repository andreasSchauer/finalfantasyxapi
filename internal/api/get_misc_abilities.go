package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMiscAbility(r *http.Request, i handlerInput[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.NamedAPIResource], id int32) (apitypes.MiscAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.MiscAbility{}, err
	}

	response := apitypes.MiscAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(E.Abilities, ability.Ability.ID),
		Description:        ability.Description,
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		LearnedBy:          namesToNamedAPIResources(E.CharacterClasses, ability.LearnedBy),
		Topmenu:            namePtrToNamedAPIResPtr(E.Topmenus, ability.Topmenu, nil),
		Submenu:            namePtrToNamedAPIResPtr(E.Submenus, ability.Submenu, nil),
		OpenSubmenu:        namePtrToNamedAPIResPtr(E.Submenus, ability.OpenSubmenu, nil),
		Cursor:             ability.Cursor,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return apitypes.MiscAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func retrieveMiscAbilities(r *http.Request, i handlerInput[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeMiscAbility

	return i.filterIDs(r, ids, []IDFilter{
		i.queryIntList(r, ids, ts.QpnRank, getTypedAbilityIDsByRank(abilityType)),
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetMiscAbilityIDsByCharClass),
		i.queryBool(r, ids, ts.QpnCopycat, getTypedAbilityIDsByCanCopycat(abilityType)),
		i.queryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(abilityType)),
		i.queryBool2(r, ids, ts.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(abilityType)),
	})
}
