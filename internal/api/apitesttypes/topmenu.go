package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Topmenu struct {
	TestGeneral
	Unique
	Submenus          []int32
	Abilities         []int32
	OverdriveCommands []int32
	Overdrives        []int32
	AeonCommands      []int32
}

func (e Topmenu) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Topmenu) Compare(test Test, got apitypes.Topmenu) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.CheckResIDsInSlice("submenus", apilookups.E.Submenus.Endpoint, e.Submenus, got.Submenus)
	test.CheckResIDsInSlice("abilities", apilookups.E.Abilities.Endpoint, e.Abilities, got.Abilities)
	test.CheckResIDsInSlice("overdrive commands", apilookups.E.OverdriveCommands.Endpoint, e.OverdriveCommands, got.OverdriveCommands)
	test.CheckResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, e.Overdrives, got.Overdrives)
	test.CheckResIDsInSlice("aeon commands", apilookups.E.AeonCommands.Endpoint, e.AeonCommands, got.AeonCommands)
}
