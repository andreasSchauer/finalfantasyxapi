package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func TestGetQuest(t *testing.T) {
	t.Parallel()
	tests := []Quest{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests/99",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "quest with provided id '99' doesn't exist. max id: 98",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests/1",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths:     map[string]int{},
			},
			Unique: Unique{
				id:   1,
				name: "monster arena",
			},
			questType:  database.QuestTypeSidequest,
			typedQuest: "/sidequests/1",
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests/47",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths:     map[string]int{},
			},
			Unique: Unique{
				id:   47,
				name: "valefor - first win",
			},
			questType:  database.QuestTypeSubquest,
			typedQuest: "/subquests/37",
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests/80",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 2,
				},
			},
			Unique: Unique{
				id:   80,
				name: "6 or 7 guardian spheres",
			},
			questType:  database.QuestTypeSubquest,
			typedQuest: "/subquests/70",
		},
	}

	testSingleResources(t, tests, "GetQuests", testSrv.HandleQuests, compareQuests)
}

func TestRetrieveQuests(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   98,
			results: []int32{1, 3, 17, 18, 25, 37, 44, 58, 65, 77, 83, 91, 98},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests?type=sidequest",
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests?repeatable=true",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{48, 54, 60, 70, 76},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests?limit=max&availability=post",
				expectedStatus: http.StatusOK,
			},
			count:   44,
			results: []int32{1, 6, 11, 24, 34, 46, 62, 80, 87},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests?availability=post&repeatable=true",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{58, 62},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests?availability=always&repeatable=false&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   43,
			results: []int32{3, 10, 38, 59, 77, 92, 98},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests?availability=pre-story",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{97},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/quests?availability=always&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   53,
			results: []int32{3, 20, 55, 68, 82, 94, 98},
		},
	}

	testIDList(t, tests, apilookups.E.Quests.Endpoint, "RetrieveQuests", testSrv.HandleQuests, compareAPIResourceLists[apiresource.Typed])
}
