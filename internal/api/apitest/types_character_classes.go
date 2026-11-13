package apitest

type CharacterClass struct {
	TestGeneral
	Unique
	category            string
	members             []int32
	defaultAbilities    []int32
	learnableAbilities  []int32
	defaultOverdrives   []int32
	learnableOverdrives []int32
	submenus            []int32
}

func (e CharacterClass) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
