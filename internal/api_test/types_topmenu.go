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
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "submenus", test.cfg.E.Submenus.Endpoint, exp.submenus, got.Submenus)
	checkResIDsInSlice(test, "abilities", test.cfg.E.Abilities.Endpoint, exp.abilities, got.Abilities)
	checkResIDsInSlice(test, "overdrive commands", test.cfg.E.OverdriveCommands.Endpoint, exp.overdriveCommands, got.OverdriveCommands)
	checkResIDsInSlice(test, "overdrives", test.cfg.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
	checkResIDsInSlice(test, "aeon commands", test.cfg.E.AeonCommands.Endpoint, exp.aeonCommands, got.AeonCommands)
}
