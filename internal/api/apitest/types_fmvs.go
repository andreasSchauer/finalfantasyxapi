package apitest

type FMV struct {
	TestGeneral
	Unique
	area int32
	song *int32
}

func (e FMV) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
