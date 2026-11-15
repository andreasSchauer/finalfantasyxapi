package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type MasterItem struct {
	TestGeneral
	Unique
	ItemType       database.ItemType
	TypedItem      string
	ObtainableFrom apitypes.ObtainableFrom
}

func (e MasterItem) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e MasterItem) Compare(test Test, got apitypes.MasterItem) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("item type", string(e.ItemType), string(got.Type))
	test.compPathAPIResource("typed item", e.TypedItem, got.TypedItem)
	test.compStructs("obtainable from", e.ObtainableFrom, got.ObtainableFrom)
}
