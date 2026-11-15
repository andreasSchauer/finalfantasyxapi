package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Quest struct {
	TestGeneral
	Unique
	questType  database.QuestType
	typedQuest string
}

func (e Quest) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Quest) Compare(test Test, got apitypes.Quest) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("quest type", string(e.questType), string(got.Type))
	test.compPathAPIResource("typed quest", e.typedQuest, got.TypedQuest)
}
