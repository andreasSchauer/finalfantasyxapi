package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

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

func compareAutoAbilities(test test, exp expAutoAbility, got apitypes.AutoAbility) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.checkResIDsInSlice("monsters drop", apilookups.E.Monsters.Endpoint, exp.monstersDrop, got.MonstersDrop)
	test.checkTestStructsInSlice("monsters items", exp.monstersItems, got.MonstersItems, compareMonItemAmts)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	test.checkResIDsInSlice("shops pre airship", apilookups.E.Shops.Endpoint, exp.shopsPreAirship, got.ShopsPreAirship)
	test.checkResIDsInSlice("shops post airship", apilookups.E.Shops.Endpoint, exp.shopsPostAirship, got.ShopsPostAirship)
	test.checkResIDsInSlice("equipment tables", apilookups.E.EquipmentTables.Endpoint, exp.equipmentTables, got.EquipmentTables)
}
