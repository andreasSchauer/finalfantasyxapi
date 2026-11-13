package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type AeonCommand struct {
	testGeneral
	Unique
	user              int32
	topmenu           *int32
	openSubmenu       *int32
	possibleAbilities []PossibleAbilityList
}

func (e AeonCommand) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAeonCommands(test test, exp AeonCommand, got apitypes.AeonCommand) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("user", apilookups.E.CharacterClasses.Endpoint, exp.user, got.User)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("open submenu", apilookups.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	test.compTestStructSlices("possible abilities", exp.possibleAbilities, got.PossibleAbilities, comparePosAbilityList)
}

type PossibleAbilityList struct {
	user      int32
	abilities []int32
}

func comparePosAbilityList(test test, fieldName string, exp PossibleAbilityList, got apitypes.PossibleAbilityList) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - user", apilookups.E.CharacterClasses.Endpoint, exp.user, got.User)
	test.checkResIDsInSlice(fieldName+" - abilities", apilookups.E.Abilities.Endpoint, exp.abilities, got.Abilities)
}
