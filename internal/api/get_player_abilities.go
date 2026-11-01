package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getPlayerAbility(r *http.Request, i handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource], id int32) (PlayerAbility, error) {
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

	battleInteractions, err := applyUser(cfg, r, i, response, QpnAbilityUser)
	if err != nil {
		return PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrievePlayerAbilities(r *http.Request, i handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, cfg.T.PlayerAbilityCategory, ids, QpnCategory, cfg.db.GetPlayerAbilityIDsByCategory),
		i.queryEnumList(cfg, r, cfg.T.DamageType, ids, QpnDamageType, getTypedAbilityIDsByDamageType(cfg, abilityType)),
		i.queryEnumList(cfg, r, cfg.T.AttackType, ids, QpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		i.queryEnumList(cfg, r, cfg.T.TargetType, ids, QpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		i.queryEnum(r, cfg.T.DamageFormula, ids, QpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		i.queryIntList(cfg, r, ids, QpnMp, cfg.db.GetPlayerAbilityIDsByMpCost),
		i.queryInt(r, ids, QpnMpMin, cfg.db.GetPlayerAbilityIDsByMpCostMin),
		i.queryInt(r, ids, QpnMpMax, cfg.db.GetPlayerAbilityIDsByMpCostMax),
		i.queryIntList(cfg, r, ids, QpnRank, getTypedAbilityIDsByRank(cfg, abilityType)),
		i.queryNameIDListNul(cfg, r, ids, QpnElement, cfg.E.Elements.resTypeSingle, cfg.L.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		i.queryNameID(r, ids, QpnUser, cfg.E.CharacterClasses.resTypeSingle, cfg.L.CharClasses, cfg.db.GetPlayerAbilityIDsByCharClass),
		i.queryID(r, ids, QpnLearnItem, cfg.L.Items, cfg.db.GetPlayerAbilityIDsByLearnItem),
		i.queryNameID(r, ids, QpnRelatedStat, cfg.E.Stats.resTypeSingle, cfg.L.Stats, cfg.db.GetPlayerAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, QpnStatusInflict, cfg.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		i.queryIDNul(r, ids, QpnStatusRemove, cfg.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		i.queryNameID(r, ids, QpnStdSg, cfg.E.Characters.resTypeSingle, cfg.L.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsStdSgChar)),
		i.queryNameID(r, ids, QpnExpSg, cfg.E.Characters.resTypeSingle, cfg.L.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsExpSgChar)),
		i.queryBool(r, ids, QpnOutsideBattle, cfg.db.GetPlayerAbilityIDsCanUseOutsideBattle),
		i.queryBool(r, ids, QpnCopycat, getTypedAbilityIDsByCanCopycat(cfg, abilityType)),
		i.queryBool(r, ids, QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType)),
		i.queryBool2(r, ids, QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(cfg, abilityType)),
		i.queryBool2(r, ids, QpnDarkable, getTypedAbilityIDsDarkable(cfg, abilityType)),
		i.queryBool2(r, ids, QpnSilenceable, getTypedAbilityIDsSilenceable(cfg, abilityType)),
		i.queryBool2(r, ids, QpnReflectable, getTypedAbilityIDsReflectable(cfg, abilityType)),
		i.queryBool2(r, ids, QpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
		i.queryBool2(r, ids, QpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType)),
		i.queryBool2(r, ids, QpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType)),
	})
}
