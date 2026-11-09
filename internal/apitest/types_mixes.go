package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
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

func compareMixes(test test, exp expMix, got apitypes.Mix) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compare("category", string(exp.category), got.Category)
	test.compIDAPIResource("overdrive", apilookups.E.Overdrives.Endpoint, exp.overdrive, got.Overdrive)
	test.checkTestStructsInSlice("combinations", exp.combinations, got.Combinations, compareMixCombinations)
}

type testMixCombination struct {
	index      int
	firstItem  int32
	secondItem int32
}

func (mc testMixCombination) GetIndex() int {
	return mc.index
}

func compareMixCombinations(test test, fieldName string, exp testMixCombination, got apitypes.MixCombination) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - first item", apilookups.E.Items.Endpoint, exp.firstItem, got.FirstItem)
	test.compIDAPIResource(fieldName+" - second item", apilookups.E.Items.Endpoint, exp.secondItem, got.SecondItem)
}
