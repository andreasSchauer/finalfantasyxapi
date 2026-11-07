package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getItemAbility(r *http.Request, i handlerinput.HandlerInput[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.NamedAPIResource], id int32) (apitypes.ItemAbility, error) {
	ability, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.ItemAbility{}, err
	}

	item, _ := lookup.GetResourceByID(id, lookup.L.ItemsID)

	response := apitypes.ItemAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Item:               handlerinput.IDToNamedAPIResource(E.Items, id),
		UntypedAbility:     handlerinput.IDToTypedAPIResource(E.Abilities, ability.Ability.ID),
		Description:        item.Description,
		Effect:             item.Effect,
		Category:           item.Category,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		RelatedStats:       handlerinput.NamesToNamedAPIResources(E.Stats, item.RelatedStats),
		Cursor:             item.Cursor,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteraction),
	}

	if item.Usability == string(database.ItemUsabilityOutsideBattle) || item.Usability == string(database.ItemUsabilityAlways) {
		response.CanUseOutsideBattle = true
	}

	return response, nil
}

func retrieveItemAbilities(r *http.Request, i handlerinput.HandlerInput[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeItemAbility

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.ItemCategory, ids, ts.QpnCategory, config.Cfg.DB.GetItemAbilityIDsByCategory),
		i.QueryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(abilityType)),
		i.QueryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(abilityType)),
		i.QueryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(abilityType)),
		i.QueryNameIDListNul(r, ids, ts.QpnElement, E.Elements.ResTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(abilityType)),
		i.QueryNameID(r, ids, ts.QpnRelatedStat, E.Stats.ResTypeSingle, lookup.L.Stats, config.Cfg.DB.GetItemAbilityIDsByRelatedStat),
		i.QueryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(abilityType)),
		i.QueryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(abilityType)),
		i.QueryBool2(r, ids, ts.QpnOutsideBattle, config.Cfg.DB.GetItemAbilityIDsCanUseOutsideBattle),
		i.QueryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(abilityType)),
		i.QueryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(abilityType)),
		i.QueryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(abilityType)),
	})
}
