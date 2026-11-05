package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getPlayerAbility(r *http.Request, i handlerInput[seedtypes.PlayerAbility, PlayerAbility, apiresource.NamedAPIResource], id int32) (PlayerAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return PlayerAbility{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(r.Context(), ability, config.Cfg.DB.GetPlayerAbilityMonsterIDs)
	if err != nil {
		return PlayerAbility{}, err
	}

	response := PlayerAbility{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Version:               ability.Version,
		Specification:         ability.Specification,
		UntypedAbility:        idToTypedAPIResource(E.Abilities, ability.Ability.ID),
		Description:           ability.Description,
		Effect:                ability.Effect,
		Rank:                  ability.Rank,
		AppearsInHelpBar:      ability.AppearsInHelpBar,
		CanCopycat:            ability.CanCopycat,
		CanUseOutsideBattle:   ability.CanUseOutsideBattle,
		MpCost:                ability.MPCost,
		Category:              ability.Category,
		AeonLearnItem:         E.AllItems.toResAmtFromNamePtr(ability.AeonLearnItem),
		LearnedBy:             namesToNamedAPIResources(E.CharacterClasses, ability.LearnedBy),
		RelatedStats:          namesToNamedAPIResources(E.Stats, ability.RelatedStats),
		StandardGridCharacter: namePtrToNamedAPIResPtr(E.Characters, ability.StandardGridPos, nil),
		ExpertGridCharacter:   namePtrToNamedAPIResPtr(E.Characters, ability.ExpertGridPos, nil),
		Topmenu:               namePtrToNamedAPIResPtr(E.Topmenus, ability.Topmenu, nil),
		Submenu:               namePtrToNamedAPIResPtr(E.Submenus, ability.Submenu, nil),
		OpenSubmenu:           namePtrToNamedAPIResPtr(E.Submenus, ability.OpenSubmenu, nil),
		Cursor:                ability.Cursor,
		Monsters:              monsters,
		BattleInteractions:    ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func retrievePlayerAbilities(r *http.Request, i handlerInput[seedtypes.PlayerAbility, PlayerAbility, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnumList(r, enums.T.PlayerAbilityCategory, ids, ts.QpnCategory, config.Cfg.DB.GetPlayerAbilityIDsByCategory),
		i.queryEnumList(r, enums.T.DamageType, ids, ts.QpnDamageType, getTypedAbilityIDsByDamageType(abilityType)),
		i.queryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(abilityType)),
		i.queryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(abilityType)),
		i.queryIntList(r, ids, ts.QpnMp, config.Cfg.DB.GetPlayerAbilityIDsByMpCost),
		i.queryInt(r, ids, ts.QpnMpMin, config.Cfg.DB.GetPlayerAbilityIDsByMpCostMin),
		i.queryInt(r, ids, ts.QpnMpMax, config.Cfg.DB.GetPlayerAbilityIDsByMpCostMax),
		i.queryIntList(r, ids, ts.QpnRank, getTypedAbilityIDsByRank(abilityType)),
		i.queryNameIDListNul(r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(abilityType)),
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetPlayerAbilityIDsByCharClass),
		i.queryID(r, ids, ts.QpnLearnItem, lookup.L.Items, config.Cfg.DB.GetPlayerAbilityIDsByLearnItem),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, config.Cfg.DB.GetPlayerAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(abilityType)),
		i.queryNameID(r, ids, ts.QpnStdSg, E.Characters.resTypeSingle, lookup.L.Characters, ToIntManyNull(config.Cfg.DB.GetPlayerAbilityIDsStdSgChar)),
		i.queryNameID(r, ids, ts.QpnExpSg, E.Characters.resTypeSingle, lookup.L.Characters, ToIntManyNull(config.Cfg.DB.GetPlayerAbilityIDsExpSgChar)),
		i.queryBool(r, ids, ts.QpnOutsideBattle, config.Cfg.DB.GetPlayerAbilityIDsCanUseOutsideBattle),
		i.queryBool(r, ids, ts.QpnCopycat, getTypedAbilityIDsByCanCopycat(abilityType)),
		i.queryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(abilityType)),
		i.queryBool2(r, ids, ts.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(abilityType)),
		i.queryBool2(r, ids, ts.QpnDarkable, getTypedAbilityIDsDarkable(abilityType)),
		i.queryBool2(r, ids, ts.QpnSilenceable, getTypedAbilityIDsSilenceable(abilityType)),
		i.queryBool2(r, ids, ts.QpnReflectable, getTypedAbilityIDsReflectable(abilityType)),
		i.queryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(abilityType)),
		i.queryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(abilityType)),
		i.queryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(abilityType)),
	})
}
