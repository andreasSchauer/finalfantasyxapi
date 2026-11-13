package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type PlayerUnit struct {
	TestGeneral
	Unique
	unitType         database.UnitType
	typedUnit        string
	area             int32
	celestialWeapon  *int32
	characterClasses []int32
}

func (e PlayerUnit) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
