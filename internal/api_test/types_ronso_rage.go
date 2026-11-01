package apitest

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
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("overdrive", test.cfg.E.Overdrives.Endpoint, exp.overdrive, got.Overdrive)
	test.checkResIDsInSlice("monsters", test.cfg.E.Monsters.Endpoint, exp.monsters, got.Monsters)
}
