package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type RonsoRage struct {
	testGeneral
	Unique
	overdrive int32
	monsters  []int32
}

func (e RonsoRage) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareRonsoRages(test test, exp RonsoRage, got apitypes.RonsoRage) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("overdrive", apilookups.E.Overdrives.Endpoint, exp.overdrive, got.Overdrive)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, exp.monsters, got.Monsters)
}
