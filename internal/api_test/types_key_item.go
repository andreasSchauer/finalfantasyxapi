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
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIDAPIResource(test, "untyped item", test.cfg.E.AllItems.Endpoint, exp.untypedItem, got.UntypedItem)
	compare(test, "category", string(exp.category), got.Category)
	compIDAPIResourcePtrs(test, "celestial weapon", test.cfg.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	compIDAPIResourcePtrs(test, "primer", test.cfg.E.Primers.Endpoint, exp.primer, got.Primer)
	checkResIDsInSlice(test, "areas", test.cfg.E.Areas.Endpoint, exp.areas, got.Areas)
	checkResIDsInSlice(test, "treasures", test.cfg.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "quests", test.cfg.E.Quests.Endpoint, exp.quests, got.Quests)
}
