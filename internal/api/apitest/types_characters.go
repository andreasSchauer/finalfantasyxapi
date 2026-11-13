package apitest

type Character struct {
	TestGeneral
	Unique
	untypedUnit            int32
	area                   int32
	weaponType             string
	celestialWeapon        *int32
	celestialFormula       *string
	overdriveCommand       *int32
	characterClasses       []int32
	agility                TestAgilityParams
	baseStats              map[string]int32
	regenAmounts           []TestRegenAmount
	defaultPlayerAbilities []int32
	stdSphereGrid          *SphereGrid
	expSphereGrid          *SphereGrid
	overdriveModes         map[string]int32
}

func (e Character) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type SphereGrid struct {
	playerAbilities []int32
}
