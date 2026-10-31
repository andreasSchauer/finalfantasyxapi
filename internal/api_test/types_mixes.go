package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type expMix struct {
	testGeneral
	expUnique
	category     database.MixCategory
	overdrive    int32
	combinations []testMixCombination
}

func (e expMix) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMixes(test test, exp expMix, got api.Mix) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "category", string(exp.category), got.Category)
	compIDAPIResource(test, "overdrive", test.cfg.E.Overdrives.Endpoint, exp.overdrive, got.Overdrive)
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

func compareMixCombinations(test test, fieldName string, exp testMixCombination, got api.MixCombination) {
	test.t.Helper()
	compIDAPIResource(test, fieldName+" - first item", test.cfg.E.Items.Endpoint, exp.firstItem, got.FirstItem)
	compIDAPIResource(test, fieldName+" - second item", test.cfg.E.Items.Endpoint, exp.secondItem, got.SecondItem)
}
