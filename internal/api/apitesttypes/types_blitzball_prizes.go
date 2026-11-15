package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type BlitzballPrize struct {
	TestGeneral
	IDOnly
	Category string
	Slot     string
	Items    []TestPossibleItem
}

func (e BlitzballPrize) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e BlitzballPrize) Compare(test Test, got apitypes.BlitzballPrize) {
	test.T.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compare("category", e.Category, got.Category)
	test.compare("slot", e.Slot, got.Slot)
	test.checkTestStructsInSlice("items", e.Items, got.Items)
}
