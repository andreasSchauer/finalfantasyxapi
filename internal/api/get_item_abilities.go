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
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}
	abilityType := database.AbilityTypeItemAbility

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.ItemCategory, ids, "category", cfg.db.GetItemAbilityIDsByCategory)),
		fidl(enumListQuery(cfg, r, i, cfg.t.AttackType, ids, "attack_type", getTypedAbilityIDsByAttackType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.TargetType, ids, "target_type", getTypedAbilityIDsByTargetType(cfg, abilityType))),
		fidl(enumQuery(r, i, cfg.t.DamageFormula, ids, "damage_formula", getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		fidl(nameIdListQueryNul(cfg, r, i, ids, "element", cfg.e.elements.resourceType, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		fidl(nameIdQuery(r, i, ids, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetItemAbilityIDsByRelatedStat)),
		fidl(idQueryNul(r, i, ids, "status_inflict", cfg.l.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		fidl(idQueryNul(r, i, ids, "status_remove", cfg.l.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "outside_battle", cfg.db.GetItemAbilityIDsCanUseOutsideBattle)),
		fidl(boolQuery2(r, i, ids, "delay", getTypedAbilityIDsDealsDelay(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "stat_changes", getTypedAbilityIDsWithStatChanges(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "mod_changes", getTypedAbilityIDsWithModifierChanges(cfg, abilityType))),
	})
}
