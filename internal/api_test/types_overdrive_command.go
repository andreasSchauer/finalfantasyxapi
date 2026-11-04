package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expOverdriveCommand struct {
	testGeneral
	expUnique
	rank        int32
	user        int32
	topmenu     *int32
	openSubmenu int32
	overdrives  []int32
}

func (e expOverdriveCommand) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareOverdriveCommands(test test, exp expOverdriveCommand, got api.OverdriveCommand) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compare("rank", exp.rank, got.Rank)
	test.compIDAPIResource("user", api.E.CharacterClasses.Endpoint, exp.user, got.User)
	test.compIDAPIResourcePtrs("topmenu", api.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.compIDAPIResource("open submenu", api.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	test.checkResIDsInSlice("overdrives", api.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
}
