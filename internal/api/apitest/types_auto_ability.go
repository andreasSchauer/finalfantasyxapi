package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type AutoAbility struct {
	testGeneral
	Unique
	monstersDrop     []int32
	monstersItems    []testMonItemAmts
	shopsPreAirship  []int32
	shopsPostAirship []int32
	treasures        []int32
	equipmentTables  []int32
}

func (e AutoAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e AutoAbility) Compare(test test, got apitypes.AutoAbility) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("monsters drop", apilookups.E.Monsters.Endpoint, e.monstersDrop, got.MonstersDrop)
	test.checkTestStructsInSlice("monsters items", e.monstersItems, got.MonstersItems)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.treasures, got.Treasures)
	test.checkResIDsInSlice("shops pre airship", apilookups.E.Shops.Endpoint, e.shopsPreAirship, got.ShopsPreAirship)
	test.checkResIDsInSlice("shops post airship", apilookups.E.Shops.Endpoint, e.shopsPostAirship, got.ShopsPostAirship)
	test.checkResIDsInSlice("equipment tables", apilookups.E.EquipmentTables.Endpoint, e.equipmentTables, got.EquipmentTables)
}
