package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type expTopmenu struct {
	testGeneral
	expUnique
	submenus          []int32
	abilities         []int32
	overdriveCommands []int32
	overdrives        []int32
	aeonCommands      []int32
}

func (e expTopmenu) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareTopmenus(test test, exp expTopmenu, got apitypes.Topmenu) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.checkResIDsInSlice("submenus", apilookups.E.Submenus.Endpoint, exp.submenus, got.Submenus)
	test.checkResIDsInSlice("abilities", apilookups.E.Abilities.Endpoint, exp.abilities, got.Abilities)
	test.checkResIDsInSlice("overdrive commands", apilookups.E.OverdriveCommands.Endpoint, exp.overdriveCommands, got.OverdriveCommands)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
	test.checkResIDsInSlice("aeon commands", apilookups.E.AeonCommands.Endpoint, exp.aeonCommands, got.AeonCommands)
}
