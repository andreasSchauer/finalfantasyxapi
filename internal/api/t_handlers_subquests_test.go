package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetSubquest(t *testing.T) {
	tests := []expSubquest{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/57",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "subquest with provided id '57' doesn't exist. max id: 56.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/10/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "endpoint /subquests doesn't have any subsections.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/5",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"completions": 1,
				},
			},
			expUnique:       newExpUnique(5, "jormungand"),
			parentSidequest: 1,
			completions: []testQuestCompletion{
				{
					index:  0,
					areas:  []int32{205},
					reward: newTestItemAmount("/items/50", 99),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/39",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"completions": 2,
				},
			},
			expUnique:       newExpUnique(39, "ixion"),
			parentSidequest: 2,
			completions: []testQuestCompletion{
				{
					index:  0,
					areas:  []int32{209},
					reward: newTestItemAmount("/items/54", 10),
				},
				{
					index:  1,
					areas:  []int32{209},
					reward: newTestItemAmount("/items/70", 8),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/53",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"completions": 2,
				},
			},
			expUnique:       newExpUnique(53, "after obtaining the airship"),
			parentSidequest: 6,
			completions: []testQuestCompletion{
				{
					index:  0,
					areas:  []int32{144, 145},
					reward: newTestItemAmount("/items/98", 1),
				},
				{
					index:  1,
					areas:  []int32{144, 145},
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

func TestSubsectionSubquests(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/1/subquests?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:          36,
			parentResource: h.GetStrPtr("/sidequests/1"),
			results:        []int32{1, 8, 12, 14, 18, 23, 31, 36},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/2/subquests",
				expectedStatus: http.StatusOK,
			},
			count:          8,
			parentResource: h.GetStrPtr("/sidequests/2"),
			results:        []int32{37, 40, 41, 43, 44},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/4/subquests",
				expectedStatus: http.StatusOK,
			},
			count:          4,
			parentResource: h.GetStrPtr("/sidequests/4"),
			results:        []int32{46, 47, 48, 49},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/9/subquests",
				expectedStatus: http.StatusOK,
			},
			count:          0,
			parentResource: h.GetStrPtr("/sidequests/9"),
			results:        []int32{},
		},
	}

	testIdList(t, tests, testCfg.e.subquests.endpoint, "SubsectionSubquests", testCfg.HandleSidequests, compareSubResourceLists[NamedAPIResource, SubquestSub])
}
