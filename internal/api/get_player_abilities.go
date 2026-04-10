package api

import (
	"net/http"

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
		Category:              newNamedAPIResourceFromEnum(cfg, cfg.e.playerAbilityCategory.endpoint, ability.Category, cfg.t.PlayerAbilityCategory),
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
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(enumListQuery(cfg, r, i, cfg.t.PlayerAbilityCategory, resources, "category", cfg.db.GetPlayerAbilityIDsByCategory)),
		frl(enumListQuery(cfg, r, i, cfg.t.DamageType, resources, "damage_type", cfg.db.GetPlayerAbilityIDsByDamageType)),
		frl(enumListQuery(cfg, r, i, cfg.t.AttackType, resources, "attack_type", cfg.db.GetPlayerAbilityIDsByAttackType)),
		frl(enumListQuery(cfg, r, i, cfg.t.TargetType, resources, "target_type", cfg.db.GetPlayerAbilityIDsByTargetType)),
		frl(enumQuery(cfg, r, i, cfg.t.DamageFormula, resources, "damage_formula", cfg.db.GetPlayerAbilityIDsByDamageFormula)),
		frl(intListQuery(cfg, r, i, resources, "mp", cfg.db.GetPlayerAbilityIDsByMpCost)),
		frl(intQuery(cfg, r, i, resources, "mp_min", cfg.db.GetPlayerAbilityIDsByMpCostMin)),
		frl(intQuery(cfg, r, i, resources, "mp_max", cfg.db.GetPlayerAbilityIDsByMpCostMax)),
		frl(intListQuery(cfg, r, i, resources, "rank", cfg.db.GetPlayerAbilityIDsByRank)),
		frl(nameIdListQueryNul(cfg, r, i, resources, "element", cfg.e.elements.resourceType, cfg.l.Elements, cfg.db.GetPlayerAbilityIDsByElement)),
		frl(nameIdQuery(cfg, r, i, resources, "user", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetPlayerAbilityIDsByCharClass)),
		frl(idQuery(cfg, r, i, resources, "learn_item", len(cfg.l.Items), cfg.db.GetPlayerAbilityIDsByLearnItem)),
		frl(nameIdQuery(cfg, r, i, resources, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetPlayerAbilityIDsByRelatedStat)),
		frl(idQueryNul(cfg, r, i, resources, "status_inflict", len(cfg.l.StatusConditions), cfg.db.GetPlayerAbilityIDsByInflictedStatus)),
		frl(idQueryNul(cfg, r, i, resources, "status_remove", len(cfg.l.StatusConditions), cfg.db.GetPlayerAbilityIDsByRemovedStatus)),
		frl(nameIdQuery(cfg, r, i, resources, "std_sg", cfg.e.characters.resourceType, cfg.l.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsStdSgChar))),
		frl(nameIdQuery(cfg, r, i, resources, "exp_sg", cfg.e.characters.resourceType, cfg.l.Characters, ToIntManyNull(cfg.db.GetPlayerAbilityIDsExpSgChar))),
		frl(boolQuery(cfg, r, i, resources, "outside_battle", cfg.db.GetPlayerAbilityIDsCanUseOutsideBattle)),
		frl(boolQuery(cfg, r, i, resources, "copycat", cfg.db.GetPlayerAbilityIDsByCanCopycat)),
		frl(boolQuery(cfg, r, i, resources, "help_bar", cfg.db.GetPlayerAbilityIDsByAppearsInHelpBar)),
		frl(boolQuery2(cfg, r, i, resources, "user_atk", cfg.db.GetPlayerAbilityIDsBasedOnUserAttack)),
		frl(boolQuery2(cfg, r, i, resources, "darkable", cfg.db.GetPlayerAbilityIDsDarkable)),
		frl(boolQuery2(cfg, r, i, resources, "silenceable", cfg.db.GetPlayerAbilityIDsSilenceable)),
		frl(boolQuery2(cfg, r, i, resources, "reflectable", cfg.db.GetPlayerAbilityIDsReflectable)),
		frl(boolQuery2(cfg, r, i, resources, "delay", cfg.db.GetPlayerAbilityIDsDealsDelay)),
		frl(boolQuery2(cfg, r, i, resources, "stat_changes", cfg.db.GetPlayerAbilityIDsWithStatChanges)),
		frl(boolQuery2(cfg, r, i, resources, "mod_changes", cfg.db.GetPlayerAbilityIDsWithModifierChanges)),
	})
}
