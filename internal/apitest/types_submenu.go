package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

type expSubmenu struct {
	testGeneral
	expUnique
	topmenu   *int32
	users     []int32
	abilities []int32
	openedBy  *expMenuOpen
}

func (e expSubmenu) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSubmenus(test test, exp expSubmenu, got apitypes.Submenu) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResourcePtrs("topmenu", api.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.checkResIDsInSlice("users", api.E.CharacterClasses.Endpoint, exp.users, got.Users)
	test.checkResIDsInSlice("abilities", api.E.Abilities.Endpoint, exp.abilities, got.Abilities)
	test.compTestStructPtrs("opened by", exp.openedBy, got.OpenedBy, compareMenuOpen)
}

type expMenuOpen struct {
	ability           *int32
	aeonCommand       *int32
	overdriveCommands []int32
}

func compareMenuOpen(test test, fieldName string, exp expMenuOpen, got apitypes.MenuOpen) {
	test.t.Helper()
	test.compIDAPIResourcePtrs(fieldName+" - ability", api.E.Abilities.Endpoint, exp.ability, got.Ability)
	test.compIDAPIResourcePtrs(fieldName+" - aeon command", api.E.AeonCommands.Endpoint, exp.aeonCommand, got.AeonCommand)
	test.checkResIDsInSlice(fieldName+" - overdrive commands", api.E.OverdriveCommands.Endpoint, exp.overdriveCommands, got.OverdriveCommands)
}
