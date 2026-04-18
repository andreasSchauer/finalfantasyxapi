package api

type expPlayerUnit struct {
	testGeneral
	expUnique
	unitType         int32
	typedUnit        string
	area             int32
	celestialWeapon  *int32
	characterClasses []int32
}

func (e expPlayerUnit) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func comparePlayerUnits(test test, exp expPlayerUnit, got PlayerUnit) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "unit type", test.cfg.e.unitType.endpoint, exp.unitType, got.Type)
	compPathApiResource(test, "typed unit", exp.typedUnit, got.TypedUnit)
	compIdApiResource(test, "area", test.cfg.e.areas.endpoint, exp.area, got.Area)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.e.celestialWeapons.endpoint, exp.celestialWeapon, got.CelestialWeapon)
	checkResIDsInSlice(test, "character classes", test.cfg.e.characterClasses.endpoint, exp.characterClasses, got.CharacterClasses)
}
