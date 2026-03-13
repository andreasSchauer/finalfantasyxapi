package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetOverdriveCommand(t *testing.T) {
	tests := []expOverdriveCommand{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-commands/8",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "overdrive command with provided id '8' doesn't exist. max id: 7.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-commands/2",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"overdrives": 8,
				},
			},
			expUnique: newExpUnique(2, "grand summon"),
			rank: 			0,
			user:			6,
			topmenu: 		h.GetInt32Ptr(3),
			openSubmenu: 	1,
			overdrives: 	[]int32{5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-commands/3",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"overdrives": 4,
				},
			},
			expUnique: newExpUnique(3, "slots"),
			rank: 			5,
			user:			7,
			topmenu: 		h.GetInt32Ptr(3),
			openSubmenu: 	14,
			overdrives: 	[]int32{13, 14, 15, 16},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-commands/7",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"overdrives": 64,
				},
			},
			expUnique: newExpUnique(7, "mix"),
			rank: 			5,
			user:			11,
			topmenu: 		h.GetInt32Ptr(3),
			openSubmenu: 	10,
			overdrives: 	[]int32{52, 67, 78, 84, 90, 103, 111, 115},
		},
	}

	testSingleResources(t, tests, "GetOverdriveCommand", testCfg.HandleOverdriveCommands, compareOverdriveCommands)
}

func TestRetrieveOverdriveCommands(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-commands",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{1, 2, 3, 4, 5, 6, 7},
		},
	}

	testIdList(t, tests, testCfg.e.overdriveCommands.endpoint, "RetrieveOverdriveCommands", testCfg.HandleOverdriveCommands, compareAPIResourceLists[NamedApiResourceList])
}
