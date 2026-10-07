package api

type testRegenAmount struct {
	index		int
	ticks		int32
	healedHP	int32
}

func (t testRegenAmount) GetIndex() int {
	return t.index
}

func compareRegenAmounts(test test, fieldName string, exp testRegenAmount, got RegenAmount) {
	test.t.Helper()
	compare(test, fieldName+" - ticks", exp.ticks, got.Ticks)
	compare(test, fieldName+" - healed hp", exp.healedHP, got.HealedHP)
}
