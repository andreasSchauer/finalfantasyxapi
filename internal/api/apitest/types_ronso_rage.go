package apitest

type RonsoRage struct {
	TestGeneral
	Unique
	overdrive int32
	monsters  []int32
}

func (e RonsoRage) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
