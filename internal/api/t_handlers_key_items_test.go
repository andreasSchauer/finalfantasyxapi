package api

import (
	"net/http"
	"testing"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetKeyItem(t *testing.T) {
	tests := []expKeyItem{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items/61",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "key-item with provided id '61' doesn't exist. max id: 60.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items/39",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"areas": 		2,
					"treasures": 	2,
					"quests": 		0,
				},
			},
			expUnique:    		newExpUnique(39, "al bhed primer v"),
			untypedItem: 		151,
			category: 			3,
			celestialWeapon: 	nil,
			primer: 			h.GetInt32Ptr(5),
			areas: 				[]int32{65, 170},
			treasures: 			[]int32{68, 220},
			quests: 			[]int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items/23",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"areas": 		1,
					"treasures": 	1,
					"quests": 		0,
				},
			},
			expUnique:    		newExpUnique(23, "mercury crest"),
			untypedItem: 		135,
			category: 			2,
			celestialWeapon: 	h.GetInt32Ptr(7),
			primer: 			nil,
			areas: 				[]int32{172},
			treasures: 			[]int32{231},
			quests: 			[]int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items/8",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"areas": 		1,
					"treasures": 	0,
					"quests": 		1,
				},
			},
			expUnique:    		newExpUnique(8, "mark of conquest"),
			untypedItem: 		120,
			category: 			5,
			celestialWeapon: 	nil,
			primer: 			nil,
			areas: 				[]int32{205},
			treasures: 			[]int32{},
			quests: 			[]int32{1},
		},
	}

	testSingleResources(t, tests, "GetKeyItem", testCfg.HandleKeyItems, compareKeyItems)
}

func TestRetrieveKeyItems(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   60,
			results: []int32{1, 16, 27, 48, 54, 60},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?availability=post&category=2",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{14, 16, 18, 24},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?method=quest",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{3, 8, 9, 16, 22, 24},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?method=a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:	"invalid value 'a' used for parameter 'method'. allowed values: 'treasure', 'shop'.",
			},
		},
	}

	testIdList(t, tests, testCfg.e.keyItems.endpoint, "RetrieveKeyItems", testCfg.HandleKeyItems, compareAPIResourceLists[NamedApiResourceList])
}
