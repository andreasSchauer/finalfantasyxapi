package api

type expAeonCommand struct {
	testGeneral
	expUnique
	user				int32
	topmenu				*int32
	openSubmenu			*int32
	possibleAbilities	[]expPossibleAbilityList
}

func (e expAeonCommand) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAeonCommands(test test, exp expAeonCommand, got AeonCommand) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "user",  test.cfg.e.characterClasses.endpoint, exp.user, got.User)
	compIdApiResourcePtrs(test, "topmenu", test.cfg.e.topmenus.endpoint, exp.topmenu, got.Topmenu)
	compIdApiResourcePtrs(test, "open submenu", test.cfg.e.submenus.endpoint, exp.openSubmenu, got.OpenSubmenu)
	compTestStructSlices(test, "possible abilities", exp.possibleAbilities, got.PossibleAbilities, comparePosAbilityList)
}


type expPossibleAbilityList struct {
	user		int32
	abilities	[]int32
}

func comparePosAbilityList(test test, fieldName string, exp expPossibleAbilityList, got PossibleAbilityList) {
	compIdApiResource(test, fieldName+" - user",  test.cfg.e.characterClasses.endpoint, exp.user, got.User)
	checkResIDsInSlice(test, fieldName+" - abilities", test.cfg.e.abilities.endpoint, exp.abilities, got.Abilities)
}