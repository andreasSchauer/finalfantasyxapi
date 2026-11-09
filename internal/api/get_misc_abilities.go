package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMiscAbility(r *http.Request, i handlerinput.Main[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.Named], id int32) (apitypes.MiscAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.MiscAbility{}, err
	}

	response := apitypes.MiscAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     handlerinput.IDToTypedAPIResource(apilookups.E.Abilities, ability.Ability.ID),
		Description:        ability.Description,
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		LearnedBy:          handlerinput.NamesToNamedAPIResources(apilookups.E.CharacterClasses, ability.LearnedBy),
		Topmenu:            handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Topmenus, ability.Topmenu, nil),
		Submenu:            handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Submenus, ability.Submenu, nil),
		OpenSubmenu:        handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Submenus, ability.OpenSubmenu, nil),
		Cursor:             ability.Cursor,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, apiconv.ConvertBattleInteraction),
	}

	battleInteractions, err := applyUser(r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return apitypes.MiscAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func retrieveMiscAbilities(r *http.Request, i handlerinput.Main[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeMiscAbility

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryIntList(r, ids, ts.QpnRank, getTypedAbilityIDsByRank(abilityType)),
		i.QueryNameID(r, ids, ts.QpnUser, apilookups.E.CharacterClasses.ResTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetMiscAbilityIDsByCharClass),
		i.QueryBool(r, ids, ts.QpnCopycat, getTypedAbilityIDsByCanCopycat(abilityType)),
		i.QueryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(abilityType)),
		i.QueryBool2(r, ids, ts.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(abilityType)),
	})
}
