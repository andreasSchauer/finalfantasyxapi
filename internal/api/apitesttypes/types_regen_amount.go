package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type TestRegenAmount struct {
	Index    int
	Ticks    int32
	HealedHP int32
}

func (e TestRegenAmount) GetIndex() int {
	return e.Index
}

func (e TestRegenAmount) Compare(test Test, fieldName string, got apitypes.RegenAmount) {
	test.T.Helper()
	test.compare(fieldName+" - ticks", e.Ticks, got.Ticks)
	test.compare(fieldName+" - healed hp", e.HealedHP, got.HealedHP)
}
