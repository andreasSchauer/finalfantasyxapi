package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestGetTopmenu(t *testing.T) {
	t.Parallel()
	tests := []Topmenu{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/topmenus/5",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "topmenu with provided id '5' doesn't exist. max id: 4",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/topmenus/1",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"submenus":           7,
					"abilities":          8,
					"overdrive commands": 0,
					"overdrives":         0,
					"aeon commands":      9,
				},
			},
			Unique:            newExpUnique(1, "main"),
			submenus:          []int32{1, 2, 3, 4, 5, 6, 12},
			abilities:         []int32{88, 89, 90, 91, 92, 93, 374, 375},
			overdriveCommands: []int32{},
			overdrives:        []int32{},
			aeonCommands:      []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/topmenus/2",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"submenus":           2,
					"abilities":          6,
					"overdrive commands": 0,
					"overdrives":         0,
					"aeon commands":      0,
				},
			},
			Unique:            newExpUnique(2, "right"),
			submenus:          []int32{7, 8},
			abilities:         []int32{377, 378, 379, 380, 381, 382},
			overdriveCommands: []int32{},
			overdrives:        []int32{},
			aeonCommands:      []int32{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/topmenus/3",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"submenus":           0,
					"abilities":          12,
					"overdrive commands": 7,
					"overdrives":         8,
					"aeon commands":      0,
				},
			},
			Unique:            newExpUnique(3, "left"),
			submenus:          []int32{},
			abilities:         []int32{362, 364, 367, 370, 372, 373},
			overdriveCommands: []int32{1, 2, 3, 4, 5, 6, 7},
			overdrives:        []int32{116, 117, 118, 119, 120, 121, 122, 123},
			aeonCommands:      []int32{},
		},
	}

	testSingleResources(t, tests, "GetTopmenu", testSrv.HandleTopmenus)
}

func TestRetrieveTopmenus(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/topmenus",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{1, 2, 3, 4},
		},
	}

	testIDList(t, tests, apilookups.E.Topmenus.Endpoint, "RetrieveTopmenus", testSrv.HandleTopmenus, compareAPIResourceLists[apiresource.Named])
}
