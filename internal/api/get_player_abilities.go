package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getPlayerAbility(r *http.Request, i handlerInput[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource], id int32) (PlayerAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return PlayerAbility{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(srv, r.Context(), ability, config.Cfg.DB.GetPlayerAbilityMonsterIDs)
	if err != nil {
		return PlayerAbility{}, err
	}

	response := PlayerAbility{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Version:               ability.Version,
		Specification:         ability.Specification,
		UntypedAbility:        idToTypedAPIResource(srv, E.Abilities, ability.Ability.ID),
		Description:           ability.Description,
		Effect:                ability.Effect,
		Rank:                  ability.Rank,
		AppearsInHelpBar:      ability.AppearsInHelpBar,
		CanCopycat:            ability.CanCopycat,
		CanUseOutsideBattle:   ability.CanUseOutsideBattle,
		MpCost:                ability.MPCost,
		Category:              ability.Category,
		AeonLearnItem:         E.AllItems.toResAmtFromNamePtr(srv, ability.AeonLearnItem),
		LearnedBy:             namesToNamedAPIResources(srv, E.CharacterClasses, ability.LearnedBy),
		RelatedStats:          namesToNamedAPIResources(srv, E.Stats, ability.RelatedStats),
		StandardGridCharacter: namePtrToNamedAPIResPtr(srv, E.Characters, ability.StandardGridPos, nil),
		ExpertGridCharacter:   namePtrToNamedAPIResPtr(srv, E.Characters, ability.ExpertGridPos, nil),
		Topmenu:               namePtrToNamedAPIResPtr(srv, E.Topmenus, ability.Topmenu, nil),
		Submenu:               namePtrToNamedAPIResPtr(srv, E.Submenus, ability.Submenu, nil),
		OpenSubmenu:           namePtrToNamedAPIResPtr(srv, E.Submenus, ability.OpenSubmenu, nil),
		Cursor:                ability.Cursor,
		Monsters:              monsters,
		BattleInteractions:    convertObjSlice(srv, ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(srv, r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (srv *Server) retrievePlayerAbilities(r *http.Request, i handlerInput[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.PlayerAbilityCategory, ids, ts.QpnCategory, config.Cfg.DB.GetPlayerAbilityIDsByCategory),
		i.queryEnumList(srv, r, enums.T.DamageType, ids, ts.QpnDamageType, getTypedAbilityIDsByDamageType(srv, abilityType)),
		i.queryEnumList(srv, r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(srv, abilityType)),
		i.queryEnumList(srv, r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(srv, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(srv, abilityType)),
		i.queryIntList(srv, r, ids, ts.QpnMp, config.Cfg.DB.GetPlayerAbilityIDsByMpCost),
		i.queryInt(r, ids, ts.QpnMpMin, config.Cfg.DB.GetPlayerAbilityIDsByMpCostMin),
		i.queryInt(r, ids, ts.QpnMpMax, config.Cfg.DB.GetPlayerAbilityIDsByMpCostMax),
		i.queryIntList(srv, r, ids, ts.QpnRank, getTypedAbilityIDsByRank(srv, abilityType)),
		i.queryNameIDListNul(srv, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(srv, abilityType)),
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetPlayerAbilityIDsByCharClass),
		i.queryID(r, ids, ts.QpnLearnItem, lookup.L.Items, config.Cfg.DB.GetPlayerAbilityIDsByLearnItem),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, config.Cfg.DB.GetPlayerAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(srv, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(srv, abilityType)),
		i.queryNameID(r, ids, ts.QpnStdSg, E.Characters.resTypeSingle, lookup.L.Characters, ToIntManyNull(config.Cfg.DB.GetPlayerAbilityIDsStdSgChar)),
		i.queryNameID(r, ids, ts.QpnExpSg, E.Characters.resTypeSingle, lookup.L.Characters, ToIntManyNull(config.Cfg.DB.GetPlayerAbilityIDsExpSgChar)),
		i.queryBool(r, ids, ts.QpnOutsideBattle, config.Cfg.DB.GetPlayerAbilityIDsCanUseOutsideBattle),
		i.queryBool(r, ids, ts.QpnCopycat, getTypedAbilityIDsByCanCopycat(srv, abilityType)),
		i.queryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnDarkable, getTypedAbilityIDsDarkable(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnSilenceable, getTypedAbilityIDsSilenceable(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnReflectable, getTypedAbilityIDsReflectable(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(srv, abilityType)),
	})
}
