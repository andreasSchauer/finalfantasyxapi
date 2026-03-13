package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetSubmenu(t *testing.T) {
	tests := []expSubmenu{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus/16",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "submenu with provided id '16' doesn't exist. max id: 15.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus/1",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"users": 		1,
					"abilities": 	8,
					"opened by - overdrive commands": 1,
				},
			},
			expUnique: 	newExpUnique(1, "summon"),
			topmenu: 	h.GetInt32Ptr(1),
			users: 		[]int32{6},
			abilities: 	[]int32{390, 391, 392, 393, 394, 395, 396, 397},
			openedBy: &expMenuOpen{
				ability: 			nil,
				aeonCommand: 		nil,
				overdriveCommands: 	[]int32{2},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus/3",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"users": 		2,
					"abilities": 	22,
				},
			},
			expUnique: 	newExpUnique(3, "special"),
			topmenu: 	h.GetInt32Ptr(1),
			users: 		[]int32{1, 3},
			abilities: 	[]int32{23, 27, 32, 33, 38, 41, 44},
			openedBy: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus/9",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"users": 		1,
					"abilities": 	0,
				},
			},
			expUnique: 	newExpUnique(9, "use"),
			topmenu: 	nil,
			users: 		[]int32{1},
			abilities: 	[]int32{},
			openedBy: &expMenuOpen{
				ability: 			h.GetInt32Ptr(24),
				aeonCommand: 		nil,
				overdriveCommands: 	nil,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/submenus/12",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"users": 		1,
					"abilities": 	0,
				},
			},
			expUnique: 	newExpUnique(12, "gil yojimbo"),
			topmenu: 	h.GetInt32Ptr(1),
			users: 		[]int32{19},
			abilities: 	[]int32{},
			openedBy: &expMenuOpen{
				ability: 			nil,
				aeonCommand: 		h.GetInt32Ptr(1),
				overdriveCommands: 	nil,
			},
		},
		
	}

	testSingleResources(t, tests, "GetSubmenu", testCfg.HandleSubmenus, compareSubmenus)
}

func TestRetrieveSubmenus(t *testing.T) {
	tests := []expListIDs{
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

	testIdList(t, tests, testCfg.e.submenus.endpoint, "RetrieveSubmenus", testCfg.HandleSubmenus, compareAPIResourceLists[NamedApiResourceList])
}
