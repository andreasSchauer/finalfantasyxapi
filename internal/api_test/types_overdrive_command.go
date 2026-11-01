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
	test.compIDAPIResource("user", test.cfg.E.CharacterClasses.Endpoint, exp.user, got.User)
	test.compIDAPIResourcePtrs("topmenu", test.cfg.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.compIDAPIResource("open submenu", test.cfg.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	test.checkResIDsInSlice("overdrives", test.cfg.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
}
