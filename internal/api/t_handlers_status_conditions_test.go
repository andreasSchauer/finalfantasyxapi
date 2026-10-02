package api

import (
	"net/http"
	"testing"
)

func TestGetStatusCondition(t *testing.T) {
	t.Parallel()
	tests := []expStatusCondition{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/46",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "status condition with provided id '46' doesn't exist. max id: 45.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/5?inflict_min=always",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities":                     5,
					"monsters":                           191,
					"inflicted by - player abilities":    1,
					"inflicted by - overdrive abilities": 0,
					"inflicted by - item abilities":      0,
					"inflicted by - misc abilities":      0,
					"inflicted by - enemy abilities":     6,
					"inflicted by - status conditions":   0,
				},
			},
			expUnique:          newExpUnique(5, "death"),
			autoAbilities:      []int32{19, 20, 54, 77, 78},
			monstersResistance: []int32{},
			inflictedBy: &testStatusInfliction{
				playerAbilities:    []int32{98},
				overdriveAbilities: []int32{},
				itemAbilities:      []int32{},
				miscAbilities:      []int32{},
				enemyAbilities:     []int32{43, 79, 166, 187, 217, 290},
				statusConditions:   []int32{},
			},
			removedBy: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/9?inflict_max=60",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities":                     6,
					"monsters":                           195,
					"inflicted by - player abilities":    0,
					"inflicted by - overdrive abilities": 0,
					"inflicted by - item abilities":      0,
					"inflicted by - misc abilities":      0,
					"inflicted by - enemy abilities":     7,
					"inflicted by - status conditions":   0,
					"removed by - player abilities":      1,
					"removed by - overdrive abilities":   4,
					"removed by - item abilities":        3,
					"removed by - enemy abilities":       4,
					"removed by - status conditions":     0,
				},
			},
			expUnique:          newExpUnique(9, "petrification"),
			autoAbilities:      []int32{17, 18, 54, 75, 76, 129},
			monstersResistance: []int32{2, 5, 55, 92, 129, 157, 189, 216, 237, 272, 295, 307},
			inflictedBy: &testStatusInfliction{
				playerAbilities:    []int32{},
				overdriveAbilities: []int32{},
				itemAbilities:      []int32{},
				miscAbilities:      []int32{},
				enemyAbilities:     []int32{21, 29, 178, 183, 378, 380, 382},
				statusConditions:   []int32{},
			},
			removedBy: &testStatusRemoval{
				playerAbilities:    []int32{53},
				overdriveAbilities: []int32{152, 153, 158, 159},
				itemAbilities:      []int32{12, 16, 21},
				enemyAbilities:     []int32{275, 276, 327, 437},
				statusConditions:   []int32{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/21?inflict_min=infinite&resistance=immune",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities":                     2,
					"monsters":                           20,
					"inflicted by - player abilities":    2,
					"inflicted by - overdrive abilities": 3,
					"inflicted by - item abilities":      2,
					"inflicted by - misc abilities":      0,
					"inflicted by - enemy abilities":     0,
					"inflicted by - status conditions":   0,
					"removed by - player abilities":      2,
					"removed by - overdrive abilities":   0,
					"removed by - item abilities":        1,
					"removed by - enemy abilities":       14,
					"removed by - status conditions":     3,
				},
			},
			expUnique:          newExpUnique(21, "haste"),
			autoAbilities:      []int32{96, 101},
			monstersResistance: []int32{3, 33, 135, 156, 199, 252, 253},
			inflictedBy: &testStatusInfliction{
				playerAbilities:    []int32{56, 57},
				overdriveAbilities: []int32{165, 166, 167},
				itemAbilities:      []int32{54, 55},
				miscAbilities:      []int32{},
				enemyAbilities:     []int32{},
				statusConditions:   []int32{},
			},
			removedBy: &testStatusRemoval{
				playerAbilities:    []int32{63, 90},
				overdriveAbilities: []int32{},
				itemAbilities:      []int32{63},
				enemyAbilities:     []int32{6, 8, 135, 247, 397},
				statusConditions:   []int32{8, 9, 14},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/31",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"auto-abilities":                     0,
					"monsters":                           6,
					"inflicted by - player abilities":    0,
					"inflicted by - overdrive abilities": 0,
					"inflicted by - item abilities":      0,
					"inflicted by - misc abilities":      1,
					"inflicted by - enemy abilities":     0,
					"inflicted by - status conditions":   0,
					"removed by - player abilities":      0,
					"removed by - overdrive abilities":   0,
					"removed by - item abilities":        0,
					"removed by - enemy abilities":       0,
					"removed by - status conditions":     2,
				},
			},
			expUnique:          newExpUnique(31, "boost"),
			autoAbilities:      []int32{},
			monstersResistance: []int32{3, 33, 69, 108, 135, 253},
			inflictedBy: &testStatusInfliction{
				playerAbilities:    []int32{},
				overdriveAbilities: []int32{},
				itemAbilities:      []int32{},
				miscAbilities:      []int32{8},
				enemyAbilities:     []int32{},
				statusConditions:   []int32{},
			},
			removedBy: &testStatusRemoval{
				playerAbilities:    []int32{},
				overdriveAbilities: []int32{},
				itemAbilities:      []int32{},
				enemyAbilities:     []int32{},
				statusConditions:   []int32{8, 9},
			},
		},
	}

	testSingleResources(t, tests, "GetStatusCondition", testCfg.HandleStatusConditions, compareStatusConditions)
}

func TestRetrieveStatusConditions(t *testing.T) {
	t.Parallel()
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   45,
			results: []int32{1, 8, 14, 16, 22, 38, 45},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/status-conditions/?category=2",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 40},
		},
	}

	testIdList(t, tests, testCfg.e.statusConditions.endpoint, "RetrieveStatusConditions", testCfg.HandleStatusConditions, compareAPIResourceLists[NamedApiResourceList])
}
