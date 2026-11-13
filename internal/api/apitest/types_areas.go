package apitest

type Area struct {
	TestGeneral
	NameVer
	displayName       string
	parentLocation    int32
	parentSublocation int32
	connectedAreas    []int32
	LocRel
}

func (e Area) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
