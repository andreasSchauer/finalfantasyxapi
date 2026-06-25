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

func (cfg *Config) retrieveItemAbilities(r *http.Request, i handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeItemAbility

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.ItemCategory, ids, qpnCategory, cfg.db.GetItemAbilityIDsByCategory)),
		fidl(enumListQuery(cfg, r, i, cfg.t.AttackType, ids, qpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.TargetType, ids, qpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType))),
		fidl(enumQuery(r, i, cfg.t.DamageFormula, ids, qpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		fidl(nameIdListQueryNul(cfg, r, i, ids, qpnElement, cfg.e.elements.resTypeSing, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		fidl(nameIdQuery(r, i, ids, qpnRelatedStat, cfg.e.stats.resTypeSing, cfg.l.Stats, cfg.db.GetItemAbilityIDsByRelatedStat)),
		fidl(idQueryNul(r, i, ids, qpnStatusInflict, cfg.l.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		fidl(idQueryNul(r, i, ids, qpnStatusRemove, cfg.l.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnOutsideBattle, cfg.db.GetItemAbilityIDsCanUseOutsideBattle)),
		fidl(boolQuery2(r, i, ids, qpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType))),
	})
}
