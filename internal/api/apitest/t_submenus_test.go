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
			TestGeneral: TestGeneral{
				RequestURL:     "/api/submenus/16",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "submenu with provided id '16' doesn't exist. max id: 15",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/submenus/1",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"users":                          1,
					"abilities":                      8,
					"opened by - overdrive commands": 1,
				},
			},
			Unique:    newExpUnique(1, "summon"),
			Topmenu:   h.GetInt32Ptr(1),
			Users:     []int32{6},
			Abilities: []int32{390, 391, 392, 393, 394, 395, 396, 397},
			OpenedBy: &MenuOpen{
				Ability:           nil,
				AeonCommand:       nil,
				OverdriveCommands: []int32{2},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/submenus/3",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"users":     2,
					"abilities": 22,
				},
			},
			Unique:    newExpUnique(3, "special"),
			Topmenu:   h.GetInt32Ptr(1),
			Users:     []int32{1, 3},
			Abilities: []int32{23, 27, 32, 33, 38, 41, 44},
			OpenedBy:  nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/submenus/9",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"users":     1,
					"abilities": 0,
				},
			},
			Unique:    newExpUnique(9, "use"),
			Topmenu:   nil,
			Users:     []int32{1},
			Abilities: []int32{},
			OpenedBy: &MenuOpen{
				Ability:           h.GetInt32Ptr(24),
				AeonCommand:       nil,
				OverdriveCommands: nil,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/submenus/12",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"users":     1,
					"abilities": 0,
				},
			},
			Unique:    newExpUnique(12, "gil yojimbo"),
			Topmenu:   h.GetInt32Ptr(1),
			Users:     []int32{19},
			Abilities: []int32{},
			OpenedBy: &MenuOpen{
				Ability:           nil,
				AeonCommand:       h.GetInt32Ptr(1),
				OverdriveCommands: nil,
			},
		},
	}

	testSingleResources(t, tests, "GetSubmenu", testSrv.HandleSubmenus)
}

func TestRetrieveSubmenus(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/submenus",
				ExpectedStatus: http.StatusOK,
			},
			Count:   15,
			Results: []int32{1, 4, 8, 9, 13, 15},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/submenus?topmenu=1",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{1, 2, 3, 4, 5, 6, 12},
		},
	}

	testIDList(t, tests, apilookups.E.Submenus.Endpoint, "RetrieveSubmenus", testSrv.HandleSubmenus, compareAPIResourceLists[apiresource.Named])
}
