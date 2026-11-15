package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Sidequest struct {
	TestGeneral
	Unique
	UntypedQuest int32
	Completion   *TestQuestCompletion
	Subquests    []int32
}

func (e Sidequest) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Sidequest) Compare(test Test, got apitypes.Sidequest) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped quest", apilookups.E.Quests.Endpoint, e.UntypedQuest, got.UntypedQuest)
	test.compTestStructPtrs("completion", e.Completion, got.Completion)
	test.CheckResIDsInSlice("subquests", apilookups.E.Subquests.Endpoint, e.Subquests, got.Subquests)
}

type Subquest struct {
	TestGeneral
	Unique
	UntypedQuest    int32
	ParentSidequest int32
	Completion      TestQuestCompletion
	ArenaCreation   *int32
}

func (e Subquest) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Subquest) Compare(test Test, got apitypes.Subquest) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped quest", apilookups.E.Quests.Endpoint, e.UntypedQuest, got.UntypedQuest)
	test.compIDAPIResource("parent sidequest", apilookups.E.Sidequests.Endpoint, e.ParentSidequest, got.ParentSidequest)
	e.Completion.Compare(test, "completion", got.Completion)
	test.compIDAPIResourcePtrs("arena creation", apilookups.E.ArenaCreations.Endpoint, e.ArenaCreation, got.ArenaCreation)
}

type TestQuestCompletion struct {
	Areas  []int32
	Reward TestResAmount[apiresource.Typed]
}

func (e TestQuestCompletion) Compare(test Test, fn string, got apitypes.QuestCompletion) {
	test.T.Helper()
	test.CheckResIDsInSlice(childFn(fn, "areas"), apilookups.E.Areas.Endpoint, e.Areas, got.Areas)
	e.Reward.Compare(test, childFn(fn, "rewards"), got.Reward)
}
