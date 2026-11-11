package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type OverdriveCommand struct {
	testGeneral
	Unique
	rank        int32
	user        int32
	topmenu     *int32
	openSubmenu int32
	overdrives  []int32
}

func (e OverdriveCommand) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareOverdriveCommands(test test, exp OverdriveCommand, got apitypes.OverdriveCommand) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("rank", exp.rank, got.Rank)
	test.compIDAPIResource("user", apilookups.E.CharacterClasses.Endpoint, exp.user, got.User)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.compIDAPIResource("open submenu", apilookups.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
}
