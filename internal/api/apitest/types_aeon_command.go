package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type AeonCommand struct {
	testGeneral
	Unique
	user              int32
	topmenu           *int32
	openSubmenu       *int32
	possibleAbilities []PossibleAbilityList
}

func (e AeonCommand) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e AeonCommand) Compare(test test, got apitypes.AeonCommand) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("user", apilookups.E.CharacterClasses.Endpoint, e.user, got.User)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("open submenu", apilookups.E.Submenus.Endpoint, e.openSubmenu, got.OpenSubmenu)
	test.compTestStructSlices("possible abilities", e.possibleAbilities, got.PossibleAbilities)
}

type PossibleAbilityList struct {
	user      int32
	abilities []int32
}

func (e PossibleAbilityList) Compare(test test, fieldName string, got apitypes.PossibleAbilityList) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - user", apilookups.E.CharacterClasses.Endpoint, e.user, got.User)
	test.checkResIDsInSlice(fieldName+" - abilities", apilookups.E.Abilities.Endpoint, e.abilities, got.Abilities)
}
