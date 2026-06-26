package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getItemAbility(r *http.Request, i handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList], id int32) (ItemAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
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

	if item.Usability == string(database.ItemUsabilityOutsideBattle) || item.Usability == string(database.ItemUsabilityAlways) {
		response.CanUseOutsideBattle = true
	}

	return response, nil
}

func (cfg *Config) retrieveItemAbilities(r *http.Request, i handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeItemAbility

	return filterIDs(cfg, r, i, ids, []IdFilter{
		enumListQuery(cfg, r, i, cfg.t.ItemCategory, ids, qpnCategory, cfg.db.GetItemAbilityIDsByCategory),
		enumListQuery(cfg, r, i, cfg.t.AttackType, ids, qpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		enumListQuery(cfg, r, i, cfg.t.TargetType, ids, qpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		enumQuery(r, i, cfg.t.DamageFormula, ids, qpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		nameIdListQueryNul(cfg, r, i, ids, qpnElement, cfg.e.elements.resTypeSing, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		nameIdQuery(r, i, ids, qpnRelatedStat, cfg.e.stats.resTypeSing, cfg.l.Stats, cfg.db.GetItemAbilityIDsByRelatedStat),
		idQueryNul(r, i, ids, qpnStatusInflict, cfg.l.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		idQueryNul(r, i, ids, qpnStatusRemove, cfg.l.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		boolQuery2(r, i, ids, qpnOutsideBattle, cfg.db.GetItemAbilityIDsCanUseOutsideBattle),
		boolQuery2(r, i, ids, qpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
		boolQuery2(r, i, ids, qpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType)),
		boolQuery2(r, i, ids, qpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType)),
	})
}
