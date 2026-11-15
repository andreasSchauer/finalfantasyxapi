package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetOverdriveCommand(t *testing.T) {
	t.Parallel()
	tests := []OverdriveCommand{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-commands/8",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "overdrive command with provided id '8' doesn't exist. max id: 7",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-commands/2",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"overdrives": 8,
				},
			},
			Unique:      newExpUnique(2, "grand summon"),
			Rank:        0,
			User:        6,
			Topmenu:     h.GetInt32Ptr(3),
			OpenSubmenu: 1,
			Overdrives:  []int32{5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-commands/3",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"overdrives": 4,
				},
			},
			Unique:      newExpUnique(3, "slots"),
			Rank:        5,
			User:        7,
			Topmenu:     h.GetInt32Ptr(3),
			OpenSubmenu: 14,
			Overdrives:  []int32{13, 14, 15, 16},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-commands/7",
				ExpectedStatus: http.StatusOK,
				ExpLengths: map[string]int{
					"overdrives": 64,
				},
			},
			Unique:      newExpUnique(7, "mix"),
			Rank:        5,
			User:        11,
			Topmenu:     h.GetInt32Ptr(3),
			OpenSubmenu: 10,
			Overdrives:  []int32{52, 67, 78, 84, 90, 103, 111, 115},
		},
	}

	testSingleResources(t, tests, "GetOverdriveCommand", testSrv.HandleOverdriveCommands)
}

func TestRetrieveOverdriveCommands(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/overdrive-commands",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{1, 2, 3, 4, 5, 6, 7},
		},
	}

	testIDList(t, tests, apilookups.E.OverdriveCommands.Endpoint, "RetrieveOverdriveCommands", testSrv.HandleOverdriveCommands, compareAPIResourceLists[apiresource.Named])
}
