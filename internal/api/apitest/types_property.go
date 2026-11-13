package apitest

type Property struct {
	TestGeneral
	Unique
	autoAbilities []int32
	monsters      []int32
}

func (e Property) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
