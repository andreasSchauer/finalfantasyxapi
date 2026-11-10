package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type expPrimer struct {
	testGeneral
	expUnique
	keyItem   int32
	areas     []int32
	treasures []int32
}

func (e expPrimer) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func comparePrimers(test test, exp expPrimer, got apitypes.Primer) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("key item", apilookups.E.KeyItems.Endpoint, exp.keyItem, got.KeyItem)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, exp.areas, got.Areas)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, exp.treasures, got.Treasures)
}
