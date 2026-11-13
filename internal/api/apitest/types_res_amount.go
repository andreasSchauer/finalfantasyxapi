package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

type TestResAmount[A apiresource.APIResource] struct {
	item   int32
	amount int32
}

func newTestResAmount[A apiresource.APIResource](itemID int32, amount int32) TestResAmount[A] {
	return TestResAmount[A]{
		item:   itemID,
		amount: amount,
	}
}

type TestPossibleItem struct {
	index int
	TestResAmount[apiresource.Typed]
	chance int32
}

func newTestPossibleItem(idx int, itemID int32, amount, chance int32) TestPossibleItem {
	return TestPossibleItem{
		index:         idx,
		TestResAmount: newTestResAmount[apiresource.Typed](itemID, amount),
		chance:        chance,
	}
}

func (t TestPossibleItem) GetIndex() int {
	return t.index
}
