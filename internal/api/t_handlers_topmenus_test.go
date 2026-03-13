package api

import (
	"net/http"
	"testing"
)

func TestGetTopmenu(t *testing.T) {
	tests := []expTopmenu{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/topmenus/5",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "topmenu with provided id '5' doesn't exist. max id: 4.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/topmenus/1",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"submenus": 			7,
					"abilities": 			8,
					"overdrive commands":	0,
					"overdrives": 			0,
					"aeon commands": 		9,
				},
			},
			expUnique: newExpUnique(1, "main"),
			submenus: 			[]int32{1, 2, 3, 4, 5, 6, 12},
			abilities: 			[]int32{88, 89, 90, 91, 92, 93, 374, 375},
			overdriveCommands: 	[]int32{},
			overdrives: 		[]int32{},
			aeonCommands: 		[]int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/topmenus/2",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"submenus": 			2,
					"abilities": 			6,
					"overdrive commands":	0,
					"overdrives": 			0,
					"aeon commands": 		0,
				},
			},
			expUnique: newExpUnique(2, "right"),
			submenus: 			[]int32{7, 8},
			abilities: 			[]int32{377, 378, 379, 380, 381, 382},
			overdriveCommands: 	[]int32{},
			overdrives: 		[]int32{},
			aeonCommands: 		[]int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/topmenus/3",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"submenus": 			0,
					"abilities": 			12,
					"overdrive commands":	7,
					"overdrives": 			8,
					"aeon commands": 		0,
				},
			},
			expUnique: newExpUnique(3, "left"),
			submenus: 			[]int32{},
			abilities: 			[]int32{362, 364, 367, 370, 372, 373},
			overdriveCommands: 	[]int32{1, 2, 3, 4, 5, 6, 7},
			overdrives: 		[]int32{116, 117, 118, 119, 120, 121, 122, 123},
			aeonCommands: 		[]int32{},
		},
	}

	testSingleResources(t, tests, "GetTopmenu", testCfg.HandleTopmenus, compareTopmenus)
}

func TestRetrieveTopmenus(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/topmenus",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{1, 2, 3, 4},
		},
	}

	testIdList(t, tests, testCfg.e.topmenus.endpoint, "RetrieveTopmenus", testCfg.HandleTopmenus, compareAPIResourceLists[NamedApiResourceList])
}
