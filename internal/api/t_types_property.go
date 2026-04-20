package api

type expProperty struct {
	testGeneral
	expUnique
	autoAbilities []int32
	monsters      []int32
}

func (e expProperty) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareProperties(test test, exp expProperty, got Property) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monsters, got.Monsters)
}