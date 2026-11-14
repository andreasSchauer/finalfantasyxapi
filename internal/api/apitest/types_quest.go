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

func (e Quest) Compare(test test, got apitypes.Quest) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("quest type", string(e.questType), string(got.Type))
	test.compPathAPIResource("typed quest", e.typedQuest, got.TypedQuest)
}
