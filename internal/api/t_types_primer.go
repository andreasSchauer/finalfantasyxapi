package api

type expPrimer struct {
	testGeneral
	expUnique
	keyItem		int32
	areas		[]int32
	treasures	[]int32
}

func (p expPrimer) GetTestGeneral() testGeneral {
	return p.testGeneral
}

func comparePrimers(test test, exp expPrimer, got Primer) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "key item", test.cfg.e.keyItems.endpoint, exp.keyItem, got.KeyItem)
	checkResIDsInSlice(test, "areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas)
	checkResIDsInSlice(test, "treasures", test.cfg.e.treasures.endpoint, exp.treasures, got.Treasures)
}