package main

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetSubquest(t *testing.T) {
	tests := []expSubquest{
		{
			testGeneral: testGeneral{
				requestURL: "/api/subquests/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/subquests/0",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{},
			},
			expUnique: newExpUnique(0, ""),
			parentSidequest: 0,
			completions: []testQuestCompletion{
				{
					index: 0,
					areas: []int32{},
					reward: newTestItemAmount("/items/0", 0),
				},
			},
		},
	}

	testSingleResources(t, tests, "GetSubquest", testCfg.HandleSubquests, compareSubquests)
}

func TestRetrieveSubquests(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   0,
			results: []int32{},
		},
	}

	testIdList(t, tests, testCfg.e.subquests.endpoint, "RetrieveSubquests", testCfg.HandleSubquests, compareAPIResourceLists[NamedApiResourceList])
}
