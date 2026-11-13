package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareTopmenus(test Test, exp Topmenu, got apitypes.Topmenu) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("submenus", apilookups.E.Submenus.Endpoint, exp.submenus, got.Submenus)
	test.checkResIDsInSlice("abilities", apilookups.E.Abilities.Endpoint, exp.abilities, got.Abilities)
	test.checkResIDsInSlice("overdrive commands", apilookups.E.OverdriveCommands.Endpoint, exp.overdriveCommands, got.OverdriveCommands)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
	test.checkResIDsInSlice("aeon commands", apilookups.E.AeonCommands.Endpoint, exp.aeonCommands, got.AeonCommands)
}
