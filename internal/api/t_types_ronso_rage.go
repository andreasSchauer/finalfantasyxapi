package api

type expRonsoRage struct {
	testGeneral
	expUnique
	overdrive int32
	monsters  []int32
}

func (e expRonsoRage) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareRonsoRages(test test, exp expRonsoRage, got RonsoRage) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "overdrive", test.cfg.e.overdrives.endpoint, exp.overdrive, got.Overdrive)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monsters, got.Monsters)
}
