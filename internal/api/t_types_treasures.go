package api

type expTreasure struct {
	testGeneral
	expIdOnly
	area            int32
	isPostAirship   bool
	isAnimaTreasure bool
	treasureType    int32
	lootType        int32
	gilAmount       *int32
	items           []testItemAmount
	equipment       *testFoundEquipment
}

func (e expTreasure) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareTreasures(test test, exp expTreasure, got Treasure) {
	compareExpIdOnly(test, exp.expIdOnly, got.ID)
	compIdApiResource(test, "area", test.cfg.e.areas.endpoint, exp.area, got.Area)
	compare(test, "is post airship", exp.isPostAirship, got.IsPostAirship)
	compare(test, "is anima treasure", exp.isAnimaTreasure, got.IsAnimaTreasure)
	compIdApiResource(test, "treasure type", test.cfg.e.treasureType.endpoint, exp.treasureType, got.TreasureType)
	compIdApiResource(test, "loot type", test.cfg.e.lootType.endpoint, exp.lootType, got.LootType)
	compare(test, "gil amount", exp.gilAmount, got.GilAmount)
	compTestStructSlices(test, "items", exp.items, got.Items, compareItemAmounts)
	compTestStructPtrs(test, "equipment", exp.equipment, got.Equipment, compareFoundEquipment)
}

type testFoundEquipment struct {
	equipmentName    int32
	abilities        []int32
	emptySlotsAmount int32
}

func compareFoundEquipment(test test, fieldName string, exp testFoundEquipment, got FoundEquipment) {
	enEndpoint := test.cfg.e.equipment.endpoint
	aaEndpoint := test.cfg.e.autoAbilities.endpoint

	compIdApiResource(test, fieldName+" - equipment name", enEndpoint, exp.equipmentName, got.EquipmentName)
	checkResIDsInSlice(test, fieldName+" - abilities", aaEndpoint, exp.abilities, got.Abilities)
	compare(test, fieldName+" - empty slots amount", exp.emptySlotsAmount, got.EmptySlotsAmount)
}
