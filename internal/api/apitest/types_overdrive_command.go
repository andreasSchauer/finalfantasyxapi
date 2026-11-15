package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type OverdriveCommand struct {
	TestGeneral
	Unique
	Rank        int32
	User        int32
	Topmenu     *int32
	OpenSubmenu int32
	Overdrives  []int32
}

func (e OverdriveCommand) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e OverdriveCommand) Compare(test Test, got apitypes.OverdriveCommand) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("rank", e.Rank, got.Rank)
	test.compIDAPIResource("user", apilookups.E.CharacterClasses.Endpoint, e.User, got.User)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.Topmenu, got.Topmenu)
	test.compIDAPIResource("open submenu", apilookups.E.Submenus.Endpoint, e.OpenSubmenu, got.OpenSubmenu)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, e.Overdrives, got.Overdrives)
}
