package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetSubmenu(t *testing.T) {
	t.Parallel()
	tests := []Submenu{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus/16",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "submenu with provided id '16' doesn't exist. max id: 15",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus/1",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"users":                          1,
					"abilities":                      8,
					"opened by - overdrive commands": 1,
				},
			},
			Unique:    newExpUnique(1, "summon"),
			topmenu:   h.GetInt32Ptr(1),
			users:     []int32{6},
			abilities: []int32{390, 391, 392, 393, 394, 395, 396, 397},
			openedBy: &MenuOpen{
				ability:           nil,
				aeonCommand:       nil,
				overdriveCommands: []int32{2},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus/3",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"users":     2,
					"abilities": 22,
				},
			},
			Unique:    newExpUnique(3, "special"),
			topmenu:   h.GetInt32Ptr(1),
			users:     []int32{1, 3},
			abilities: []int32{23, 27, 32, 33, 38, 41, 44},
			openedBy:  nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus/9",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"users":     1,
					"abilities": 0,
				},
			},
			Unique:    newExpUnique(9, "use"),
			topmenu:   nil,
			users:     []int32{1},
			abilities: []int32{},
			openedBy: &MenuOpen{
				ability:           h.GetInt32Ptr(24),
				aeonCommand:       nil,
				overdriveCommands: nil,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus/12",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"users":     1,
					"abilities": 0,
				},
			},
			Unique:    newExpUnique(12, "gil yojimbo"),
			topmenu:   h.GetInt32Ptr(1),
			users:     []int32{19},
			abilities: []int32{},
			openedBy: &MenuOpen{
				ability:           nil,
				aeonCommand:       h.GetInt32Ptr(1),
				overdriveCommands: nil,
			},
		},
	}

	testSingleResources(t, tests, "GetSubmenu", testSrv.HandleSubmenus)
}

func TestRetrieveSubmenus(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus",
				expectedStatus: http.StatusOK,
			},
			count:   15,
			results: []int32{1, 4, 8, 9, 13, 15},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus?topmenu=1",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{1, 2, 3, 4, 5, 6, 12},
		},
	}

	testIDList(t, tests, apilookups.E.Submenus.Endpoint, "RetrieveSubmenus", testSrv.HandleSubmenus, compareAPIResourceLists[apiresource.Named])
}
