package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Modifier struct {
	TestGeneral
	Unique
	AutoAbilities      []int32
	PlayerAbilities    []int32
	OverdriveAbilities []int32
	ItemAbilities      []int32
	TriggerCommands    []int32
	EnemyAbilities     []int32
	StatusConditions   []int32
	Properties         []int32
}

func (e Modifier) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Modifier) Compare(test Test, got apitypes.Modifier) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.CheckResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.CheckResIDsInSlice("player abilities", apilookups.E.PlayerAbilities.Endpoint, e.PlayerAbilities, got.PlayerAbilities)
	test.CheckResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, e.OverdriveAbilities, got.OverdriveAbilities)
	test.CheckResIDsInSlice("item abilities", apilookups.E.ItemAbilities.Endpoint, e.ItemAbilities, got.ItemAbilities)
	test.CheckResIDsInSlice("trigger commands", apilookups.E.TriggerCommands.Endpoint, e.TriggerCommands, got.TriggerCommands)
	test.CheckResIDsInSlice("enemy abilities", apilookups.E.EnemyAbilities.Endpoint, e.EnemyAbilities, got.EnemyAbilities)
	test.CheckResIDsInSlice("status conditions", apilookups.E.StatusConditions.Endpoint, e.StatusConditions, got.StatusConditions)
	test.CheckResIDsInSlice("properties", apilookups.E.Properties.Endpoint, e.Properties, got.Properties)
}
