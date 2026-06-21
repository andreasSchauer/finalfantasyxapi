package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getPlayerAbility(r *http.Request, i handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList], id int32) (PlayerAbility, error) {
	ability, err := verifyParamsAndGet(cfg, r, i, id)
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

	battleInteractions, err := applyUser(cfg, r, i, response, "ability_user")
	if err != nil {
		return PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrievePlayerAbilities(r *http.Request, i handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.PlayerAbilityCategory, ids, "category", cfg.db.GetPlayerAbilityIDsByCategory)),
		fidl(enumListQuery(cfg, r, i, cfg.t.DamageType, ids, "damage_type", getTypedAbilityIDsByDamageType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.AttackType, ids, "attack_type", getTypedAbilityIDsByAttackType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.TargetType, ids, "target_type", getTypedAbilityIDsByTargetType(cfg, abilityType))),
		fidl(enumQuery(r, i, cfg.t.DamageFormula, ids, "damage_formula", getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		fidl(intListQuery(cfg, r, i, ids, "mp", cfg.db.GetPlayerAbilityIDsByMpCost)),
		fidl(intQuery(r, i, ids, "mp_min", cfg.db.GetPlayerAbilityIDsByMpCostMin)),
		fidl(intQuery(r, i, ids, "mp_max", cfg.db.GetPlayerAbilityIDsByMpCostMax)),
		fidl(intListQuery(cfg, r, i, ids, "rank", getTypedAbilityIDsByRank(cfg, abilityType))),
		fidl(nameIdListQueryNul(cfg, r, i, ids, "element", cfg.e.elements.resourceType, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		fidl(nameIdQuery(r, i, ids, "user", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetPlayerAbilityIDsByCharClass)),
		fidl(idQuery(r, i, ids, "learn_item", cfg.l.Items, cfg.db.GetPlayerAbilityIDsByLearnItem)),
		fidl(nameIdQuery(r, i, ids, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetPlayerAbilityIDsByRelatedStat)),
		fidl(idQueryNul(r, i, ids, "status_inflict", cfg.l.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		fidl(idQueryNul(r, i, ids, "status_remove", cfg.l.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		fidl(nameIdQuery(r, i, ids, "std_sg", cfg.e.characters.resourceType, cfg.l.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsStdSgChar))),
		fidl(nameIdQuery(r, i, ids, "exp_sg", cfg.e.characters.resourceType, cfg.l.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsExpSgChar))),
		fidl(boolQuery(r, i, ids, "outside_battle", cfg.db.GetPlayerAbilityIDsCanUseOutsideBattle)),
		fidl(boolQuery(r, i, ids, "copycat", getTypedAbilityIDsByCanCopycat(cfg, abilityType))),
		fidl(boolQuery(r, i, ids, "help_bar", getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "user_atk", getTypedAbilityIDsBasedOnUserAttack(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "darkable", getTypedAbilityIDsDarkable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "silenceable", getTypedAbilityIDsSilenceable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "reflectable", getTypedAbilityIDsReflectable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "delay", getTypedAbilityIDsDealsDelay(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "stat_changes", getTypedAbilityIDsWithStatChanges(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "mod_changes", getTypedAbilityIDsWithModifierChanges(cfg, abilityType))),
	})
}
