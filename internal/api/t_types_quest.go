package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"

type expQuest struct {
	testGeneral
	expUnique
	questType  database.QuestType
	typedQuest string
}

func (e expQuest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareQuests(test test, exp expQuest, got Quest) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "quest type", string(exp.questType), string(got.Type))
	compPathApiResource(test, "typed quest", exp.typedQuest, got.TypedQuest)
}
