package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type testResAmount[A apiresource.APIResource] struct {
	item   int32
	amount int32
}

func (e testResAmount[A]) Compare(test test, fieldName string, got apiresource.ResourceAmount[A]) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - resource", apilookups.E.AllItems.Endpoint, e.item, got)
	test.compare(fieldName+" - amount", e.amount, got.Amount)
}

func newTestResAmount[A apiresource.APIResource](itemID int32, amount int32) testResAmount[A] {
	return testResAmount[A]{
		item:   itemID,
		amount: amount,
	}
}

type testPossibleItem struct {
	index int
	testResAmount[apiresource.Typed]
	chance int32
}

func newTestPossibleItem(idx int, itemID int32, amount, chance int32) testPossibleItem {
	return testPossibleItem{
		index:         idx,
		testResAmount: newTestResAmount[apiresource.Typed](itemID, amount),
		chance:        chance,
	}
}

func (e testPossibleItem) GetIndex() int {
	return e.index
}

func (e testPossibleItem) Compare(test test, fieldName string, got apitypes.PossibleItem) {
	test.t.Helper()
	e.testResAmount.Compare(test, fieldName+" - itemAmount", got.ResourceAmount)
	test.compare(fieldName+" - chance", e.chance, got.Chance)
}
