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
				RequestURL:     "/api/quests/99",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "quest with provided id '99' doesn't exist. max id: 98",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests/1",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			Unique: Unique{
				ID:   1,
				Name: "monster arena",
			},
			questType:  database.QuestTypeSidequest,
			typedQuest: "/sidequests/1",
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests/47",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			Unique: Unique{
				ID:   47,
				Name: "valefor - first win",
			},
			questType:  database.QuestTypeSubquest,
			typedQuest: "/subquests/37",
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests/80",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 2,
				},
			},
			Unique: Unique{
				ID:   80,
				Name: "6 or 7 guardian spheres",
			},
			questType:  database.QuestTypeSubquest,
			typedQuest: "/subquests/70",
		},
	}

	testSingleResources(t, tests, "GetQuests", testSrv.HandleQuests)
}

func TestRetrieveQuests(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   98,
			Results: []int32{1, 3, 17, 18, 25, 37, 44, 58, 65, 77, 83, 91, 98},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests?type=sidequest",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests?repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{48, 54, 60, 70, 76},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests?limit=max&availability=post",
				ExpectedStatus: http.StatusOK,
			},
			Count:   44,
			Results: []int32{1, 6, 11, 24, 34, 46, 62, 80, 87},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests?availability=post&repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{58, 62},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests?availability=always&repeatable=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   43,
			Results: []int32{3, 10, 38, 59, 77, 92, 98},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests?availability=pre-story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{97},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/quests?availability=always&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   53,
			Results: []int32{3, 20, 55, 68, 82, 94, 98},
		},
	}

	testIDList(t, tests, apilookups.E.Quests.Endpoint, "RetrieveQuests", testSrv.HandleQuests, compareAPIResourceLists[apiresource.Typed])
}
