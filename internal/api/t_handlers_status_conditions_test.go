package api

import (
	"net/http"
	"testing"
)

func TestGetStatusCondition(t *testing.T) {
	tests := []expStatusCondition{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/47",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "status condition with provided id '47' doesn't exist. max id: 46.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/5?inflict_min=always",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 						5,
					"monsters":								191,
					"inflicted by - player abilities": 		1,
					"inflicted by - overdrive abilities": 	0,
					"inflicted by - item abilities": 		0,
					"inflicted by - unspecified abilities": 0,
					"inflicted by - enemy abilities": 		6,
					"inflicted by - status conditions": 	0,
				},
			},
			expUnique: 			newExpUnique(5, "death"),
			autoAbilities: 		[]int32{19, 20, 54, 77, 78},
			monstersResistance: []int32{},
			inflictedBy: &testStatusInteractions{
				playerAbilities: 		[]int32{98},
				overdriveAbilities: 	[]int32{},
				itemAbilities: 			[]int32{},
				unspecifiedAbilities: 	[]int32{},
				enemyAbilities: 		[]int32{43, 79, 166, 187, 217, 290},
				statusConditions: 		[]int32{},
			},
			removedBy: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 						0,
					"monsters":								0,
					"inflicted by - player abilities": 		0,
					"inflicted by - overdrive abilities": 	0,
					"inflicted by - item abilities": 		0,
					"inflicted by - unspecified abilities": 0,
					"inflicted by - enemy abilities": 		0,
					"inflicted by - status conditions": 	0,
					"removed by - player abilities": 		0,
					"removed by - overdrive abilities": 	0,
					"removed by - item abilities": 			0,
					"removed by - unspecified abilities": 	0,
					"removed by - enemy abilities": 		0,
					"removed by - status conditions": 		0,
				},
			},
			expUnique: 			newExpUnique(0, ""),
			autoAbilities: 		[]int32{},
			monstersResistance: []int32{},
			inflictedBy: &testStatusInteractions{
				playerAbilities: 		[]int32{},
				overdriveAbilities: 	[]int32{},
				itemAbilities: 			[]int32{},
				unspecifiedAbilities: 	[]int32{},
				enemyAbilities: 		[]int32{},
				statusConditions: 		[]int32{},
			},
			removedBy: &testStatusInteractions{
				playerAbilities: 		[]int32{},
				overdriveAbilities: 	[]int32{},
				itemAbilities: 			[]int32{},
				unspecifiedAbilities: 	[]int32{},
				enemyAbilities: 		[]int32{},
				statusConditions: 		[]int32{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 						0,
					"monsters":								0,
					"inflicted by - player abilities": 		0,
					"inflicted by - overdrive abilities": 	0,
					"inflicted by - item abilities": 		0,
					"inflicted by - unspecified abilities": 0,
					"inflicted by - enemy abilities": 		0,
					"inflicted by - status conditions": 	0,
					"removed by - player abilities": 		0,
					"removed by - overdrive abilities": 	0,
					"removed by - item abilities": 			0,
					"removed by - unspecified abilities": 	0,
					"removed by - enemy abilities": 		0,
					"removed by - status conditions": 		0,
				},
			},
			expUnique: 			newExpUnique(0, ""),
			autoAbilities: 		[]int32{},
			monstersResistance: []int32{},
			inflictedBy: &testStatusInteractions{
				playerAbilities: 		[]int32{},
				overdriveAbilities: 	[]int32{},
				itemAbilities: 			[]int32{},
				unspecifiedAbilities: 	[]int32{},
				enemyAbilities: 		[]int32{},
				statusConditions: 		[]int32{},
			},
			removedBy: &testStatusInteractions{
				playerAbilities: 		[]int32{},
				overdriveAbilities: 	[]int32{},
				itemAbilities: 			[]int32{},
				unspecifiedAbilities: 	[]int32{},
				enemyAbilities: 		[]int32{},
				statusConditions: 		[]int32{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities": 						0,
					"monsters":								0,
					"inflicted by - player abilities": 		0,
					"inflicted by - overdrive abilities": 	0,
					"inflicted by - item abilities": 		0,
					"inflicted by - unspecified abilities": 0,
					"inflicted by - enemy abilities": 		0,
					"inflicted by - status conditions": 	0,
					"removed by - player abilities": 		0,
					"removed by - overdrive abilities": 	0,
					"removed by - item abilities": 			0,
					"removed by - unspecified abilities": 	0,
					"removed by - enemy abilities": 		0,
					"removed by - status conditions": 		0,
				},
			},
			expUnique: 			newExpUnique(0, ""),
			autoAbilities: 		[]int32{},
			monstersResistance: []int32{},
			inflictedBy: &testStatusInteractions{
				playerAbilities: 		[]int32{},
				overdriveAbilities: 	[]int32{},
				itemAbilities: 			[]int32{},
				unspecifiedAbilities: 	[]int32{},
				enemyAbilities: 		[]int32{},
				statusConditions: 		[]int32{},
			},
			removedBy: &testStatusInteractions{
				playerAbilities: 		[]int32{},
				overdriveAbilities: 	[]int32{},
				itemAbilities: 			[]int32{},
				unspecifiedAbilities: 	[]int32{},
				enemyAbilities: 		[]int32{},
				statusConditions: 		[]int32{},
			},
		},
		
	}

	testSingleResources(t, tests, "GetStatusCondition", testCfg.HandleStatusConditions, compareStatusConditions)
}

func TestRetrieveStatusConditions(t *testing.T) {
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

	testIdList(t, tests, testCfg.e.modifiers.endpoint, "RetrieveStatusConditions", testCfg.HandleStatusConditions, compareAPIResourceLists[NamedApiResourceList])
}
