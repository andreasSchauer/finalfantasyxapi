package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareSidequests(test test, exp expSidequest, got api.Sidequest) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIDAPIResource(test, "untyped quest", test.cfg.E.Quests.Endpoint, exp.untypedQuest, got.UntypedQuest)
	compTestStructPtrs(test, "completion", exp.completion, got.Completion, compareQuestCompletions)
	checkResIDsInSlice(test, "subquests", test.cfg.E.Subquests.Endpoint, exp.subquests, got.Subquests)
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

func compareSubquests(test test, exp expSubquest, got api.Subquest) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIDAPIResource(test, "untyped quest", test.cfg.E.Quests.Endpoint, exp.untypedQuest, got.UntypedQuest)
	compIDAPIResource(test, "parent sidequest", test.cfg.E.Sidequests.Endpoint, exp.parentSidequest, got.ParentSidequest)
	compareQuestCompletions(test, "completion", exp.completion, got.Completion)
	compIDAPIResourcePtrs(test, "arena creation", test.cfg.E.ArenaCreations.Endpoint, exp.arenaCreation, got.ArenaCreation)
}

type testQuestCompletion struct {
	areas  []int32
	reward testResAmount[api.TypedAPIResource]
}

func compareQuestCompletions(test test, fieldName string, exp testQuestCompletion, got api.QuestCompletion) {
	test.t.Helper()
	checkResIDsInSlice(test, fieldName+" - areas", test.cfg.E.Areas.Endpoint, exp.areas, got.Areas)
	compareResAmounts(test, fieldName+" - rewards", exp.reward, got.Reward)
}
