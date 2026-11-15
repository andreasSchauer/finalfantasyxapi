package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Quest struct {
	TestGeneral
	Unique
	QuestType  database.QuestType
	TypedQuest string
}

func (e Quest) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Quest) Compare(test Test, got apitypes.Quest) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("quest type", string(e.QuestType), string(got.Type))
	test.compPathAPIResource("typed quest", e.TypedQuest, got.TypedQuest)
}
