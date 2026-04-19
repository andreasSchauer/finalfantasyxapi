package api

import (
	"net/http"
	"testing"
)

func TestGetStat(t *testing.T) {
	tests := []expStat{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/stats/11",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "stat with provided id '11' doesn't exist. max id: 10.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/stats/2?changes_only=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"spheres": 				5,
					"auto-abilities": 		4,
					"player abilities": 	0,
					"overdrive abilities": 	3,
					"item abilities": 		2,
					"trigger commands": 	0,
					"status conditions": 	0,
					"properties": 			0,
				},
			},
			expUnique: 			newExpUnique(2, "mp"),
			spheres: 			[]int32{2, 6, 11, 17, 26},
			autoAbilities: 		[]int32{115, 116, 117, 118},
			playerAbilities: 	[]int32{},
			overdriveAbilities: []int32{171, 172, 173},
			itemAbilities: 		[]int32{65, 67},
			triggerCommands: 	[]int32{},
			statusConditions: 	[]int32{},
			properties: 		[]int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/stats/3",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"spheres": 				5,
					"auto-abilities": 		5,
					"player abilities": 	2,
					"overdrive abilities": 	3,
					"item abilities": 		0,
					"trigger commands": 	1,
					"status conditions": 	3,
					"properties": 			0,
				},
			},
			expUnique: 			newExpUnique(3, "strength"),
			spheres: 			[]int32{1, 6, 11, 18, 26},
			autoAbilities: 		[]int32{29, 30, 31, 32, 51},
			playerAbilities: 	[]int32{11, 27},
			overdriveAbilities: []int32{162, 163, 170},
			itemAbilities: 		[]int32{},
			triggerCommands: 	[]int32{4},
			statusConditions: 	[]int32{1, 18, 35},
			properties: 		[]int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/stats/4",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"spheres": true,
					"auto-abilities": true,
					"player abilities": true,
					"overdrive abilities": true,
					"item abilities": true,
					"trigger commands": true,
					"status conditions": true,
				},
				expLengths: map[string]int{
					"properties": 			2,
				},
			},
			expUnique: 			newExpUnique(4, "defense"),
			properties: 		[]int32{1, 4},
		},
	}

	testSingleResources(t, tests, "GetStat", testCfg.HandleStats, compareStats)
}

func TestRetrieveStats(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/stats",
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	testIdList(t, tests, testCfg.e.stats.endpoint, "RetrieveStats", testCfg.HandleStats, compareAPIResourceLists[NamedApiResourceList])
}
