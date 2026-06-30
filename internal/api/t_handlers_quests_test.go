package api

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func TestGetQuest(t *testing.T) {
	t.Parallel()
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
				expLengths:     map[string]int{},
			},
			expUnique: expUnique{
				id:   1,
				name: "monster arena",
			},
			questType:  database.QuestTypeSidequest,
			typedQuest: "/sidequests/1",
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests/47",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths:     map[string]int{},
			},
			expUnique: expUnique{
				id:   47,
				name: "valefor - first win",
			},
			questType:  database.QuestTypeSubquest,
			typedQuest: "/subquests/37",
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests/80",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 2,
				},
			},
			expUnique: expUnique{
				id:   80,
				name: "6 or 7 guardian spheres",
			},
			questType:  database.QuestTypeSubquest,
			typedQuest: "/subquests/70",
		},
	}

	testSingleResources(t, tests, "GetQuests", testCfg.HandleQuests, compareQuests)
}

func TestRetrieveQuests(t *testing.T) {
	t.Parallel()
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
			results: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?repeatable=true",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{48, 54, 60, 70, 76},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?limit=max&availability=post",
				expectedStatus: http.StatusOK,
			},
			count:   44,
			results: []int32{1, 6, 11, 24, 34, 46, 62, 80, 87},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?availability=post&repeatable=true",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{58, 62},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?availability=always&repeatable=false&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   43,
			results: []int32{3, 10, 38, 59, 77, 92, 98},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?availability=pre-story",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{97},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/quests?availability=always&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   53,
			results: []int32{3, 20, 55, 68, 82, 94, 98},
		},
	}

	testIdList(t, tests, testCfg.e.quests.endpoint, "RetrieveQuests", testCfg.HandleQuests, compareAPIResourceLists[TypedAPIResourceList])
}
