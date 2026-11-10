package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type expCharacter struct {
	testGeneral
	expUnique
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
	stdSphereGrid          *expSphereGrid
	expSphereGrid          *expSphereGrid
	overdriveModes         map[string]int32
}

func (e expCharacter) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareCharacters(test test, exp expCharacter, got apitypes.Character) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
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

type expSphereGrid struct {
	playerAbilities []int32
}

func compareSphereGrids(test test, fieldName string, exp expSphereGrid, got apitypes.SphereGrid) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" abilities", apilookups.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
}
