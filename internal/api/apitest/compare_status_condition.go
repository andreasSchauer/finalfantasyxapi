package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareStatusConditions(test Test, exp StatusCondition, got apitypes.StatusCondition) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.compTestStructPtrs("inflicted by", exp.inflictedBy, got.InflictedBy, compareStatusInflictions)
	test.compTestStructPtrs("removed by", exp.removedBy, got.RemovedBy, compareStatusRemovals)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, exp.monstersResistance, got.MonstersResistance)
}

func compareStatusInflictions(test Test, fieldName string, exp TestStatusInfliction, got apitypes.StatusInfliction) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - player abilities", apilookups.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice(fieldName+" - overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice(fieldName+" - item abilities", apilookups.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice(fieldName+" - misc abilities", apilookups.E.MiscAbilities.Endpoint, exp.miscAbilities, got.MiscAbilities)
	test.checkResIDsInSlice(fieldName+" - enemy abilities", apilookups.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
}

func compareStatusRemovals(test Test, fieldName string, exp TestStatusRemoval, got apitypes.StatusRemoval) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - player abilities", apilookups.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice(fieldName+" - overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice(fieldName+" - item abilities", apilookups.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice(fieldName+" - enemy abilities", apilookups.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice(fieldName+" - status conditions", apilookups.E.StatusConditions.Endpoint, exp.statusConditions, got.StatusConditions)
}
