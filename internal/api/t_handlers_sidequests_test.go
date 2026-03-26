package api

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetSidequest(t *testing.T) {
	tests := []expSidequest{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/10",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sidequest with provided id '10' doesn't exist. max id: 9.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/10/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sidequest with provided id '10' doesn't exist. max id: 9.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sidequest not found: 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/9/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "subsection 'a' does not exist for endpoint /sidequests. supported subsections: 'subquests'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/2",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"completion - areas": 1,
					"subquests":          8,
				},
			},
			expUnique:    newExpUnique(2, "remiem temple"),
			untypedQuest: 38,
			completion: &testQuestCompletion{
				areas:  []int32{209},
				reward: newTestResAmount[TypedAPIResource](126, 1),
			},
			subquests: []int32{37, 39, 42, 44},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/4",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"subquests": 4,
				},
			},
			expUnique:    newExpUnique(4, "chocobo training"),
			untypedQuest: 49,
			completion:   nil,
			subquests:    []int32{46, 47, 48, 49},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/9",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"completion - areas": 2,
					"subquests":          0,
				},
			},
			expUnique:    newExpUnique(9, "al bhed primers"),
			untypedQuest: 65,
			completion: &testQuestCompletion{
				areas:  []int32{185, 182},
				reward: newTestResAmount[TypedAPIResource](111, 99),
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
				requestURL:     "/api/sidequests",
				expectedStatus: http.StatusOK,
			},
			count:   9,
			results: []int32{1, 2, 3, 5, 6, 9},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests?post_airship=true",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{1, 2, 5, 6, 9},
		},
	}

	testIdList(t, tests, testCfg.e.sidequests.endpoint, "RetrieveSidequests", testCfg.HandleSidequests, compareAPIResourceLists[NamedApiResourceList])
}
