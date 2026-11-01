package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getItemAbility(r *http.Request, i handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource], id int32) (ItemAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
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

func (cfg *Config) retrieveItemAbilities(r *http.Request, i handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeItemAbility

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, cfg.T.ItemCategory, ids, QpnCategory, cfg.db.GetItemAbilityIDsByCategory),
		i.queryEnumList(cfg, r, cfg.T.AttackType, ids, QpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		i.queryEnumList(cfg, r, cfg.T.TargetType, ids, QpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		i.queryEnum(r, cfg.T.DamageFormula, ids, QpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		i.queryNameIDListNul(cfg, r, ids, QpnElement, cfg.E.Elements.resTypeSingle, cfg.L.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		i.queryNameID(r, ids, QpnRelatedStat, cfg.E.Stats.resTypeSingle, cfg.L.Stats, cfg.db.GetItemAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, QpnStatusInflict, cfg.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		i.queryIDNul(r, ids, QpnStatusRemove, cfg.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		i.queryBool2(r, ids, QpnOutsideBattle, cfg.db.GetItemAbilityIDsCanUseOutsideBattle),
		i.queryBool2(r, ids, QpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
		i.queryBool2(r, ids, QpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType)),
		i.queryBool2(r, ids, QpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType)),
	})
}
