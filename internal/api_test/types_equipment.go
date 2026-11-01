package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expEquipment struct {
	testGeneral
	expUnique
	equipmentTable          int32
	priority                *int32
	celestialWeapon         *int32
	requiredAutoAbilities   []int32
	selectableAutoAbilities []testAbilityPool
	requiredSlots           *int32
	treasures               []int32
	shops                   []int32
}

func (e expEquipment) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareEquipment(test test, exp expEquipment, got api.EquipmentName) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("equipment table", test.cfg.E.EquipmentTables.Endpoint, exp.equipmentTable, got.EquipmentTable)
	test.compare("priority", exp.priority, got.Priority)
	test.compIDAPIResourcePtrs("celestial weapon", test.cfg.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	test.checkResIDsInSlice("required auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.requiredAutoAbilities, got.RequiredAutoAbilities)
	test.compTestStructSlices("selectable auto-abilities", exp.selectableAutoAbilities, got.SelectableAutoAbilities, compareAbilityPools)
	test.compare("required slots", exp.requiredSlots, got.RequiredSlots)
	test.checkResIDsInSlice("treasures", test.cfg.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	test.checkResIDsInSlice("shops", test.cfg.E.Shops.Endpoint, exp.shops, got.Shops)
}
