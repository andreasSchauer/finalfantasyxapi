package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareQuests(test Test, exp Quest, got apitypes.Quest) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("quest type", string(exp.questType), string(got.Type))
	test.compPathAPIResource("typed quest", exp.typedQuest, got.TypedQuest)
}

func compareSidequests(test Test, exp Sidequest, got apitypes.Sidequest) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped quest", apilookups.E.Quests.Endpoint, exp.untypedQuest, got.UntypedQuest)
	test.compTestStructPtrs("completion", exp.completion, got.Completion, compareQuestCompletions)
	test.checkResIDsInSlice("subquests", apilookups.E.Subquests.Endpoint, exp.subquests, got.Subquests)
}

func compareSubquests(test Test, exp Subquest, got apitypes.Subquest) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped quest", apilookups.E.Quests.Endpoint, exp.untypedQuest, got.UntypedQuest)
	test.compIDAPIResource("parent sidequest", apilookups.E.Sidequests.Endpoint, exp.parentSidequest, got.ParentSidequest)
	compareQuestCompletions(test, "completion", exp.completion, got.Completion)
	test.compIDAPIResourcePtrs("arena creation", apilookups.E.ArenaCreations.Endpoint, exp.arenaCreation, got.ArenaCreation)
}

func compareQuestCompletions(test Test, fieldName string, exp TestQuestCompletion, got apitypes.QuestCompletion) {
	test.t.Helper()
	test.checkResIDsInSlice(fieldName+" - areas", apilookups.E.Areas.Endpoint, exp.areas, got.Areas)
	compareResAmounts(test, fieldName+" - rewards", exp.reward, got.Reward)
}
