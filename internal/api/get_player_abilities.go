package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getPlayerAbility(r *http.Request, i handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedAPIResourceList], id int32) (PlayerAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return PlayerAbility{}, err
	}

	monsters, err := getResourcesDBItem(cfg, r.Context(), cfg.E.Monsters, ability, cfg.db.GetPlayerAbilityMonsterIDs)
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
		AeonLearnItem:         nameAmountPtrToResAmtPtr(cfg, cfg.E.AllItems, ability.AeonLearnItem),
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

func (cfg *Config) retrievePlayerAbilities(r *http.Request, i handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return filterIDs(cfg, r, i, ids, []IDFilter{
		enumListQuery(cfg, r, i, cfg.T.PlayerAbilityCategory, ids, QpnCategory, cfg.db.GetPlayerAbilityIDsByCategory),
		enumListQuery(cfg, r, i, cfg.T.DamageType, ids, QpnDamageType, getTypedAbilityIDsByDamageType(cfg, abilityType)),
		enumListQuery(cfg, r, i, cfg.T.AttackType, ids, QpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		enumListQuery(cfg, r, i, cfg.T.TargetType, ids, QpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		enumQuery(r, i, cfg.T.DamageFormula, ids, QpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		intListQuery(cfg, r, i, ids, QpnMp, cfg.db.GetPlayerAbilityIDsByMpCost),
		intQuery(r, i, ids, QpnMpMin, cfg.db.GetPlayerAbilityIDsByMpCostMin),
		intQuery(r, i, ids, QpnMpMax, cfg.db.GetPlayerAbilityIDsByMpCostMax),
		intListQuery(cfg, r, i, ids, QpnRank, getTypedAbilityIDsByRank(cfg, abilityType)),
		nameIDListQueryNul(cfg, r, i, ids, QpnElement, cfg.E.Elements.resTypeSingle, cfg.L.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		nameIDQuery(r, i, ids, QpnUser, cfg.E.CharacterClasses.resTypeSingle, cfg.L.CharClasses, cfg.db.GetPlayerAbilityIDsByCharClass),
		idQuery(r, i, ids, QpnLearnItem, cfg.L.Items, cfg.db.GetPlayerAbilityIDsByLearnItem),
		nameIDQuery(r, i, ids, QpnRelatedStat, cfg.E.Stats.resTypeSingle, cfg.L.Stats, cfg.db.GetPlayerAbilityIDsByRelatedStat),
		idQueryNul(r, i, ids, QpnStatusInflict, cfg.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		idQueryNul(r, i, ids, QpnStatusRemove, cfg.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		nameIDQuery(r, i, ids, QpnStdSg, cfg.E.Characters.resTypeSingle, cfg.L.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsStdSgChar)),
		nameIDQuery(r, i, ids, QpnExpSg, cfg.E.Characters.resTypeSingle, cfg.L.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsExpSgChar)),
		boolQuery(r, i, ids, QpnOutsideBattle, cfg.db.GetPlayerAbilityIDsCanUseOutsideBattle),
		boolQuery(r, i, ids, QpnCopycat, getTypedAbilityIDsByCanCopycat(cfg, abilityType)),
		boolQuery(r, i, ids, QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnDarkable, getTypedAbilityIDsDarkable(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnSilenceable, getTypedAbilityIDsSilenceable(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnReflectable, getTypedAbilityIDsReflectable(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType)),
	})
}
