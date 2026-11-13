package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type MasterItem struct {
	TestGeneral
	Unique
	itemType       database.ItemType
	typedItem      string
	obtainableFrom apitypes.ObtainableFrom
}

func (e MasterItem) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
