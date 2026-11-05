package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type expKeyItem struct {
	testGeneral
	expUnique
	untypedItem     int32
	category        database.KeyItemCategory
	celestialWeapon *int32
	primer          *int32
	areas           []int32
	treasures       []int32
	quests          []int32
}

func (e expKeyItem) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareKeyItems(test test, exp expKeyItem, got api.KeyItem) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("untyped item", api.E.AllItems.Endpoint, exp.untypedItem, got.UntypedItem)
	test.compare("category", string(exp.category), got.Category)
	test.compIDAPIResourcePtrs("celestial weapon", api.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	test.compIDAPIResourcePtrs("primer", api.E.Primers.Endpoint, exp.primer, got.Primer)
	test.checkResIDsInSlice("areas", api.E.Areas.Endpoint, exp.areas, got.Areas)
	test.checkResIDsInSlice("treasures", api.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	test.checkResIDsInSlice("quests", api.E.Quests.Endpoint, exp.quests, got.Quests)
}
