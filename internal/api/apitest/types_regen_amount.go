package apitest

type TestRegenAmount struct {
	index    int
	ticks    int32
	healedHP int32
}

func (t TestRegenAmount) GetIndex() int {
	return t.index
}
