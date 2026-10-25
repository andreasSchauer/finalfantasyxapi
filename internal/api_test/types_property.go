package apiTest

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
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "monsters", test.cfg.E.Monsters.Endpoint, exp.monsters, got.Monsters)
}
