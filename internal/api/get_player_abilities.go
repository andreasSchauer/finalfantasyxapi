package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getPlayerAbility(r *http.Request, i handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList], id int32) (PlayerAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return PlayerAbility{}, err
	}

	monsters, err := getResourcesDbItem(cfg, r, cfg.e.monsters, ability, cfg.db.GetPlayerAbilityMonsterIDs)
	if err != nil {
		return PlayerAbility{}, err
	}

	response := PlayerAbility{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Version:               ability.Version,
		Specification:         ability.Specification,
		UntypedAbility:        idToTypedAPIResource(cfg, cfg.e.abilities, ability.Ability.ID),
		Description:           ability.Description,
		Effect:                ability.Effect,
		Rank:                  ability.Rank,
		AppearsInHelpBar:      ability.AppearsInHelpBar,
		CanCopycat:            ability.CanCopycat,
		CanUseOutsideBattle:   ability.CanUseOutsideBattle,
		MpCost:                ability.MPCost,
		Category:              enumToNamedAPIResource(cfg, cfg.e.playerAbilityCategory.endpoint, ability.Category, cfg.t.PlayerAbilityCategory),
		AeonLearnItem:         nameAmountPtrToResAmtPtr(cfg, cfg.e.allItems, ability.AeonLearnItem),
		LearnedBy:             namesToNamedAPIResources(cfg, cfg.e.characterClasses, ability.LearnedBy),
		RelatedStats:          namesToNamedAPIResources(cfg, cfg.e.stats, ability.RelatedStats),
		StandardGridCharacter: namePtrToNamedAPIResPtr(cfg, cfg.e.characters, ability.StandardGridPos, nil),
		ExpertGridCharacter:   namePtrToNamedAPIResPtr(cfg, cfg.e.characters, ability.ExpertGridPos, nil),
		Topmenu:               namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, ability.Topmenu, nil),
		Submenu:               namePtrToNamedAPIResPtr(cfg, cfg.e.submenus, ability.Submenu, nil),
		OpenSubmenu:           namePtrToNamedAPIResPtr(cfg, cfg.e.submenus, ability.OpenSubmenu, nil),
		Cursor:                ability.Cursor,
		Monsters:              monsters,
		BattleInteractions:    convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(cfg, r, i, response, qpnAbilityUser)
	if err != nil {
		return PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrievePlayerAbilities(r *http.Request, i handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.PlayerAbilityCategory, ids, qpnCategory, cfg.db.GetPlayerAbilityIDsByCategory)),
		fidl(enumListQuery(cfg, r, i, cfg.t.DamageType, ids, qpnDamageType, getTypedAbilityIDsByDamageType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.AttackType, ids, qpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.TargetType, ids, qpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType))),
		fidl(enumQuery(r, i, cfg.t.DamageFormula, ids, qpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		fidl(intListQuery(cfg, r, i, ids, qpnMp, cfg.db.GetPlayerAbilityIDsByMpCost)),
		fidl(intQuery(r, i, ids, qpnMpMin, cfg.db.GetPlayerAbilityIDsByMpCostMin)),
		fidl(intQuery(r, i, ids, qpnMpMax, cfg.db.GetPlayerAbilityIDsByMpCostMax)),
		fidl(intListQuery(cfg, r, i, ids, qpnRank, getTypedAbilityIDsByRank(cfg, abilityType))),
		fidl(nameIdListQueryNul(cfg, r, i, ids, qpnElement, cfg.e.elements.resTypeSing, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		fidl(nameIdQuery(r, i, ids, qpnUser, cfg.e.characterClasses.resTypeSing, cfg.l.CharClasses, cfg.db.GetPlayerAbilityIDsByCharClass)),
		fidl(idQuery(r, i, ids, qpnLearnItem, cfg.l.Items, cfg.db.GetPlayerAbilityIDsByLearnItem)),
		fidl(nameIdQuery(r, i, ids, qpnRelatedStat, cfg.e.stats.resTypeSing, cfg.l.Stats, cfg.db.GetPlayerAbilityIDsByRelatedStat)),
		fidl(idQueryNul(r, i, ids, qpnStatusInflict, cfg.l.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		fidl(idQueryNul(r, i, ids, qpnStatusRemove, cfg.l.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		fidl(nameIdQuery(r, i, ids, qpnStdSg, cfg.e.characters.resTypeSing, cfg.l.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsStdSgChar))),
		fidl(nameIdQuery(r, i, ids, qpnExpSg, cfg.e.characters.resTypeSing, cfg.l.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsExpSgChar))),
		fidl(boolQuery(r, i, ids, qpnOutsideBattle, cfg.db.GetPlayerAbilityIDsCanUseOutsideBattle)),
		fidl(boolQuery(r, i, ids, qpnCopycat, getTypedAbilityIDsByCanCopycat(cfg, abilityType))),
		fidl(boolQuery(r, i, ids, qpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnDarkable, getTypedAbilityIDsDarkable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnSilenceable, getTypedAbilityIDsSilenceable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnReflectable, getTypedAbilityIDsReflectable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType))),
	})
}
