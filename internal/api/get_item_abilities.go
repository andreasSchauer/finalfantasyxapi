package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getItemAbility(r *http.Request, i handlerInput[seeding.Item, ItemAbility, NamedAPIResource, NamedApiResourceList], id int32) (ItemAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return ItemAbility{}, err
	}

	item, _ := seeding.GetResourceByID(id, cfg.l.ItemsID)

	category, _ := newNamedAPIResourceFromType(cfg, cfg.e.itemCategory.endpoint, ability.Category, cfg.t.ItemCategory)

	

	response := ItemAbility{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Item: 				   idToNamedAPIResource(cfg, cfg.e.items, id),
		Description:           ability.Description,
		Effect:                ability.Effect,
		Category:              category,
		Rank:                  ability.Rank,
		AppearsInHelpBar:      ability.AppearsInHelpBar,
		CanCopycat:            ability.CanCopycat,
		RelatedStats:          namesToNamedAPIResources(cfg, cfg.e.stats, ability.RelatedStats),
		Cursor:                item.Cursor,
		BattleInteractions:    convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	if item.Usability == "outside-battle" || item.Usability == "always" {
		response.CanUseOutsideBattle = true
	}

	return response, nil
}



func (cfg *Config) retrieveItemAbilities(r *http.Request, i handlerInput[seeding.Item, ItemAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.ItemCategory, resources, "category", cfg.db.GetItemAbilityIDsByCategory)),
		frl(typeQuery(cfg, r, i, cfg.t.AttackType, resources, "attack_type", cfg.db.GetItemAbilityIDsByAttackType)),
		frl(typeQuery(cfg, r, i, cfg.t.DamageFormula, resources, "damage_formula", cfg.db.GetItemAbilityIDsByDamageFormula)),
		frl(nameOrIdQuery(cfg, r, i, resources, "element", cfg.e.elements.resourceType, cfg.l.Elements, cfg.db.GetItemAbilityIDsByElement)),
		frl(nameOrIdQuery(cfg, r, i, resources, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetItemAbilityIDsByRelatedStat)),
		frl(idQueryWrapper(cfg, r, i, resources, "status_inflict", len(cfg.l.StatusConditions), getItemAbilitiesInflictedStatus)),
		frl(idQuery(cfg, r, i, resources, "status_remove", len(cfg.l.StatusConditions), cfg.db.GetItemAbilityIDsByRemovedStatus)),
		frl(boolQuery2(cfg, r, i, resources, "outside_battle", cfg.db.GetItemAbilityIDsCanUseOutsideBattle)),
		frl(boolQuery2(cfg, r, i, resources, "delay", cfg.db.GetItemAbilityIDsDealsDelay)),
		frl(boolQuery2(cfg, r, i, resources, "stat_changes", cfg.db.GetItemAbilityIDsWithStatChanges)),
		frl(boolQuery2(cfg, r, i, resources, "mod_changes", cfg.db.GetItemAbilityIDsWithModifierChanges)),
	})
}
