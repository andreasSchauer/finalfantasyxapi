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
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/filters"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetOverdriveAbility(r *http.Request, i handlerinput.Main[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.Named], id int32) (apitypes.OverdriveAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.OverdriveAbility{}, err
	}

	rel, err := assemble.GetOverdriveAbilityRelationships(r, ability)
	if err != nil {
		return apitypes.OverdriveAbility{}, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive", err)
	}

	response := apitypes.OverdriveAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     handlerinput.IDToTypedAPIResource(apilookups.E.Abilities, ability.Ability.ID),
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Overdrives:         rel.Overdrives,
		OverdriveCommand:   rel.OverdriveCommand,
		User:               rel.User,
		RelatedStats:       handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, ability.RelatedStats),
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, apiconv.ConvertBattleInteraction),
	}

	return response, nil
}

func RetrieveOverdriveAbilities(r *http.Request, i handlerinput.Main[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeOverdriveAbility

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, filters.GetTypedAbilityIDsByAttackType(abilityType)),
		i.QueryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, filters.GetTypedAbilityIDsByTargetType(abilityType)),
		i.QueryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, filters.GetTypedAbilityIDsByDamageFormula(abilityType)),
		i.QueryIntList(r, ids, ts.QpnRank, filters.GetTypedAbilityIDsByRank(abilityType)),
		i.QueryNameIDListNul(r, ids, ts.QpnElement, apilookups.E.Elements.ResTypeSingle, lookup.L.Elements, filters.GetTypedAbilityIDsByElement(abilityType)),
		i.QueryNameID(r, ids, ts.QpnUser, apilookups.E.CharacterClasses.ResTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetOverdriveAbilityIDsByCharClass),
		i.QueryNameID(r, ids, ts.QpnRelatedStat, apilookups.E.Stats.ResTypeSingle, lookup.L.Stats, config.Cfg.DB.GetOverdriveAbilityIDsByRelatedStat),
		i.QueryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, filters.GetTypedAbilityIDsByInflictedStatus(abilityType)),
		i.QueryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, filters.GetTypedAbilityIDsByRemovedStatus(abilityType)),
		i.QueryBool2(r, ids, ts.QpnCanCrit, filters.GetTypedAbilityIDsCanCrit(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDelay, filters.GetTypedAbilityIDsDealsDelay(abilityType)),
		i.QueryBool2(r, ids, ts.QpnStatChanges, filters.GetTypedAbilityIDsWithStatChanges(abilityType)),
		i.QueryBool2(r, ids, ts.QpnModChanges, filters.GetTypedAbilityIDsWithModifierChanges(abilityType)),
	})
}
