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
			testGeneral: testGeneral{
				requestURL:     "/api/primers/27",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "primer with provided id '27' doesn't exist. max id: 26",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers/1",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"areas":     2,
					"treasures": 2,
				},
			},
			Unique:    newExpUnique(1, "al bhed primer i"),
			keyItem:   35,
			areas:     []int32{15, 167},
			treasures: []int32{13, 215},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers/22",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"areas":     1,
					"treasures": 1,
				},
			},
			Unique:    newExpUnique(22, "al bhed primer xxii"),
			keyItem:   56,
			areas:     []int32{195},
			treasures: []int32{252},
		},
	}

	testSingleResources(t, tests, "GetPrimer", testSrv.HandlePrimers)
}

func TestRetrievePrimers(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   26,
			results: []int32{1, 5, 17, 20, 26},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers?availability=pre-story",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{19, 20, 21, 22},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers?availability=always&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   21,
			results: []int32{1, 5, 10, 15, 17, 23, 25},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers?availability=pre-story",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{19, 20, 21, 22},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers?availability=post",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{26},
		},
	}

	testIDList(t, tests, apilookups.E.Primers.Endpoint, "RetrievePrimers", testSrv.HandlePrimers, compareAPIResourceLists[apiresource.Named])
}
