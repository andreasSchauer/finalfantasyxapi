package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expEquipmentTable struct {
	testGeneral
	expIDOnly
	celestialWeapon         *int32
	specificCharacter       *int32
	requiredAutoAbilities   []int32
	selectableAutoAbilities []testAbilityPool
	requiredSlots           *int32
	equipment               []int32
}

func (e expEquipmentTable) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareEquipmentTables(test test, exp expEquipmentTable, got api.EquipmentTable) {
	test.t.Helper()
	test.compareExpIDOnly(exp.expIDOnly, got.ID)
	test.compIDAPIResourcePtrs("celestial weapon", api.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	test.compIDAPIResourcePtrs("specific character", api.E.Characters.Endpoint, exp.specificCharacter, got.SpecificCharacter)
	test.checkResIDsInSlice("required auto-abilities", api.E.AutoAbilities.Endpoint, exp.requiredAutoAbilities, got.RequiredAutoAbilities)
	test.compTestStructSlices("selectable auto-abilities", exp.selectableAutoAbilities, got.SelectableAutoAbilities, compareAbilityPools)
	test.compare("required slots", exp.requiredSlots, got.RequiredSlots)
	test.checkResIDsInSlice("equipment", api.E.Equipment.Endpoint, exp.equipment, got.Equipment)
}

type testAbilityPool struct {
	index         int
	autoAbilities []int32
	reqAmount     int32
}

func compareAbilityPools(test test, fieldName string, exp testAbilityPool, got api.AbilityPool) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - auto-abilities", api.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.compare(fieldName+" - required amount", exp.reqAmount, got.ReqAmount)
}
