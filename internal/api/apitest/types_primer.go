package apitest

type Primer struct {
	TestGeneral
	Unique
	keyItem   int32
	areas     []int32
	treasures []int32
}

func (e Primer) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
