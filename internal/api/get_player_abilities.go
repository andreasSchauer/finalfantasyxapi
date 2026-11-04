package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getPlayerAbility(r *http.Request, i handlerInput[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource], id int32) (PlayerAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return PlayerAbility{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(s, r.Context(), ability, s.Cfg.DB.GetPlayerAbilityMonsterIDs)
	if err != nil {
		return PlayerAbility{}, err
	}

	response := PlayerAbility{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Version:               ability.Version,
		Specification:         ability.Specification,
		UntypedAbility:        idToTypedAPIResource(s, E.Abilities, ability.Ability.ID),
		Description:           ability.Description,
		Effect:                ability.Effect,
		Rank:                  ability.Rank,
		AppearsInHelpBar:      ability.AppearsInHelpBar,
		CanCopycat:            ability.CanCopycat,
		CanUseOutsideBattle:   ability.CanUseOutsideBattle,
		MpCost:                ability.MPCost,
		Category:              ability.Category,
		AeonLearnItem:         E.AllItems.toResAmtFromNamePtr(s, ability.AeonLearnItem),
		LearnedBy:             namesToNamedAPIResources(s, E.CharacterClasses, ability.LearnedBy),
		RelatedStats:          namesToNamedAPIResources(s, E.Stats, ability.RelatedStats),
		StandardGridCharacter: namePtrToNamedAPIResPtr(s, E.Characters, ability.StandardGridPos, nil),
		ExpertGridCharacter:   namePtrToNamedAPIResPtr(s, E.Characters, ability.ExpertGridPos, nil),
		Topmenu:               namePtrToNamedAPIResPtr(s, E.Topmenus, ability.Topmenu, nil),
		Submenu:               namePtrToNamedAPIResPtr(s, E.Submenus, ability.Submenu, nil),
		OpenSubmenu:           namePtrToNamedAPIResPtr(s, E.Submenus, ability.OpenSubmenu, nil),
		Cursor:                ability.Cursor,
		Monsters:              monsters,
		BattleInteractions:    convertObjSlice(s, ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(s, r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (s *Server) retrievePlayerAbilities(r *http.Request, i handlerInput[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryEnumList(s, r, enums.T.PlayerAbilityCategory, ids, ts.QpnCategory, s.Cfg.DB.GetPlayerAbilityIDsByCategory),
		i.queryEnumList(s, r, enums.T.DamageType, ids, ts.QpnDamageType, getTypedAbilityIDsByDamageType(s, abilityType)),
		i.queryEnumList(s, r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(s, abilityType)),
		i.queryEnumList(s, r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(s, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(s, abilityType)),
		i.queryIntList(s, r, ids, ts.QpnMp, s.Cfg.DB.GetPlayerAbilityIDsByMpCost),
		i.queryInt(r, ids, ts.QpnMpMin, s.Cfg.DB.GetPlayerAbilityIDsByMpCostMin),
		i.queryInt(r, ids, ts.QpnMpMax, s.Cfg.DB.GetPlayerAbilityIDsByMpCostMax),
		i.queryIntList(s, r, ids, ts.QpnRank, getTypedAbilityIDsByRank(s, abilityType)),
		i.queryNameIDListNul(s, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(s, abilityType)),
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, s.Cfg.DB.GetPlayerAbilityIDsByCharClass),
		i.queryID(r, ids, ts.QpnLearnItem, lookup.L.Items, s.Cfg.DB.GetPlayerAbilityIDsByLearnItem),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, s.Cfg.DB.GetPlayerAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(s, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(s, abilityType)),
		i.queryNameID(r, ids, ts.QpnStdSg, E.Characters.resTypeSingle, lookup.L.Characters, ToIntManyNull(s.Cfg.DB.GetPlayerAbilityIDsStdSgChar)),
		i.queryNameID(r, ids, ts.QpnExpSg, E.Characters.resTypeSingle, lookup.L.Characters, ToIntManyNull(s.Cfg.DB.GetPlayerAbilityIDsExpSgChar)),
		i.queryBool(r, ids, ts.QpnOutsideBattle, s.Cfg.DB.GetPlayerAbilityIDsCanUseOutsideBattle),
		i.queryBool(r, ids, ts.QpnCopycat, getTypedAbilityIDsByCanCopycat(s, abilityType)),
		i.queryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnDarkable, getTypedAbilityIDsDarkable(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnSilenceable, getTypedAbilityIDsSilenceable(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnReflectable, getTypedAbilityIDsReflectable(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(s, abilityType)),
	})
}
