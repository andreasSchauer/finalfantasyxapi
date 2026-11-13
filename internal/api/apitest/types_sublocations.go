package apitest

type Sublocation struct {
	TestGeneral
	Unique
	parentLocation        int32
	connectedSublocations []int32
	areas                 []int32
	LocRel
}

func (e Sublocation) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
