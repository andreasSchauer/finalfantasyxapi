package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type StatusCondition struct {
	testGeneral
	Unique
	autoAbilities      []int32
	inflictedBy        *testStatusInfliction
	removedBy          *testStatusRemoval
	monstersResistance []int32
}

func (e StatusCondition) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareStatusConditions(test test, exp StatusCondition, got apitypes.StatusCondition) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.compTestStructPtrs("inflicted by", exp.inflictedBy, got.InflictedBy, compareStatusInflictions)
	test.compTestStructPtrs("removed by", exp.removedBy, got.RemovedBy, compareStatusRemovals)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, exp.monstersResistance, got.MonstersResistance)
}

type testStatusInfliction struct {
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	miscAbilities      []int32
	enemyAbilities     []int32
}

func compareStatusInflictions(test test, fieldName string, exp testStatusInfliction, got apitypes.StatusInfliction) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - player abilities", apilookups.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice(fieldName+" - overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice(fieldName+" - item abilities", apilookups.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice(fieldName+" - misc abilities", apilookups.E.MiscAbilities.Endpoint, exp.miscAbilities, got.MiscAbilities)
	test.checkResIDsInSlice(fieldName+" - enemy abilities", apilookups.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
}

type testStatusRemoval struct {
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	enemyAbilities     []int32
	statusConditions   []int32
}

func compareStatusRemovals(test test, fieldName string, exp testStatusRemoval, got apitypes.StatusRemoval) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - player abilities", apilookups.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice(fieldName+" - overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice(fieldName+" - item abilities", apilookups.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice(fieldName+" - enemy abilities", apilookups.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice(fieldName+" - status conditions", apilookups.E.StatusConditions.Endpoint, exp.statusConditions, got.StatusConditions)
}
