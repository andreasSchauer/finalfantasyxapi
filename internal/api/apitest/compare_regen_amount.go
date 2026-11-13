package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareRegenAmounts(test Test, fieldName string, exp TestRegenAmount, got apitypes.RegenAmount) {
	test.t.Helper()
	test.compare(fieldName+" - ticks", exp.ticks, got.Ticks)
	test.compare(fieldName+" - healed hp", exp.healedHP, got.HealedHP)
}
