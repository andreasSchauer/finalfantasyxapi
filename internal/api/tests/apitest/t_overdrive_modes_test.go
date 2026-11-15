package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetOverdriveMode(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.OverdriveMode{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes/ally/2",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. usage: '/api/overdrive-modes', '/api/overdrive-modes/{id}', '/api/overdrive-modes/{name}'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes/18",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "overdrive mode with provided id '18' doesn't exist. max id: 17",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes/a",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "overdrive mode not found: 'a'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes/ally/",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"effect": true,
				},
				ExpLengths: map[string]int{
					"actions": 7,
				},
			},
			Unique:      apitesttypes.NewExpUnique(14, "ally"),
			Description: "Charges on character's turn.",
			ModeType:    "per-action",
			FillRate:    h.GetFloat32Ptr(0.03),
			ActionsAmount: map[string]int32{
				"tidus":   600,
				"yuna":    500,
				"wakka":   350,
				"lulu":    480,
				"kimahri": 300,
				"auron":   450,
				"rikku":   320,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes/3",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"description": true,
				},
			},
			Unique:   apitesttypes.NewExpUnique(3, "comrade"),
			Effect:   "The gauge fills when an ally takes damage.",
			ModeType: "formula",
			FillRate: nil,
			ActionsAmount: map[string]int32{
				"tidus": 300,
				"wakka": 100,
				"rikku": 100,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes/1/",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"description": true,
					"effect":      true,
					"fill rate":   true,
				},
				ExpLengths: map[string]int{
					"actions": 0,
				},
			},
			Unique:   apitesttypes.NewExpUnique(1, "stoic"),
			ModeType: "formula",
		},
	}

	testSingleResources(t, tests, "GetOverdriveMode", testSrv.HandleOverdriveModes)
}

func TestRetrieveOverdriveModes(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes?type=f",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid enum value 'f' used for parameter 'type'. use /api/overdrive-modes/parameters to see allowed values",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes/",
				ExpectedStatus: http.StatusOK,
			},
			Count:    17,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 8, 17},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes?type=formula",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"previous": true,
					"next":     true,
				},
			},
			Count:   4,
			Results: []int32{1, 2, 3, 4},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes?type=2",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"previous": true,
					"next":     true,
				},
			},
			Count:   13,
			Results: []int32{5, 12, 17},
		},
	}

	testIDList(t, tests, apilookups.E.OverdriveModes.Endpoint, "RetrieveOverdriveModes", testSrv.HandleOverdriveModes, compareAPIResourceLists[apiresource.Named])
}
