package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Aeon struct {
	TestGeneral
	Unique
	UntypedUnit            int32
	Area                   int32
	BattlesToRegen         int32
	Agility                testAgilityParams
	CelestialWeapon        *int32
	CharacterClasses       []int32
	BaseStats              map[string]int32
	RegenAmounts           []testRegenAmount
	AeonCommands           []int32
	DefaultPlayerAbilities []int32
	Overdrives             []int32
	WeaponAbilities        []AeonEquipment
	ArmorAbilities         []AeonEquipment
}

func (e Aeon) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Aeon) Compare(test Test, got apitypes.Aeon) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped unit", apilookups.E.PlayerUnits.Endpoint, e.UntypedUnit, got.UntypedUnit)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, e.Area, got.Area)
	test.compare("battles to regenerate", e.BattlesToRegen, got.BattlesToRegenerate)
	e.Agility.Compare(test, "agility params", got.AgilityParameters)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.CelestialWeapon, got.CelestialWeapon)
	test.checkResIDsInSlice("character classes", apilookups.E.CharacterClasses.Endpoint, e.CharacterClasses, got.CharacterClasses)
	test.checkResAmtTypes("base stats", e.BaseStats, got.BaseStats)
	test.checkResIDsInSlice("aeon commands", apilookups.E.AeonCommands.Endpoint, e.AeonCommands, got.AeonCommands)
	test.checkResIDsInSlice("default abilities", apilookups.E.PlayerAbilities.Endpoint, e.DefaultPlayerAbilities, got.DefaultPlayerAbilities)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, e.Overdrives, got.Overdrives)
	test.compTestStructSlices("weapon abilities", e.WeaponAbilities, got.WeaponAbilities)
	test.compTestStructSlices("armor abilities", e.ArmorAbilities, got.ArmorAbilities)
}

type AeonEquipment struct {
	AutoAbility     int32
	CelestialWeapon bool
}

func (e AeonEquipment) Compare(test Test, _ string, got apitypes.AeonEquipment) {
	test.T.Helper()
	test.compIDAPIResource("auto ability", apilookups.E.AutoAbilities.Endpoint, e.AutoAbility, got.AutoAbility)
	test.compare("celestial weapon", e.CelestialWeapon, got.CelestialWeapon)
}
