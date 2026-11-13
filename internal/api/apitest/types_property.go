package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Property struct {
	testGeneral
	Unique
	autoAbilities []int32
	monsters      []int32
}

func (e Property) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareProperties(test test, exp Property, got apitypes.Property) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, exp.monsters, got.Monsters)
}
