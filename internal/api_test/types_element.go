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
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIDAPIResourcePtrs(test, "status protection", test.cfg.E.StatusConditions.Endpoint, exp.statusProtection, got.StatusProtection)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "player abilities", test.cfg.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, "overdrive abilities", test.cfg.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, "item abilities", test.cfg.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, "enemy abilities", test.cfg.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	checkResIDsInSlice(test, "monsters weak", test.cfg.E.Monsters.Endpoint, exp.monstersWeak, got.MonstersWeak)
	checkResIDsInSlice(test, "monsters halved", test.cfg.E.Monsters.Endpoint, exp.monstersHalved, got.MonstersHalved)
	checkResIDsInSlice(test, "monsters immune", test.cfg.E.Monsters.Endpoint, exp.monstersImmune, got.MonstersImmune)
	checkResIDsInSlice(test, "monsters absorb", test.cfg.E.Monsters.Endpoint, exp.monstersAbsorb, got.MonstersAbsorb)
}
