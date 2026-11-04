package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getAbility(r *http.Request, i handlerInput[seedtypes.Ability, Ability, TypedAPIResource], id int32) (Ability, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Ability{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(s, r.Context(), ability, s.Cfg.DB.GetAbilityMonsterIDs)
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
		TypedAbility:       refToNamedAPIResource(s, ability.GetAbilityRef()),
		Monsters:           monsters,
		BattleInteractions: convertObjSlice(s, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (s *Server) retrieveAbilities(r *http.Request, i handlerInput[seedtypes.Ability, Ability, TypedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryEnumList(s, r, enums.T.AbilityType, ids, ts.QpnType, s.Cfg.DB.GetAbilityIDsByType),
		i.queryEnumList(s, r, enums.T.DamageType, ids, ts.QpnDamageType, s.Cfg.DB.GetAbilityIDsByDamageType),
		i.queryEnumList(s, r, enums.T.AttackType, ids, ts.QpnAttackType, s.Cfg.DB.GetAbilityIDsByAttackType),
		i.queryEnumList(s, r, enums.T.TargetType, ids, ts.QpnTargetType, s.Cfg.DB.GetAbilityIDsByTargetType),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, s.Cfg.DB.GetAbilityIDsByDamageFormula),
		i.queryIntList(s, r, ids, ts.QpnRank, s.Cfg.DB.GetAbilityIDsByRank),
		i.queryNameIDListNul(s, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, s.Cfg.DB.GetAbilityIDsByElement),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, s.Cfg.DB.GetAbilityIDsByInflictedStatus),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, s.Cfg.DB.GetAbilityIDsByRemovedStatus),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.MonsterFormations, s.Cfg.DB.GetAbilityIDsByMonster),
		i.queryBool(r, ids, ts.QpnCopycat, s.Cfg.DB.GetAbilityIDsByCanCopycat),
		i.queryBool(r, ids, ts.QpnHelpBar, s.Cfg.DB.GetAbilityIDsByAppearsInHelpBar),
		i.queryBool2(r, ids, ts.QpnCanCrit, s.Cfg.DB.GetAbilityIDsCanCrit),
		i.queryBool2(r, ids, ts.QpnBDL, s.Cfg.DB.GetAbilityIDsBreakDmgLimit),
		i.queryBool2(r, ids, ts.QpnUserAtk, s.Cfg.DB.GetAbilityIDsBasedOnUserAttack),
		i.queryBool2(r, ids, ts.QpnDarkable, s.Cfg.DB.GetAbilityIDsDarkable),
		i.queryBool2(r, ids, ts.QpnSilenceable, s.Cfg.DB.GetAbilityIDsSilenceable),
		i.queryBool2(r, ids, ts.QpnReflectable, s.Cfg.DB.GetAbilityIDsReflectable),
		i.queryBool2(r, ids, ts.QpnDelay, s.Cfg.DB.GetAbilityIDsDealsDelay),
		i.queryBool2(r, ids, ts.QpnStatChanges, s.Cfg.DB.GetAbilityIDsWithStatChanges),
		i.queryBool2(r, ids, ts.QpnModChanges, s.Cfg.DB.GetAbilityIDsWithModifierChanges),
	})
}
