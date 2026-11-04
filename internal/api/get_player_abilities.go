package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getPlayerAbility(r *http.Request, i handlerInput[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource], id int32) (PlayerAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return PlayerAbility{}, err
	}

	monsters, err := cfg.E.Monsters.dbGetRelations(cfg, r.Context(), ability, cfg.db.GetPlayerAbilityMonsterIDs)
	if err != nil {
		return PlayerAbility{}, err
	}

	response := PlayerAbility{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Version:               ability.Version,
		Specification:         ability.Specification,
		UntypedAbility:        idToTypedAPIResource(cfg, cfg.E.Abilities, ability.Ability.ID),
		Description:           ability.Description,
		Effect:                ability.Effect,
		Rank:                  ability.Rank,
		AppearsInHelpBar:      ability.AppearsInHelpBar,
		CanCopycat:            ability.CanCopycat,
		CanUseOutsideBattle:   ability.CanUseOutsideBattle,
		MpCost:                ability.MPCost,
		Category:              ability.Category,
		AeonLearnItem:         cfg.E.AllItems.toResAmtFromNamePtr(cfg, ability.AeonLearnItem),
		LearnedBy:             namesToNamedAPIResources(cfg, cfg.E.CharacterClasses, ability.LearnedBy),
		RelatedStats:          namesToNamedAPIResources(cfg, cfg.E.Stats, ability.RelatedStats),
		StandardGridCharacter: namePtrToNamedAPIResPtr(cfg, cfg.E.Characters, ability.StandardGridPos, nil),
		ExpertGridCharacter:   namePtrToNamedAPIResPtr(cfg, cfg.E.Characters, ability.ExpertGridPos, nil),
		Topmenu:               namePtrToNamedAPIResPtr(cfg, cfg.E.Topmenus, ability.Topmenu, nil),
		Submenu:               namePtrToNamedAPIResPtr(cfg, cfg.E.Submenus, ability.Submenu, nil),
		OpenSubmenu:           namePtrToNamedAPIResPtr(cfg, cfg.E.Submenus, ability.OpenSubmenu, nil),
		Cursor:                ability.Cursor,
		Monsters:              monsters,
		BattleInteractions:    convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(cfg, r, i, response, typestrings.QpnAbilityUser)
	if err != nil {
		return PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrievePlayerAbilities(r *http.Request, i handlerInput[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, enums.T.PlayerAbilityCategory, ids, typestrings.QpnCategory, cfg.db.GetPlayerAbilityIDsByCategory),
		i.queryEnumList(cfg, r, enums.T.DamageType, ids, typestrings.QpnDamageType, getTypedAbilityIDsByDamageType(cfg, abilityType)),
		i.queryEnumList(cfg, r, enums.T.AttackType, ids, typestrings.QpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		i.queryEnumList(cfg, r, enums.T.TargetType, ids, typestrings.QpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, typestrings.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		i.queryIntList(cfg, r, ids, typestrings.QpnMp, cfg.db.GetPlayerAbilityIDsByMpCost),
		i.queryInt(r, ids, typestrings.QpnMpMin, cfg.db.GetPlayerAbilityIDsByMpCostMin),
		i.queryInt(r, ids, typestrings.QpnMpMax, cfg.db.GetPlayerAbilityIDsByMpCostMax),
		i.queryIntList(cfg, r, ids, typestrings.QpnRank, getTypedAbilityIDsByRank(cfg, abilityType)),
		i.queryNameIDListNul(cfg, r, ids, typestrings.QpnElement, cfg.E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		i.queryNameID(r, ids, typestrings.QpnUser, cfg.E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, cfg.db.GetPlayerAbilityIDsByCharClass),
		i.queryID(r, ids, typestrings.QpnLearnItem, lookup.L.Items, cfg.db.GetPlayerAbilityIDsByLearnItem),
		i.queryNameID(r, ids, typestrings.QpnRelatedStat, cfg.E.Stats.resTypeSingle, lookup.L.Stats, cfg.db.GetPlayerAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, typestrings.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		i.queryIDNul(r, ids, typestrings.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		i.queryNameID(r, ids, typestrings.QpnStdSg, cfg.E.Characters.resTypeSingle, lookup.L.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsStdSgChar)),
		i.queryNameID(r, ids, typestrings.QpnExpSg, cfg.E.Characters.resTypeSingle, lookup.L.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsExpSgChar)),
		i.queryBool(r, ids, typestrings.QpnOutsideBattle, cfg.db.GetPlayerAbilityIDsCanUseOutsideBattle),
		i.queryBool(r, ids, typestrings.QpnCopycat, getTypedAbilityIDsByCanCopycat(cfg, abilityType)),
		i.queryBool(r, ids, typestrings.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnDarkable, getTypedAbilityIDsDarkable(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnSilenceable, getTypedAbilityIDsSilenceable(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnReflectable, getTypedAbilityIDsReflectable(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType)),
	})
}
