package apitest

type Aeon struct {
	TestGeneral
	Unique
	untypedUnit            int32
	area                   int32
	battlesToRegen         int32
	agility                TestAgilityParams
	celestialWeapon        *int32
	characterClasses       []int32
	baseStats              map[string]int32
	regenAmounts           []TestRegenAmount
	aeonCommands           []int32
	defaultPlayerAbilities []int32
	overdrives             []int32
	weaponAbilities        []AeonEquipment
	armorAbilities         []AeonEquipment
}

func (e Aeon) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type AeonEquipment struct {
	autoAbility     int32
	celestialWeapon bool
}
