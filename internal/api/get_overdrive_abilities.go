package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getOverdriveAbility(r *http.Request, i handlerInput[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource], id int32) (OverdriveAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return OverdriveAbility{}, err
	}

	overdriveIDs, err := srv.Cfg.DB.GetOverdriveAbilityOverdriveIDs(r.Context(), id)
	if err != nil {
		return OverdriveAbility{}, newHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive", err)
	}
	overdrive, _ := lookup.GetResourceByID(overdriveIDs[0], lookup.L.OverdrivesID)

	response := OverdriveAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(srv, E.Abilities, ability.Ability.ID),
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Overdrives:         E.Overdrives.toResources(srv, overdriveIDs),
		OverdriveCommand:   namePtrToNamedAPIResPtr(srv, E.OverdriveCommands, overdrive.OverdriveCommand, nil),
		User:               nameToNamedAPIResource(srv, E.CharacterClasses, overdrive.User, nil),
		RelatedStats:       namesToNamedAPIResources(srv, E.Stats, ability.RelatedStats),
		BattleInteractions: convertObjSlice(srv, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (srv *Server) retrieveOverdriveAbilities(r *http.Request, i handlerInput[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeOverdriveAbility

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.AttackType, ids, ts.QpnAttackType, getTypedAbilityIDsByAttackType(srv, abilityType)),
		i.queryEnumList(srv, r, enums.T.TargetType, ids, ts.QpnTargetType, getTypedAbilityIDsByTargetType(srv, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(srv, abilityType)),
		i.queryIntList(srv, r, ids, ts.QpnRank, getTypedAbilityIDsByRank(srv, abilityType)),
		i.queryNameIDListNul(srv, r, ids, ts.QpnElement, E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(srv, abilityType)),
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, srv.Cfg.DB.GetOverdriveAbilityIDsByCharClass),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, srv.Cfg.DB.GetOverdriveAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(srv, abilityType)),
		i.queryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnCanCrit, getTypedAbilityIDsCanCrit(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnDelay, getTypedAbilityIDsDealsDelay(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnStatChanges, getTypedAbilityIDsWithStatChanges(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnModChanges, getTypedAbilityIDsWithModifierChanges(srv, abilityType)),
	})
}
