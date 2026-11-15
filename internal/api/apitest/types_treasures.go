package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Treasure struct {
	TestGeneral
	IDOnly
	area            int32
	availability    database.AvailabilityType
	isAnimaTreasure bool
	treasureType    database.TreasureType
	lootType        database.LootType
	gilAmount       *int32
	items           []testResAmount[apiresource.Typed]
	equipment       *testFoundEquipment
}

func (e Treasure) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Treasure) Compare(test Test, got apitypes.Treasure) {
	test.T.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, e.area, got.Area)
	test.compare("availability", string(e.availability), got.Availability)
	test.compare("is anima treasure", e.isAnimaTreasure, got.IsAnimaTreasure)
	test.compare("treasure type", string(e.treasureType), got.TreasureType)
	test.compare("loot type", string(e.lootType), got.LootType)
	test.compare("gil amount", e.gilAmount, got.GilAmount)
	test.compTestStructSlices("items", e.items, got.Items)
	test.compTestStructPtrs("equipment", e.equipment, got.Equipment)
}

type testFoundEquipment struct {
	equipmentName    int32
	abilities        []int32
	emptySlotsAmount int32
}

func (e testFoundEquipment) Compare(test Test, fieldName string, got apitypes.FoundEquipment) {
	test.T.Helper()
	enEndpoint := apilookups.E.Equipment.Endpoint
	aaEndpoint := apilookups.E.AutoAbilities.Endpoint

	test.compIDAPIResource(fieldName+" - equipment name", enEndpoint, e.equipmentName, got.EquipmentName)
	test.checkResIDsInSlice(fieldName+" - abilities", aaEndpoint, e.abilities, got.Abilities)
	test.compare(fieldName+" - empty slots amount", e.emptySlotsAmount, got.EmptySlotsAmount)
}
