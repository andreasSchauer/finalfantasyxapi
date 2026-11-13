package apitest

type Location struct {
	TestGeneral
	Unique
	connectedLocations []int32
	sublocations       []int32
	LocRel
}

func (e Location) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
