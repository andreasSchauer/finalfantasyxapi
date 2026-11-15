package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Submenu struct {
	TestGeneral
	Unique
	Topmenu   *int32
	Users     []int32
	Abilities []int32
	OpenedBy  *MenuOpen
}

func (e Submenu) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Submenu) Compare(test Test, got apitypes.Submenu) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.Topmenu, got.Topmenu)
	test.checkResIDsInSlice("users", apilookups.E.CharacterClasses.Endpoint, e.Users, got.Users)
	test.checkResIDsInSlice("abilities", apilookups.E.Abilities.Endpoint, e.Abilities, got.Abilities)
	test.compTestStructPtrs("opened by", e.OpenedBy, got.OpenedBy)
}

type MenuOpen struct {
	Ability           *int32
	AeonCommand       *int32
	OverdriveCommands []int32
}

func (e MenuOpen) Compare(test Test, fieldName string, got apitypes.MenuOpen) {
	test.T.Helper()
	test.compIDAPIResourcePtrs(fieldName+" - ability", apilookups.E.Abilities.Endpoint, e.Ability, got.Ability)
	test.compIDAPIResourcePtrs(fieldName+" - aeon command", apilookups.E.AeonCommands.Endpoint, e.AeonCommand, got.AeonCommand)
	test.checkResIDsInSlice(fieldName+" - overdrive commands", apilookups.E.OverdriveCommands.Endpoint, e.OverdriveCommands, got.OverdriveCommands)
}
