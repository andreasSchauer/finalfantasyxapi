package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareSubmenus(test Test, exp Submenu, got apitypes.Submenu) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.checkResIDsInSlice("users", apilookups.E.CharacterClasses.Endpoint, exp.users, got.Users)
	test.checkResIDsInSlice("abilities", apilookups.E.Abilities.Endpoint, exp.abilities, got.Abilities)
	test.compTestStructPtrs("opened by", exp.openedBy, got.OpenedBy, compareMenuOpen)
}

func compareMenuOpen(test Test, fieldName string, exp MenuOpen, got apitypes.MenuOpen) {
	test.t.Helper()
	test.compIDAPIResourcePtrs(fieldName+" - ability", apilookups.E.Abilities.Endpoint, exp.ability, got.Ability)
	test.compIDAPIResourcePtrs(fieldName+" - aeon command", apilookups.E.AeonCommands.Endpoint, exp.aeonCommand, got.AeonCommand)
	test.checkResIDsInSlice(fieldName+" - overdrive commands", apilookups.E.OverdriveCommands.Endpoint, exp.overdriveCommands, got.OverdriveCommands)
}
