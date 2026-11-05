package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getItemAbility(r *http.Request, i handlerInput[seedtypes.ItemAbility, ItemAbility, NamedAPIResource], id int32) (ItemAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return ItemAbility{}, err
	}

	item, _ := lookup.GetResourceByID(id, lookup.L.ItemsID)

	response := ItemAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Item:               idToNamedAPIResource(srv, E.Items, id),
		UntypedAbility:     idToTypedAPIResource(srv, E.Abilities, ability.Ability.ID),
		Description:        item.Description,
		Effect:             item.Effect,
		Category:           item.Category,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		RelatedStats:       namesToNamedAPIResources(srv, E.Stats, item.RelatedStats),
		Cursor:             item.Cursor,
		BattleInteractions: convertObjSlice(srv, ability.BattleInteractions, convertBattleInteraction),
	}

	if item.Usability == string(database.ItemUsabilityOutsideBattle) || item.Usability == string(database.ItemUsabilityAlways) {
		response.CanUseOutsideBattle = true
	}

	return response, nil
}

func (srv *Server) retrieveItemAbilities(r *http.Request, i handlerInput[seedtypes.ItemAbility, ItemAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeItemAbility

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.ItemCategory, ids, ts.QpnCategory, config.Cfg.DB.GetItemAbilityIDsByCategory),
		i.queryEnumList(srv, r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(srv, abilityType)),
		i.queryEnumList(srv, r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(srv, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(srv, abilityType)),
		i.queryNameIDListNul(srv, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(srv, abilityType)),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, config.Cfg.DB.GetItemAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(srv, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnOutsideBattle, config.Cfg.DB.GetItemAbilityIDsCanUseOutsideBattle),
		i.queryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(srv, abilityType)),
	})
}
