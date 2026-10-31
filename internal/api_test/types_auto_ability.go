package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expAutoAbility struct {
	testGeneral
	expUnique
	monstersDrop     []int32
	monstersItems    []testMonItemAmts
	shopsPreAirship  []int32
	shopsPostAirship []int32
	treasures        []int32
	equipmentTables  []int32
}

func (e expAutoAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAutoAbilities(test test, exp expAutoAbility, got api.AutoAbility) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "monsters drop", test.cfg.E.Monsters.Endpoint, exp.monstersDrop, got.MonstersDrop)
	checkTestStructsInSlice(test, "monsters items", exp.monstersItems, got.MonstersItems, compareMonItemAmts)
	checkResIDsInSlice(test, "treasures", test.cfg.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "shops pre airship", test.cfg.E.Shops.Endpoint, exp.shopsPreAirship, got.ShopsPreAirship)
	checkResIDsInSlice(test, "shops post airship", test.cfg.E.Shops.Endpoint, exp.shopsPostAirship, got.ShopsPostAirship)
	checkResIDsInSlice(test, "equipment tables", test.cfg.E.EquipmentTables.Endpoint, exp.equipmentTables, got.EquipmentTables)
}
