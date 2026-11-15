package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestGetPrimer(t *testing.T) {
	t.Parallel()
	tests := []Primer{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/primers/27",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "primer with provided id '27' doesn't exist. max id: 26",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/primers/1",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"areas":     2,
					"treasures": 2,
				},
			},
			Unique:    newExpUnique(1, "al bhed primer i"),
			KeyItem:   35,
			Areas:     []int32{15, 167},
			Treasures: []int32{13, 215},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/primers/22",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"areas":     1,
					"treasures": 1,
				},
			},
			Unique:    newExpUnique(22, "al bhed primer xxii"),
			KeyItem:   56,
			Areas:     []int32{195},
			Treasures: []int32{252},
		},
	}

	testSingleResources(t, tests, "GetPrimer", testSrv.HandlePrimers)
}

func TestRetrievePrimers(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/primers?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   26,
			Results: []int32{1, 5, 17, 20, 26},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/primers?availability=pre-story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{19, 20, 21, 22},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/primers?availability=always&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   21,
			Results: []int32{1, 5, 10, 15, 17, 23, 25},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/primers?availability=pre-story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{19, 20, 21, 22},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/primers?availability=post",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{26},
		},
	}

	testIDList(t, tests, apilookups.E.Primers.Endpoint, "RetrievePrimers", testSrv.HandlePrimers, compareAPIResourceLists[apiresource.Named])
}
