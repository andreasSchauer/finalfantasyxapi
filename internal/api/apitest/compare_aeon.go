package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareAeons(test Test, exp Aeon, got apitypes.Aeon) {
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

func compareAeonEquipment(test Test, _ string, exp AeonEquipment, got apitypes.AeonEquipment) {
	test.t.Helper()
	test.compIDAPIResource("auto ability", apilookups.E.AutoAbilities.Endpoint, exp.autoAbility, got.AutoAbility)
	test.compare("celestial weapon", exp.celestialWeapon, got.CelestialWeapon)
}
