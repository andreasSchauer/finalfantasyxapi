package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expSubquest struct {
	testGeneral
	expUnique
	parentSidequest	int32
	completions		[]testQuestCompletion
}

func (e expSubquest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSubquests(test test, exp expSubquest, got Subquest) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "parent sidequest", testCfg.e.sidequests.endpoint, exp.parentSidequest, got.ParentSidequest)
	compTestStructSlices(test, "completions", exp.completions, got.Completions, compareQuestCompletions)
}

func TestGetSubquest(t *testing.T) {
	tests := []expSubquest{}

	testSingleResources(t, tests, "GetSubquest", testCfg.HandleSubquests, compareSubquests)
}

func TestRetrieveSubquests(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.subquests.endpoint, "RetrieveSubquests", testCfg.HandleSubquests, compareAPIResourceLists[NamedApiResourceList])
}
