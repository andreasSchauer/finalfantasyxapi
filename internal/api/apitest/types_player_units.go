package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type PlayerUnit struct {
	TestGeneral
	Unique
	UnitType         database.UnitType
	TypedUnit        string
	Area             int32
	CelestialWeapon  *int32
	CharacterClasses []int32
}

func (e PlayerUnit) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e PlayerUnit) Compare(test Test, got apitypes.PlayerUnit) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("unit type", string(e.UnitType), string(got.Type))
	test.compPathAPIResource("typed unit", e.TypedUnit, got.TypedUnit)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, e.Area, got.Area)
	test.compIDAPIResourcePtrs("celestial weapon", apilookups.E.CelestialWeapons.Endpoint, e.CelestialWeapon, got.CelestialWeapon)
	test.checkResIDsInSlice("character classes", apilookups.E.CharacterClasses.Endpoint, e.CharacterClasses, got.CharacterClasses)
}
