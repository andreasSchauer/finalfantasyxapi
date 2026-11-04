package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getOverdriveAbility(r *http.Request, i handlerInput[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource], id int32) (OverdriveAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return OverdriveAbility{}, err
	}

	overdriveIDs, err := s.Cfg.DB.GetOverdriveAbilityOverdriveIDs(r.Context(), id)
	if err != nil {
		return OverdriveAbility{}, newHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive", err)
	}
	overdrive, _ := lookup.GetResourceByID(overdriveIDs[0], lookup.L.OverdrivesID)

	response := OverdriveAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(s, E.Abilities, ability.Ability.ID),
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Overdrives:         E.Overdrives.toResources(s, overdriveIDs),
		OverdriveCommand:   namePtrToNamedAPIResPtr(s, E.OverdriveCommands, overdrive.OverdriveCommand, nil),
		User:               nameToNamedAPIResource(s, E.CharacterClasses, overdrive.User, nil),
		RelatedStats:       namesToNamedAPIResources(s, E.Stats, ability.RelatedStats),
		BattleInteractions: convertObjSlice(s, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (s *Server) retrieveOverdriveAbilities(r *http.Request, i handlerInput[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeOverdriveAbility

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryEnumList(s, r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(s, abilityType)),
		i.queryEnumList(s, r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(s, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(s, abilityType)),
		i.queryIntList(s, r, ids, ts.QpnRank, getTypedAbilityIDsByRank(s, abilityType)),
		i.queryNameIDListNul(s, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(s, abilityType)),
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, s.Cfg.DB.GetOverdriveAbilityIDsByCharClass),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, s.Cfg.DB.GetOverdriveAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(s, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnCanCrit, getTypedAbilityIDsCanCrit(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(s, abilityType)),
	})
}
