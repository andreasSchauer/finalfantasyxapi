package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type EquipmentTable struct {
	testGeneral
	IDOnly
	celestialWeapon         *int32
	specificCharacter       *int32
	requiredAutoAbilities   []int32
	selectableAutoAbilities []testAbilityPool
	requiredSlots           *int32
	equipment               []int32
}

func (e EquipmentTable) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e EquipmentTable) Compare(test test, got apitypes.EquipmentTable) {
	test.t.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.celestialWeapon, got.CelestialWeapon)
	test.compIDAPIResourcePtrs("specific character", apilookups.E.Characters.Endpoint, e.specificCharacter, got.SpecificCharacter)
	test.checkResIDsInSlice("required auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.requiredAutoAbilities, got.RequiredAutoAbilities)
	test.compTestStructSlices("selectable auto-abilities", e.selectableAutoAbilities, got.SelectableAutoAbilities)
	test.compare("required slots", e.requiredSlots, got.RequiredSlots)
	test.checkResIDsInSlice("equipment", apilookups.E.Equipment.Endpoint, e.equipment, got.Equipment)
}

type testAbilityPool struct {
	index         int
	autoAbilities []int32
	reqAmount     int32
}

func (e testAbilityPool) Compare(test test, fieldName string, got apitypes.AbilityPool) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.autoAbilities, got.AutoAbilities)
	test.compare(fieldName+" - required amount", e.reqAmount, got.ReqAmount)
}
