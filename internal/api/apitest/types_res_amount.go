package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type testResAmount[A apiresource.APIResource] struct {
	Item   int32
	Amount int32
}

func (e testResAmount[A]) Compare(test Test, fieldName string, got apiresource.ResourceAmount[A]) {
	test.T.Helper()
	test.compIDAPIResource(fieldName+" - resource", apilookups.E.AllItems.Endpoint, e.Item, got)
	test.compare(fieldName+" - amount", e.Amount, got.Amount)
}

func newTestResAmount[A apiresource.APIResource](itemID int32, amount int32) testResAmount[A] {
	return testResAmount[A]{
		Item:   itemID,
		Amount: amount,
	}
}

type testPossibleItem struct {
	Index int
	testResAmount[apiresource.Typed]
	Chance int32
}

func newTestPossibleItem(idx int, itemID int32, amount, chance int32) testPossibleItem {
	return testPossibleItem{
		Index:         idx,
		testResAmount: newTestResAmount[apiresource.Typed](itemID, amount),
		Chance:        chance,
	}
}

func (e testPossibleItem) GetIndex() int {
	return e.Index
}

func (e testPossibleItem) Compare(test Test, fieldName string, got apitypes.PossibleItem) {
	test.T.Helper()
	e.testResAmount.Compare(test, fieldName+" - itemAmount", got.ResourceAmount)
	test.compare(fieldName+" - chance", e.Chance, got.Chance)
}
