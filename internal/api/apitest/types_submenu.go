package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Submenu struct {
	testGeneral
	Unique
	topmenu   *int32
	users     []int32
	abilities []int32
	openedBy  *MenuOpen
}

func (e Submenu) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Submenu) Compare(test test, got apitypes.Submenu) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.topmenu, got.Topmenu)
	test.checkResIDsInSlice("users", apilookups.E.CharacterClasses.Endpoint, e.users, got.Users)
	test.checkResIDsInSlice("abilities", apilookups.E.Abilities.Endpoint, e.abilities, got.Abilities)
	test.compTestStructPtrs("opened by", e.openedBy, got.OpenedBy)
}

type MenuOpen struct {
	ability           *int32
	aeonCommand       *int32
	overdriveCommands []int32
}

func (e MenuOpen) Compare(test test, fieldName string, got apitypes.MenuOpen) {
	test.t.Helper()
	test.compIDAPIResourcePtrs(fieldName+" - ability", apilookups.E.Abilities.Endpoint, e.ability, got.Ability)
	test.compIDAPIResourcePtrs(fieldName+" - aeon command", apilookups.E.AeonCommands.Endpoint, e.aeonCommand, got.AeonCommand)
	test.checkResIDsInSlice(fieldName+" - overdrive commands", apilookups.E.OverdriveCommands.Endpoint, e.overdriveCommands, got.OverdriveCommands)
}
