package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func comparePrimers(test test, exp expPrimer, got api.Primer) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIDAPIResource(test, "key item", test.cfg.E.KeyItems.Endpoint, exp.keyItem, got.KeyItem)
	checkResIDsInSlice(test, "areas", test.cfg.E.Areas.Endpoint, exp.areas, got.Areas)
	checkResIDsInSlice(test, "treasures", test.cfg.E.Treasures.Endpoint, exp.treasures, got.Treasures)
}
