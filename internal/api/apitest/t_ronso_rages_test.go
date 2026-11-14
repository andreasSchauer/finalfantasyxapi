package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestGetRonsoRage(t *testing.T) {
	t.Parallel()
	tests := []RonsoRage{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/ronso-rages/13",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "ronso rage with provided id '13' doesn't exist. max id: 12",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/ronso-rages/4",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"monsters": 5,
				},
			},
			Unique:    newExpUnique(4, "self destruct"),
			overdrive: 39,
			monsters:  []int32{38, 109, 167, 176, 245},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/ronso-rages/8",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"monsters": 3,
				},
			},
			Unique:    newExpUnique(8, "doom"),
			overdrive: 43,
			monsters:  []int32{154, 167, 210},
		},
	}

	testSingleResources(t, tests, "GetRonsoRage", testSrv.HandleRonsoRages)
}

func TestRetrieveRonsoRages(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/ronso-rages",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{1, 3, 8, 9, 11, 12},
		},
	}

	testIDList(t, tests, apilookups.E.RonsoRages.Endpoint, "RetrieveRonsoRages", testSrv.HandleRonsoRages, compareAPIResourceLists[apiresource.Named])
}
