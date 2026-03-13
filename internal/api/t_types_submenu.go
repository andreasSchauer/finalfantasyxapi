package api

type expSubmenu struct {
	testGeneral
	expUnique
	topmenu			*int32
	users			[]int32
	abilities		[]int32
	openedBy		*expMenuOpen
}

func (e expSubmenu) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSubmenus(test test, exp expSubmenu, got Submenu) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResourcePtrs(test, "topmenu", test.cfg.e.topmenus.endpoint, exp.topmenu, got.Topmenu)
	checkResIDsInSlice(test, "users", test.cfg.e.characterClasses.endpoint, exp.users, got.Users)
	checkResIDsInSlice(test, "abilities", test.cfg.e.abilities.endpoint, exp.abilities, got.Abilities)
	compTestStructPtrs(test, "opened by", exp.openedBy, got.OpenedBy, compareMenuOpen)
}

type expMenuOpen struct {
	ability				*int32
	aeonCommand			*int32
	overdriveCommands	[]int32
}

func compareMenuOpen(test test, fieldName string, exp expMenuOpen, got MenuOpen) {
	compIdApiResourcePtrs(test, fieldName+" - ability", test.cfg.e.abilities.endpoint, exp.ability, got.Ability)
	compIdApiResourcePtrs(test, fieldName+" - aeon command", test.cfg.e.aeonCommands.endpoint, exp.aeonCommand, got.AeonCommand)
	checkResIDsInSlice(test, fieldName+" - overdrive commands", test.cfg.e.overdriveCommands.endpoint, exp.overdriveCommands, got.OverdriveCommands)
}