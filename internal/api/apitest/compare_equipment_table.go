package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareEquipmentTables(test Test, exp EquipmentTable, got apitypes.EquipmentTable) {
	test.t.Helper()
	test.compareExpIDOnly(exp.IDOnly, got.ID)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	test.compIDAPIResourcePtrs("specific character", apilookups.E.Characters.Endpoint, exp.specificCharacter, got.SpecificCharacter)
	test.checkResIDsInSlice("required auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.requiredAutoAbilities, got.RequiredAutoAbilities)
	test.compTestStructSlices("selectable auto-abilities", exp.selectableAutoAbilities, got.SelectableAutoAbilities, compareAbilityPools)
	test.compare("required slots", exp.requiredSlots, got.RequiredSlots)
	test.checkResIDsInSlice("equipment", apilookups.E.Equipment.Endpoint, exp.equipment, got.Equipment)
}

func compareAbilityPools(test Test, fieldName string, exp TestAbilityPool, got apitypes.AbilityPool) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.compare(fieldName+" - required amount", exp.reqAmount, got.ReqAmount)
}
