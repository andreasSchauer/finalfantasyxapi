package apiTest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type testResAmount[A api.APIResource] struct {
	item   int32
	amount int32
}

func compareResAmounts[A api.APIResource](test test, fieldName string, exp testResAmount[A], got api.ResourceAmount[A]) {
	test.t.Helper()
	compIdApiResource(test, fieldName+" - resource", test.cfg.E.AllItems.Endpoint, exp.item, got)
	compare(test, fieldName+" - amount", exp.amount, got.Amount)
}

func newTestResAmount[A api.APIResource](itemID int32, amount int32) testResAmount[A] {
	return testResAmount[A]{
		item:   itemID,
		amount: amount,
	}
}

type testPossibleItem struct {
	index int
	testResAmount[api.TypedAPIResource]
	chance int32
}

func newTestPossibleItem(idx int, itemID int32, amount, chance int32) testPossibleItem {
	return testPossibleItem{
		index:         idx,
		testResAmount: newTestResAmount[api.TypedAPIResource](itemID, amount),
		chance:        chance,
	}
}

func (t testPossibleItem) GetIndex() int {
	return t.index
}

func comparePossibleItems(test test, fieldName string, exp testPossibleItem, got api.PossibleItem) {
	test.t.Helper()
	compareResAmounts(test, fieldName+" - itemAmount", exp.testResAmount, got.ResourceAmount)
	compare(test, fieldName+" - chance", exp.chance, got.Chance)
}
