package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareElements(test test, exp expElement, got api.Element) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResourcePtrs("status protection", test.cfg.E.StatusConditions.Endpoint, exp.statusProtection, got.StatusProtection)
	test.checkResIDsInSlice("auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("player abilities", test.cfg.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice("overdrive abilities", test.cfg.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice("item abilities", test.cfg.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice("enemy abilities", test.cfg.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice("monsters weak", test.cfg.E.Monsters.Endpoint, exp.monstersWeak, got.MonstersWeak)
	test.checkResIDsInSlice("monsters halved", test.cfg.E.Monsters.Endpoint, exp.monstersHalved, got.MonstersHalved)
	test.checkResIDsInSlice("monsters immune", test.cfg.E.Monsters.Endpoint, exp.monstersImmune, got.MonstersImmune)
	test.checkResIDsInSlice("monsters absorb", test.cfg.E.Monsters.Endpoint, exp.monstersAbsorb, got.MonstersAbsorb)
}
