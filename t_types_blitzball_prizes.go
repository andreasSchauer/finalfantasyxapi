package main

type expBlitzballPrize struct {
	testGeneral
	expIdOnly
	category		string
	slot			string
	items			[]testPossibleItem
}

func (e expBlitzballPrize) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareBlitzballPrizes(test test, exp expBlitzballPrize, got BlitzballPrize) {
	compareExpIdOnly(test, exp.expIdOnly, got.ID)
	compare(test, "category", exp.category, got.Category)
	compare(test, "slot", exp.slot, got.Slot)
	checkTestStructsInSlice(test, "items", exp.items, got.Items, comparePossibleItems)
}