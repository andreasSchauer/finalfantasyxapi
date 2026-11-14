package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Modifier struct {
	testGeneral
	Unique
	autoAbilities      []int32
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	triggerCommands    []int32
	enemyAbilities     []int32
	statusConditions   []int32
	properties         []int32
}

func (e Modifier) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Modifier) Compare(test test, got apitypes.Modifier) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("player abilities", apilookups.E.PlayerAbilities.Endpoint, e.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, e.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice("item abilities", apilookups.E.ItemAbilities.Endpoint, e.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice("trigger commands", apilookups.E.TriggerCommands.Endpoint, e.triggerCommands, got.TriggerCommands)
	test.checkResIDsInSlice("enemy abilities", apilookups.E.EnemyAbilities.Endpoint, e.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice("status conditions", apilookups.E.StatusConditions.Endpoint, e.statusConditions, got.StatusConditions)
	test.checkResIDsInSlice("properties", apilookups.E.Properties.Endpoint, e.properties, got.Properties)
}
