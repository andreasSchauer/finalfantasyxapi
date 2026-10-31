package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getItemAbility(r *http.Request, i handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedAPIResourceList], id int32) (ItemAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return ItemAbility{}, err
	}

	item, _ := seeding.GetResourceByID(id, cfg.L.ItemsID)

	response := ItemAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Item:               idToNamedAPIResource(cfg, cfg.E.Items, id),
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.E.Abilities, ability.Ability.ID),
		Description:        item.Description,
		Effect:             item.Effect,
		Category:           item.Category,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		RelatedStats:       namesToNamedAPIResources(cfg, cfg.E.Stats, item.RelatedStats),
		Cursor:             item.Cursor,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	if item.Usability == string(database.ItemUsabilityOutsideBattle) || item.Usability == string(database.ItemUsabilityAlways) {
		response.CanUseOutsideBattle = true
	}

	return response, nil
}

func (cfg *Config) retrieveItemAbilities(r *http.Request, i handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeItemAbility

	return filterIDs(cfg, r, i, ids, []IDFilter{
		enumListQuery(cfg, r, i, cfg.T.ItemCategory, ids, QpnCategory, cfg.db.GetItemAbilityIDsByCategory),
		enumListQuery(cfg, r, i, cfg.T.AttackType, ids, QpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		enumListQuery(cfg, r, i, cfg.T.TargetType, ids, QpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		enumQuery(r, i, cfg.T.DamageFormula, ids, QpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		nameIDListQueryNul(cfg, r, i, ids, QpnElement, cfg.E.Elements.resTypeSingle, cfg.L.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		nameIDQuery(r, i, ids, QpnRelatedStat, cfg.E.Stats.resTypeSingle, cfg.L.Stats, cfg.db.GetItemAbilityIDsByRelatedStat),
		idQueryNul(r, i, ids, QpnStatusInflict, cfg.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		idQueryNul(r, i, ids, QpnStatusRemove, cfg.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnOutsideBattle, cfg.db.GetItemAbilityIDsCanUseOutsideBattle),
		boolQuery2(r, i, ids, QpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType)),
	})
}
