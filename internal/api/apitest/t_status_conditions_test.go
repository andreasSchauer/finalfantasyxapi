package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestGetStatusCondition(t *testing.T) {
	t.Parallel()
	tests := []StatusCondition{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/status-conditions/46",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "status condition with provided id '46' doesn't exist. max id: 45",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/status-conditions/5?inflict_min=always",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
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
			Unique:             NewExpUnique(5, "death"),
			AutoAbilities:      []int32{19, 20, 54, 77, 78},
			MonstersResistance: []int32{},
			InflictedBy: &TestStatusInfliction{
				PlayerAbilities:    []int32{98},
				OverdriveAbilities: []int32{},
				ItemAbilities:      []int32{},
				MiscAbilities:      []int32{},
				EnemyAbilities:     []int32{43, 79, 166, 187, 217, 289},
			},
			RemovedBy: nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/status-conditions/9?inflict_max=60",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
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
			Unique:             NewExpUnique(9, "petrification"),
			AutoAbilities:      []int32{17, 18, 54, 75, 76, 129},
			MonstersResistance: []int32{2, 5, 55, 92, 129, 157, 189, 216, 237, 272, 295, 307},
			InflictedBy: &TestStatusInfliction{
				PlayerAbilities:    []int32{},
				OverdriveAbilities: []int32{},
				ItemAbilities:      []int32{},
				MiscAbilities:      []int32{},
				EnemyAbilities:     []int32{21, 29, 178, 183, 377, 379, 381},
			},
			RemovedBy: &TestStatusRemoval{
				PlayerAbilities:    []int32{53},
				OverdriveAbilities: []int32{152, 153, 158, 159},
				ItemAbilities:      []int32{12, 16, 21},
				EnemyAbilities:     []int32{275, 275, 326, 437},
				StatusConditions:   []int32{},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/status-conditions/21?inflict_min=infinite&resistance=immune",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
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
			Unique:             NewExpUnique(21, "haste"),
			AutoAbilities:      []int32{96, 101},
			MonstersResistance: []int32{3, 33, 135, 156, 199, 252, 253},
			InflictedBy: &TestStatusInfliction{
				PlayerAbilities:    []int32{56, 57},
				OverdriveAbilities: []int32{165, 166, 167},
				ItemAbilities:      []int32{54, 55},
				MiscAbilities:      []int32{},
				EnemyAbilities:     []int32{},
			},
			RemovedBy: &TestStatusRemoval{
				PlayerAbilities:    []int32{63, 90},
				OverdriveAbilities: []int32{},
				ItemAbilities:      []int32{63},
				EnemyAbilities:     []int32{6, 8, 135, 247, 396},
				StatusConditions:   []int32{8, 9, 14},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/status-conditions/31",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
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
			Unique:             NewExpUnique(31, "boost"),
			AutoAbilities:      []int32{},
			MonstersResistance: []int32{3, 33, 69, 108, 135, 253},
			InflictedBy: &TestStatusInfliction{
				PlayerAbilities:    []int32{},
				OverdriveAbilities: []int32{},
				ItemAbilities:      []int32{},
				MiscAbilities:      []int32{8},
				EnemyAbilities:     []int32{},
			},
			RemovedBy: &TestStatusRemoval{
				PlayerAbilities:    []int32{},
				OverdriveAbilities: []int32{},
				ItemAbilities:      []int32{},
				EnemyAbilities:     []int32{},
				StatusConditions:   []int32{8, 9},
			},
		},
	}

	testSingleResources(t, tests, "GetStatusCondition", testSrv.HandleStatusConditions)
}

func TestRetrieveStatusConditions(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/status-conditions?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   45,
			Results: []int32{1, 8, 14, 16, 22, 38, 45},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/status-conditions/?category=2",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 40},
		},
	}

	testIDList(t, tests, apilookups.E.StatusConditions.Endpoint, "RetrieveStatusConditions", testSrv.HandleStatusConditions, compareAPIResourceLists[apiresource.Named])
}
