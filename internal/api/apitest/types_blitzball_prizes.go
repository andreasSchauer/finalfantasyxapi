package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type BlitzballPrize struct {
	testGeneral
	IDOnly
	category string
	slot     string
	items    []testPossibleItem
}

func (e BlitzballPrize) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e BlitzballPrize) Compare(test test, got apitypes.BlitzballPrize) {
	test.t.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compare("category", e.category, got.Category)
	test.compare("slot", e.slot, got.Slot)
	test.checkTestStructsInSlice("items", e.items, got.Items)
}
