package api

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetSphere(t *testing.T) {
	tests := []expSphere{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres/31",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sphere with provided id '31' doesn't exist. max id: 30.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres/4?rel_availability=pre-story&repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         10,
					"treasures":        0,
					"shops":            0,
					"quests":           0,
					"blitzball prizes": 1,
				},
			},
			expUnique:   newExpUnique(4, "ability sphere"),
			item:        73,
			createdNode: nil,
			monsters: []testMonItemAmts{
				{
					index:      3,
					monster:    96,
					dropCommon: 1,
					dropRare:   1,
				},
				{
					index:      9,
					monster:    123,
					dropCommon: 1,
					dropRare:   1,
				},
			},
			treasures: nil,
			shops:     []int32{},
			quests:    nil,
			blitzballPrizes: map[int32]int32{
				3: 1,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres/16",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         3,
					"treasures":        4,
					"shops":            0,
					"quests":           1,
					"blitzball prizes": 0,
				},
			},
			expUnique: newExpUnique(16, "hp sphere"),
			item:      85,
			createdNode: &CreatedNode{
				Node:  "hp",
				Value: 300,
			},
			monsters: []testMonItemAmts{
				{
					index:    0,
					monster:  28,
					dropRare: 1,
				},
				{
					index:      1,
					monster:    196,
					dropCommon: 1,
					dropRare:   1,
				},
				{
					index:      2,
					monster:    282,
					dropCommon: 1,
				},
			},
			treasures: map[int32]int32{
				78:  1,
				253: 1,
				288: 1,
				313: 1,
			},
			shops: []int32{},
			quests: map[int32]int32{
				89: 3,
			},
			blitzballPrizes: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres/26",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         0,
					"treasures":        0,
					"shops":            1,
					"quests":           0,
					"blitzball prizes": 0,
				},
			},
			expUnique:       newExpUnique(26, "clear sphere"),
			item:            95,
			createdNode:     nil,
			monsters:        []testMonItemAmts{},
			treasures:       nil,
			shops:           []int32{33},
			quests:          nil,
			blitzballPrizes: nil,
		},
	}

	testSingleResources(t, tests, "GetSphere", testCfg.HandleSpheres, compareSpheres)
}

func TestRetrieveSpheres(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   30,
			results: []int32{1, 13, 15, 19, 23, 27, 30},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?color=yellow&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{6, 7, 8, 9, 10, 11},
		},
	}

	testIdList(t, tests, testCfg.e.spheres.endpoint, "RetrieveSpheres", testCfg.HandleSpheres, compareAPIResourceLists[NamedApiResourceList])
}
