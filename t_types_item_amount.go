package main

type testItemAmount struct {
	item   string
	amount int32
}

func compareItemAmounts(test test, exp testItemAmount, got ItemAmount) {
	compPathApiResource(test, "item amount - item", exp.item, got)
	compare(test, "item amount - amount", exp.amount, got.Amount)
}

func newTestItemAmount(itemPath string, amount int32) testItemAmount {
	return testItemAmount{
		item: 	itemPath,
		amount: amount,
	}
}

type testPossibleItem struct {
	index			int
	testItemAmount
	chance 			int32
}

func newTestPossibleItem(idx int, itemPath string, amount, chance int32) testPossibleItem {
	return testPossibleItem{
		index: 			idx,
		testItemAmount: newTestItemAmount(itemPath, amount),
		chance: 		chance,
	}
}

func (t testPossibleItem) GetIndex() int {
	return t.index
}

func comparePossibleItems(test test, exp testPossibleItem, got PossibleItem) {
	compareItemAmounts(test, exp.testItemAmount, got.ItemAmount)
	compare(test, "possible item - chance", exp.chance, got.Chance)
}