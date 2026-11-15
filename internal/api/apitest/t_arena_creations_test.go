package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetArenaCreation(t *testing.T) {
	t.Parallel()
	tests := []ArenaCreation{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/arena-creations/36",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "arena creation with provided id '36' doesn't exist. max id: 35",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/arena-creations/a",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "arena creation not found: 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/arena-creations/a/2",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. usage: '/api/arena-creations', '/api/arena-creations/{id}', '/api/arena-creations/{name}'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/arena-creations/13",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required monsters": 10,
				},
			},
			Unique:              NewExpUnique(13, "vorban"),
			Category:            "area",
			Monster:             268,
			ParentSubquest:      13,
			Reward:              NewTestResAmount[apiresource.Typed](107, 60),
			RequiredCatchAmount: 1,
			RequiredMonsters:    []int32{239, 243, 245, 248},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/arena-creations/one-eye",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required monsters": 5,
				},
			},
			Unique:              NewExpUnique(19, "one-eye"),
			Category:            "species",
			Monster:             274,
			ParentSubquest:      19,
			Reward:              NewTestResAmount[apiresource.Typed](64, 60),
			RequiredCatchAmount: 4,
			RequiredMonsters:    []int32{40, 73, 87, 170, 240},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/arena-creations/31",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"required monsters": 0,
				},
			},
			Unique:                    NewExpUnique(31, "th'uban"),
			Category:                  "original",
			Monster:                   286,
			ParentSubquest:            31,
			Reward:                    NewTestResAmount[apiresource.Typed](110, 99),
			RequiredCatchAmount:       6,
			UnlockedCreationsCategory: h.GetStrPtr("species"),
			RequiredMonsters:          nil,
		},
	}

	testSingleResources(t, tests, "GetArenaCreation", testSrv.HandleArenaCreations)
}

func TestRetrieveArenaCreations(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/arena-creations?category=asd",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid enum value 'asd' used for parameter 'category'. use /api/arena-creations/parameters to see allowed values",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/arena-creations?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   35,
			Results: []int32{1, 11, 14, 20, 27, 31, 35},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/arena-creations?category=species&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   14,
			Results: []int32{14, 15, 18, 19, 23, 27},
		},
	}

	testIDList(t, tests, apilookups.E.ArenaCreations.Endpoint, "RetrieveArenaCreations", testSrv.HandleArenaCreations, compareAPIResourceLists[apiresource.Named])
}
