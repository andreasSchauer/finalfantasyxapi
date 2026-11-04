package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareTopmenus(test test, exp expTopmenu, got api.Topmenu) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.checkResIDsInSlice("submenus", api.E.Submenus.Endpoint, exp.submenus, got.Submenus)
	test.checkResIDsInSlice("abilities", api.E.Abilities.Endpoint, exp.abilities, got.Abilities)
	test.checkResIDsInSlice("overdrive commands", api.E.OverdriveCommands.Endpoint, exp.overdriveCommands, got.OverdriveCommands)
	test.checkResIDsInSlice("overdrives", api.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
	test.checkResIDsInSlice("aeon commands", api.E.AeonCommands.Endpoint, exp.aeonCommands, got.AeonCommands)
}
