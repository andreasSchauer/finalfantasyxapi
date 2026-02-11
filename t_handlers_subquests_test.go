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
				requestURL: 	"/api/subquests/57",
				expectedStatus: http.StatusNotFound,
				expectedErr: 	"subquest with provided id '57' doesn't exist. max id: 56.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/subquests/10/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: 	"endpoint 'subquests' doesn't have any subsections.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/subquests/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: 	"invalid id 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/subquests/5",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{
					"completions": 1,
				},
			},
			expUnique: newExpUnique(5, "jormungand"),
			parentSidequest: 1,
			completions: []testQuestCompletion{
				{
					index: 0,
					areas: []int32{205},
					reward: newTestItemAmount("/items/50", 99),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/subquests/39",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{
					"completions": 2,
				},
			},
			expUnique: newExpUnique(39, "ixion"),
			parentSidequest: 2,
			completions: []testQuestCompletion{
				{
					index: 0,
					areas: []int32{209},
					reward: newTestItemAmount("/items/54", 10),
				},
				{
					index: 1,
					areas: []int32{209},
					reward: newTestItemAmount("/items/70", 8),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/subquests/53",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{
					"completions": 2,
				},
			},
			expUnique: newExpUnique(53, "after obtaining the airship"),
			parentSidequest: 6,
			completions: []testQuestCompletion{
				{
					index: 0,
					areas: []int32{144, 145},
					reward: newTestItemAmount("/items/98", 1),
				},
				{
					index: 1,
					areas: []int32{144, 145},
					reward: newTestItemAmount("/key-items/18", 1),
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
			count:   56,
			results: []int32{1, 5, 17, 28, 30, 34, 44, 51, 56},
		},
	}

	testIdList(t, tests, testCfg.e.subquests.endpoint, "RetrieveSubquests", testCfg.HandleSubquests, compareAPIResourceLists[NamedApiResourceList])
}
