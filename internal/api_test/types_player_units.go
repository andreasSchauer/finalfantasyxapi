package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type expPlayerUnit struct {
	testGeneral
	expUnique
	unitType         database.UnitType
	typedUnit        string
	area             int32
	celestialWeapon  *int32
	characterClasses []int32
}

func (e expPlayerUnit) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func comparePlayerUnits(test test, exp expPlayerUnit, got api.PlayerUnit) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "unit type", string(exp.unitType), string(got.Type))
	compPathAPIResource(test, "typed unit", exp.typedUnit, got.TypedUnit)
	compIDAPIResource(test, "area", test.cfg.E.Areas.Endpoint, exp.area, got.Area)
	compIDAPIResourcePtrs(test, "celestial weapon", test.cfg.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	checkResIDsInSlice(test, "character classes", test.cfg.E.CharacterClasses.Endpoint, exp.characterClasses, got.CharacterClasses)
}
