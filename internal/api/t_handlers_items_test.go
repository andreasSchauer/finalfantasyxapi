package api

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetItem(t *testing.T) {
	tests := []expItem{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/113",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "item with provided id '113' doesn't exist. max id: 112.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/69",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"monsters": 			5,
					"treasures": 			0,
					"shops": 				0,
					"quests": 				2,
					"blitzball prizes": 	1,
					"aeon learn abilities": 1,
					"auto abilities": 		2,
					"mixes": 				21,
				},
			},
			expUnique:    	newExpUnique(69, "three stars"),
			untypedItem: 	69,
			category: 		3,
			monsters: []testMonItemAmts{
				{
					index: 			0,
					monster: 		203,
					steal: 			nil,
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			14,
					other: 			0,
				},
				{
					index: 			2,
					monster: 		266,
					steal: 			nil,
					drop: 			&CommonRareAmount{
						Common: 1,
					},
					secondaryDrop: 	nil,
					bribe: 			0,
					other: 			0,
				},
				{
					index: 			4,
					monster: 		301,
					steal: 			&CommonRareAmount{
						Common: 2,
					},
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			0,
					other: 			0,
				},
			},
			treasures: nil,
			shops: []int32{},
			quests: map[int32]int32{
				29: 60,
				61: 60,
			},
			blitzballPrizes: map[int32]int32{
				1: 1,
			},
			aeonLearnAbilities: map[int32]int32{
				43: 5,
			},
			autoAbilities: map[string]int32{
				"one mp cost": 20,
				"break mp limit": 30,
			},
			mixes: []int32{6, 12, 26, 49, 62, 64},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/69?repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"untyped item": true,
					"category": true,
					"treasures": true,
					"shops": true,
					"blitzball prizes": true,
					"auto abilities": true,
					"aeon learn abilities": true,
					"mixes": true,
				},
				expLengths: 	map[string]int{
					"monsters": 			4,
					"quests": 				0,
				},
			},
			expUnique:    	newExpUnique(69, "three stars"),
			monsters: []testMonItemAmts{
				{
					index: 			0,
					monster: 		203,
					steal: 			nil,
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			14,
					other: 			0,
				},
				{
					index: 			2,
					monster: 		266,
					steal: 			nil,
					drop: 			&CommonRareAmount{
						Common: 1,
					},
					secondaryDrop: 	nil,
					bribe: 			0,
					other: 			0,
				},
				{
					index: 			3,
					monster: 		291,
					steal: 			&CommonRareAmount{
						Rare: 1,
					},
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			0,
					other: 			0,
				},
			},
			quests: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/26",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"monsters": 			5,
					"treasures": 			0,
					"shops": 				0,
					"quests": 				0,
					"blitzball prizes": 	0,
					"aeon learn abilities": 1,
					"auto abilities": 		1,
					"mixes": 				24,
				},
			},
			expUnique:    	newExpUnique(26, "fire gem"),
			untypedItem: 	26,
			category: 		2,
			monsters: []testMonItemAmts{
				{
					index: 			0,
					monster: 		109,
					steal: 			nil,
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			14,
					other: 			0,
				},
				{
					index: 			1,
					monster: 		141,
					steal: 			&CommonRareAmount{
						Common: 1,
						Rare: 	2,
					},
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			10,
					other: 			0,
				},
			},
			treasures: nil,
			shops: []int32{},
			quests: nil,
			blitzballPrizes: nil,
			aeonLearnAbilities: map[int32]int32{
				77: 4,
			},
			autoAbilities: map[string]int32{
				"fire eater": 20,
			},
			mixes: []int32{1, 6, 15, 31, 55, 63},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/2",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"monsters": 			37,
					"treasures": 			26,
					"shops": 				22,
					"quests": 				0,
					"blitzball prizes": 	2,
					"aeon learn abilities": 1,
					"auto abilities": 		0,
					"mixes": 				16,
				},
			},
			expUnique:    	newExpUnique(2, "hi-potion"),
			untypedItem: 	2,
			category: 		1,
			monsters: []testMonItemAmts{
				{
					index: 			1,
					monster: 		39,
					steal: 			&CommonRareAmount{
						Rare: 1,
					},
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			60,
					other: 			0,
				},
				{
					index: 			36,
					monster: 		242,
					steal: 			&CommonRareAmount{
						Common: 2,
					},
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			0,
					other: 			0,
				},
			},
			treasures: map[int32]int32{
				4: 1,
				33: 1,
				73: 2,
				142: 2,
				219: 4,
				232: 8,
			},
			shops: []int32{1, 7, 12, 20, 25, 28, 32, 36},
			quests: nil,
			blitzballPrizes: map[int32]int32{
				2: 2,
				5: 1,
			},
			aeonLearnAbilities: map[int32]int32{
				45: 99,
			},
			autoAbilities: nil,
			mixes: []int32{34, 40, 51, 56, 61},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/64",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"monsters": 			10,
					"treasures": 			0,
					"shops": 				0,
					"quests": 				1,
					"blitzball prizes": 	0,
					"aeon learn abilities": 0,
					"auto abilities": 		2,
					"mixes": 				9,
				},
			},
			expUnique:    	newExpUnique(64, "stamina tablet"),
			untypedItem: 	64,
			category: 		3,
			monsters: []testMonItemAmts{
				{
					index: 			0,
					monster: 		69,
					steal: 			nil,
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			0,
					other: 			1,
				},
				{
					index: 			2,
					monster: 		156,
					steal: 			nil,
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			0,
					other: 			1,
				},
			},
			treasures: nil,
			shops: []int32{},
			quests: map[int32]int32{
				20: 60,
			},
			blitzballPrizes: nil,
			aeonLearnAbilities: nil,
			autoAbilities: map[string]int32{
				"auto-potion": 4,
				"hp stroll": 2,
			},
			mixes: []int32{42, 56, 60, 63, 64},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/64?rel_availability=story",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"monsters": 			1,
					"treasures": 			0,
					"shops": 				0,
					"quests": 				0,
					"blitzball prizes": 	0,
					"aeon learn abilities": 0,
					"auto abilities": 		2,
					"mixes": 				9,
				},
			},
			expUnique:    	newExpUnique(64, "stamina tablet"),
			untypedItem: 	64,
			category: 		3,
			monsters: []testMonItemAmts{
				{
					index: 			0,
					monster: 		195,
					steal: 			&CommonRareAmount{
						Common: 1,
					},
					drop: 			nil,
					secondaryDrop: 	nil,
					bribe: 			0,
					other: 			0,
				},
			},
			treasures: nil,
			shops: []int32{},
			quests: nil,
			blitzballPrizes: nil,
			aeonLearnAbilities: nil,
			autoAbilities: map[string]int32{
				"auto-potion": 4,
				"hp stroll": 2,
			},
			mixes: []int32{42, 56, 60, 63, 64},
		},
	}

	testSingleResources(t, tests, "GetItem", testCfg.HandleItems, compareItems)
}

func TestRetrieveItems(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   112,
			results: []int32{1, 42, 67, 78, 90, 110, 112},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?category=healing,support&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   31,
			results: []int32{1, 5, 11, 21, 23, 54, 59, 64, 69},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?category=healing,support&limit=max&flip=true",
				expectedStatus: http.StatusOK,
			},
			count:   81,
			results: []int32{17, 29, 41, 49, 73, 86, 106, 112},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?method=treasure,shop",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{1, 2, 9, 11, 12, 13, 16, 21},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?has_ability=true&related_stat=mp",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{65, 67, 68, 69},
		},
	}

	testIdList(t, tests, testCfg.e.items.endpoint, "RetrieveItems", testCfg.HandleItems, compareAPIResourceLists[NamedApiResourceList])
}
