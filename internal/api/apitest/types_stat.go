package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Stat struct {
	TestGeneral
	Unique
	Spheres            []int32
	AutoAbilities      []int32
	PlayerAbilities    []int32
	OverdriveAbilities []int32
	ItemAbilities      []int32
	TriggerCommands    []int32
	StatusConditions   []int32
	Properties         []int32
}

func (e Stat) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Stat) Compare(test Test, got apitypes.Stat) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("spheres", apilookups.E.Spheres.Endpoint, e.Spheres, got.Spheres)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("player abilities", apilookups.E.PlayerAbilities.Endpoint, e.PlayerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, e.OverdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice("item abilities", apilookups.E.ItemAbilities.Endpoint, e.ItemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice("trigger commands", apilookups.E.TriggerCommands.Endpoint, e.TriggerCommands, got.TriggerCommands)
	test.checkResIDsInSlice("status conditions", apilookups.E.StatusConditions.Endpoint, e.StatusConditions, got.StatusConditions)
	test.checkResIDsInSlice("properties", apilookups.E.Properties.Endpoint, e.Properties, got.Properties)
}
