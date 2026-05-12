package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getItemAbility(r *http.Request, i handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList], id int32) (ItemAbility, error) {
	ability, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return ItemAbility{}, err
	}

	item, _ := seeding.GetResourceByID(id, cfg.l.ItemsID)

	response := ItemAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Item:               idToNamedAPIResource(cfg, cfg.e.items, id),
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.e.abilities, ability.Ability.ID),
		Description:        item.Description,
		Effect:             item.Effect,
		Category:           enumToNamedAPIResource(cfg, cfg.e.itemCategory.endpoint, item.Category, cfg.t.ItemCategory),
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		RelatedStats:       namesToNamedAPIResources(cfg, cfg.e.stats, item.RelatedStats),
		Cursor:             item.Cursor,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	if item.Usability == "outside-battle" || item.Usability == "always" {
		response.CanUseOutsideBattle = true
	}

	return response, nil
}

func (cfg *Config) retrieveItemAbilities(r *http.Request, i handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}
	abilityType := database.AbilityTypeItemAbility

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(enumListQuery(cfg, r, i, cfg.t.ItemCategory, resources, "category", cfg.db.GetItemAbilityIDsByCategory)),
		frl(enumListQuery(cfg, r, i, cfg.t.AttackType, resources, "attack_type", getTypedAbilityIDsByAttackType(cfg, abilityType))),
		frl(enumListQuery(cfg, r, i, cfg.t.TargetType, resources, "target_type", getTypedAbilityIDsByTargetType(cfg, abilityType))),
		frl(enumQuery(cfg, r, i, cfg.t.DamageFormula, resources, "damage_formula", getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		frl(nameIdListQueryNul(cfg, r, i, resources, "element", cfg.e.elements.resourceType, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		frl(nameIdQuery(cfg, r, i, resources, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetItemAbilityIDsByRelatedStat)),
		frl(idQueryNul(cfg, r, i, resources, "status_inflict", len(cfg.l.StatusConditions), getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		frl(idQueryNul(cfg, r, i, resources, "status_remove", len(cfg.l.StatusConditions), getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "outside_battle", cfg.db.GetItemAbilityIDsCanUseOutsideBattle)),
		frl(boolQuery2(cfg, r, i, resources, "delay", getTypedAbilityIDsDealsDelay(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "stat_changes", getTypedAbilityIDsWithStatChanges(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "mod_changes", getTypedAbilityIDsWithModifierChanges(cfg, abilityType))),
	})
}
