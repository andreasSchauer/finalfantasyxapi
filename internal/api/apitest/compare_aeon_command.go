package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareAeonCommands(test Test, exp AeonCommand, got apitypes.AeonCommand) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("user", apilookups.E.CharacterClasses.Endpoint, exp.user, got.User)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("open submenu", apilookups.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	test.compTestStructSlices("possible abilities", exp.possibleAbilities, got.PossibleAbilities, comparePosAbilityList)
}

func comparePosAbilityList(test Test, fieldName string, exp PossibleAbilityList, got apitypes.PossibleAbilityList) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - user", apilookups.E.CharacterClasses.Endpoint, exp.user, got.User)
	test.checkResIDsInSlice(fieldName+" - abilities", apilookups.E.Abilities.Endpoint, exp.abilities, got.Abilities)
}
