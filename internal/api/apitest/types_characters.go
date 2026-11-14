package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Character struct {
	testGeneral
	Unique
	untypedUnit            int32
	area                   int32
	weaponType             string
	celestialWeapon        *int32
	celestialFormula       *string
	overdriveCommand       *int32
	characterClasses       []int32
	agility                testAgilityParams
	baseStats              map[string]int32
	regenAmounts           []testRegenAmount
	defaultPlayerAbilities []int32
	stdSphereGrid          *SphereGrid
	expSphereGrid          *SphereGrid
	overdriveModes         map[string]int32
}

func (e Character) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Character) Compare(test test, got apitypes.Character) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped unit", apilookups.E.PlayerUnits.Endpoint, e.untypedUnit, got.UntypedUnit)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, e.area, got.Area)
	test.compare("weapon type", e.weaponType, got.WeaponType)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.celestialWeapon, got.CelestialWeapon)
	test.compare("celestial formula", e.celestialFormula, got.CelestialFormula)
	test.compIDAPIResourcePtrs("overdrive command", apilookups.E.OverdriveCommands.Endpoint, e.overdriveCommand, got.OverdriveCommand)
	test.checkResIDsInSlice("character classes", apilookups.E.CharacterClasses.Endpoint, e.characterClasses, got.CharacterClasses)
	e.agility.Compare(test, "agility params", got.AgilityParameters)
	test.checkResAmtTypes("base stats", e.baseStats, got.BaseStats)
	test.checkTestStructsInSlice("regen amounts", e.regenAmounts, got.RegenAmounts)
	test.checkResIDsInSlice("default abilities", apilookups.E.PlayerAbilities.Endpoint, e.defaultPlayerAbilities, got.DefaultPlayerAbilities)
	test.compTestStructPtrs("std sg", e.stdSphereGrid, got.StdSphereGrid)
	test.compTestStructPtrs("exp sg", e.expSphereGrid, got.ExpSphereGrid)
	test.checkResAmts("overdrive modes", e.overdriveModes, got.OverdriveModes)
}

type SphereGrid struct {
	playerAbilities []int32
}

func (e SphereGrid) Compare(test test, fieldName string, got apitypes.SphereGrid) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" abilities", apilookups.E.PlayerAbilities.Endpoint, e.playerAbilities, got.PlayerAbilities)
}
