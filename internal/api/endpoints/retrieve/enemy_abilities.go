package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func EnemyAbilities(r *http.Request, i handlerinput.Main[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.Named]) ([]int32, error) {
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
		i.QueryNameIDListNul(r, ids, ts.QpnElement, apilookups.E.Elements.ResTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(abilityType)),
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
