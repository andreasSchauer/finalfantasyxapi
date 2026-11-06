package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

type expSidequest struct {
	testGeneral
	expUnique
	untypedQuest int32
	completion   *testQuestCompletion
	subquests    []int32
}

func (e expSidequest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSidequests(test test, exp expSidequest, got apitypes.Sidequest) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("untyped quest", api.E.Quests.Endpoint, exp.untypedQuest, got.UntypedQuest)
	test.compTestStructPtrs("completion", exp.completion, got.Completion, compareQuestCompletions)
	test.checkResIDsInSlice("subquests", api.E.Subquests.Endpoint, exp.subquests, got.Subquests)
}

type expSubquest struct {
	testGeneral
	expUnique
	untypedQuest    int32
	parentSidequest int32
	completion      testQuestCompletion
	arenaCreation   *int32
}

func (e expSubquest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSubquests(test test, exp expSubquest, got apitypes.Subquest) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("untyped quest", api.E.Quests.Endpoint, exp.untypedQuest, got.UntypedQuest)
	test.compIDAPIResource("parent sidequest", api.E.Sidequests.Endpoint, exp.parentSidequest, got.ParentSidequest)
	compareQuestCompletions(test, "completion", exp.completion, got.Completion)
	test.compIDAPIResourcePtrs("arena creation", api.E.ArenaCreations.Endpoint, exp.arenaCreation, got.ArenaCreation)
}

type testQuestCompletion struct {
	areas  []int32
	reward testResAmount[apiresource.TypedAPIResource]
}

func compareQuestCompletions(test test, fieldName string, exp testQuestCompletion, got apitypes.QuestCompletion) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - areas", api.E.Areas.Endpoint, exp.areas, got.Areas)
	compareResAmounts(test, fieldName+" - rewards", exp.reward, got.Reward)
}
