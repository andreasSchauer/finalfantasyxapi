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

func PlayerAbilities(r *http.Request, i handlerinput.Main[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.PlayerAbilityCategory, ids, ts.QpnCategory, config.Cfg.DB.GetPlayerAbilityIDsByCategory),
		i.QueryEnumList(r, enums.T.DamageType, ids, ts.QpnDamageType, getTypedAbilityIDsByDamageType(abilityType)),
		i.QueryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(abilityType)),
		i.QueryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(abilityType)),
		i.QueryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(abilityType)),
		i.QueryIntList(r, ids, ts.QpnMp, config.Cfg.DB.GetPlayerAbilityIDsByMpCost),
		i.QueryInt(r, ids, ts.QpnMpMin, config.Cfg.DB.GetPlayerAbilityIDsByMpCostMin),
		i.QueryInt(r, ids, ts.QpnMpMax, config.Cfg.DB.GetPlayerAbilityIDsByMpCostMax),
		i.QueryIntList(r, ids, ts.QpnRank, getTypedAbilityIDsByRank(abilityType)),
		i.QueryNameIDListNul(r, ids, ts.QpnElement, apilookups.E.Elements.ResTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(abilityType)),
		i.QueryNameID(r, ids, ts.QpnUser, apilookups.E.CharacterClasses.ResTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetPlayerAbilityIDsByCharClass),
		i.QueryID(r, ids, ts.QpnLearnItem, lookup.L.Items, config.Cfg.DB.GetPlayerAbilityIDsByLearnItem),
		i.QueryNameID(r, ids, ts.QpnRelatedStat, apilookups.E.Stats.ResTypeSingle, lookup.L.Stats, config.Cfg.DB.GetPlayerAbilityIDsByRelatedStat),
		i.QueryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(abilityType)),
		i.QueryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(abilityType)),
		i.QueryNameID(r, ids, ts.QpnStdSg, apilookups.E.Characters.ResTypeSingle, lookup.L.Characters, handlerinput.ToIntManyNull(config.Cfg.DB.GetPlayerAbilityIDsStdSgChar)),
		i.QueryNameID(r, ids, ts.QpnExpSg, apilookups.E.Characters.ResTypeSingle, lookup.L.Characters, handlerinput.ToIntManyNull(config.Cfg.DB.GetPlayerAbilityIDsExpSgChar)),
		i.QueryBool(r, ids, ts.QpnOutsideBattle, config.Cfg.DB.GetPlayerAbilityIDsCanUseOutsideBattle),
		i.QueryBool(r, ids, ts.QpnCopycat, getTypedAbilityIDsByCanCopycat(abilityType)),
		i.QueryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(abilityType)),
		i.QueryBool2(r, ids, ts.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDarkable, getTypedAbilityIDsDarkable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnSilenceable, getTypedAbilityIDsSilenceable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnReflectable, getTypedAbilityIDsReflectable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(abilityType)),
		i.QueryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(abilityType)),
		i.QueryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(abilityType)),
	})
}
