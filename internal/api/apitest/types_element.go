package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Element struct {
	testGeneral
	Unique
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

func (e Element) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Element) Compare(test test, got apitypes.Element) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResourcePtrs("status protection", apilookups.E.StatusConditions.Endpoint, e.statusProtection, got.StatusProtection)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("player abilities", apilookups.E.PlayerAbilities.Endpoint, e.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, e.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice("item abilities", apilookups.E.ItemAbilities.Endpoint, e.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice("enemy abilities", apilookups.E.EnemyAbilities.Endpoint, e.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice("monsters weak", apilookups.E.Monsters.Endpoint, e.monstersWeak, got.MonstersWeak)
	test.checkResIDsInSlice("monsters halved", apilookups.E.Monsters.Endpoint, e.monstersHalved, got.MonstersHalved)
	test.checkResIDsInSlice("monsters immune", apilookups.E.Monsters.Endpoint, e.monstersImmune, got.MonstersImmune)
	test.checkResIDsInSlice("monsters absorb", apilookups.E.Monsters.Endpoint, e.monstersAbsorb, got.MonstersAbsorb)
}
