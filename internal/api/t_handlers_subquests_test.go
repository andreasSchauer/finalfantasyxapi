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
				requestURL:     "/api/subquests/89",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "subquest with provided id '89' doesn't exist. max id: 88.",
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
			untypedQuest:    6,
			parentSidequest: 1,
			completion: testQuestCompletion{
				areas:  []int32{205},
				reward: newTestResAmount[TypedAPIResource](50, 99),
			},
			arenaCreation: h.GetInt32Ptr(5),
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
			expUnique:       newExpUnique(39, "ifrit - first win"),
			untypedQuest:    41,
			parentSidequest: 2,
			completion: testQuestCompletion{
				areas:  []int32{209},
				reward: newTestResAmount[TypedAPIResource](3, 30),
			},
			arenaCreation: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/76",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"completions": 2,
				},
			},
			expUnique:       newExpUnique(76, "first hunt after acquiring the airship"),
			untypedQuest:    82,
			parentSidequest: 6,
			completion: testQuestCompletion{
				areas:  []int32{144, 145},
				reward: newTestResAmount[TypedAPIResource](98, 1),
			},
			arenaCreation: nil,
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
			count:   88,
			results: []int32{1, 5, 17, 28, 30, 34, 44, 51, 56, 66, 72, 79, 88},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests?availability=post&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   39,
			results: []int32{1, 8, 12, 23, 33, 52, 68, 71, 76},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests?repeatable=true&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{38, 42, 44, 50, 52, 60, 64, 66},
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
			count:          16,
			parentResource: h.GetStrPtr("/sidequests/2"),
			results:        []int32{37, 40, 41, 43, 44, 46, 49, 52},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/4/subquests",
				expectedStatus: http.StatusOK,
			},
			count:          9,
			parentResource: h.GetStrPtr("/sidequests/4"),
			results:        []int32{59, 62, 66, 67},
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

	testIdList(t, tests, testCfg.e.subquests.endpoint, "SubsectionSubquests", testCfg.HandleSidequests, compareSimpleResourceLists[NamedAPIResource, SubquestSimple])
}
