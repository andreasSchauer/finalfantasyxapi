package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getPlayerAbility(r *http.Request, i HandlerInput[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.NamedAPIResource], id int32) (apitypes.PlayerAbility, error) {
	ability, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}

	monsters, err := E.Monsters.DBGetRelations(r.Context(), ability, config.Cfg.DB.GetPlayerAbilityMonsterIDs)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}

	response := apitypes.PlayerAbility{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Version:               ability.Version,
		Specification:         ability.Specification,
		UntypedAbility:        IDToTypedAPIResource(E.Abilities, ability.Ability.ID),
		Description:           ability.Description,
		Effect:                ability.Effect,
		Rank:                  ability.Rank,
		AppearsInHelpBar:      ability.AppearsInHelpBar,
		CanCopycat:            ability.CanCopycat,
		CanUseOutsideBattle:   ability.CanUseOutsideBattle,
		MpCost:                ability.MPCost,
		Category:              ability.Category,
		AeonLearnItem:         E.AllItems.ToResAmtFromNamePtr(ability.AeonLearnItem),
		LearnedBy:             NamesToNamedAPIResources(E.CharacterClasses, ability.LearnedBy),
		RelatedStats:          NamesToNamedAPIResources(E.Stats, ability.RelatedStats),
		StandardGridCharacter: NamePtrToNamedAPIResPtr(E.Characters, ability.StandardGridPos, nil),
		ExpertGridCharacter:   NamePtrToNamedAPIResPtr(E.Characters, ability.ExpertGridPos, nil),
		Topmenu:               NamePtrToNamedAPIResPtr(E.Topmenus, ability.Topmenu, nil),
		Submenu:               NamePtrToNamedAPIResPtr(E.Submenus, ability.Submenu, nil),
		OpenSubmenu:           NamePtrToNamedAPIResPtr(E.Submenus, ability.OpenSubmenu, nil),
		Cursor:                ability.Cursor,
		Monsters:              monsters,
		BattleInteractions:    ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func retrievePlayerAbilities(r *http.Request, i HandlerInput[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return i.FilterIDs(r, ids, []IDFilter{
		i.QueryEnumList(r, enums.T.PlayerAbilityCategory, ids, ts.QpnCategory, config.Cfg.DB.GetPlayerAbilityIDsByCategory),
		i.QueryEnumList(r, enums.T.DamageType, ids, ts.QpnDamageType, getTypedAbilityIDsByDamageType(abilityType)),
		i.QueryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(abilityType)),
		i.QueryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(abilityType)),
		i.QueryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(abilityType)),
		i.QueryIntList(r, ids, ts.QpnMp, config.Cfg.DB.GetPlayerAbilityIDsByMpCost),
		i.QueryInt(r, ids, ts.QpnMpMin, config.Cfg.DB.GetPlayerAbilityIDsByMpCostMin),
		i.QueryInt(r, ids, ts.QpnMpMax, config.Cfg.DB.GetPlayerAbilityIDsByMpCostMax),
		i.QueryIntList(r, ids, ts.QpnRank, getTypedAbilityIDsByRank(abilityType)),
		i.QueryNameIDListNul(r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(abilityType)),
		i.QueryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetPlayerAbilityIDsByCharClass),
		i.QueryID(r, ids, ts.QpnLearnItem, lookup.L.Items, config.Cfg.DB.GetPlayerAbilityIDsByLearnItem),
		i.QueryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, config.Cfg.DB.GetPlayerAbilityIDsByRelatedStat),
		i.QueryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(abilityType)),
		i.QueryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(abilityType)),
		i.QueryNameID(r, ids, ts.QpnStdSg, E.Characters.resTypeSingle, lookup.L.Characters, ToIntManyNull(config.Cfg.DB.GetPlayerAbilityIDsStdSgChar)),
		i.QueryNameID(r, ids, ts.QpnExpSg, E.Characters.resTypeSingle, lookup.L.Characters, ToIntManyNull(config.Cfg.DB.GetPlayerAbilityIDsExpSgChar)),
		i.QueryBool(r, ids, ts.QpnOutsideBattle, config.Cfg.DB.GetPlayerAbilityIDsCanUseOutsideBattle),
		i.QueryBool(r, ids, ts.QpnCopycat, getTypedAbilityIDsByCanCopycat(abilityType)),
		i.QueryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(abilityType)),
		i.QueryBool2(r, ids, ts.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDarkable, getTypedAbilityIDsDarkable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnSilenceable, getTypedAbilityIDsSilenceable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnReflectable, getTypedAbilityIDsReflectable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(abilityType)),
		i.QueryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(abilityType)),
		i.QueryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(abilityType)),
	})
}
