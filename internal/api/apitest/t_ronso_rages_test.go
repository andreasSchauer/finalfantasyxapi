package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
)

func TestGetRonsoRage(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.RonsoRage{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/ronso-rages/13",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "ronso rage with provided id '13' doesn't exist. max id: 12",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/ronso-rages/4",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"monsters": 5,
				},
			},
			Unique:    apitesttypes.NewExpUnique(4, "self destruct"),
			Overdrive: 39,
			Monsters:  []int32{38, 109, 167, 176, 245},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/ronso-rages/8",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"monsters": 3,
				},
			},
			Unique:    apitesttypes.NewExpUnique(8, "doom"),
			Overdrive: 43,
			Monsters:  []int32{154, 167, 210},
		},
	}

	testSingleResources(t, tests, "GetRonsoRage", testSrv.HandleRonsoRages)
}

func TestRetrieveRonsoRages(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/ronso-rages",
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{1, 3, 8, 9, 11, 12},
		},
	}

	testIDList(t, tests, apilookups.E.RonsoRages.Endpoint, "RetrieveRonsoRages", testSrv.HandleRonsoRages, compareAPIResourceLists[apiresource.Named])
}
