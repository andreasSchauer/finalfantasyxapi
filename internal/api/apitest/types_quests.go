package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
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

type Sidequest struct {
	TestGeneral
	Unique
	untypedQuest int32
	completion   *TestQuestCompletion
	subquests    []int32
}

func (e Sidequest) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type Subquest struct {
	TestGeneral
	Unique
	untypedQuest    int32
	parentSidequest int32
	completion      TestQuestCompletion
	arenaCreation   *int32
}

func (e Subquest) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestQuestCompletion struct {
	areas  []int32
	reward TestResAmount[apiresource.Typed]
}
