package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
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
	test.checkResIDsInSlice("monsters drop", api.E.Monsters.Endpoint, exp.monstersDrop, got.MonstersDrop)
	test.checkTestStructsInSlice("monsters items", exp.monstersItems, got.MonstersItems, compareMonItemAmts)
	test.checkResIDsInSlice("treasures", api.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	test.checkResIDsInSlice("shops pre airship", api.E.Shops.Endpoint, exp.shopsPreAirship, got.ShopsPreAirship)
	test.checkResIDsInSlice("shops post airship", api.E.Shops.Endpoint, exp.shopsPostAirship, got.ShopsPostAirship)
	test.checkResIDsInSlice("equipment tables", api.E.EquipmentTables.Endpoint, exp.equipmentTables, got.EquipmentTables)
}
