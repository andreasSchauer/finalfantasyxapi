package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Character struct {
	TestGeneral
	Unique
	UntypedUnit            int32
	Area                   int32
	WeaponType             string
	CelestialWeapon        *int32
	CelestialFormula       *string
	OverdriveCommand       *int32
	CharacterClasses       []int32
	Agility                TestAgilityParams
	BaseStats              map[string]int32
	RegenAmounts           []TestRegenAmount
	DefaultPlayerAbilities []int32
	StdSphereGrid          *SphereGrid
	ExpSphereGrid          *SphereGrid
	OverdriveModes         map[string]int32
}

func (e Character) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Character) Compare(test Test, got apitypes.Character) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped unit", apilookups.E.PlayerUnits.Endpoint, e.UntypedUnit, got.UntypedUnit)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, e.Area, got.Area)
	test.compare("weapon type", e.WeaponType, got.WeaponType)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.CelestialWeapon, got.CelestialWeapon)
	test.compare("celestial formula", e.CelestialFormula, got.CelestialFormula)
	test.compIDAPIResourcePtrs("overdrive command", apilookups.E.OverdriveCommands.Endpoint, e.OverdriveCommand, got.OverdriveCommand)
	test.checkResIDsInSlice("character classes", apilookups.E.CharacterClasses.Endpoint, e.CharacterClasses, got.CharacterClasses)
	e.Agility.Compare(test, "agility params", got.AgilityParameters)
	test.checkResAmtTypes("base stats", e.BaseStats, got.BaseStats)
	test.checkTestStructsInSlice("regen amounts", e.RegenAmounts, got.RegenAmounts)
	test.checkResIDsInSlice("default abilities", apilookups.E.PlayerAbilities.Endpoint, e.DefaultPlayerAbilities, got.DefaultPlayerAbilities)
	test.compTestStructPtrs("std sg", e.StdSphereGrid, got.StdSphereGrid)
	test.compTestStructPtrs("exp sg", e.ExpSphereGrid, got.ExpSphereGrid)
	test.checkResAmts("overdrive modes", e.OverdriveModes, got.OverdriveModes)
}

type SphereGrid struct {
	PlayerAbilities []int32
}

func (e SphereGrid) Compare(test Test, fieldName string, got apitypes.SphereGrid) {
	test.T.Helper()
	test.checkResIDsInSlice(fieldName+" abilities", apilookups.E.PlayerAbilities.Endpoint, e.PlayerAbilities, got.PlayerAbilities)
}
