package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Primer struct {
	testGeneral
	Unique
	keyItem   int32
	areas     []int32
	treasures []int32
}

func (e Primer) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Primer) Compare(test test, got apitypes.Primer) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("key item", apilookups.E.KeyItems.Endpoint, e.keyItem, got.KeyItem)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, e.areas, got.Areas)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.treasures, got.Treasures)
}
