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

func (e RonsoRage) Compare(test test, got apitypes.RonsoRage) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("overdrive", apilookups.E.Overdrives.Endpoint, e.overdrive, got.Overdrive)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.monsters, got.Monsters)
}
