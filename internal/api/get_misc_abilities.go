package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getMiscAbility(r *http.Request, i handlerInput[seedtypes.MiscAbility, MiscAbility, NamedAPIResource], id int32) (MiscAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return MiscAbility{}, err
	}

	response := MiscAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(srv, E.Abilities, ability.Ability.ID),
		Description:        ability.Description,
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		LearnedBy:          namesToNamedAPIResources(srv, E.CharacterClasses, ability.LearnedBy),
		Topmenu:            namePtrToNamedAPIResPtr(srv, E.Topmenus, ability.Topmenu, nil),
		Submenu:            namePtrToNamedAPIResPtr(srv, E.Submenus, ability.Submenu, nil),
		OpenSubmenu:        namePtrToNamedAPIResPtr(srv, E.Submenus, ability.OpenSubmenu, nil),
		Cursor:             ability.Cursor,
		BattleInteractions: convertObjSlice(srv, ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(srv, r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return MiscAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (srv *Server) retrieveMiscAbilities(r *http.Request, i handlerInput[seedtypes.MiscAbility, MiscAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeMiscAbility

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryIntList(srv, r, ids, ts.QpnRank, getTypedAbilityIDsByRank(srv, abilityType)),
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, srv.Cfg.DB.GetMiscAbilityIDsByCharClass),
		i.queryBool(r, ids, ts.QpnCopycat, getTypedAbilityIDsByCanCopycat(srv, abilityType)),
		i.queryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(srv, abilityType)),
		i.queryBool2(r, ids, ts.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(srv, abilityType)),
	})
}
