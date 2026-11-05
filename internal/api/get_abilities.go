package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api_helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAbility(r *http.Request, i handlerInput[seedtypes.Ability, Ability, TypedAPIResource], id int32) (Ability, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Ability{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(r.Context(), ability, config.Cfg.DB.GetAbilityMonsterIDs)
	if err != nil {
		return Ability{}, err
	}

	response := Ability{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Type:               ability.Type,
		Rank:               ability.Rank,
		CanCopycat:         ability.CanCopycat,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		TypedAbility:       refToNamedAPIResource(ability.GetAbilityRef()),
		Monsters:           monsters,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func retrieveAbilities(r *http.Request, i handlerInput[seedtypes.Ability, Ability, TypedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnumList(r, enums.T.AbilityType, ids, ts.QpnType, config.Cfg.DB.GetAbilityIDsByType),
		i.queryEnumList(r, enums.T.DamageType, ids, ts.QpnDamageType, config.Cfg.DB.GetAbilityIDsByDamageType),
		i.queryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, config.Cfg.DB.GetAbilityIDsByAttackType),
		i.queryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, config.Cfg.DB.GetAbilityIDsByTargetType),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, config.Cfg.DB.GetAbilityIDsByDamageFormula),
		i.queryIntList(r, ids, ts.QpnRank, config.Cfg.DB.GetAbilityIDsByRank),
		i.queryNameIDListNul(r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, config.Cfg.DB.GetAbilityIDsByElement),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, config.Cfg.DB.GetAbilityIDsByInflictedStatus),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, config.Cfg.DB.GetAbilityIDsByRemovedStatus),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.MonsterFormations, config.Cfg.DB.GetAbilityIDsByMonster),
		i.queryBool(r, ids, ts.QpnCopycat, config.Cfg.DB.GetAbilityIDsByCanCopycat),
		i.queryBool(r, ids, ts.QpnHelpBar, config.Cfg.DB.GetAbilityIDsByAppearsInHelpBar),
		i.queryBool2(r, ids, ts.QpnCanCrit, config.Cfg.DB.GetAbilityIDsCanCrit),
		i.queryBool2(r, ids, ts.QpnBDL, config.Cfg.DB.GetAbilityIDsBreakDmgLimit),
		i.queryBool2(r, ids, ts.QpnUserAtk, config.Cfg.DB.GetAbilityIDsBasedOnUserAttack),
		i.queryBool2(r, ids, ts.QpnDarkable, config.Cfg.DB.GetAbilityIDsDarkable),
		i.queryBool2(r, ids, ts.QpnSilenceable, config.Cfg.DB.GetAbilityIDsSilenceable),
		i.queryBool2(r, ids, ts.QpnReflectable, config.Cfg.DB.GetAbilityIDsReflectable),
		i.queryBool2(r, ids, ts.QpnDelay, config.Cfg.DB.GetAbilityIDsDealsDelay),
		i.queryBool2(r, ids, ts.QpnStatChanges, config.Cfg.DB.GetAbilityIDsWithStatChanges),
		i.queryBool2(r, ids, ts.QpnModChanges, config.Cfg.DB.GetAbilityIDsWithModifierChanges),
	})
}
