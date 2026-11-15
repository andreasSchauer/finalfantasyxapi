package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Mix struct {
	TestGeneral
	Unique
	category     database.MixCategory
	overdrive    int32
	combinations []testMixCombination
}

func (e Mix) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Mix) Compare(test Test, got apitypes.Mix) {
	test.T.Helper()
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

func (e testMixCombination) Compare(test Test, fieldName string, got apitypes.MixCombination) {
	test.T.Helper()
	test.compIDAPIResource(fieldName+" - first item", apilookups.E.Items.Endpoint, e.firstItem, got.FirstItem)
	test.compIDAPIResource(fieldName+" - second item", apilookups.E.Items.Endpoint, e.secondItem, got.SecondItem)
}
