package api

type testResAmount[A APIResource] struct {
	item   int32
	amount int32
}

func compareResAmounts[A APIResource](test test, fieldName string, exp testResAmount[A], got ResourceAmount[A]) {
	compIdApiResource(test, fieldName+" - resource", test.cfg.e.allItems.endpoint, exp.item, got)
	compare(test, fieldName+" - amount", exp.amount, got.Amount)
}

func newTestResAmount[A APIResource](itemID int32, amount int32) testResAmount[A] {
	return testResAmount[A]{
		item:   itemID,
		amount: amount,
	}
}

type testPossibleItem struct {
	index int
	testResAmount[TypedAPIResource]
	chance int32
}

func newTestPossibleItem(idx int, itemID int32, amount, chance int32) testPossibleItem {
	return testPossibleItem{
		index:         idx,
		testResAmount: newTestResAmount[TypedAPIResource](itemID, amount),
		chance:        chance,
	}
}

func (t testPossibleItem) GetIndex() int {
	return t.index
}

func comparePossibleItems(test test, fieldName string, exp testPossibleItem, got PossibleItem) {
	compareResAmounts(test, fieldName+" - itemAmount", exp.testResAmount, got.ResourceAmount)
	compare(test, fieldName+" - chance", exp.chance, got.Chance)
}
