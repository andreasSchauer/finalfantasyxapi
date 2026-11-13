package apitest

type AeonCommand struct {
	TestGeneral
	Unique
	user              int32
	topmenu           *int32
	openSubmenu       *int32
	possibleAbilities []PossibleAbilityList
}

func (e AeonCommand) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type PossibleAbilityList struct {
	user      int32
	abilities []int32
}
