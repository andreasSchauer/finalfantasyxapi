package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getItemAbility(r *http.Request, i handlerInput[seedtypes.ItemAbility, ItemAbility, NamedAPIResource], id int32) (ItemAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return ItemAbility{}, err
	}

	item, _ := lookup.GetResourceByID(id, lookup.L.ItemsID)

	response := ItemAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Item:               idToNamedAPIResource(s, E.Items, id),
		UntypedAbility:     idToTypedAPIResource(s, E.Abilities, ability.Ability.ID),
		Description:        item.Description,
		Effect:             item.Effect,
		Category:           item.Category,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		RelatedStats:       namesToNamedAPIResources(s, E.Stats, item.RelatedStats),
		Cursor:             item.Cursor,
		BattleInteractions: convertObjSlice(s, ability.BattleInteractions, convertBattleInteraction),
	}

	if item.Usability == string(database.ItemUsabilityOutsideBattle) || item.Usability == string(database.ItemUsabilityAlways) {
		response.CanUseOutsideBattle = true
	}

	return response, nil
}

func (s *Server) retrieveItemAbilities(r *http.Request, i handlerInput[seedtypes.ItemAbility, ItemAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeItemAbility

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryEnumList(s, r, enums.T.ItemCategory, ids, ts.QpnCategory, s.Cfg.DB.GetItemAbilityIDsByCategory),
		i.queryEnumList(s, r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(s, abilityType)),
		i.queryEnumList(s, r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(s, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(s, abilityType)),
		i.queryNameIDListNul(s, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(s, abilityType)),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, s.Cfg.DB.GetItemAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(s, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnOutsideBattle, s.Cfg.DB.GetItemAbilityIDsCanUseOutsideBattle),
		i.queryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(s, abilityType)),
	})
}
