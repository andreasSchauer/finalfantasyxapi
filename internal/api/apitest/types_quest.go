package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Quest struct {
	testGeneral
	Unique
	questType  database.QuestType
	typedQuest string
}

func (e Quest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareQuests(test test, exp Quest, got apitypes.Quest) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("quest type", string(exp.questType), string(got.Type))
	test.compPathAPIResource("typed quest", exp.typedQuest, got.TypedQuest)
}
