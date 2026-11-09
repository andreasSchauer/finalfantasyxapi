package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAbility(r *http.Request, i handlerinput.Main[seedtypes.Ability, apitypes.Ability, apiresource.Typed], id int32) (apitypes.Ability, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Ability{}, err
	}

	rel, err := GetAbilityRelationships(r, ability)
	if err != nil {
		return apitypes.Ability{}, err
	}

	response := apitypes.Ability{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Type:               ability.Type,
		Rank:               ability.Rank,
		CanCopycat:         ability.CanCopycat,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		TypedAbility:       apiconv.RefToNamedAPIResource(ability.GetAbilityRef()),
		Monsters:           rel.Monsters,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, apiconv.ConvertBattleInteraction),
	}

	return response, nil
}

func retrieveAbilities(r *http.Request, i handlerinput.Main[seedtypes.Ability, apitypes.Ability, apiresource.Typed]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.AbilityType, ids, ts.QpnType, config.Cfg.DB.GetAbilityIDsByType),
		i.QueryEnumList(r, enums.T.DamageType, ids, ts.QpnDamageType, config.Cfg.DB.GetAbilityIDsByDamageType),
		i.QueryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, config.Cfg.DB.GetAbilityIDsByAttackType),
		i.QueryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, config.Cfg.DB.GetAbilityIDsByTargetType),
		i.QueryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, config.Cfg.DB.GetAbilityIDsByDamageFormula),
		i.QueryIntList(r, ids, ts.QpnRank, config.Cfg.DB.GetAbilityIDsByRank),
		i.QueryNameIDListNul(r, ids, ts.QpnElement, apilookups.E.Elements.ResTypeSingle, lookup.L.Elements, config.Cfg.DB.GetAbilityIDsByElement),
		i.QueryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, config.Cfg.DB.GetAbilityIDsByInflictedStatus),
		i.QueryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, config.Cfg.DB.GetAbilityIDsByRemovedStatus),
		i.QueryID(r, ids, ts.QpnMonster, lookup.L.MonsterFormations, config.Cfg.DB.GetAbilityIDsByMonster),
		i.QueryBool(r, ids, ts.QpnCopycat, config.Cfg.DB.GetAbilityIDsByCanCopycat),
		i.QueryBool(r, ids, ts.QpnHelpBar, config.Cfg.DB.GetAbilityIDsByAppearsInHelpBar),
		i.QueryBool2(r, ids, ts.QpnCanCrit, config.Cfg.DB.GetAbilityIDsCanCrit),
		i.QueryBool2(r, ids, ts.QpnBDL, config.Cfg.DB.GetAbilityIDsBreakDmgLimit),
		i.QueryBool2(r, ids, ts.QpnUserAtk, config.Cfg.DB.GetAbilityIDsBasedOnUserAttack),
		i.QueryBool2(r, ids, ts.QpnDarkable, config.Cfg.DB.GetAbilityIDsDarkable),
		i.QueryBool2(r, ids, ts.QpnSilenceable, config.Cfg.DB.GetAbilityIDsSilenceable),
		i.QueryBool2(r, ids, ts.QpnReflectable, config.Cfg.DB.GetAbilityIDsReflectable),
		i.QueryBool2(r, ids, ts.QpnDelay, config.Cfg.DB.GetAbilityIDsDealsDelay),
		i.QueryBool2(r, ids, ts.QpnStatChanges, config.Cfg.DB.GetAbilityIDsWithStatChanges),
		i.QueryBool2(r, ids, ts.QpnModChanges, config.Cfg.DB.GetAbilityIDsWithModifierChanges),
	})
}
