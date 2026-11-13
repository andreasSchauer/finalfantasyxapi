package apitest

type Stat struct {
	TestGeneral
	Unique
	spheres            []int32
	autoAbilities      []int32
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	triggerCommands    []int32
	statusConditions   []int32
	properties         []int32
}

func (e Stat) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
