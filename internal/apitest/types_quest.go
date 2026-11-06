package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type expQuest struct {
	testGeneral
	expUnique
	questType  database.QuestType
	typedQuest string
}

func (e expQuest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareQuests(test test, exp expQuest, got apitypes.Quest) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compare("quest type", string(exp.questType), string(got.Type))
	test.compPathAPIResource("typed quest", exp.typedQuest, got.TypedQuest)
}
