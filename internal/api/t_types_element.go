package api

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

func compareElements(test test, exp expElement, got Element) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResourcePtrs(test, "status protection", test.cfg.e.statusConditions.endpoint, exp.statusProtection, got.StatusProtection)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "player abilities", test.cfg.e.playerAbilities.endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, "overdrive abilities", test.cfg.e.overdriveAbilities.endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, "item abilities", test.cfg.e.itemAbilities.endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, "enemy abilities", test.cfg.e.enemyAbilities.endpoint, exp.enemyAbilities, got.EnemyAbilities)
	checkResIDsInSlice(test, "monsters weak", test.cfg.e.monsters.endpoint, exp.monstersWeak, got.MonstersWeak)
	checkResIDsInSlice(test, "monsters halved", test.cfg.e.monsters.endpoint, exp.monstersHalved, got.MonstersHalved)
	checkResIDsInSlice(test, "monsters immune", test.cfg.e.monsters.endpoint, exp.monstersImmune, got.MonstersImmune)
	checkResIDsInSlice(test, "monsters absorb", test.cfg.e.monsters.endpoint, exp.monstersAbsorb, got.MonstersAbsorb)
}