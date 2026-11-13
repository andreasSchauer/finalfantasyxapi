package apitest

type EquipmentTable struct {
	TestGeneral
	IDOnly
	celestialWeapon         *int32
	specificCharacter       *int32
	requiredAutoAbilities   []int32
	selectableAutoAbilities []TestAbilityPool
	requiredSlots           *int32
	equipment               []int32
}

func (e EquipmentTable) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestAbilityPool struct {
	index         int
	autoAbilities []int32
	reqAmount     int32
}
