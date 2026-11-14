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

func (e OverdriveCommand) Compare(test test, got apitypes.OverdriveCommand) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("rank", e.rank, got.Rank)
	test.compIDAPIResource("user", apilookups.E.CharacterClasses.Endpoint, e.user, got.User)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.topmenu, got.Topmenu)
	test.compIDAPIResource("open submenu", apilookups.E.Submenus.Endpoint, e.openSubmenu, got.OpenSubmenu)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, e.overdrives, got.Overdrives)
}
