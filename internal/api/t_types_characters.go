package api

type expCharacter struct {
	testGeneral
	expUnique
	untypedUnit            int32
	area                   int32
	weaponType             string
	celestialWeapon        *int32
	overdriveCommand       *int32
	characterClasses       []int32
	baseStats              map[string]int32
	defaultPlayerAbilities []int32
	stdSgPlayerAbilities   []int32
	expSgPlayerAbilities   []int32
	overdriveModes         map[string]int32
}

func (e expCharacter) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareCharacters(test test, exp expCharacter, got Character) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "untyped unit", test.cfg.e.playerUnits.endpoint, exp.untypedUnit, got.UntypedUnit)
	compIdApiResource(test, "area", test.cfg.e.areas.endpoint, exp.area, got.Area)
	compare(test, "weapon type", exp.weaponType, got.WeaponType)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.e.celestialWeapons.endpoint, exp.celestialWeapon, got.CelestialWeapon)
	compIdApiResourcePtrs(test, "overdrive command", test.cfg.e.overdriveCommands.endpoint, exp.overdriveCommand, got.OverdriveCommand)
	checkResIDsInSlice(test, "character classes", test.cfg.e.characterClasses.endpoint, exp.characterClasses, got.CharacterClasses)
	checkResAmtTypes(test, "base stats", exp.baseStats, got.BaseStats)
	checkResIDsInSlice(test, "default abilities", test.cfg.e.playerAbilities.endpoint, exp.defaultPlayerAbilities, got.DefaultPlayerAbilities)
	checkResIDsInSlice(test, "standard sg abilities", test.cfg.e.playerAbilities.endpoint, exp.stdSgPlayerAbilities, got.StdSgPlayerAbilities)
	checkResIDsInSlice(test, "expert sg abilities", test.cfg.e.playerAbilities.endpoint, exp.expSgPlayerAbilities, got.ExpSgPlayerAbilities)
	checkResAmts(test, "overdrive modes", exp.overdriveModes, got.OverdriveModes)
}
