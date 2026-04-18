package api

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

func compareOverdriveCommands(test test, exp expOverdriveCommand, got OverdriveCommand) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "rank", exp.rank, got.Rank)
	compIdApiResource(test, "user", test.cfg.e.characterClasses.endpoint, exp.user, got.User)
	compIdApiResourcePtrs(test, "topmenu", test.cfg.e.topmenus.endpoint, exp.topmenu, got.Topmenu)
	compIdApiResource(test, "open submenu", test.cfg.e.submenus.endpoint, exp.openSubmenu, got.OpenSubmenu)
	checkResIDsInSlice(test, "overdrives", test.cfg.e.overdrives.endpoint, exp.overdrives, got.Overdrives)
}
