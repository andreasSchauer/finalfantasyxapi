package apiTest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expEquipmentTable struct {
	testGeneral
	expIdOnly
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
	compareExpIdOnly(test, exp.expIdOnly, got.ID)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	compIdApiResourcePtrs(test, "specific character", test.cfg.E.Characters.Endpoint, exp.specificCharacter, got.SpecificCharacter)
	checkResIDsInSlice(test, "required auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.requiredAutoAbilities, got.RequiredAutoAbilities)
	compTestStructSlices(test, "selectable auto-abilities", exp.selectableAutoAbilities, got.SelectableAutoAbilities, compareAbilityPools)
	compare(test, "required slots", exp.requiredSlots, got.RequiredSlots)
	checkResIDsInSlice(test, "equipment", test.cfg.E.Equipment.Endpoint, exp.equipment, got.Equipment)
}

type testAbilityPool struct {
	index         int
	autoAbilities []int32
	reqAmount     int32
}

func compareAbilityPools(test test, fieldName string, exp testAbilityPool, got api.AbilityPool) {
	test.t.Helper()
	checkResIDsInSlice(test, fieldName+" - auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	compare(test, fieldName+" - required amount", exp.reqAmount, got.ReqAmount)
}
