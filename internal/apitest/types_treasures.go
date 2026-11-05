package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type expTreasure struct {
	testGeneral
	expIDOnly
	area            int32
	availability    database.AvailabilityType
	isAnimaTreasure bool
	treasureType    database.TreasureType
	lootType        database.LootType
	gilAmount       *int32
	items           []testResAmount[apiresource.TypedAPIResource]
	equipment       *testFoundEquipment
}

func (e expTreasure) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareTreasures(test test, exp expTreasure, got api.Treasure) {
	test.t.Helper()
	test.compareExpIDOnly(exp.expIDOnly, got.ID)
	test.compIDAPIResource("area", api.E.Areas.Endpoint, exp.area, got.Area)
	test.compare("availability", string(exp.availability), got.Availability)
	test.compare("is anima treasure", exp.isAnimaTreasure, got.IsAnimaTreasure)
	test.compare("treasure type", string(exp.treasureType), got.TreasureType)
	test.compare("loot type", string(exp.lootType), got.LootType)
	test.compare("gil amount", exp.gilAmount, got.GilAmount)
	test.compTestStructSlices("items", exp.items, got.Items, compareResAmounts)
	test.compTestStructPtrs("equipment", exp.equipment, got.Equipment, compareFoundEquipment)
}

type testFoundEquipment struct {
	equipmentName    int32
	abilities        []int32
	emptySlotsAmount int32
}

func compareFoundEquipment(test test, fieldName string, exp testFoundEquipment, got api.FoundEquipment) {
	test.t.Helper()
	enEndpoint := api.E.Equipment.Endpoint
	aaEndpoint := api.E.AutoAbilities.Endpoint

	test.compIDAPIResource(fieldName+" - equipment name", enEndpoint, exp.equipmentName, got.EquipmentName)
	test.checkResIDsInSlice(fieldName+" - abilities", aaEndpoint, exp.abilities, got.Abilities)
	test.compare(fieldName+" - empty slots amount", exp.emptySlotsAmount, got.EmptySlotsAmount)
}
