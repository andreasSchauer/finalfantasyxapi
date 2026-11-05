package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expStatusCondition struct {
	testGeneral
	expUnique
	autoAbilities      []int32
	inflictedBy        *testStatusInfliction
	removedBy          *testStatusRemoval
	monstersResistance []int32
}

func (e expStatusCondition) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareStatusConditions(test test, exp expStatusCondition, got api.StatusCondition) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.checkResIDsInSlice("auto-abilities", api.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.compTestStructPtrs("inflicted by", exp.inflictedBy, got.InflictedBy, compareStatusInflictions)
	test.compTestStructPtrs("removed by", exp.removedBy, got.RemovedBy, compareStatusRemovals)
	test.checkResIDsInSlice("monsters", api.E.Monsters.Endpoint, exp.monstersResistance, got.MonstersResistance)
}

type testStatusInfliction struct {
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	miscAbilities      []int32
	enemyAbilities     []int32
}

func compareStatusInflictions(test test, fieldName string, exp testStatusInfliction, got api.StatusInfliction) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - player abilities", api.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice(fieldName+" - overdrive abilities", api.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice(fieldName+" - item abilities", api.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice(fieldName+" - misc abilities", api.E.MiscAbilities.Endpoint, exp.miscAbilities, got.MiscAbilities)
	test.checkResIDsInSlice(fieldName+" - enemy abilities", api.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
}

type testStatusRemoval struct {
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	enemyAbilities     []int32
	statusConditions   []int32
}

func compareStatusRemovals(test test, fieldName string, exp testStatusRemoval, got api.StatusRemoval) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - player abilities", api.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice(fieldName+" - overdrive abilities", api.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice(fieldName+" - item abilities", api.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice(fieldName+" - enemy abilities", api.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice(fieldName+" - status conditions", api.E.StatusConditions.Endpoint, exp.statusConditions, got.StatusConditions)
}
