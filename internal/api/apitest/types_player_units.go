package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type PlayerUnit struct {
	testGeneral
	Unique
	unitType         database.UnitType
	typedUnit        string
	area             int32
	celestialWeapon  *int32
	characterClasses []int32
}

func (e PlayerUnit) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func comparePlayerUnits(test test, exp PlayerUnit, got apitypes.PlayerUnit) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("unit type", string(exp.unitType), string(got.Type))
	test.compPathAPIResource("typed unit", exp.typedUnit, got.TypedUnit)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, exp.area, got.Area)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, exp.celestialWeapon, got.CelestialWeapon)
	test.checkResIDsInSlice("character classes", apilookups.E.CharacterClasses.Endpoint, exp.characterClasses, got.CharacterClasses)
}
