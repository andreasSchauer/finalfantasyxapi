package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareAutoAbilities(test Test, exp AutoAbility, got apitypes.AutoAbility) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("monsters drop", apilookups.E.Monsters.Endpoint, exp.monstersDrop, got.MonstersDrop)
	test.checkTestStructsInSlice("monsters items", exp.monstersItems, got.MonstersItems, compareMonItemAmts)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	test.checkResIDsInSlice("shops pre airship", apilookups.E.Shops.Endpoint, exp.shopsPreAirship, got.ShopsPreAirship)
	test.checkResIDsInSlice("shops post airship", apilookups.E.Shops.Endpoint, exp.shopsPostAirship, got.ShopsPostAirship)
	test.checkResIDsInSlice("equipment tables", apilookups.E.EquipmentTables.Endpoint, exp.equipmentTables, got.EquipmentTables)
}
