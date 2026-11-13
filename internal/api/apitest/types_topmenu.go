package apitest

type Topmenu struct {
	TestGeneral
	Unique
	submenus          []int32
	abilities         []int32
	overdriveCommands []int32
	overdrives        []int32
	aeonCommands      []int32
}

func (e Topmenu) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
