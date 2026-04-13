package api

type expAutoAbility struct {
	testGeneral
	expUnique
	monstersDrop 		[]int32
	monstersItems 		[]testMonItemAmts
	shopsPreAirship		[]int32
	shopsPostAirship	[]int32
	treasures			[]int32
	equipmentTables		[]int32
}

func (a expAutoAbility) GetTestGeneral() testGeneral {
	return a.testGeneral
}

func compareAutoAbilities(test test, exp expAutoAbility, got AutoAbility) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "monsters drop", test.cfg.e.monsters.endpoint, exp.monstersDrop, got.MonstersDrop)
	checkTestStructsInSlice(test, "monsters items", exp.monstersItems, got.MonstersItems, compareMonItemAmts)
	checkResIDsInSlice(test, "treasures", test.cfg.e.treasures.endpoint, exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "shops pre airship", test.cfg.e.shops.endpoint, exp.shopsPreAirship, got.ShopsPreAirship)
	checkResIDsInSlice(test, "shops post airship", test.cfg.e.shops.endpoint, exp.shopsPostAirship, got.ShopsPostAirship)
	checkResIDsInSlice(test, "equipment tables", test.cfg.e.equipmentTables.endpoint, exp.equipmentTables, got.EquipmentTables)
}
