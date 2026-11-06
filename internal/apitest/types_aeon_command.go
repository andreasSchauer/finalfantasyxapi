package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

type expAeonCommand struct {
	testGeneral
	expUnique
	user              int32
	topmenu           *int32
	openSubmenu       *int32
	possibleAbilities []expPossibleAbilityList
}

func (e expAeonCommand) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAeonCommands(test test, exp expAeonCommand, got apitypes.AeonCommand) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("user", api.E.CharacterClasses.Endpoint, exp.user, got.User)
	test.compIDAPIResourcePtrs("topmenu", api.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("open submenu", api.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	test.compTestStructSlices("possible abilities", exp.possibleAbilities, got.PossibleAbilities, comparePosAbilityList)
}

type expPossibleAbilityList struct {
	user      int32
	abilities []int32
}

func comparePosAbilityList(test test, fieldName string, exp expPossibleAbilityList, got apitypes.PossibleAbilityList) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - user", api.E.CharacterClasses.Endpoint, exp.user, got.User)
	test.checkResIDsInSlice(fieldName+" - abilities", api.E.Abilities.Endpoint, exp.abilities, got.Abilities)
}
