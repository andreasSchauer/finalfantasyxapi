package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type expRonsoRage struct {
	testGeneral
	expUnique
	overdrive int32
	monsters  []int32
}

func (e expRonsoRage) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareRonsoRages(test test, exp expRonsoRage, got apitypes.RonsoRage) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("overdrive", apilookups.E.Overdrives.Endpoint, exp.overdrive, got.Overdrive)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, exp.monsters, got.Monsters)
}
