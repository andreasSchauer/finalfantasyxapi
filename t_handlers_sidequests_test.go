package main

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetSidequest(t *testing.T) {
	tests := []expSidequest{
		{
			testGeneral: testGeneral{
				requestURL: "/api/sidequests/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/sidequests/0",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{},
			},
			expUnique: newExpUnique(0, ""),
			completion: &testQuestCompletion{
				index: 	0,
				areas: 	[]int32{},
				reward: newTestItemAmount("items/0", 0),
			},
			subquests: []int32{},
		},
	}

	testSingleResources(t, tests, "GetSidequest", testCfg.HandleSidequests, compareSidequests)
}

func TestRetrieveSidequests(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   0,
			results: []int32{},
		},
	}

	testIdList(t, tests, testCfg.e.sidequests.endpoint, "RetrieveSidequests", testCfg.HandleSidequests, compareAPIResourceLists[NamedApiResourceList])
}
