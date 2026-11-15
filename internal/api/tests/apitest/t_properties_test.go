package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
)

func TestGetProperty(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Property{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/properties/14",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "property with provided id '14' doesn't exist. max id: 13",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/properties/1",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"auto-abilities": 0,
					"monsters":       53,
				},
			},
			Unique:        apitesttypes.NewExpUnique(1, "armored"),
			AutoAbilities: []int32{},
			Monsters:      []int32{5, 29, 79, 107, 197, 277, 292},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/properties/4",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"auto-abilities": 1,
					"monsters":       0,
				},
			},
			Unique:        apitesttypes.NewExpUnique(4, "piercing"),
			AutoAbilities: []int32{2},
			Monsters:      []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/properties/11",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"auto-abilities": 1,
					"monsters":       123,
				},
			},
			Unique:        apitesttypes.NewExpUnique(11, "delayproof"),
			AutoAbilities: []int32{54},
			Monsters:      []int32{5, 33, 70, 136, 193, 235, 262, 282, 307},
		},
	}

	testSingleResources(t, tests, "GetProperty", testSrv.HandleProperties)
}

func TestRetrieveProperties(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/properties",
				ExpectedStatus: http.StatusOK,
			},
			Count:   13,
			Results: []int32{1, 4, 5, 8, 9, 10, 12, 13},
		},
	}

	testIDList(t, tests, apilookups.E.Properties.Endpoint, "RetrieveProperties", testSrv.HandleProperties, compareAPIResourceLists[apiresource.Named])
}
