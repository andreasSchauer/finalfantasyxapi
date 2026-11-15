package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Mix struct {
	TestGeneral
	Unique
	Category     database.MixCategory
	Overdrive    int32
	Combinations []testMixCombination
}

func (e Mix) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Mix) Compare(test Test, got apitypes.Mix) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("category", string(e.Category), got.Category)
	test.compIDAPIResource("overdrive", apilookups.E.Overdrives.Endpoint, e.Overdrive, got.Overdrive)
	test.checkTestStructsInSlice("combinations", e.Combinations, got.Combinations)
}

type testMixCombination struct {
	Index      int
	FirstItem  int32
	SecondItem int32
}

func (e testMixCombination) GetIndex() int {
	return e.Index
}

func (e testMixCombination) Compare(test Test, fieldName string, got apitypes.MixCombination) {
	test.T.Helper()
	test.compIDAPIResource(fieldName+" - first item", apilookups.E.Items.Endpoint, e.FirstItem, got.FirstItem)
	test.compIDAPIResource(fieldName+" - second item", apilookups.E.Items.Endpoint, e.SecondItem, got.SecondItem)
}
