package api

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetAutoAbility(t *testing.T) {
	t.Parallel()
	tests := []expAutoAbility{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/131",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "auto-ability with provided id '131' doesn't exist. max id: 130.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/50",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters drop":       1,
					"monsters items":      4,
					"shops pre airship":   0,
					"shops post airship":  0,
					"treasures":           0,
					"equipment tables":    2,
				},
			},
			expUnique:    newExpUnique(50, "gillionaire"),
			monstersDrop: []int32{302},
			monstersItems: []testMonItemAmts{
				{
					index:   0,
					monster: 190,
					bribe:   5,
				},
				{
					index:     1,
					monster:   261,
					stealRare: 1,
				},
				{
					index:     2,
					monster:   265,
					stealRare: 1,
				},
				{
					index:      3,
					monster:    285,
					dropCommon: 1,
				},
			},
			shopsPreAirship:  []int32{},
			shopsPostAirship: []int32{},
			treasures:        []int32{},
			equipmentTables:  []int32{7, 30},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/2?rel_availability=post",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters drop":       22,
					"monsters items":      2,
					"shops pre airship":   0,
					"shops post airship":  1,
					"treasures":           1,
					"equipment tables":    2,
				},
			},
			expUnique:    newExpUnique(2, "piercing"),
			monstersDrop: []int32{201, 206, 210, 243, 248, 287},
			monstersItems: []testMonItemAmts{
				{
					index:   	0,
					monster: 	247,
					dropRare:   1,
				},
				{
					index:     	1,
					monster:   	285,
					stealRare: 	1,
				},
			},
			shopsPreAirship:  []int32{},
			shopsPostAirship: []int32{27},
			treasures:        []int32{31},
			equipmentTables:  []int32{76, 78},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/2?rel_availability=always",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters drop":       100,
					"monsters items":      3,
					"shops pre airship":   5,
					"shops post airship":  5,
					"treasures":           3,
					"equipment tables":    2,
				},
			},
			expUnique:    newExpUnique(2, "piercing"),
			monstersDrop: []int32{16, 25, 41, 54, 78, 139, 163, 183, 192},
			monstersItems: []testMonItemAmts{
				{
					index:   	0,
					monster: 	161,
					dropRare:   1,
				},
				{
					index:     	1,
					monster:   	173,
					bribe: 		30,
				},
				{
					index:     	2,
					monster:   	190,
					dropCommon: 1,
				},
			},
			shopsPreAirship:  []int32{2, 12, 16, 17, 32},
			shopsPostAirship: []int32{2, 12, 16, 17, 32},
			treasures:        []int32{76, 101, 105},
			equipmentTables:  []int32{76, 78},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/45?rel_availability=always",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters drop":       0,
					"monsters items":      2,
					"shops pre airship":   0,
					"shops post airship":  0,
					"treasures":           0,
					"equipment tables":    2,
				},
			},
			expUnique:    newExpUnique(45, "triple ap"),
			monstersDrop: []int32{},
			monstersItems: []testMonItemAmts{
				{
					index:   0,
					monster: 142,
					bribe:   4,
				},
				{
					index:      1,
					monster:    291,
					dropCommon: 1,
				},
			},
			shopsPreAirship:  []int32{},
			shopsPostAirship: []int32{},
			treasures:        []int32{},
			equipmentTables:  []int32{15, 20},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/55?rel_availability=story&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters drop":       3,
					"monsters items":      0,
					"shops pre airship":   2,
					"shops post airship":  0,
					"treasures":           0,
					"equipment tables":    1,
				},
			},
			expUnique:        newExpUnique(55, "fire ward"),
			monstersDrop:     []int32{109, 118, 119},
			monstersItems:    []testMonItemAmts{},
			shopsPreAirship:  []int32{4, 10},
			shopsPostAirship: []int32{},
			treasures:        []int32{},
			equipmentTables:  []int32{137},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/55?rel_availability=post-game&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters drop":       8,
					"monsters items":      2,
					"shops pre airship":   0,
					"shops post airship":  0,
					"treasures":           0,
					"equipment tables":    1,
				},
			},
			expUnique:        newExpUnique(55, "fire ward"),
			monstersDrop:     []int32{38, 53, 141, 151, 176, 191, 192, 239},
			monstersItems:    []testMonItemAmts{
				{
					index:       0,
					monster:     38,
					stealCommon: 2,
					stealRare:   3,
				},
				{
					index:       1,
					monster:     53,
					stealCommon: 1,
					stealRare:   2,
				},
			},
			shopsPreAirship:  []int32{},
			shopsPostAirship: []int32{},
			treasures:        []int32{},
			equipmentTables:  []int32{137},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/10",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters drop":       9,
					"monsters items":      9,
					"shops pre airship":   0,
					"shops post airship":  1,
					"treasures":           0,
					"equipment tables":    3,
				},
			},
			expUnique:    newExpUnique(10, "darkstrike"),
			monstersDrop: []int32{185, 200, 257, 279, 285},
			monstersItems: []testMonItemAmts{
				{
					index:     0,
					monster:   11,
					stealRare: 1,
				},
				{
					index:       5,
					monster:     99,
					stealCommon: 1,
					stealRare:   2,
				},
				{
					index:       8,
					monster:     271,
					stealCommon: 4,
				},
			},
			shopsPreAirship:  []int32{},
			shopsPostAirship: []int32{7},
			treasures:        []int32{},
			equipmentTables:  []int32{25, 32, 46},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/89?rel_availability=pre-airship",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters drop":       9,
					"monsters items":      4,
					"shops pre airship":   1,
					"shops post airship":  1,
					"treasures":           2,
					"equipment tables":    2,
				},
			},
			expUnique:    newExpUnique(89, "sos nulblaze"),
			monstersDrop: []int32{31, 71, 102, 143, 158, 181},
			monstersItems: []testMonItemAmts{
				{
					index:     	0,
					monster:   	38,
					bribe: 		16,
				},
				{
					index:     	1,
					monster:   	53,
					bribe: 		8,
				},
				{
					index:     	2,
					monster:   	68,
					other: 		2,
				},
				{
					index:       3,
					monster:     109,
					stealCommon: 2,
					stealRare:   3,
				},
			},
			shopsPreAirship:  []int32{14},
			shopsPostAirship: []int32{14},
			treasures:        []int32{61, 291},
			equipmentTables:  []int32{102, 142},
		},
	}

	testSingleResources(t, tests, "GetAutoAbility", testCfg.HandleAutoAbilities, compareAutoAbilities)
}

func TestRetrieveAutoAbilities(t *testing.T) {
	t.Parallel()
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   130,
			results: []int32{1, 33, 67, 89, 101, 111, 121, 130},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?category=1&type=armor",
				expectedStatus: http.StatusOK,
			},
			count:   16,
			results: []int32{103, 106, 110, 112, 114, 118},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?monster=294",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{5, 51, 57, 128, 129},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?monster_items=294",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{51, 113, 117, 121, 129},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?availability=post&req_item=true",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{32, 35, 36, 42, 88, 105, 118},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?availability=post&req_item=false&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   22,
			results: []int32{3, 38, 42, 50, 84, 98, 110, 121, 127},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?availability=always&monster=85&character=1",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{5, 6, 7, 8, 27, 73, 116},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?availability=pre-story&character=5&area=215",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{2, 10, 93},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?availability=always&sublocation=25&character=2&methods=treasure",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{69, 83},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?availability=post&location=4",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{3, 38, 43, 51, 59, 102, 128, 129},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities?availability=post&area=239&repeatable=false",
				expectedStatus: http.StatusOK,
			},
			count:   9,
			results: []int32{37, 39, 60, 74, 127},
		},
	}

	testIdList(t, tests, testCfg.e.autoAbilities.endpoint, "RetrieveAutoAbilities", testCfg.HandleAutoAbilities, compareAPIResourceLists[NamedApiResourceList])
}
