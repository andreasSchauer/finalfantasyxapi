package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"

type expTreasure struct {
	testGeneral
	expIdOnly
	area            int32
	availability    database.AvailabilityType
	isAnimaTreasure bool
	treasureType    database.TreasureType
	lootType        database.LootType
	gilAmount       *int32
	items           []testResAmount[TypedAPIResource]
	equipment       *testFoundEquipment
}

func (e expTreasure) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareTreasures(test test, exp expTreasure, got Treasure) {
	test.t.Helper()
	compareExpIdOnly(test, exp.expIdOnly, got.ID)
	compIdApiResource(test, "area", test.cfg.e.areas.endpoint, exp.area, got.Area)
	compare(test, "availability", string(exp.availability), got.Availability)
	compare(test, "is anima treasure", exp.isAnimaTreasure, got.IsAnimaTreasure)
	compare(test, "treasure type", string(exp.treasureType), got.TreasureType)
	compare(test, "loot type", string(exp.lootType), got.LootType)
	compare(test, "gil amount", exp.gilAmount, got.GilAmount)
	compTestStructSlices(test, "items", exp.items, got.Items, compareResAmounts)
	compTestStructPtrs(test, "equipment", exp.equipment, got.Equipment, compareFoundEquipment)
}

type testFoundEquipment struct {
	equipmentName    int32
	abilities        []int32
	emptySlotsAmount int32
}

func compareFoundEquipment(test test, fieldName string, exp testFoundEquipment, got FoundEquipment) {
	test.t.Helper()
	enEndpoint := test.cfg.e.equipment.endpoint
	aaEndpoint := test.cfg.e.autoAbilities.endpoint

	compIdApiResource(test, fieldName+" - equipment name", enEndpoint, exp.equipmentName, got.EquipmentName)
	checkResIDsInSlice(test, fieldName+" - abilities", aaEndpoint, exp.abilities, got.Abilities)
	compare(test, fieldName+" - empty slots amount", exp.emptySlotsAmount, got.EmptySlotsAmount)
}
