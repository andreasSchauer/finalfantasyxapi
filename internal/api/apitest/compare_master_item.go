package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareMasterItems(test Test, exp MasterItem, got apitypes.MasterItem) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("item type", string(exp.itemType), string(got.Type))
	test.compPathAPIResource("typed item", exp.typedItem, got.TypedItem)
	test.compStructs("obtainable from", exp.obtainableFrom, got.ObtainableFrom)
}
