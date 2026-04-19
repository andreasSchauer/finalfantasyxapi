package api

import (
	"net/http"
	"testing"
)

func TestGetModifier(t *testing.T) {
	tests := []expModifier{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/modifiers/32",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "modifier with provided id '32' doesn't exist. max id: 31.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/modifiers/3",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 		0,
					"player abilities": 	1,
					"overdrive abilities": 	3,
					"item abilities": 		0,
					"trigger commands": 	0,
					"enemy abilities": 		0,
					"status conditions": 	0,
					"properties": 			0,
				},
			},
			expUnique: 			newExpUnique(3, "buff-factor-str-based"),
			autoAbilities: 		[]int32{},
			playerAbilities: 	[]int32{27},
			overdriveAbilities: []int32{162, 163, 170},
			itemAbilities: 		[]int32{},
			triggerCommands: 	[]int32{},
			enemyAbilities: 	[]int32{},
			statusConditions: 	[]int32{},
			properties: 		[]int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/modifiers/16",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 		3,
					"player abilities": 	0,
					"overdrive abilities": 	2,
					"item abilities": 		2,
					"trigger commands": 	0,
					"enemy abilities": 		0,
					"status conditions": 	0,
					"properties": 			0,
				},
			},
			expUnique: 			newExpUnique(16, "mp-cost"),
			autoAbilities: 		[]int32{40, 42, 43},
			playerAbilities: 	[]int32{},
			overdriveAbilities: []int32{174, 175},
			itemAbilities: 		[]int32{68, 69},
			triggerCommands: 	[]int32{},
			enemyAbilities: 	[]int32{},
			statusConditions: 	[]int32{},
			properties: 		[]int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/modifiers/18",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 		0,
					"player abilities": 	0,
					"overdrive abilities": 	0,
					"item abilities": 		0,
					"trigger commands": 	1,
					"enemy abilities": 		0,
					"status conditions": 	0,
					"properties": 			0,
				},
			},
			expUnique: 			newExpUnique(18, "overdrive-gauge"),
			autoAbilities: 		[]int32{},
			playerAbilities: 	[]int32{},
			overdriveAbilities: []int32{},
			itemAbilities: 		[]int32{},
			triggerCommands: 	[]int32{6},
			enemyAbilities: 	[]int32{},
			statusConditions: 	[]int32{},
			properties: 		[]int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/modifiers/20",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 		4,
					"player abilities": 	0,
					"overdrive abilities": 	0,
					"item abilities": 		0,
					"trigger commands": 	0,
					"enemy abilities": 		1,
					"status conditions": 	6,
					"properties": 			1,
				},
			},
			expUnique: 			newExpUnique(20, "physical-damage-taken"),
			autoAbilities: 		[]int32{103, 104, 105, 106},
			playerAbilities: 	[]int32{},
			overdriveAbilities: []int32{},
			itemAbilities: 		[]int32{},
			triggerCommands: 	[]int32{},
			enemyAbilities: 	[]int32{86},
			statusConditions: 	[]int32{10, 27, 32, 34, 44},
			properties: 		[]int32{1},
		},
	}

	testSingleResources(t, tests, "GetModifier", testCfg.HandleModifiers, compareModifiers)
}

func TestRetrieveModifiers(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/modifiers?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   31,
			results: []int32{1, 7, 12, 17, 20, 25, 31},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/modifiers?category=percentage",
				expectedStatus: http.StatusOK,
			},
			count:   9,
			results: []int32{5, 7, 9, 25, 26, 29},
		},
	}

	testIdList(t, tests, testCfg.e.modifiers.endpoint, "RetrieveModifiers", testCfg.HandleModifiers, compareAPIResourceLists[NamedApiResourceList])
}
