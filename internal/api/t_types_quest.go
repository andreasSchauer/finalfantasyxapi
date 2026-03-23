package api

type expQuest struct {
	testGeneral
	expUnique
	questType			int32
	typedQuest			string
}

func (e expQuest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareQuests(test test, exp expQuest, got Quest) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "quest type", test.cfg.e.questType.endpoint, exp.questType, got.Type)
	compPathApiResource(test, "typed quest", exp.typedQuest, got.TypedQuest)
}