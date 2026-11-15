package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Treasure struct {
	TestGeneral
	IDOnly
	Area            int32
	Availability    database.AvailabilityType
	IsAnimaTreasure bool
	TreasureType    database.TreasureType
	LootType        database.LootType
	GilAmount       *int32
	Items           []TestResAmount[apiresource.Typed]
	Equipment       *TestFoundEquipment
}

func (e Treasure) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Treasure) Compare(test Test, got apitypes.Treasure) {
	test.T.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, e.Area, got.Area)
	test.compare("availability", string(e.Availability), got.Availability)
	test.compare("is anima treasure", e.IsAnimaTreasure, got.IsAnimaTreasure)
	test.compare("treasure type", string(e.TreasureType), got.TreasureType)
	test.compare("loot type", string(e.LootType), got.LootType)
	test.compare("gil amount", e.GilAmount, got.GilAmount)
	test.compTestStructSlices("items", e.Items, got.Items)
	test.compTestStructPtrs("equipment", e.Equipment, got.Equipment)
}

type TestFoundEquipment struct {
	EquipmentName    int32
	Abilities        []int32
	EmptySlotsAmount int32
}

func (e TestFoundEquipment) Compare(test Test, fieldName string, got apitypes.FoundEquipment) {
	test.T.Helper()
	enEndpoint := apilookups.E.Equipment.Endpoint
	aaEndpoint := apilookups.E.AutoAbilities.Endpoint

	test.compIDAPIResource(fieldName+" - equipment name", enEndpoint, e.EquipmentName, got.EquipmentName)
	test.CheckResIDsInSlice(fieldName+" - abilities", aaEndpoint, e.Abilities, got.Abilities)
	test.compare(fieldName+" - empty slots amount", e.EmptySlotsAmount, got.EmptySlotsAmount)
}
