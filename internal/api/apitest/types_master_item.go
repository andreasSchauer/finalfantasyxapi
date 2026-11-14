package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type MasterItem struct {
	testGeneral
	Unique
	itemType       database.ItemType
	typedItem      string
	obtainableFrom apitypes.ObtainableFrom
}

func (e MasterItem) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e MasterItem) Compare(test test, got apitypes.MasterItem) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("item type", string(e.itemType), string(got.Type))
	test.compPathAPIResource("typed item", e.typedItem, got.TypedItem)
	test.compStructs("obtainable from", e.obtainableFrom, got.ObtainableFrom)
}
