package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareBlitzballPrizes(test test, exp expBlitzballPrize, got api.BlitzballPrize) {
	test.t.Helper()
	compareExpIDOnly(test, exp.expIDOnly, got.ID)
	compare(test, "category", exp.category, got.Category)
	compare(test, "slot", exp.slot, got.Slot)
	checkTestStructsInSlice(test, "items", exp.items, got.Items, comparePossibleItems)
}
