package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expAeon struct {
	testGeneral
	expUnique
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
	weaponAbilities        []expAeonEquipment
	armorAbilities         []expAeonEquipment
}

func (e expAeon) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAeons(test test, exp expAeon, got api.Aeon) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("untyped unit", api.E.PlayerUnits.Endpoint, exp.untypedUnit, got.UntypedUnit)
	test.compIDAPIResource("area", api.E.Areas.Endpoint, exp.area, got.Area)
	test.compare("battles to regenerate", exp.battlesToRegen, got.BattlesToRegenerate)
	compareAgilityParams(test, "agility params", exp.agility, got.AgilityParameters)
	test.compIDAPIResourcePtrs("celestial weapon", api.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	test.checkResIDsInSlice("character classes", api.E.CharacterClasses.Endpoint, exp.characterClasses, got.CharacterClasses)
	test.checkResAmtTypes("base stats", exp.baseStats, got.BaseStats)
	test.checkResIDsInSlice("aeon commands", api.E.AeonCommands.Endpoint, exp.aeonCommands, got.AeonCommands)
	test.checkResIDsInSlice("default abilities", api.E.PlayerAbilities.Endpoint, exp.defaultPlayerAbilities, got.DefaultPlayerAbilities)
	test.checkResIDsInSlice("overdrives", api.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
	test.compTestStructSlices("weapon abilities", exp.weaponAbilities, got.WeaponAbilities, compareAeonEquipment)
	test.compTestStructSlices("armor abilities", exp.armorAbilities, got.ArmorAbilities, compareAeonEquipment)
}

type expAeonEquipment struct {
	autoAbility     int32
	celestialWeapon bool
}

func compareAeonEquipment(test test, _ string, exp expAeonEquipment, got api.AeonEquipment) {
	test.t.Helper()
	test.compIDAPIResource("auto ability", api.E.AutoAbilities.Endpoint, exp.autoAbility, got.AutoAbility)
	test.compare("celestial weapon", exp.celestialWeapon, got.CelestialWeapon)
}
