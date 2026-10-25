package apiTest

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
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "untyped unit", test.cfg.E.PlayerUnits.Endpoint, exp.untypedUnit, got.UntypedUnit)
	compIdApiResource(test, "area", test.cfg.E.Areas.Endpoint, exp.area, got.Area)
	compare(test, "battles to regenerate", exp.battlesToRegen, got.BattlesToRegenerate)
	compareAgilityParams(test, "agility params", exp.agility, got.AgilityParameters)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	checkResIDsInSlice(test, "character classes", test.cfg.E.CharacterClasses.Endpoint, exp.characterClasses, got.CharacterClasses)
	checkResAmtTypes(test, "base stats", exp.baseStats, got.BaseStats)
	checkResIDsInSlice(test, "aeon commands", test.cfg.E.AeonCommands.Endpoint, exp.aeonCommands, got.AeonCommands)
	checkResIDsInSlice(test, "default abilities", test.cfg.E.PlayerAbilities.Endpoint, exp.defaultPlayerAbilities, got.DefaultPlayerAbilities)
	checkResIDsInSlice(test, "overdrives", test.cfg.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
	compTestStructSlices(test, "weapon abilities", exp.weaponAbilities, got.WeaponAbilities, compareAeonEquipment)
	compTestStructSlices(test, "armor abilities", exp.armorAbilities, got.ArmorAbilities, compareAeonEquipment)
}

type expAeonEquipment struct {
	autoAbility     int32
	celestialWeapon bool
}

func compareAeonEquipment(test test, _ string, exp expAeonEquipment, got api.AeonEquipment) {
	test.t.Helper()
	compIdApiResource(test, "auto ability", test.cfg.E.AutoAbilities.Endpoint, exp.autoAbility, got.AutoAbility)
	compare(test, "celestial weapon", exp.celestialWeapon, got.CelestialWeapon)
}
