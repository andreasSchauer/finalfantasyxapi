package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetPlayerUnit(t *testing.T) {
	tests := []expPlayerUnit{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-units/19",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "player unit with provided id '19' doesn't exist. max id: 18.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-units/1",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 2,
				},
			},
			expUnique: expUnique{
				id:   1,
				name: "tidus",
			},
			unitType: 			1,
			typedUnit: 			"/characters/1",
			area:           	1,
			celestialWeapon:  	h.GetInt32Ptr(1),
			characterClasses: 	[]int32{1, 5},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-units/8",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 1,
				},
			},
			expUnique: expUnique{
				id:   8,
				name: "seymour",
			},
			unitType: 			1,
			typedUnit: 			"/characters/8",
			area:           	103,
			celestialWeapon:  	nil,
			characterClasses: 	[]int32{12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-units/10",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 3,
				},
			},
			expUnique: expUnique{
				id:   10,
				name: "ifrit",
			},
			unitType: 			2,
			typedUnit: 			"/aeons/2",
			area:           	61,
			celestialWeapon:  	h.GetInt32Ptr(3),
			characterClasses: 	[]int32{2, 3, 14},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-units/15",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 2,
				},
			},
			expUnique: expUnique{
				id:   15,
				name: "yojimbo",
			},
			unitType: 			2,
			typedUnit: 			"/aeons/7",
			area:           	212,
			celestialWeapon:  	h.GetInt32Ptr(6),
			characterClasses: 	[]int32{2, 19},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-units/16",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"character classes": 3,
				},
			},
			expUnique: expUnique{
				id:   16,
				name: "cindy",
			},
			unitType: 			2,
			typedUnit: 			"/aeons/8",
			area:           	210,
			celestialWeapon:  	nil,
			characterClasses: 	[]int32{2, 4, 20},
		},
	}

	testSingleResources(t, tests, "GetPlayerUnits", testCfg.HandlePlayerUnits, comparePlayerUnits)
}

func TestRetrievePlayerUnits(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-units",
				expectedStatus: http.StatusOK,
			},
			count:   18,
			results: []int32{1, 3, 4, 8, 12, 17, 18},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/player-units?type=aeon",
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{9, 11, 13, 14, 15, 16, 18},
		},
	}

	testIdList(t, tests, testCfg.e.playerUnits.endpoint, "RetrievePlayerUnits", testCfg.HandlePlayerUnits, compareAPIResourceLists[TypedAPIResourceList])
}
