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
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests/11",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "sidequest with provided id '11' doesn't exist. max id: 10",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests/11/a",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "sidequest with provided id '11' doesn't exist. max id: 10",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests/a",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "sidequest not found: 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests/10/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "subsection 'a' does not exist for endpoint /sidequests. supported subsections: 'subquests'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests/2",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"completion - areas": 1,
					"subquests":          16,
				},
			},
			Unique:       newExpUnique(2, "remiem temple"),
			UntypedQuest: 2,
			Completion: &testQuestCompletion{
				Areas:  []int32{207},
				Reward: newTestResAmount[apiresource.Typed](126, 1),
			},
			Subquests: []int32{37, 39, 42, 44, 45, 49, 52},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests/4",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"subquests": 9,
				},
			},
			Unique:       newExpUnique(4, "chocobo training"),
			UntypedQuest: 4,
			Completion:   nil,
			Subquests:    []int32{59, 62, 64, 65, 67},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests/9",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"completion - areas": 2,
					"subquests":          0,
				},
			},
			Unique:       newExpUnique(9, "al bhed primers"),
			UntypedQuest: 9,
			Completion: &testQuestCompletion{
				Areas:  []int32{183, 180},
				Reward: newTestResAmount[apiresource.Typed](111, 99),
			},
			Subquests: []int32{},
		},
	}

	testSingleResources(t, tests, "GetSidequest", testSrv.HandleSidequests)
}

func TestRetrieveSidequests(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{1, 2, 3, 5, 6, 9, 10},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests?availability=post",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{1, 2, 5, 6, 9},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests?availability=always",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{3, 4, 7, 8, 10},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sidequests?availability=post",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{1, 2, 5, 6, 9},
		},
	}

	testIDList(t, tests, apilookups.E.Sidequests.Endpoint, "RetrieveSidequests", testSrv.HandleSidequests, compareAPIResourceLists[apiresource.Named])
}
