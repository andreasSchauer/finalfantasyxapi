package apiTest

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
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "equipment table", test.cfg.E.EquipmentTables.Endpoint, exp.equipmentTable, got.EquipmentTable)
	compare(test, "priority", exp.priority, got.Priority)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	checkResIDsInSlice(test, "required auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.requiredAutoAbilities, got.RequiredAutoAbilities)
	compTestStructSlices(test, "selectable auto-abilities", exp.selectableAutoAbilities, got.SelectableAutoAbilities, compareAbilityPools)
	compare(test, "required slots", exp.requiredSlots, got.RequiredSlots)
	checkResIDsInSlice(test, "treasures", test.cfg.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "shops", test.cfg.E.Shops.Endpoint, exp.shops, got.Shops)
}
