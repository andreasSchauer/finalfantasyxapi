package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareKeyItems(test Test, exp KeyItem, got apitypes.KeyItem) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped item", apilookups.E.AllItems.Endpoint, exp.untypedItem, got.UntypedItem)
	test.compare("category", string(exp.category), got.Category)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	test.compIDAPIResourcePtrs("primer", apilookups.E.Primers.Endpoint, exp.primer, got.Primer)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, exp.areas, got.Areas)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	test.checkResIDsInSlice("quests", apilookups.E.Quests.Endpoint, exp.quests, got.Quests)
}
