package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type AeonCommand struct {
	TestGeneral
	Unique
	User              int32
	Topmenu           *int32
	OpenSubmenu       *int32
	PossibleAbilities []PossibleAbilityList
}

func (e AeonCommand) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e AeonCommand) Compare(test Test, got apitypes.AeonCommand) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("user", apilookups.E.CharacterClasses.Endpoint, e.User, got.User)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.Topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("open submenu", apilookups.E.Submenus.Endpoint, e.OpenSubmenu, got.OpenSubmenu)
	test.compTestStructSlices("possible abilities", e.PossibleAbilities, got.PossibleAbilities)
}

type PossibleAbilityList struct {
	User      int32
	Abilities []int32
}

func (e PossibleAbilityList) Compare(test Test, fieldName string, got apitypes.PossibleAbilityList) {
	test.T.Helper()
	test.compIDAPIResource(fieldName+" - user", apilookups.E.CharacterClasses.Endpoint, e.User, got.User)
	test.checkResIDsInSlice(fieldName+" - abilities", apilookups.E.Abilities.Endpoint, e.Abilities, got.Abilities)
}
