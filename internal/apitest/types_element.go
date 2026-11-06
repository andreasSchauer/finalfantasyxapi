package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
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
	test.compIDAPIResourcePtrs("status protection", api.E.StatusConditions.Endpoint, exp.statusProtection, got.StatusProtection)
	test.checkResIDsInSlice("auto-abilities", api.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("player abilities", api.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice("overdrive abilities", api.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice("item abilities", api.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice("enemy abilities", api.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice("monsters weak", api.E.Monsters.Endpoint, exp.monstersWeak, got.MonstersWeak)
	test.checkResIDsInSlice("monsters halved", api.E.Monsters.Endpoint, exp.monstersHalved, got.MonstersHalved)
	test.checkResIDsInSlice("monsters immune", api.E.Monsters.Endpoint, exp.monstersImmune, got.MonstersImmune)
	test.checkResIDsInSlice("monsters absorb", api.E.Monsters.Endpoint, exp.monstersAbsorb, got.MonstersAbsorb)
}
