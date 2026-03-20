package api

type testItemAmount struct {
	item   int32
	amount int32
}

func compareItemAmounts(test test, fieldName string, exp testItemAmount, got ItemAmount) {
	compIdApiResource(test, fieldName+" - item", test.cfg.e.masterItems.endpoint, exp.item, got)
	compare(test, fieldName+" - amount", exp.amount, got.Amount)
}

func newTestItemAmount(itemID int32, amount int32) testItemAmount {
	return testItemAmount{
		item:   itemID,
		amount: amount,
	}
}

type testPossibleItem struct {
	index int
	testItemAmount
	chance int32
}

func newTestPossibleItem(idx int, itemID int32, amount, chance int32) testPossibleItem {
	return testPossibleItem{
		index:          idx,
		testItemAmount: newTestItemAmount(itemID, amount),
		chance:         chance,
	}
}

func (t testPossibleItem) GetIndex() int {
	return t.index
}

func comparePossibleItems(test test, fieldName string, exp testPossibleItem, got PossibleItem) {
	compareItemAmounts(test, fieldName+" - itemAmount", exp.testItemAmount, got.ItemAmount)
	compare(test, fieldName+" - chance", exp.chance, got.Chance)
}
