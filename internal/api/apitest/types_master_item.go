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

func (e MasterItem) Compare(test Test, got apitypes.MasterItem) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("item type", string(e.itemType), string(got.Type))
	test.compPathAPIResource("typed item", e.typedItem, got.TypedItem)
	test.compStructs("obtainable from", e.obtainableFrom, got.ObtainableFrom)
}
