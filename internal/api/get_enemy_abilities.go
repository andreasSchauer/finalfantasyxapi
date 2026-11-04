package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getEnemyAbility(r *http.Request, i handlerInput[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource], id int32) (EnemyAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return EnemyAbility{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(srv, r.Context(), ability, srv.Cfg.DB.GetEnemyAbilityMonsterIDs)
	if err != nil {
		return EnemyAbility{}, err
	}

	response := EnemyAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(srv, E.Abilities, ability.Ability.ID),
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Monsters:           monsters,
		BattleInteractions: convertObjSlice(srv, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (srv *Server) retrieveEnemyAbilities(r *http.Request, i handlerInput[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource]) ([]int32, error) {
	abilityType := database.AbilityTypeEnemyAbility
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, srv.Cfg.DB.GetEnemyAbilityIDsByMonster),
		i.queryEnumList(srv, r, enums.T.DamageType, ids, ts.QpnDamageType, getTypedAbilityIDsByDamageType(srv, abilityType)),
		i.queryEnumList(srv, r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(srv, abilityType)),
		i.queryEnumList(srv, r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(srv, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(srv, abilityType)),
		i.queryIntList(srv, r, ids, ts.QpnRank, getTypedAbilityIDsByRank(srv, abilityType)),
		i.queryNameIDListNul(srv, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(srv, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(srv, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(srv, abilityType)),
		i.queryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnCanCrit, getTypedAbilityIDsCanCrit(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnBDL, getTypedAbilityIDsBreakDmgLimit(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnDarkable, getTypedAbilityIDsDarkable(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnSilenceable, getTypedAbilityIDsSilenceable(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnReflectable, getTypedAbilityIDsReflectable(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(srv, abilityType)),
	})
}
