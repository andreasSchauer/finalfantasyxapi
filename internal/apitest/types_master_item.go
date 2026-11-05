package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type expMasterItem struct {
	testGeneral
	expUnique
	itemType       database.ItemType
	typedItem      string
	obtainableFrom api.ObtainableFrom
}

func (e expMasterItem) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMasterItems(test test, exp expMasterItem, got api.MasterItem) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compare("item type", string(exp.itemType), string(got.Type))
	test.compPathAPIResource("typed item", exp.typedItem, got.TypedItem)
	test.compStructs("obtainable from", exp.obtainableFrom, got.ObtainableFrom)
}
