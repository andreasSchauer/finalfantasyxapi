package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Topmenu struct {
	TestGeneral
	Unique
	submenus          []int32
	abilities         []int32
	overdriveCommands []int32
	overdrives        []int32
	aeonCommands      []int32
}

func (e Topmenu) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Topmenu) Compare(test Test, got apitypes.Topmenu) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("submenus", apilookups.E.Submenus.Endpoint, e.submenus, got.Submenus)
	test.checkResIDsInSlice("abilities", apilookups.E.Abilities.Endpoint, e.abilities, got.Abilities)
	test.checkResIDsInSlice("overdrive commands", apilookups.E.OverdriveCommands.Endpoint, e.overdriveCommands, got.OverdriveCommands)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, e.overdrives, got.Overdrives)
	test.checkResIDsInSlice("aeon commands", apilookups.E.AeonCommands.Endpoint, e.aeonCommands, got.AeonCommands)
}
