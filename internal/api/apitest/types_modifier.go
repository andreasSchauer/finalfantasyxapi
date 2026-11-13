package apitest

type Modifier struct {
	TestGeneral
	Unique
	autoAbilities      []int32
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	triggerCommands    []int32
	enemyAbilities     []int32
	statusConditions   []int32
	properties         []int32
}

func (e Modifier) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
