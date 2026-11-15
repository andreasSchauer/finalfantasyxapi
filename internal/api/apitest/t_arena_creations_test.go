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
			Unique:              newExpUnique(13, "vorban"),
			category:            "area",
			monster:             268,
			parentSubquest:      13,
			reward:              newTestResAmount[apiresource.Typed](107, 60),
			requiredCatchAmount: 1,
			requiredMonsters:    []int32{239, 243, 245, 248},
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
			Unique:              newExpUnique(19, "one-eye"),
			category:            "species",
			monster:             274,
			parentSubquest:      19,
			reward:              newTestResAmount[apiresource.Typed](64, 60),
			requiredCatchAmount: 4,
			requiredMonsters:    []int32{40, 73, 87, 170, 240},
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
			Unique:                    newExpUnique(31, "th'uban"),
			category:                  "original",
			monster:                   286,
			parentSubquest:            31,
			reward:                    newTestResAmount[apiresource.Typed](110, 99),
			requiredCatchAmount:       6,
			unlockedCreationsCategory: h.GetStrPtr("species"),
			requiredMonsters:          nil,
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
