package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/filters"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetEnemyAbility(r *http.Request, i handlerinput.Main[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.Named], id int32) (apitypes.EnemyAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.EnemyAbility{}, err
	}

	rel, err := assemble.GetEnemyAbilityRelationships(r, ability)
	if err != nil {
		return apitypes.EnemyAbility{}, err
	}

	response := apitypes.EnemyAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     handlerinput.IDToTypedAPIResource(apilookups.E.Abilities, ability.Ability.ID),
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Monsters:           rel.Monsters,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, apiconv.ConvertBattleInteraction),
	}

	return response, nil
}

func RetrieveEnemyAbilities(r *http.Request, i handlerinput.Main[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.Named]) ([]int32, error) {
	abilityType := database.AbilityTypeEnemyAbility
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetEnemyAbilityIDsByMonster),
		i.QueryEnumList(r, enums.T.DamageType, ids, ts.QpnDamageType, filters.GetTypedAbilityIDsByDamageType(abilityType)),
		i.QueryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, filters.GetTypedAbilityIDsByAttackType(abilityType)),
		i.QueryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, filters.GetTypedAbilityIDsByTargetType(abilityType)),
		i.QueryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, filters.GetTypedAbilityIDsByDamageFormula(abilityType)),
		i.QueryIntList(r, ids, ts.QpnRank, filters.GetTypedAbilityIDsByRank(abilityType)),
		i.QueryNameIDListNul(r, ids, ts.QpnElement, apilookups.E.Elements.ResTypeSingle, lookup.L.Elements, filters.GetTypedAbilityIDsByElement(abilityType)),
		i.QueryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, filters.GetTypedAbilityIDsByInflictedStatus(abilityType)),
		i.QueryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, filters.GetTypedAbilityIDsByRemovedStatus(abilityType)),
		i.QueryBool(r, ids, ts.QpnHelpBar, filters.GetTypedAbilityIDsByAppearsInHelpBar(abilityType)),
		i.QueryBool2(r, ids, ts.QpnCanCrit, filters.GetTypedAbilityIDsCanCrit(abilityType)),
		i.QueryBool2(r, ids, ts.QpnBDL, filters.GetTypedAbilityIDsBreakDmgLimit(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDarkable, filters.GetTypedAbilityIDsDarkable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnSilenceable, filters.GetTypedAbilityIDsSilenceable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnReflectable, filters.GetTypedAbilityIDsReflectable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDelay, filters.GetTypedAbilityIDsDealsDelay(abilityType)),
	})
}
