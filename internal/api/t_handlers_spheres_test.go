package api

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetSphere(t *testing.T) {
	t.Parallel()
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
				requestURL:     "/api/spheres/4?rel_availability=pre-story&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         10,
					"treasures":        0,
					"shops":            0,
					"quests":           0,
					"blitzball prizes": 0,
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
			blitzballPrizes: nil,
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
				92: 3,
			},
			blitzballPrizes: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres/16?rel_availability=post-game&rel_repeatable=false",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         1,
					"treasures":        3,
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
			},
			treasures: map[int32]int32{
				78:  1,
				288: 1,
				313: 1,
			},
			shops: []int32{},
			quests: map[int32]int32{
				92: 3,
			},
			blitzballPrizes: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres/26?rel_availability=post-game&rel_repeatable=true",
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
			expUnique: newExpUnique(26, "clear sphere"),
			item:      95,
			createdNode: nil,
			monsters: []testMonItemAmts{},
			treasures: nil,
			shops: []int32{33},
			quests: nil,
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
	t.Parallel()
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
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=always&location=12&repeatable=true",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{1, 2, 3},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=post&repeatable=true",
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{7, 10, 16, 21, 24, 30},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=post-game&repeatable=true&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   30,
			results: []int32{1, 5, 10, 20, 25, 29, 30},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=always&repeatable=true",
				expectedStatus: http.StatusOK,
			},
			count:   16,
			results: []int32{1, 5, 11, 15, 27, 29},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=always&repeatable=false",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{16, 17, 18, 20, 21, 25},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=always&repeatable=false&methods=treasure",
				expectedStatus: http.StatusOK,
			},
			count:   15,
			results: []int32{2, 4, 12, 15, 21, 25, 29},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=always&repeatable=false&methods=monster",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{16, 17},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=always&repeatable=false&methods=monster,treasure",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{16, 17, 20, 21, 25, 29},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=story&sublocation=27",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{4, 7, 10},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=story&repeatable=true&sublocation=27",
				expectedStatus: http.StatusOK,
			},
			count:   0,
			results: []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/spheres?availability=post&repeatable=true&area=205",
				expectedStatus: http.StatusOK,
			},
			count:   17,
			results: []int32{13, 15, 17, 19, 22, 25, 26, 30},
		},
	}

	testIdList(t, tests, testCfg.e.spheres.endpoint, "RetrieveSpheres", testCfg.HandleSpheres, compareAPIResourceLists[NamedApiResourceList])
}
