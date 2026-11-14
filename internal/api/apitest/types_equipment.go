package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Equipment struct {
	testGeneral
	Unique
	equipmentTable          int32
	priority                *int32
	celestialWeapon         *int32
	requiredAutoAbilities   []int32
	selectableAutoAbilities []testAbilityPool
	requiredSlots           *int32
	treasures               []int32
	shops                   []int32
}

func (e Equipment) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Equipment) Compare(test test, got apitypes.EquipmentName) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("equipment table", apilookups.E.EquipmentTables.Endpoint, e.equipmentTable, got.EquipmentTable)
	test.compare("priority", e.priority, got.Priority)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.celestialWeapon, got.CelestialWeapon)
	test.checkResIDsInSlice("required auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.requiredAutoAbilities, got.RequiredAutoAbilities)
	test.compTestStructSlices("selectable auto-abilities", e.selectableAutoAbilities, got.SelectableAutoAbilities)
	test.compare("required slots", e.requiredSlots, got.RequiredSlots)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.treasures, got.Treasures)
	test.checkResIDsInSlice("shops", apilookups.E.Shops.Endpoint, e.shops, got.Shops)
}
