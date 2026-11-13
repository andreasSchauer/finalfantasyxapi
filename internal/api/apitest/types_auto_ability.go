package apitest

type AutoAbility struct {
	TestGeneral
	Unique
	monstersDrop     []int32
	monstersItems    []TestMonItemAmts
	shopsPreAirship  []int32
	shopsPostAirship []int32
	treasures        []int32
	equipmentTables  []int32
}

func (e AutoAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
