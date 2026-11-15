package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type StatusCondition struct {
	TestGeneral
	Unique
	AutoAbilities      []int32
	InflictedBy        *TestStatusInfliction
	RemovedBy          *TestStatusRemoval
	MonstersResistance []int32
}

func (e StatusCondition) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e StatusCondition) Compare(test Test, got apitypes.StatusCondition) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.CheckResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.compTestStructPtrs("inflicted by", e.InflictedBy, got.InflictedBy)
	test.compTestStructPtrs("removed by", e.RemovedBy, got.RemovedBy)
	test.CheckResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.MonstersResistance, got.MonstersResistance)
}

type TestStatusInfliction struct {
	PlayerAbilities    []int32
	OverdriveAbilities []int32
	ItemAbilities      []int32
	MiscAbilities      []int32
	EnemyAbilities     []int32
}

func (e TestStatusInfliction) Compare(test Test, fn string, got apitypes.StatusInfliction) {
	test.T.Helper()
	test.CheckResIDsInSlice(childFn(fn, "player abilities"), apilookups.E.PlayerAbilities.Endpoint, e.PlayerAbilities, got.PlayerAbilities)
	test.CheckResIDsInSlice(childFn(fn, "overdrive abilities"), apilookups.E.OverdriveAbilities.Endpoint, e.OverdriveAbilities, got.OverdriveAbilities)
	test.CheckResIDsInSlice(childFn(fn, "item abilities"), apilookups.E.ItemAbilities.Endpoint, e.ItemAbilities, got.ItemAbilities)
	test.CheckResIDsInSlice(childFn(fn, "misc abilities"), apilookups.E.MiscAbilities.Endpoint, e.MiscAbilities, got.MiscAbilities)
	test.CheckResIDsInSlice(childFn(fn, "enemy abilities"), apilookups.E.EnemyAbilities.Endpoint, e.EnemyAbilities, got.EnemyAbilities)
}

type TestStatusRemoval struct {
	PlayerAbilities    []int32
	OverdriveAbilities []int32
	ItemAbilities      []int32
	EnemyAbilities     []int32
	StatusConditions   []int32
}

func (e TestStatusRemoval) Compare(test Test, fn string, got apitypes.StatusRemoval) {
	test.T.Helper()
	test.CheckResIDsInSlice(childFn(fn, "player abilities"), apilookups.E.PlayerAbilities.Endpoint, e.PlayerAbilities, got.PlayerAbilities)
	test.CheckResIDsInSlice(childFn(fn, "overdrive abilities"), apilookups.E.OverdriveAbilities.Endpoint, e.OverdriveAbilities, got.OverdriveAbilities)
	test.CheckResIDsInSlice(childFn(fn, "item abilities"), apilookups.E.ItemAbilities.Endpoint, e.ItemAbilities, got.ItemAbilities)
	test.CheckResIDsInSlice(childFn(fn, "enemy abilities"), apilookups.E.EnemyAbilities.Endpoint, e.EnemyAbilities, got.EnemyAbilities)
	test.CheckResIDsInSlice(childFn(fn, "status conditions"), apilookups.E.StatusConditions.Endpoint, e.StatusConditions, got.StatusConditions)
}
