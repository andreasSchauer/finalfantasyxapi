package apitest

type CelestialWeapon struct {
	TestGeneral
	Unique
	equipment     int32
	autoAbilities []int32
	crest         int32
	sigil         int32
	wpnTreasure   int32
	crestTreasure int32
	sigilQuest    int32
}

func (e CelestialWeapon) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
