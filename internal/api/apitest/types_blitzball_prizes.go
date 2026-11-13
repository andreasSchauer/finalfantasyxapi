package apitest

type BlitzballPrize struct {
	TestGeneral
	IDOnly
	category string
	slot     string
	items    []TestPossibleItem
}

func (e BlitzballPrize) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
