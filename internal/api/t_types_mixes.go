package api

type expMix struct {
	testGeneral
	expUnique
	category     int32
	overdrive    int32
	combinations []testMixCombination
}

func (e expMix) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMixes(test test, exp expMix, got Mix) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "category", test.cfg.e.mixCategory.endpoint, exp.category, got.Category)
	compIdApiResource(test, "overdrive", test.cfg.e.overdrives.endpoint, exp.overdrive, got.Overdrive)
	checkTestStructsInSlice(test, "combinations", exp.combinations, got.Combinations, compareMixCombinations)
}

type testMixCombination struct {
	index      int
	firstItem  int32
	secondItem int32
}

func (mc testMixCombination) GetIndex() int {
	return mc.index
}

func compareMixCombinations(test test, fieldName string, exp testMixCombination, got MixCombination) {
	compIdApiResource(test, fieldName+" - first item", test.cfg.e.items.endpoint, exp.firstItem, got.FirstItem)
	compIdApiResource(test, fieldName+" - second item", test.cfg.e.items.endpoint, exp.secondItem, got.SecondItem)
}
