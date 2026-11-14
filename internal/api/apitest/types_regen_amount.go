package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type testRegenAmount struct {
	index    int
	ticks    int32
	healedHP int32
}

func (e testRegenAmount) GetIndex() int {
	return e.index
}

func (e testRegenAmount) Compare(test test, fieldName string, got apitypes.RegenAmount) {
	test.t.Helper()
	test.compare(fieldName+" - ticks", e.ticks, got.Ticks)
	test.compare(fieldName+" - healed hp", e.healedHP, got.HealedHP)
}
