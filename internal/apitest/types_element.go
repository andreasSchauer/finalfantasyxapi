package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

type expElement struct {
	testGeneral
	expUnique
	statusProtection   *int32
	autoAbilities      []int32
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	enemyAbilities     []int32
	monstersWeak       []int32
	monstersHalved     []int32
	monstersImmune     []int32
	monstersAbsorb     []int32
}

func (e expElement) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareElements(test test, exp expElement, got apitypes.Element) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResourcePtrs("status protection", apilookups.E.StatusConditions.Endpoint, exp.statusProtection, got.StatusProtection)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("player abilities", apilookups.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice("item abilities", apilookups.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice("enemy abilities", apilookups.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice("monsters weak", apilookups.E.Monsters.Endpoint, exp.monstersWeak, got.MonstersWeak)
	test.checkResIDsInSlice("monsters halved", apilookups.E.Monsters.Endpoint, exp.monstersHalved, got.MonstersHalved)
	test.checkResIDsInSlice("monsters immune", apilookups.E.Monsters.Endpoint, exp.monstersImmune, got.MonstersImmune)
	test.checkResIDsInSlice("monsters absorb", apilookups.E.Monsters.Endpoint, exp.monstersAbsorb, got.MonstersAbsorb)
}
