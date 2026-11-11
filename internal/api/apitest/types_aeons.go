package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Aeon struct {
	testGeneral
	Unique
	untypedUnit            int32
	area                   int32
	battlesToRegen         int32
	agility                testAgilityParams
	celestialWeapon        *int32
	characterClasses       []int32
	baseStats              map[string]int32
	regenAmounts           []testRegenAmount
	aeonCommands           []int32
	defaultPlayerAbilities []int32
	overdrives             []int32
	weaponAbilities        []AeonEquipment
	armorAbilities         []AeonEquipment
}

func (e Aeon) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAeons(test test, exp Aeon, got apitypes.Aeon) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped unit", apilookups.E.PlayerUnits.Endpoint, exp.untypedUnit, got.UntypedUnit)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, exp.area, got.Area)
	test.compare("battles to regenerate", exp.battlesToRegen, got.BattlesToRegenerate)
	compareAgilityParams(test, "agility params", exp.agility, got.AgilityParameters)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	test.checkResIDsInSlice("character classes", apilookups.E.CharacterClasses.Endpoint, exp.characterClasses, got.CharacterClasses)
	test.checkResAmtTypes("base stats", exp.baseStats, got.BaseStats)
	test.checkResIDsInSlice("aeon commands", apilookups.E.AeonCommands.Endpoint, exp.aeonCommands, got.AeonCommands)
	test.checkResIDsInSlice("default abilities", apilookups.E.PlayerAbilities.Endpoint, exp.defaultPlayerAbilities, got.DefaultPlayerAbilities)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
	test.compTestStructSlices("weapon abilities", exp.weaponAbilities, got.WeaponAbilities, compareAeonEquipment)
	test.compTestStructSlices("armor abilities", exp.armorAbilities, got.ArmorAbilities, compareAeonEquipment)
}

type AeonEquipment struct {
	autoAbility     int32
	celestialWeapon bool
}

func compareAeonEquipment(test test, _ string, exp AeonEquipment, got apitypes.AeonEquipment) {
	test.t.Helper()
	test.compIDAPIResource("auto ability", apilookups.E.AutoAbilities.Endpoint, exp.autoAbility, got.AutoAbility)
	test.compare("celestial weapon", exp.celestialWeapon, got.CelestialWeapon)
}
