package apitest

type OverdriveMode struct {
	TestGeneral
	Unique
	description   string
	effect        string
	modeType      string
	fillRate      *float32
	actionsAmount map[string]int32
}

func (e OverdriveMode) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
