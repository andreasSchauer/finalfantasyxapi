package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type AutoAbility struct {
	TestGeneral
	Unique
	MonstersDrop     []int32
	MonstersItems    []TestMonItemAmts
	ShopsPreAirship  []int32
	ShopsPostAirship []int32
	Treasures        []int32
	EquipmentTables  []int32
}

func (e AutoAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e AutoAbility) Compare(test Test, got apitypes.AutoAbility) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.CheckResIDsInSlice("monsters drop", apilookups.E.Monsters.Endpoint, e.MonstersDrop, got.MonstersDrop)
	test.checkTestStructsInSlice("monsters items", e.MonstersItems, got.MonstersItems)
	test.CheckResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.Treasures, got.Treasures)
	test.CheckResIDsInSlice("shops pre airship", apilookups.E.Shops.Endpoint, e.ShopsPreAirship, got.ShopsPreAirship)
	test.CheckResIDsInSlice("shops post airship", apilookups.E.Shops.Endpoint, e.ShopsPostAirship, got.ShopsPostAirship)
	test.CheckResIDsInSlice("equipment tables", apilookups.E.EquipmentTables.Endpoint, e.EquipmentTables, got.EquipmentTables)
}
