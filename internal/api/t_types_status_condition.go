package api

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

func compareStatusConditions(test test, exp expStatusCondition, got StatusCondition) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	compTestStructPtrs(test, "inflicted by", exp.inflictedBy, got.InflictedBy, compareStatusInflictions)
	compTestStructPtrs(test, "removed by", exp.removedBy, got.RemovedBy, compareStatusRemovals)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monstersResistance, got.MonstersResistance)
}

type testStatusInfliction struct {
	playerAbilities      []int32
	overdriveAbilities   []int32
	itemAbilities        []int32
	unspecifiedAbilities []int32
	enemyAbilities       []int32
	statusConditions     []int32
}

func compareStatusInflictions(test test, fieldName string, exp testStatusInfliction, got StatusInfliction) {
	checkResIDsInSlice(test, fieldName+" - player abilities", test.cfg.e.playerAbilities.endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, fieldName+" - overdrive abilities", test.cfg.e.overdriveAbilities.endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, fieldName+" - item abilities", test.cfg.e.itemAbilities.endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, fieldName+" - unspecified abilities", test.cfg.e.unspecifiedAbilities.endpoint, exp.unspecifiedAbilities, got.UnspecifiedAbilities)
	checkResIDsInSlice(test, fieldName+" - enemy abilities", test.cfg.e.enemyAbilities.endpoint, exp.enemyAbilities, got.EnemyAbilities)
	checkResIDsInSlice(test, fieldName+" - status conditions", test.cfg.e.statusConditions.endpoint, exp.statusConditions, got.StatusConditions)
}

type testStatusRemoval struct {
	playerAbilities      []int32
	overdriveAbilities   []int32
	itemAbilities        []int32
	enemyAbilities       []int32
	statusConditions     []int32
}

func compareStatusRemovals(test test, fieldName string, exp testStatusRemoval, got StatusRemoval) {
	checkResIDsInSlice(test, fieldName+" - player abilities", test.cfg.e.playerAbilities.endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, fieldName+" - overdrive abilities", test.cfg.e.overdriveAbilities.endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, fieldName+" - item abilities", test.cfg.e.itemAbilities.endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, fieldName+" - enemy abilities", test.cfg.e.enemyAbilities.endpoint, exp.enemyAbilities, got.EnemyAbilities)
	checkResIDsInSlice(test, fieldName+" - status conditions", test.cfg.e.statusConditions.endpoint, exp.statusConditions, got.StatusConditions)
}