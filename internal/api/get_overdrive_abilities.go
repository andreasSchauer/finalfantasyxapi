package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getOverdriveAbility(r *http.Request, i handlerInput[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource], id int32) (OverdriveAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return OverdriveAbility{}, err
	}

	overdriveIDs, err := cfg.db.GetOverdriveAbilityOverdriveIDs(r.Context(), id)
	if err != nil {
		return OverdriveAbility{}, newHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive", err)
	}
	overdrive, _ := lookup.GetResourceByID(overdriveIDs[0], lookup.L.OverdrivesID)

	response := OverdriveAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(cfg, E.Abilities, ability.Ability.ID),
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Overdrives:         E.Overdrives.toResources(cfg, overdriveIDs),
		OverdriveCommand:   namePtrToNamedAPIResPtr(cfg, E.OverdriveCommands, overdrive.OverdriveCommand, nil),
		User:               nameToNamedAPIResource(cfg, E.CharacterClasses, overdrive.User, nil),
		RelatedStats:       namesToNamedAPIResources(cfg, E.Stats, ability.RelatedStats),
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (cfg *Config) retrieveOverdriveAbilities(r *http.Request, i handlerInput[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeOverdriveAbility

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		i.queryEnumList(cfg, r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		i.queryIntList(cfg, r, ids, ts.QpnRank, getTypedAbilityIDsByRank(cfg, abilityType)),
		i.queryNameIDListNul(cfg, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, cfg.db.GetOverdriveAbilityIDsByCharClass),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, cfg.db.GetOverdriveAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		i.queryBool2(r, ids, ts.QpnCanCrit, getTypedAbilityIDsCanCrit(cfg, abilityType)),
		i.queryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
		i.queryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType)),
		i.queryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType)),
	})
}
