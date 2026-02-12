package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetTreasure(t *testing.T) {
	tests := []expTreasure{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/344",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "treasure with provided id '344' doesn't exist. max id: 343.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/3/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "endpoint /treasures doesn't have any subsections.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/a/3",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. usage: '/api/treasures', '/api/treasures/{id}'",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/2/2",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid subsection '2'. subsection can't be an integer. use /api/treasures/sections for valid subsections.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/200",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths:     map[string]int{},
			},
			expIdOnly:       newExpIdOnly(200),
			area:            160,
			isPostAirship:   false,
			isAnimaTreasure: false,
			treasureType:    1,
			lootType:        1,
			gilAmount:       nil,
			items: []testItemAmount{
				newTestItemAmount("/items/82", 1),
			},
			equipment: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/13",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths:     map[string]int{},
			},
			expIdOnly:       newExpIdOnly(13),
			area:            15,
			isPostAirship:   false,
			isAnimaTreasure: false,
			treasureType:    3,
			lootType:        1,
			gilAmount:       nil,
			items: []testItemAmount{
				newTestItemAmount("/key-items/35", 1),
			},
			equipment: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/62",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths:     map[string]int{},
			},
			expIdOnly:       newExpIdOnly(62),
			area:            61,
			isPostAirship:   false,
			isAnimaTreasure: true,
			treasureType:    1,
			lootType:        2,
			gilAmount:       nil,
			items:           nil,
			equipment: &testFoundEquipment{
				equipmentName:    877,
				abilities:        []int32{55, 64, 58},
				emptySlotsAmount: 0,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/285",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths:     map[string]int{},
			},
			expIdOnly:       newExpIdOnly(285),
			area:            214,
			isPostAirship:   false,
			isAnimaTreasure: false,
			treasureType:    1,
			lootType:        3,
			gilAmount:       h.GetInt32Ptr(20000),
			items:           nil,
			equipment:       nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/45",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths:     map[string]int{},
			},
			expIdOnly:       newExpIdOnly(45),
			area:            42,
			isPostAirship:   true,
			isAnimaTreasure: false,
			treasureType:    2,
			lootType:        1,
			gilAmount:       nil,
			items: []testItemAmount{
				newTestItemAmount("/items/97", 1),
			},
			equipment: nil,
		},
	}

	testSingleResources(t, tests, "GetTreasure", testCfg.HandleTreasures, compareTreasures)
}

func TestRetrieveTreasures(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   343,
			results: []int32{1, 28, 53, 84, 123, 127, 167, 191, 214, 286, 300, 343},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures?location=4&loot_type=item&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   21,
			results: []int32{15, 18, 20, 23, 35, 37, 41, 44},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures?sublocation=41&treasure_type=chest",
				expectedStatus: http.StatusOK,
			},
			count:   15,
			results: []int32{329, 330, 335, 337, 339, 343},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures?area=214",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{285, 286, 287, 288, 289, 290},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures?airship=true&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   69,
			results: []int32{5, 11, 29, 30, 45, 63, 81, 109, 153, 213, 303, 318, 341},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures?anima=true",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{39, 62, 150, 211, 254, 303},
		},
	}

	testIdList(t, tests, testCfg.e.treasures.endpoint, "RetrieveTreasures", testCfg.HandleTreasures, compareAPIResourceLists[UnnamedApiResourceList])
}

func TestSubsectionTreasures(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/6/treasures",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleLocations,
			},
			count:          17,
			parentResource: h.GetStrPtr("/locations/6"),
			results:        []int32{50, 53, 55, 59, 62, 65, 66},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/20/treasures",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleSublocations,
			},
			count:          5,
			parentResource: h.GetStrPtr("/sublocations/20"),
			results:        []int32{167, 168, 169, 170, 171},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/225/treasures",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          3,
			parentResource: h.GetStrPtr("/areas/225"),
			results:        []int32{299, 300, 301},
		},
	}

	testIdList(t, tests, testCfg.e.treasures.endpoint, "SubsectionTreasures", nil, compareSubResourceLists[UnnamedAPIResource, TreasureSub])
}
