package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestGetProperty(t *testing.T) {
	t.Parallel()
	tests := []Property{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/properties/14",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "property with provided id '14' doesn't exist. max id: 13",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/properties/1",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 0,
					"monsters":       53,
				},
			},
			Unique:     newExpUnique(1, "armored"),
			autoAbilities: []int32{},
			monsters:      []int32{5, 29, 79, 107, 197, 277, 292},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/properties/4",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 1,
					"monsters":       0,
				},
			},
			Unique:     newExpUnique(4, "piercing"),
			autoAbilities: []int32{2},
			monsters:      []int32{},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/properties/11",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 1,
					"monsters":       123,
				},
			},
			Unique:     newExpUnique(11, "delayproof"),
			autoAbilities: []int32{54},
			monsters:      []int32{5, 33, 70, 136, 193, 235, 262, 282, 307},
		},
	}

	testSingleResources(t, tests, "GetProperty", testSrv.HandleProperties, compareProperties)
}

func TestRetrieveProperties(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/properties",
				expectedStatus: http.StatusOK,
			},
			count:   13,
			results: []int32{1, 4, 5, 8, 9, 10, 12, 13},
		},
	}

	testIDList(t, tests, apilookups.E.Properties.Endpoint, "RetrieveProperties", testSrv.HandleProperties, compareAPIResourceLists[apiresource.Named])
}
