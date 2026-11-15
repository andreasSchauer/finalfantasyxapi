package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Equipment struct {
	TestGeneral
	Unique
	EquipmentTable          int32
	Priority                *int32
	CelestialWeapon         *int32
	RequiredAutoAbilities   []int32
	SelectableAutoAbilities []testAbilityPool
	RequiredSlots           *int32
	Treasures               []int32
	Shops                   []int32
}

func (e Equipment) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Equipment) Compare(test Test, got apitypes.EquipmentName) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("equipment table", apilookups.E.EquipmentTables.Endpoint, e.EquipmentTable, got.EquipmentTable)
	test.compare("priority", e.Priority, got.Priority)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.CelestialWeapon, got.CelestialWeapon)
	test.checkResIDsInSlice("required auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.RequiredAutoAbilities, got.RequiredAutoAbilities)
	test.compTestStructSlices("selectable auto-abilities", e.SelectableAutoAbilities, got.SelectableAutoAbilities)
	test.compare("required slots", e.RequiredSlots, got.RequiredSlots)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.Treasures, got.Treasures)
	test.checkResIDsInSlice("shops", apilookups.E.Shops.Endpoint, e.Shops, got.Shops)
}
