package api

import (
	"net/http"
	"testing"
)

func TestGetQuest(t *testing.T) {
	tests := []expQuest{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests/99",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "quest with provided id '99' doesn't exist. max id: 98.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests/1",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{},
			},
			expUnique: expUnique{
				id:   1,
				name: "monster arena",
			},
			questType: 	1,
			typedQuest: "/sidequests/1",
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests/39",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{},
			},
			expUnique: expUnique{
				id:   39,
				name: "valefor - first win",
			},
			questType: 	2,
			typedQuest: "/subquests/37",
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests/75",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 2,
				},
			},
			expUnique: expUnique{
				id:   75,
				name: "6 or 7 guardian spheres",
			},
			questType: 	2,
			typedQuest: "/subquests/70",
		},
	}

	testSingleResources(t, tests, "GetQuests", testCfg.HandleQuests, compareQuests)
}

func TestRetrieveQuests(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   98,
			results: []int32{1, 3, 17, 18, 25, 37, 44, 58, 65, 77, 83, 91, 98},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?type=sidequest",
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{1, 38, 55, 62, 72, 77, 84, 94, 97, 98},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?limit=max&availability=post",
				expectedStatus: http.StatusOK,
			},
			count:   44,
			results: []int32{1, 18, 28, 33, 53, 72, 77, 83, 97},
		},
	}

	testIdList(t, tests, testCfg.e.quests.endpoint, "RetrieveQuests", testCfg.HandleQuests, compareAPIResourceLists[TypedAPIResourceList])
}
