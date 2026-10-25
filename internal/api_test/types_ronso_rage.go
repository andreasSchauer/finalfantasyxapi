package apiTest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expRonsoRage struct {
	testGeneral
	expUnique
	overdrive int32
	monsters  []int32
}

func (e expRonsoRage) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareRonsoRages(test test, exp expRonsoRage, got api.RonsoRage) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "overdrive", test.cfg.E.Overdrives.Endpoint, exp.overdrive, got.Overdrive)
	checkResIDsInSlice(test, "monsters", test.cfg.E.Monsters.Endpoint, exp.monsters, got.Monsters)
}
