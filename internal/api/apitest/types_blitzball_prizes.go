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

func compareBlitzballPrizes(test test, exp BlitzballPrize, got apitypes.BlitzballPrize) {
	test.t.Helper()
	test.compareExpIDOnly(exp.IDOnly, got.ID)
	test.compare("category", exp.category, got.Category)
	test.compare("slot", exp.slot, got.Slot)
	test.checkTestStructsInSlice("items", exp.items, got.Items, comparePossibleItems)
}
