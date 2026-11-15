package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type KeyItem struct {
	TestGeneral
	Unique
	UntypedItem     int32
	Category        database.KeyItemCategory
	CelestialWeapon *int32
	Primer          *int32
	Areas           []int32
	Treasures       []int32
	Quests          []int32
}

func (e KeyItem) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e KeyItem) Compare(test Test, got apitypes.KeyItem) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped item", apilookups.E.AllItems.Endpoint, e.UntypedItem, got.UntypedItem)
	test.compare("category", string(e.Category), got.Category)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.CelestialWeapon, got.CelestialWeapon)
	test.compIDAPIResourcePtrs("primer", apilookups.E.Primers.Endpoint, e.Primer, got.Primer)
	test.CheckResIDsInSlice("areas", apilookups.E.Areas.Endpoint, e.Areas, got.Areas)
	test.CheckResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.Treasures, got.Treasures)
	test.CheckResIDsInSlice("quests", apilookups.E.Quests.Endpoint, e.Quests, got.Quests)
}
