package apitest

type Equipment struct {
	TestGeneral
	Unique
	equipmentTable          int32
	priority                *int32
	celestialWeapon         *int32
	requiredAutoAbilities   []int32
	selectableAutoAbilities []TestAbilityPool
	requiredSlots           *int32
	treasures               []int32
	shops                   []int32
}

func (e Equipment) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
