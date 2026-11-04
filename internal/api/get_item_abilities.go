package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getItemAbility(r *http.Request, i handlerInput[seedtypes.ItemAbility, ItemAbility, NamedAPIResource], id int32) (ItemAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return ItemAbility{}, err
	}

	item, _ := lookup.GetResourceByID(id, lookup.L.ItemsID)

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

func (cfg *Config) retrieveItemAbilities(r *http.Request, i handlerInput[seedtypes.ItemAbility, ItemAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeItemAbility

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, enums.T.ItemCategory, ids, typestrings.QpnCategory, cfg.db.GetItemAbilityIDsByCategory),
		i.queryEnumList(cfg, r, enums.T.AttackType, ids, typestrings.QpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		i.queryEnumList(cfg, r, enums.T.TargetType, ids, typestrings.QpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, typestrings.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		i.queryNameIDListNul(cfg, r, ids, typestrings.QpnElement, cfg.E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		i.queryNameID(r, ids, typestrings.QpnRelatedStat, cfg.E.Stats.resTypeSingle, lookup.L.Stats, cfg.db.GetItemAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, typestrings.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		i.queryIDNul(r, ids, typestrings.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnOutsideBattle, cfg.db.GetItemAbilityIDsCanUseOutsideBattle),
		i.queryBool2(r, ids, typestrings.QpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType)),
	})
}
