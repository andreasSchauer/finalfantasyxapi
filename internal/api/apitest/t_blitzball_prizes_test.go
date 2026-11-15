package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestGetBlitzballPrize(t *testing.T) {
	t.Parallel()
	tests := []BlitzballPrize{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/blitzball-prizes/9",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "blitzball prize table with provided id '9' doesn't exist. max id: 8",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/blitzball-prizes/9/2",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid subsection '2'. subsection can't be an integer. use /api/blitzball-prizes/sections for valid subsections",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/blitzball-prizes/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid id 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/blitzball-prizes/1",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"items": 8,
				},
			},
			IDOnly:   newExpIDOnly(1),
			Category: "league",
			Slot:     "1st",
			Items: []testPossibleItem{
				newTestPossibleItem(1, 69, 1, 15),
				newTestPossibleItem(4, 53, 1, 10),
				newTestPossibleItem(6, 96, 1, 10),
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/blitzball-prizes/8",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"items": 3,
				},
			},
			IDOnly:   newExpIDOnly(8),
			Category: "tournament",
			Slot:     "top-scorer",
			Items: []testPossibleItem{
				newTestPossibleItem(0, 5, 1, 20),
				newTestPossibleItem(2, 101, 1, 20),
			},
		},
	}

	testSingleResources(t, tests, "GetBlitzballPrize", testSrv.HandleBlitzballPrizes)
}

func TestRetrieveBlitzballPrizes(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/blitzball-prizes?category=asd",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid enum value 'asd' used for parameter 'category'. use /api/blitzball-prizes/parameters to see allowed values",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/blitzball-prizes",
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/blitzball-prizes?category=league",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{1, 2, 3, 4},
		},
	}

	testIDList(t, tests, apilookups.E.BlitzballPrizes.Endpoint, "RetrieveBlitzballPositions", testSrv.HandleBlitzballPrizes, compareAPIResourceLists[apiresource.Unnamed])
}
