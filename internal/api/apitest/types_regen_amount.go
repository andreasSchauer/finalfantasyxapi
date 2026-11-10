package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type testRegenAmount struct {
	index    int
	ticks    int32
	healedHP int32
}

func (t testRegenAmount) GetIndex() int {
	return t.index
}

func compareRegenAmounts(test test, fieldName string, exp testRegenAmount, got apitypes.RegenAmount) {
	test.t.Helper()
	test.compare(fieldName+" - ticks", exp.ticks, got.Ticks)
	test.compare(fieldName+" - healed hp", exp.healedHP, got.HealedHP)
}
