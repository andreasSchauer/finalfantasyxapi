package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
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

func RetrieveItemAbilities(r *http.Request, i handlerinput.Main[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeItemAbility

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.ItemCategory, ids, ts.QpnCategory, config.Cfg.DB.GetItemAbilityIDsByCategory),
		i.QueryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(abilityType)),
		i.QueryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(abilityType)),
		i.QueryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(abilityType)),
		i.QueryNameIDListNul(r, ids, ts.QpnElement, apilookups.E.Elements.ResTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(abilityType)),
		i.QueryNameID(r, ids, ts.QpnRelatedStat, apilookups.E.Stats.ResTypeSingle, lookup.L.Stats, config.Cfg.DB.GetItemAbilityIDsByRelatedStat),
		i.QueryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(abilityType)),
		i.QueryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(abilityType)),
		i.QueryBool2(r, ids, ts.QpnOutsideBattle, config.Cfg.DB.GetItemAbilityIDsCanUseOutsideBattle),
		i.QueryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(abilityType)),
		i.QueryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(abilityType)),
		i.QueryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(abilityType)),
	})
}
