package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareCharacters(test Test, exp Character, got apitypes.Character) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped unit", apilookups.E.PlayerUnits.Endpoint, exp.untypedUnit, got.UntypedUnit)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, exp.area, got.Area)
	test.compare("weapon type", exp.weaponType, got.WeaponType)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	test.compare("celestial formula", exp.celestialFormula, got.CelestialFormula)
	test.compIDAPIResourcePtrs("overdrive command", apilookups.E.OverdriveCommands.Endpoint, exp.overdriveCommand, got.OverdriveCommand)
	test.checkResIDsInSlice("character classes", apilookups.E.CharacterClasses.Endpoint, exp.characterClasses, got.CharacterClasses)
	compareAgilityParams(test, "agility params", exp.agility, got.AgilityParameters)
	test.checkResAmtTypes("base stats", exp.baseStats, got.BaseStats)
	test.checkTestStructsInSlice("regen amounts", exp.regenAmounts, got.RegenAmounts, compareRegenAmounts)
	test.checkResIDsInSlice("default abilities", apilookups.E.PlayerAbilities.Endpoint, exp.defaultPlayerAbilities, got.DefaultPlayerAbilities)
	test.compTestStructPtrs("std sg", exp.stdSphereGrid, got.StdSphereGrid, compareSphereGrids)
	test.compTestStructPtrs("exp sg", exp.expSphereGrid, got.ExpSphereGrid, compareSphereGrids)
	test.checkResAmts("overdrive modes", exp.overdriveModes, got.OverdriveModes)
}

func compareSphereGrids(test Test, fieldName string, exp SphereGrid, got apitypes.SphereGrid) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" abilities", apilookups.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
}
