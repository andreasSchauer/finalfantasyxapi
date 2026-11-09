package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
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

func comparePlayerUnits(test test, exp expPlayerUnit, got apitypes.PlayerUnit) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compare("unit type", string(exp.unitType), string(got.Type))
	test.compPathAPIResource("typed unit", exp.typedUnit, got.TypedUnit)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, exp.area, got.Area)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	test.checkResIDsInSlice("character classes", apilookups.E.CharacterClasses.Endpoint, exp.characterClasses, got.CharacterClasses)
}
