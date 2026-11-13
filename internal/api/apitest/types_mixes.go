package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Mix struct {
	TestGeneral
	Unique
	category     database.MixCategory
	overdrive    int32
	combinations []TestMixCombination
}

func (e Mix) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestMixCombination struct {
	index      int
	firstItem  int32
	secondItem int32
}

func (mc TestMixCombination) GetIndex() int {
	return mc.index
}
