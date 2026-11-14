package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type KeyItem struct {
	testGeneral
	Unique
	untypedItem     int32
	category        database.KeyItemCategory
	celestialWeapon *int32
	primer          *int32
	areas           []int32
	treasures       []int32
	quests          []int32
}

func (e KeyItem) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e KeyItem) Compare(test test, got apitypes.KeyItem) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped item", apilookups.E.AllItems.Endpoint, e.untypedItem, got.UntypedItem)
	test.compare("category", string(e.category), got.Category)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.celestialWeapon, got.CelestialWeapon)
	test.compIDAPIResourcePtrs("primer", apilookups.E.Primers.Endpoint, e.primer, got.Primer)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, e.areas, got.Areas)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.treasures, got.Treasures)
	test.checkResIDsInSlice("quests", apilookups.E.Quests.Endpoint, e.quests, got.Quests)
}
