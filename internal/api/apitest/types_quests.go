package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Sidequest struct {
	testGeneral
	Unique
	untypedQuest int32
	completion   *testQuestCompletion
	subquests    []int32
}

func (e Sidequest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSidequests(test test, exp Sidequest, got apitypes.Sidequest) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped quest", apilookups.E.Quests.Endpoint, exp.untypedQuest, got.UntypedQuest)
	test.compTestStructPtrs("completion", exp.completion, got.Completion, compareQuestCompletions)
	test.checkResIDsInSlice("subquests", apilookups.E.Subquests.Endpoint, exp.subquests, got.Subquests)
}

type Subquest struct {
	testGeneral
	Unique
	untypedQuest    int32
	parentSidequest int32
	completion      testQuestCompletion
	arenaCreation   *int32
}

func (e Subquest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSubquests(test test, exp Subquest, got apitypes.Subquest) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped quest", apilookups.E.Quests.Endpoint, exp.untypedQuest, got.UntypedQuest)
	test.compIDAPIResource("parent sidequest", apilookups.E.Sidequests.Endpoint, exp.parentSidequest, got.ParentSidequest)
	compareQuestCompletions(test, "completion", exp.completion, got.Completion)
	test.compIDAPIResourcePtrs("arena creation", apilookups.E.ArenaCreations.Endpoint, exp.arenaCreation, got.ArenaCreation)
}

type testQuestCompletion struct {
	areas  []int32
	reward testResAmount[apiresource.Typed]
}

func compareQuestCompletions(test test, fieldName string, exp testQuestCompletion, got apitypes.QuestCompletion) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - areas", apilookups.E.Areas.Endpoint, exp.areas, got.Areas)
	compareResAmounts(test, fieldName+" - rewards", exp.reward, got.Reward)
}
