package apitest

type OverdriveCommand struct {
	TestGeneral
	Unique
	rank        int32
	user        int32
	topmenu     *int32
	openSubmenu int32
	overdrives  []int32
}

func (e OverdriveCommand) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
