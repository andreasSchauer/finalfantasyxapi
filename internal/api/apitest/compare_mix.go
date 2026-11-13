package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareMixes(test Test, exp Mix, got apitypes.Mix) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("category", string(exp.category), got.Category)
	test.compIDAPIResource("overdrive", apilookups.E.Overdrives.Endpoint, exp.overdrive, got.Overdrive)
	test.checkTestStructsInSlice("combinations", exp.combinations, got.Combinations, compareMixCombinations)
}

func compareMixCombinations(test Test, fieldName string, exp TestMixCombination, got apitypes.MixCombination) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - first item", apilookups.E.Items.Endpoint, exp.firstItem, got.FirstItem)
	test.compIDAPIResource(fieldName+" - second item", apilookups.E.Items.Endpoint, exp.secondItem, got.SecondItem)
}
