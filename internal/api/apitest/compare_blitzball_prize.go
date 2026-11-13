package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareBlitzballPrizes(test Test, exp BlitzballPrize, got apitypes.BlitzballPrize) {
	test.t.Helper()
	test.compareExpIDOnly(exp.IDOnly, got.ID)
	test.compare("category", exp.category, got.Category)
	test.compare("slot", exp.slot, got.Slot)
	test.checkTestStructsInSlice("items", exp.items, got.Items, comparePossibleItems)
}
