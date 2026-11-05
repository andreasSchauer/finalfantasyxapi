package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expModifier struct {
	testGeneral
	expUnique
	autoAbilities      []int32
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	triggerCommands    []int32
	enemyAbilities     []int32
	statusConditions   []int32
	properties         []int32
}

func (e expModifier) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareModifiers(test test, exp expModifier, got api.Modifier) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.checkResIDsInSlice("auto-abilities", api.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("player abilities", api.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice("overdrive abilities", api.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice("item abilities", api.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice("trigger commands", api.E.TriggerCommands.Endpoint, exp.triggerCommands, got.TriggerCommands)
	test.checkResIDsInSlice("enemy abilities", api.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice("status conditions", api.E.StatusConditions.Endpoint, exp.statusConditions, got.StatusConditions)
	test.checkResIDsInSlice("properties", api.E.Properties.Endpoint, exp.properties, got.Properties)
}
