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

func compareMasterItems(test test, exp MasterItem, got apitypes.MasterItem) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("item type", string(exp.itemType), string(got.Type))
	test.compPathAPIResource("typed item", exp.typedItem, got.TypedItem)
	test.compStructs("obtainable from", exp.obtainableFrom, got.ObtainableFrom)
}
