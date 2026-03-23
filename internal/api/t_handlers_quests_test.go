package api

import (
	"net/http"
	"testing"
)

func TestGetQuest(t *testing.T) {
	tests := []expQuest{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests/66",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "quest with provided id '66' doesn't exist. max id: 65.",
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
				name: "valefor",
			},
			questType: 	2,
			typedQuest: "/subquests/37",
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests/55",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 2,
				},
			},
			expUnique: expUnique{
				id:   55,
				name: "guardian spheres",
			},
			questType: 	2,
			typedQuest: "/subquests/50",
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
			count:   65,
			results: []int32{1, 3, 17, 18, 25, 37, 44, 58, 65},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?type=sidequest",
				expectedStatus: http.StatusOK,
			},
			count:   9,
			results: []int32{1, 38, 47, 49, 54, 56, 60, 63, 65},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?limit=max&post_airship=true",
				expectedStatus: http.StatusOK,
			},
			count:   38,
			results: []int32{1, 9, 15, 22, 31, 46, 54, 59, 65},
		},
	}

	testIdList(t, tests, testCfg.e.quests.endpoint, "RetrieveQuests", testCfg.HandleQuests, compareAPIResourceLists[TypedAPIResourceList])
}
