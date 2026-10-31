package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareCharacters(test test, exp expCharacter, got api.Character) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIDAPIResource(test, "untyped unit", test.cfg.E.PlayerUnits.Endpoint, exp.untypedUnit, got.UntypedUnit)
	compIDAPIResource(test, "area", test.cfg.E.Areas.Endpoint, exp.area, got.Area)
	compare(test, "weapon type", exp.weaponType, got.WeaponType)
	compIDAPIResourcePtrs(test, "celestial weapon", test.cfg.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	compare(test, "celestial formula", exp.celestialFormula, got.CelestialFormula)
	compIDAPIResourcePtrs(test, "overdrive command", test.cfg.E.OverdriveCommands.Endpoint, exp.overdriveCommand, got.OverdriveCommand)
	checkResIDsInSlice(test, "character classes", test.cfg.E.CharacterClasses.Endpoint, exp.characterClasses, got.CharacterClasses)
	compareAgilityParams(test, "agility params", exp.agility, got.AgilityParameters)
	checkResAmtTypes(test, "base stats", exp.baseStats, got.BaseStats)
	checkTestStructsInSlice(test, "regen amounts", exp.regenAmounts, got.RegenAmounts, compareRegenAmounts)
	checkResIDsInSlice(test, "default abilities", test.cfg.E.PlayerAbilities.Endpoint, exp.defaultPlayerAbilities, got.DefaultPlayerAbilities)
	compTestStructPtrs(test, "std sg", exp.stdSphereGrid, got.StdSphereGrid, compareSphereGrids)
	compTestStructPtrs(test, "exp sg", exp.expSphereGrid, got.ExpSphereGrid, compareSphereGrids)
	checkResAmts(test, "overdrive modes", exp.overdriveModes, got.OverdriveModes)
}

type expSphereGrid struct {
	playerAbilities []int32
}

func compareSphereGrids(test test, fieldName string, exp expSphereGrid, got api.SphereGrid) {
	test.t.Helper()
	checkResIDsInSlice(test, fieldName+" abilities", test.cfg.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
}
