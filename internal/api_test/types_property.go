package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expProperty struct {
	testGeneral
	expUnique
	autoAbilities []int32
	monsters      []int32
}

func (e expProperty) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareProperties(test test, exp expProperty, got api.Property) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.checkResIDsInSlice("auto-abilities", api.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("monsters", api.E.Monsters.Endpoint, exp.monsters, got.Monsters)
}
