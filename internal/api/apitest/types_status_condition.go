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

func (e StatusCondition) Compare(test test, got apitypes.StatusCondition) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.autoAbilities, got.AutoAbilities)
	test.compTestStructPtrs("inflicted by", e.inflictedBy, got.InflictedBy)
	test.compTestStructPtrs("removed by", e.removedBy, got.RemovedBy)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.monstersResistance, got.MonstersResistance)
}

type testStatusInfliction struct {
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	miscAbilities      []int32
	enemyAbilities     []int32
}

func (e testStatusInfliction) Compare(test test, fieldName string, got apitypes.StatusInfliction) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - player abilities", apilookups.E.PlayerAbilities.Endpoint, e.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice(fieldName+" - overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, e.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice(fieldName+" - item abilities", apilookups.E.ItemAbilities.Endpoint, e.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice(fieldName+" - misc abilities", apilookups.E.MiscAbilities.Endpoint, e.miscAbilities, got.MiscAbilities)
	test.checkResIDsInSlice(fieldName+" - enemy abilities", apilookups.E.EnemyAbilities.Endpoint, e.enemyAbilities, got.EnemyAbilities)
}

type testStatusRemoval struct {
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	enemyAbilities     []int32
	statusConditions   []int32
}

func (e testStatusRemoval) Compare(test test, fieldName string, got apitypes.StatusRemoval) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - player abilities", apilookups.E.PlayerAbilities.Endpoint, e.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice(fieldName+" - overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, e.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice(fieldName+" - item abilities", apilookups.E.ItemAbilities.Endpoint, e.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice(fieldName+" - enemy abilities", apilookups.E.EnemyAbilities.Endpoint, e.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice(fieldName+" - status conditions", apilookups.E.StatusConditions.Endpoint, e.statusConditions, got.StatusConditions)
}
