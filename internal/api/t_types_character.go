package api

type expCharacter struct {
	testGeneral
	expUnique
	area				int32
	weaponType			string
	celestialWeapon		*int32
	overdriveCommand	*int32
	characterClasses	[]int32
	baseStats			map[string]int32
	defaultAbilities	[]int32
	stdSgAbilities		[]int32
	expSgAbilities		[]int32
	overdriveModes		map[string]int32
}

func (e expCharacter) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareCharacters(test test, exp expCharacter, got Character) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "area", test.cfg.e.areas.endpoint, exp.area, got.Area)
	compare(test, "weapon type", exp.weaponType, got.WeaponType)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.e.celestialWeapons.endpoint, exp.celestialWeapon, got.CelestialWeapon)
	compIdApiResourcePtrs(test, "overdrive command", test.cfg.e.overdriveCommands.endpoint, exp.overdriveCommand, got.OverdriveCommand)
	checkResIDsInSlice(test, "character classes", test.cfg.e.characterClasses.endpoint, exp.characterClasses, got.CharacterClasses)
	checkResAmtsNameVals(test, "base stats", exp.baseStats, got.BaseStats)
	checkResIDsInSlice(test, "default abilities", test.cfg.e.playerAbilities.endpoint, exp.defaultAbilities, got.DefaultAbilities)
	checkResIDsInSlice(test, "standard sg abilities", test.cfg.e.playerAbilities.endpoint, exp.stdSgAbilities, got.StdSphereGridAbilities)
	checkResIDsInSlice(test, "expert sg abilities", test.cfg.e.playerAbilities.endpoint, exp.expSgAbilities, got.ExpSphereGridAbilities)
	checkResAmtsNameVals(test, "overdrive modes", exp.overdriveModes, got.OverdriveModes)
}
