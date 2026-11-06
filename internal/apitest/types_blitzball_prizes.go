package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

type expBlitzballPrize struct {
	testGeneral
	expIDOnly
	category string
	slot     string
	items    []testPossibleItem
}

func (e expBlitzballPrize) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareBlitzballPrizes(test test, exp expBlitzballPrize, got apitypes.BlitzballPrize) {
	test.t.Helper()
	test.compareExpIDOnly(exp.expIDOnly, got.ID)
	test.compare("category", exp.category, got.Category)
	test.compare("slot", exp.slot, got.Slot)
	test.checkTestStructsInSlice("items", exp.items, got.Items, comparePossibleItems)
}