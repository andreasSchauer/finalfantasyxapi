package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getMiscAbility(r *http.Request, i handlerInput[seedtypes.MiscAbility, MiscAbility, NamedAPIResource], id int32) (MiscAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return MiscAbility{}, err
	}

	response := MiscAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(s, E.Abilities, ability.Ability.ID),
		Description:        ability.Description,
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		LearnedBy:          namesToNamedAPIResources(s, E.CharacterClasses, ability.LearnedBy),
		Topmenu:            namePtrToNamedAPIResPtr(s, E.Topmenus, ability.Topmenu, nil),
		Submenu:            namePtrToNamedAPIResPtr(s, E.Submenus, ability.Submenu, nil),
		OpenSubmenu:        namePtrToNamedAPIResPtr(s, E.Submenus, ability.OpenSubmenu, nil),
		Cursor:             ability.Cursor,
		BattleInteractions: convertObjSlice(s, ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(s, r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return MiscAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (s *Server) retrieveMiscAbilities(r *http.Request, i handlerInput[seedtypes.MiscAbility, MiscAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeMiscAbility

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryIntList(s, r, ids, ts.QpnRank, getTypedAbilityIDsByRank(s, abilityType)),
		i.queryNameID(r, ids, ts.QpnUser, E.CharacterClasses.resTypeSingle, lookup.L.CharClasses, s.Cfg.DB.GetMiscAbilityIDsByCharClass),
		i.queryBool(r, ids, ts.QpnCopycat, getTypedAbilityIDsByCanCopycat(s, abilityType)),
		i.queryBool(r, ids, ts.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(s, abilityType)),
		i.queryBool2(r, ids, ts.QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(s, abilityType)),
	})
}
