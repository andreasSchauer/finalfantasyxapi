package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expSidequest struct {
	testGeneral
	expUnique
	completion		*testQuestCompletion
	subquests		[]int32
}

func (e expSidequest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSidequests(test test, exp expSidequest, got Sidequest) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compTestStructPtrs(test, "completion", exp.completion, got.Completion, compareQuestCompletions)
	compareResListTest(test, rltIDs("subquests", test.cfg.e.subquests.endpoint, exp.subquests, got.Subquests))
}

type testQuestCompletion struct {
	index	int
	areas	[]int32
	reward	testItemAmount
}

func (t testQuestCompletion) GetIndex() int {
	return t.index
}

func compareQuestCompletions(test test, exp testQuestCompletion, got QuestCompletion) {
	compareResListTest(test, rltIDs("quest completion - areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas))
	compareItemAmounts(test, exp.reward, got.Reward)
}

func TestGetSidequest(t *testing.T) {
	tests := []expSidequest{}

	testSingleResources(t, tests, "GetSidequest", testCfg.HandleSidequests, compareSidequests)
}

func TestRetrieveSidequests(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.sidequests.endpoint, "RetrieveSidequests", testCfg.HandleSidequests, compareAPIResourceLists[NamedApiResourceList])
}
