package api

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetAutoAbility(t *testing.T) {
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
					"monstersDrop":         1,
					"monstersItems":		4,
					"shops pre airship":    0,
					"shops post airship":   0,
					"treasures":            0,
					"equipment tables": 	2,
				},
			},
			expUnique:   newExpUnique(50, "gillionaire"),
			monstersDrop: 		[]int32{302},
			monstersItems: 		[]testMonItemAmts{
				{
					index:   	0,
					monster: 	190,
					bribe:   	5,
				},
				{
					index:   	1,
					monster: 	261,
					stealRare: 	1,
				},
				{
					index:   	2,
					monster: 	265,
					stealRare:  1,
				},
				{
					index:   	3,
					monster: 	285,
					dropCommon: 1,
				},
			},
			shopsPreAirship:    []int32{},
			shopsPostAirship:   []int32{},
			treasures: 			[]int32{},
			equipmentTables: 	[]int32{7, 30},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/45?rel_availability=always",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monstersDrop":         0,
					"monstersItems":		2,
					"shops pre airship":    0,
					"shops post airship":   0,
					"treasures":            0,
					"equipment tables": 	2,
				},
			},
			expUnique:   newExpUnique(45, "triple ap"),
			monstersDrop: 		[]int32{},
			monstersItems: 		[]testMonItemAmts{
				{
					index:   	0,
					monster: 	142,
					bribe:   	4,
				},
				{
					index:   	1,
					monster: 	291,
					dropCommon: 1,
				},
			},
			shopsPreAirship:    []int32{},
			shopsPostAirship:   []int32{},
			treasures: 			[]int32{},
			equipmentTables: 	[]int32{15, 20},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/55?rel_availability=story&repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monstersDrop":         3,
					"monstersItems":		0,
					"shops pre airship":    2,
					"shops post airship":   0,
					"treasures":            1,
					"equipment tables": 	1,
				},
			},
			expUnique:   newExpUnique(55, "fire ward"),
			monstersDrop: 		[]int32{109, 118, 119},
			monstersItems: 		[]testMonItemAmts{},
			shopsPreAirship:    []int32{4, 10},
			shopsPostAirship:   []int32{},
			treasures: 			[]int32{89},
			equipmentTables: 	[]int32{137},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/auto-abilities/10",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monstersDrop":         9,
					"monstersItems":		9,
					"shops pre airship":    0,
					"shops post airship":   1,
					"treasures":            0,
					"equipment tables": 	3,
				},
			},
			expUnique:   newExpUnique(10, "darkstrike"),
			monstersDrop: 		[]int32{185, 200, 257, 279, 285},
			monstersItems: 		[]testMonItemAmts{
				{
					index:   	 0,
					monster: 	 11,
					stealRare: 	 1,
				},
				{
					index:   	 5,
					monster: 	 99,
					stealCommon: 1,
					stealRare: 	 2,
				},
				{
					index:   	 8,
					monster: 	 271,
					stealCommon: 4,
				},
			},
			shopsPreAirship:    []int32{},
			shopsPostAirship:   []int32{7},
			treasures: 			[]int32{},
			equipmentTables: 	[]int32{25, 32, 46},
		},
	}

	testSingleResources(t, tests, "GetAutoAbility", testCfg.HandleAutoAbilities, compareAutoAbilities)
}

func TestRetrieveAutoAbilities(t *testing.T) {
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
	}

	testIdList(t, tests, testCfg.e.autoAbilities.endpoint, "RetrieveAutoAbilities", testCfg.HandleAutoAbilities, compareAPIResourceLists[NamedApiResourceList])
}
