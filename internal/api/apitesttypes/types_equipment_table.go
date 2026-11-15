package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type EquipmentTable struct {
	TestGeneral
	IDOnly
	CelestialWeapon         *int32
	SpecificCharacter       *int32
	RequiredAutoAbilities   []int32
	SelectableAutoAbilities []TestAbilityPool
	RequiredSlots           *int32
	Equipment               []int32
}

func (e EquipmentTable) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e EquipmentTable) Compare(test Test, got apitypes.EquipmentTable) {
	test.T.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.CelestialWeapon, got.CelestialWeapon)
	test.compIDAPIResourcePtrs("specific character", apilookups.E.Characters.Endpoint, e.SpecificCharacter, got.SpecificCharacter)
	test.CheckResIDsInSlice("required auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.RequiredAutoAbilities, got.RequiredAutoAbilities)
	test.compTestStructSlices("selectable auto-abilities", e.SelectableAutoAbilities, got.SelectableAutoAbilities)
	test.compare("required slots", e.RequiredSlots, got.RequiredSlots)
	test.CheckResIDsInSlice("equipment", apilookups.E.Equipment.Endpoint, e.Equipment, got.Equipment)
}

type TestAbilityPool struct {
	Index         int
	AutoAbilities []int32
	ReqAmount     int32
}

func (e TestAbilityPool) Compare(test Test, fieldName string, got apitypes.AbilityPool) {
	test.T.Helper()
	test.CheckResIDsInSlice(fieldName+" - auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.compare(fieldName+" - required amount", e.ReqAmount, got.ReqAmount)
}
