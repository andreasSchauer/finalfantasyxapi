package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expAeonCommand struct {
	testGeneral
	expUnique
	user              int32
	topmenu           *int32
	openSubmenu       *int32
	possibleAbilities []expPossibleAbilityList
}

func (e expAeonCommand) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAeonCommands(test test, exp expAeonCommand, got api.AeonCommand) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIDAPIResource(test, "user", test.cfg.E.CharacterClasses.Endpoint, exp.user, got.User)
	compIDAPIResourcePtrs(test, "topmenu", test.cfg.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	compIDAPIResourcePtrs(test, "open submenu", test.cfg.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	compTestStructSlices(test, "possible abilities", exp.possibleAbilities, got.PossibleAbilities, comparePosAbilityList)
}

type expPossibleAbilityList struct {
	user      int32
	abilities []int32
}

func comparePosAbilityList(test test, fieldName string, exp expPossibleAbilityList, got api.PossibleAbilityList) {
	test.t.Helper()
	compIDAPIResource(test, fieldName+" - user", test.cfg.E.CharacterClasses.Endpoint, exp.user, got.User)
	checkResIDsInSlice(test, fieldName+" - abilities", test.cfg.E.Abilities.Endpoint, exp.abilities, got.Abilities)
}
