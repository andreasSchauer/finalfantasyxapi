package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getAbility(r *http.Request, i handlerInput[seedtypes.Ability, Ability, TypedAPIResource], id int32) (Ability, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Ability{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(srv, r.Context(), ability, srv.Cfg.DB.GetAbilityMonsterIDs)
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
		TypedAbility:       refToNamedAPIResource(srv, ability.GetAbilityRef()),
		Monsters:           monsters,
		BattleInteractions: convertObjSlice(srv, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (srv *Server) retrieveAbilities(r *http.Request, i handlerInput[seedtypes.Ability, Ability, TypedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.AbilityType, ids, ts.QpnType, srv.Cfg.DB.GetAbilityIDsByType),
		i.queryEnumList(srv, r, enums.T.DamageType, ids, ts.QpnDamageType, srv.Cfg.DB.GetAbilityIDsByDamageType),
		i.queryEnumList(srv, r, enums.T.AttackType, ids, ts.QpnAttackType, srv.Cfg.DB.GetAbilityIDsByAttackType),
		i.queryEnumList(srv, r, enums.T.TargetType, ids, ts.QpnTargetType, srv.Cfg.DB.GetAbilityIDsByTargetType),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, srv.Cfg.DB.GetAbilityIDsByDamageFormula),
		i.queryIntList(srv, r, ids, ts.QpnRank, srv.Cfg.DB.GetAbilityIDsByRank),
		i.queryNameIDListNul(srv, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, srv.Cfg.DB.GetAbilityIDsByElement),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, srv.Cfg.DB.GetAbilityIDsByInflictedStatus),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, srv.Cfg.DB.GetAbilityIDsByRemovedStatus),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.MonsterFormations, srv.Cfg.DB.GetAbilityIDsByMonster),
		i.queryBool(r, ids, ts.QpnCopycat, srv.Cfg.DB.GetAbilityIDsByCanCopycat),
		i.queryBool(r, ids, ts.QpnHelpBar, srv.Cfg.DB.GetAbilityIDsByAppearsInHelpBar),
		i.queryBool2(r, ids, ts.QpnCanCrit, srv.Cfg.DB.GetAbilityIDsCanCrit),
		i.queryBool2(r, ids, ts.QpnBDL, srv.Cfg.DB.GetAbilityIDsBreakDmgLimit),
		i.queryBool2(r, ids, ts.QpnUserAtk, srv.Cfg.DB.GetAbilityIDsBasedOnUserAttack),
		i.queryBool2(r, ids, ts.QpnDarkable, srv.Cfg.DB.GetAbilityIDsDarkable),
		i.queryBool2(r, ids, ts.QpnSilenceable, srv.Cfg.DB.GetAbilityIDsSilenceable),
		i.queryBool2(r, ids, ts.QpnReflectable, srv.Cfg.DB.GetAbilityIDsReflectable),
		i.queryBool2(r, ids, ts.QpnDelay, srv.Cfg.DB.GetAbilityIDsDealsDelay),
		i.queryBool2(r, ids, ts.QpnStatChanges, srv.Cfg.DB.GetAbilityIDsWithStatChanges),
		i.queryBool2(r, ids, ts.QpnModChanges, srv.Cfg.DB.GetAbilityIDsWithModifierChanges),
	})
}
