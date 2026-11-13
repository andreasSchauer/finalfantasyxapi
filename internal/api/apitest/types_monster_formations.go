package apitest

type MonsterFormation struct {
	TestGeneral
	IDOnly
	category        string
	isForcedAmbush  bool
	canEscape       bool
	bossMusic       *int32
	ap              int32
	apOverkill      int32
	gil             int32
	monsters        map[string]int32
	areas           []int32
	triggerCommands []TestFormationTC
}

func (e MonsterFormation) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestFormationTC struct {
	Ability int32
	Users   []int32
}
