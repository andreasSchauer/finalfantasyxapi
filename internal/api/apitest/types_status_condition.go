package apitest

type StatusCondition struct {
	TestGeneral
	Unique
	autoAbilities      []int32
	inflictedBy        *TestStatusInfliction
	removedBy          *TestStatusRemoval
	monstersResistance []int32
}

func (e StatusCondition) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestStatusInfliction struct {
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	miscAbilities      []int32
	enemyAbilities     []int32
}

type TestStatusRemoval struct {
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	enemyAbilities     []int32
	statusConditions   []int32
}
