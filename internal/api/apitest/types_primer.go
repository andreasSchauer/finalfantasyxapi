package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Primer struct {
	TestGeneral
	Unique
	KeyItem   int32
	Areas     []int32
	Treasures []int32
}

func (e Primer) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Primer) Compare(test Test, got apitypes.Primer) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("key item", apilookups.E.KeyItems.Endpoint, e.KeyItem, got.KeyItem)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, e.Areas, got.Areas)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.Treasures, got.Treasures)
}
