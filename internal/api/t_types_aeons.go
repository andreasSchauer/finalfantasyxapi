package api


type expAeon struct {
	testGeneral
	expUnique
	area				int32
	battlesToRegen		int32
	agility				AgilityParams
	celestialWeapon		*int32
	characterClasses	[]int32
	baseStats			map[string]int32
	aeonCommands		[]int32
	defaultAbilities	[]int32
	overdrives			[]int32
	weaponAbilities		[]expAeonEquipment
	armorAbilities		[]expAeonEquipment
}

func (e expAeon) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAeons(test test, exp expAeon, got Aeon) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "area", test.cfg.e.areas.endpoint, exp.area, got.Area)
	compare(test, "battles to regenerate", exp.battlesToRegen, got.BattlesToRegenerate)
	compStructs(test, "agility params", exp.agility, got.AgilityParameters)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.e.celestialWeapons.endpoint, exp.celestialWeapon, got.CelestialWeapon)
	checkResIDsInSlice(test, "character classes", test.cfg.e.characterClasses.endpoint, exp.characterClasses, got.CharacterClasses)
	checkResAmtsNameVals(test, "base stats", exp.baseStats, got.BaseStats)
	checkResIDsInSlice(test, "aeon commands", test.cfg.e.aeonCommands.endpoint, exp.aeonCommands, got.AeonCommands)
	checkResIDsInSlice(test, "default abilities", test.cfg.e.playerAbilities.endpoint, exp.defaultAbilities, got.DefaultAbilities)
	checkResIDsInSlice(test, "overdrives", test.cfg.e.overdrives.endpoint, exp.overdrives, got.Overdrives)
	compTestStructSlices(test, "weapon abilities", exp.weaponAbilities, got.WeaponAbilities, compareAeonEquipment)
	compTestStructSlices(test, "armor abilities", exp.armorAbilities, got.ArmorAbilities, compareAeonEquipment)
}

type expAeonEquipment struct {
	autoAbility			int32
	celestialWeapon		bool
}

func compareAeonEquipment(test test, _ string, exp expAeonEquipment, got AeonEquipment) {
	compIdApiResource(test, "auto ability", test.cfg.e.autoAbilities.endpoint, exp.autoAbility, got.AutoAbility)
	compare(test, "celestial weapon", exp.celestialWeapon, got.CelestialWeapon)
}