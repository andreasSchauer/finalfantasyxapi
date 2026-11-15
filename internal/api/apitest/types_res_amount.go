package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type TestResAmount[A apiresource.APIResource] struct {
	Item   int32
	Amount int32
}

func (e TestResAmount[A]) Compare(test Test, fieldName string, got apiresource.ResourceAmount[A]) {
	test.T.Helper()
	test.compIDAPIResource(fieldName+" - resource", apilookups.E.AllItems.Endpoint, e.Item, got)
	test.compare(fieldName+" - amount", e.Amount, got.Amount)
}

func NewTestResAmount[A apiresource.APIResource](itemID int32, amount int32) TestResAmount[A] {
	return TestResAmount[A]{
		Item:   itemID,
		Amount: amount,
	}
}

type TestPossibleItem struct {
	Index int
	TestResAmount[apiresource.Typed]
	Chance int32
}

func NewTestPossibleItem(idx int, itemID int32, amount, chance int32) TestPossibleItem {
	return TestPossibleItem{
		Index:         idx,
		TestResAmount: NewTestResAmount[apiresource.Typed](itemID, amount),
		Chance:        chance,
	}
}

func (e TestPossibleItem) GetIndex() int {
	return e.Index
}

func (e TestPossibleItem) Compare(test Test, fieldName string, got apitypes.PossibleItem) {
	test.T.Helper()
	e.TestResAmount.Compare(test, fieldName+" - itemAmount", got.ResourceAmount)
	test.compare(fieldName+" - chance", e.Chance, got.Chance)
}
