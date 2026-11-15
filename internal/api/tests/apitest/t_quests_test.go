package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func TestGetQuest(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Quest{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests/99",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "quest with provided id '99' doesn't exist. max id: 98",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests/1",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			Unique: apitesttypes.Unique{
				ID:   1,
				Name: "monster arena",
			},
			QuestType:  database.QuestTypeSidequest,
			TypedQuest: "/sidequests/1",
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests/47",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			Unique: apitesttypes.Unique{
				ID:   47,
				Name: "valefor - first win",
			},
			QuestType:  database.QuestTypeSubquest,
			TypedQuest: "/subquests/37",
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests/80",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 2,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   80,
				Name: "6 or 7 guardian spheres",
			},
			QuestType:  database.QuestTypeSubquest,
			TypedQuest: "/subquests/70",
		},
	}

	testSingleResources(t, tests, "GetQuests", testSrv.HandleQuests)
}

func TestRetrieveQuests(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   98,
			Results: []int32{1, 3, 17, 18, 25, 37, 44, 58, 65, 77, 83, 91, 98},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests?type=sidequest",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests?repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{48, 54, 60, 70, 76},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests?limit=max&availability=post",
				ExpectedStatus: http.StatusOK,
			},
			Count:   44,
			Results: []int32{1, 6, 11, 24, 34, 46, 62, 80, 87},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests?availability=post&repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{58, 62},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests?availability=always&repeatable=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   43,
			Results: []int32{3, 10, 38, 59, 77, 92, 98},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests?availability=pre-story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{97},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/quests?availability=always&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   53,
			Results: []int32{3, 20, 55, 68, 82, 94, 98},
		},
	}

	testIDList(t, tests, apilookups.E.Quests.Endpoint, "RetrieveQuests", testSrv.HandleQuests, compareAPIResourceLists[apiresource.Typed])
}
