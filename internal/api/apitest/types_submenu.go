package apitest

type Submenu struct {
	TestGeneral
	Unique
	topmenu   *int32
	users     []int32
	abilities []int32
	openedBy  *MenuOpen
}

func (e Submenu) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type MenuOpen struct {
	ability           *int32
	aeonCommand       *int32
	overdriveCommands []int32
}
