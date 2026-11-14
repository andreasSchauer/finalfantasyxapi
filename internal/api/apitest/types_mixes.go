package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Mix struct {
	testGeneral
	Unique
	category     database.MixCategory
	overdrive    int32
	combinations []testMixCombination
}

func (e Mix) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Mix) Compare(test test, got apitypes.Mix) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("category", string(e.category), got.Category)
	test.compIDAPIResource("overdrive", apilookups.E.Overdrives.Endpoint, e.overdrive, got.Overdrive)
	test.checkTestStructsInSlice("combinations", e.combinations, got.Combinations)
}

type testMixCombination struct {
	index      int
	firstItem  int32
	secondItem int32
}

func (e testMixCombination) GetIndex() int {
	return e.index
}

func (e testMixCombination) Compare(test test, fieldName string, got apitypes.MixCombination) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - first item", apilookups.E.Items.Endpoint, e.firstItem, got.FirstItem)
	test.compIDAPIResource(fieldName+" - second item", apilookups.E.Items.Endpoint, e.secondItem, got.SecondItem)
}
