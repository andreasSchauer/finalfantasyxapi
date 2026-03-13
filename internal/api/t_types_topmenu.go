package api

type expTopmenu struct {
	testGeneral
	expUnique
	submenus			[]int32
	abilities			[]int32
	overdriveCommands	[]int32
	overdrives			[]int32
	aeonCommands		[]int32
}

func (e expTopmenu) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareTopmenus(test test, exp expTopmenu, got Topmenu) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "submenus", test.cfg.e.submenus.endpoint, exp.submenus, got.Submenus)
	checkResIDsInSlice(test, "abilities", test.cfg.e.abilities.endpoint, exp.abilities, got.Abilities)
	checkResIDsInSlice(test, "overdrive commands", test.cfg.e.overdriveCommands.endpoint, exp.overdriveCommands, got.OverdriveCommands)
	checkResIDsInSlice(test, "overdrives", test.cfg.e.overdrives.endpoint, exp.overdrives, got.Overdrives)
	checkResIDsInSlice(test, "aeon commands", test.cfg.e.aeonCommands.endpoint, exp.aeonCommands, got.AeonCommands)
}