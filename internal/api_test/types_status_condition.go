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
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	compTestStructPtrs(test, "inflicted by", exp.inflictedBy, got.InflictedBy, compareStatusInflictions)
	compTestStructPtrs(test, "removed by", exp.removedBy, got.RemovedBy, compareStatusRemovals)
	checkResIDsInSlice(test, "monsters", test.cfg.E.Monsters.Endpoint, exp.monstersResistance, got.MonstersResistance)
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
	checkResIDsInSlice(test, fieldName+" - player abilities", test.cfg.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, fieldName+" - overdrive abilities", test.cfg.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, fieldName+" - item abilities", test.cfg.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, fieldName+" - misc abilities", test.cfg.E.MiscAbilities.Endpoint, exp.miscAbilities, got.MiscAbilities)
	checkResIDsInSlice(test, fieldName+" - enemy abilities", test.cfg.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
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
	checkResIDsInSlice(test, fieldName+" - player abilities", test.cfg.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, fieldName+" - overdrive abilities", test.cfg.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, fieldName+" - item abilities", test.cfg.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, fieldName+" - enemy abilities", test.cfg.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	checkResIDsInSlice(test, fieldName+" - status conditions", test.cfg.E.StatusConditions.Endpoint, exp.statusConditions, got.StatusConditions)
}
