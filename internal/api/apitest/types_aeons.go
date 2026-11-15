package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Aeon struct {
	TestGeneral
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

func (e Aeon) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Aeon) Compare(test Test, got apitypes.Aeon) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped unit", apilookups.E.PlayerUnits.Endpoint, e.untypedUnit, got.UntypedUnit)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, e.area, got.Area)
	test.compare("battles to regenerate", e.battlesToRegen, got.BattlesToRegenerate)
	e.agility.Compare(test, "agility params", got.AgilityParameters)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.celestialWeapon, got.CelestialWeapon)
	test.checkResIDsInSlice("character classes", apilookups.E.CharacterClasses.Endpoint, e.characterClasses, got.CharacterClasses)
	test.checkResAmtTypes("base stats", e.baseStats, got.BaseStats)
	test.checkResIDsInSlice("aeon commands", apilookups.E.AeonCommands.Endpoint, e.aeonCommands, got.AeonCommands)
	test.checkResIDsInSlice("default abilities", apilookups.E.PlayerAbilities.Endpoint, e.defaultPlayerAbilities, got.DefaultPlayerAbilities)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, e.overdrives, got.Overdrives)
	test.compTestStructSlices("weapon abilities", e.weaponAbilities, got.WeaponAbilities)
	test.compTestStructSlices("armor abilities", e.armorAbilities, got.ArmorAbilities)
}

type AeonEquipment struct {
	autoAbility     int32
	celestialWeapon bool
}

func (e AeonEquipment) Compare(test Test, _ string, got apitypes.AeonEquipment) {
	test.T.Helper()
	test.compIDAPIResource("auto ability", apilookups.E.AutoAbilities.Endpoint, e.autoAbility, got.AutoAbility)
	test.compare("celestial weapon", e.celestialWeapon, got.CelestialWeapon)
}
