package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getEnemyAbility(r *http.Request, i handlerInput[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource], id int32) (EnemyAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return EnemyAbility{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(s, r.Context(), ability, s.Cfg.DB.GetEnemyAbilityMonsterIDs)
	if err != nil {
		return EnemyAbility{}, err
	}

	response := EnemyAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(s, E.Abilities, ability.Ability.ID),
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Monsters:           monsters,
		BattleInteractions: convertObjSlice(s, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (s *Server) retrieveEnemyAbilities(r *http.Request, i handlerInput[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource]) ([]int32, error) {
	abilityType := database.AbilityTypeEnemyAbility
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, s.Cfg.DB.GetEnemyAbilityIDsByMonster),
		i.queryEnumList(s, r, enums.T.DamageType, ids, ts.QpnDamageType, getTypedAbilityIDsByDamageType(s, abilityType)),
		i.queryEnumList(s, r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(s, abilityType)),
		i.queryEnumList(s, r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(s, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(s, abilityType)),
		i.queryIntList(s, r, ids, ts.QpnRank, getTypedAbilityIDsByRank(s, abilityType)),
		i.queryNameIDListNul(s, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(s, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(s, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(s, abilityType)),
		i.queryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnCanCrit, getTypedAbilityIDsCanCrit(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnBDL, getTypedAbilityIDsBreakDmgLimit(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnDarkable, getTypedAbilityIDsDarkable(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnSilenceable, getTypedAbilityIDsSilenceable(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnReflectable, getTypedAbilityIDsReflectable(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(s, abilityType)),
	})
}
