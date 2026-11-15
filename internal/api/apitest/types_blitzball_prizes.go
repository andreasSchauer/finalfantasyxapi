package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type BlitzballPrize struct {
	TestGeneral
	IDOnly
	category string
	slot     string
	items    []testPossibleItem
}

func (e BlitzballPrize) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e BlitzballPrize) Compare(test Test, got apitypes.BlitzballPrize) {
	test.T.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compare("category", e.category, got.Category)
	test.compare("slot", e.slot, got.Slot)
	test.checkTestStructsInSlice("items", e.items, got.Items)
}
