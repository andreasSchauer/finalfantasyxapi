package api

type expKeyItem struct {
	testGeneral
	expUnique
	untypedItem     int32
	category        int32
	celestialWeapon *int32
	primer          *int32
	areas           []int32
	treasures       []int32
	quests          []int32
}

func (e expKeyItem) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareKeyItems(test test, exp expKeyItem, got KeyItem) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "untyped item", test.cfg.e.allItems.endpoint, exp.untypedItem, got.UntypedItem)
	compIdApiResource(test, "category", test.cfg.e.keyItemCategory.endpoint, exp.category, got.Category)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.e.celestialWeapons.endpoint, exp.celestialWeapon, got.CelestialWeapon)
	compIdApiResourcePtrs(test, "primer", test.cfg.e.primers.endpoint, exp.primer, got.Primer)
	checkResIDsInSlice(test, "areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas)
	checkResIDsInSlice(test, "treasures", test.cfg.e.treasures.endpoint, exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "quests", test.cfg.e.quests.endpoint, exp.quests, got.Quests)
}
