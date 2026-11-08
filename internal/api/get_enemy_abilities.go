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

func getEnemyAbility(r *http.Request, i handlerinput.Main[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.Named], id int32) (apitypes.EnemyAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.EnemyAbility{}, err
	}

	monsters, err := E.Monsters.DBGetRelations(r.Context(), ability, config.Cfg.DB.GetEnemyAbilityMonsterIDs)
	if err != nil {
		return apitypes.EnemyAbility{}, err
	}

	response := apitypes.EnemyAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     handlerinput.IDToTypedAPIResource(E.Abilities, ability.Ability.ID),
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Monsters:           monsters,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func retrieveEnemyAbilities(r *http.Request, i handlerinput.Main[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.Named]) ([]int32, error) {
	abilityType := database.AbilityTypeEnemyAbility
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetEnemyAbilityIDsByMonster),
		i.QueryEnumList(r, enums.T.DamageType, ids, ts.QpnDamageType, getTypedAbilityIDsByDamageType(abilityType)),
		i.QueryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(abilityType)),
		i.QueryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(abilityType)),
		i.QueryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(abilityType)),
		i.QueryIntList(r, ids, ts.QpnRank, getTypedAbilityIDsByRank(abilityType)),
		i.QueryNameIDListNul(r, ids, ts.QpnElement, E.Elements.ResTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(abilityType)),
		i.QueryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(abilityType)),
		i.QueryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(abilityType)),
		i.QueryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(abilityType)),
		i.QueryBool2(r, ids, ts.QpnCanCrit, getTypedAbilityIDsCanCrit(abilityType)),
		i.QueryBool2(r, ids, ts.QpnBDL, getTypedAbilityIDsBreakDmgLimit(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDarkable, getTypedAbilityIDsDarkable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnSilenceable, getTypedAbilityIDsSilenceable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnReflectable, getTypedAbilityIDsReflectable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(abilityType)),
	})
}
