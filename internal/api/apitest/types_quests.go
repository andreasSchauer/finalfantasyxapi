package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Sidequest struct {
	TestGeneral
	Unique
	untypedQuest int32
	completion   *testQuestCompletion
	subquests    []int32
}

func (e Sidequest) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Sidequest) Compare(test Test, got apitypes.Sidequest) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped quest", apilookups.E.Quests.Endpoint, e.untypedQuest, got.UntypedQuest)
	test.compTestStructPtrs("completion", e.completion, got.Completion)
	test.checkResIDsInSlice("subquests", apilookups.E.Subquests.Endpoint, e.subquests, got.Subquests)
}

type Subquest struct {
	TestGeneral
	Unique
	untypedQuest    int32
	parentSidequest int32
	completion      testQuestCompletion
	arenaCreation   *int32
}

func (e Subquest) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Subquest) Compare(test Test, got apitypes.Subquest) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped quest", apilookups.E.Quests.Endpoint, e.untypedQuest, got.UntypedQuest)
	test.compIDAPIResource("parent sidequest", apilookups.E.Sidequests.Endpoint, e.parentSidequest, got.ParentSidequest)
	e.completion.Compare(test, "completion", got.Completion)
	test.compIDAPIResourcePtrs("arena creation", apilookups.E.ArenaCreations.Endpoint, e.arenaCreation, got.ArenaCreation)
}

type testQuestCompletion struct {
	areas  []int32
	reward testResAmount[apiresource.Typed]
}

func (e testQuestCompletion) Compare(test Test, fieldName string, got apitypes.QuestCompletion) {
	test.T.Helper()
	test.checkResIDsInSlice(fieldName+" - areas", apilookups.E.Areas.Endpoint, e.areas, got.Areas)
	e.reward.Compare(test, fieldName+" - rewards", got.Reward)
}
