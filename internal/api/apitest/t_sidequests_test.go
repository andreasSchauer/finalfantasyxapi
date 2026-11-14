package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestGetSidequest(t *testing.T) {
	t.Parallel()
	tests := []Sidequest{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/11",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sidequest with provided id '11' doesn't exist. max id: 10",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/11/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sidequest with provided id '11' doesn't exist. max id: 10",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sidequest not found: 'a'",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/10/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "subsection 'a' does not exist for endpoint /sidequests. supported subsections: 'subquests'",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/2",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"completion - areas": 1,
					"subquests":          16,
				},
			},
			Unique:       newExpUnique(2, "remiem temple"),
			untypedQuest: 2,
			completion: &testQuestCompletion{
				areas:  []int32{207},
				reward: newTestResAmount[apiresource.Typed](126, 1),
			},
			subquests: []int32{37, 39, 42, 44, 45, 49, 52},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/4",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"subquests": 9,
				},
			},
			Unique:       newExpUnique(4, "chocobo training"),
			untypedQuest: 4,
			completion:   nil,
			subquests:    []int32{59, 62, 64, 65, 67},
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
			Unique:       newExpUnique(9, "al bhed primers"),
			untypedQuest: 9,
			completion: &testQuestCompletion{
				areas:  []int32{183, 180},
				reward: newTestResAmount[apiresource.Typed](111, 99),
			},
			subquests: []int32{},
		},
	}

	testSingleResources(t, tests, "GetSidequest", testSrv.HandleSidequests)
}

func TestRetrieveSidequests(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests",
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{1, 2, 3, 5, 6, 9, 10},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests?availability=post",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{1, 2, 5, 6, 9},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests?availability=always",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{3, 4, 7, 8, 10},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests?availability=post",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{1, 2, 5, 6, 9},
		},
	}

	testIDList(t, tests, apilookups.E.Sidequests.Endpoint, "RetrieveSidequests", testSrv.HandleSidequests, compareAPIResourceLists[apiresource.Named])
}
