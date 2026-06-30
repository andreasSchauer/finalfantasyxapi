package api

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetItem(t *testing.T) {
	t.Parallel()
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
				expLengths: map[string]int{
					"monsters":             5,
					"treasures":            0,
					"shops":                0,
					"quests":               2,
					"blitzball prizes":     1,
					"aeon learn abilities": 1,
					"auto abilities":       2,
					"mixes":                21,
				},
			},
			expUnique:   newExpUnique(69, "three stars"),
			untypedItem: 69,
			category:    database.ItemCategorySupport,
			monsters: []testMonItemAmts{
				{
					index:   0,
					monster: 203,
					bribe:   14,
				},
				{
					index:      2,
					monster:    266,
					dropCommon: 1,
				},
				{
					index:       4,
					monster:     301,
					stealCommon: 2,
				},
			},
			treasures: nil,
			shops:     []int32{},
			quests: map[int32]int32{
				38: 60,
				68: 60,
			},
			blitzballPrizes: map[int32]int32{
				1: 1,
			},
			aeonLearnAbilities: map[int32]int32{
				43: 5,
			},
			autoAbilities: map[string]int32{
				"one mp cost":    20,
				"break mp limit": 30,
			},
			mixes: []int32{6, 12, 26, 49, 62, 64},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/69?rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"untyped item":         true,
					"category":             true,
					"treasures":            true,
					"shops":                true,
					"blitzball prizes":     true,
					"auto abilities":       true,
					"aeon learn abilities": true,
					"mixes":                true,
				},
				expLengths: map[string]int{
					"monsters": 4,
					"quests":   0,
				},
			},
			expUnique: newExpUnique(69, "three stars"),
			monsters: []testMonItemAmts{
				{
					index:   0,
					monster: 203,
					bribe:   14,
				},
				{
					index:      2,
					monster:    266,
					dropCommon: 1,
				},
				{
					index:     3,
					monster:   291,
					stealRare: 1,
				},
			},
			quests: nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/26",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":             5,
					"treasures":            0,
					"shops":                0,
					"quests":               0,
					"blitzball prizes":     0,
					"aeon learn abilities": 1,
					"auto abilities":       1,
					"mixes":                24,
				},
			},
			expUnique:   newExpUnique(26, "fire gem"),
			untypedItem: 26,
			category:    database.ItemCategoryOffensive,
			monsters: []testMonItemAmts{
				{
					index:   0,
					monster: 109,
					bribe:   14,
				},
				{
					index:       1,
					monster:     141,
					stealCommon: 1,
					stealRare:   2,
					bribe:       10,
				},
			},
			treasures:       nil,
			shops:           []int32{},
			quests:          nil,
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
				expLengths: map[string]int{
					"monsters":             37,
					"treasures":            26,
					"shops":                22,
					"quests":               0,
					"blitzball prizes":     2,
					"aeon learn abilities": 1,
					"auto abilities":       0,
					"mixes":                16,
				},
			},
			expUnique:   newExpUnique(2, "hi-potion"),
			untypedItem: 2,
			category:    database.ItemCategoryHealing,
			monsters: []testMonItemAmts{
				{
					index:     1,
					monster:   39,
					stealRare: 1,
					bribe:     60,
				},
				{
					index:       36,
					monster:     242,
					stealCommon: 2,
				},
			},
			treasures: map[int32]int32{
				4:   1,
				33:  1,
				73:  2,
				142: 2,
				219: 4,
				232: 8,
			},
			shops:  []int32{1, 7, 12, 20, 25, 28, 32, 36},
			quests: nil,
			blitzballPrizes: map[int32]int32{
				2: 2,
				5: 1,
			},
			aeonLearnAbilities: map[int32]int32{
				45: 99,
			},
			autoAbilities: nil,
			mixes:         []int32{34, 40, 51, 56, 61},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/64",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":             10,
					"treasures":            0,
					"shops":                0,
					"quests":               1,
					"blitzball prizes":     0,
					"aeon learn abilities": 0,
					"auto abilities":       2,
					"mixes":                9,
				},
			},
			expUnique:   newExpUnique(64, "stamina tablet"),
			untypedItem: 64,
			category:    database.ItemCategorySupport,
			monsters: []testMonItemAmts{
				{
					index:   0,
					monster: 69,
					other:   1,
				},
				{
					index:   2,
					monster: 156,
					other:   1,
				},
			},
			treasures: nil,
			shops:     []int32{},
			quests: map[int32]int32{
				29: 60,
			},
			blitzballPrizes:    nil,
			aeonLearnAbilities: nil,
			autoAbilities: map[string]int32{
				"auto-potion": 4,
				"hp stroll":   2,
			},
			mixes: []int32{42, 56, 60, 63, 64},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/64?rel_availability=pre-story",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":             1,
					"treasures":            0,
					"shops":                0,
					"quests":               0,
					"blitzball prizes":     0,
					"aeon learn abilities": 0,
					"auto abilities":       2,
					"mixes":                9,
				},
			},
			expUnique:   newExpUnique(64, "stamina tablet"),
			untypedItem: 64,
			category:    database.ItemCategorySupport,
			monsters: []testMonItemAmts{
				{
					index:       0,
					monster:     195,
					stealCommon: 1,
				},
			},
			treasures:          nil,
			shops:              []int32{},
			quests:             nil,
			blitzballPrizes:    nil,
			aeonLearnAbilities: nil,
			autoAbilities: map[string]int32{
				"auto-potion": 4,
				"hp stroll":   2,
			},
			mixes: []int32{42, 56, 60, 63, 64},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/55?rel_availability=post",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"untyped item": true,
					"category": true,
					"aeon learn abilities": true,
					"auto abilities": true,
					"mixes": true,
				},
				expLengths: map[string]int{
					"monsters":             3,
					"treasures":            0,
					"shops":                0,
					"quests":               0,
					"blitzball prizes":     0,
				},
			},
			expUnique:   newExpUnique(55, "chocobo wing"),
			monsters: []testMonItemAmts{
				{
					index:       0,
					monster:     242,
					bribe: 		 60,
				},
				{
					index:       1,
					monster:     269,
					stealRare: 	 1,
				},
				{
					index:       2,
					monster:     270,
					stealRare:   1,
				},
			},
			treasures:          nil,
			shops:              []int32{},
			quests:             nil,
			blitzballPrizes:    nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/55?rel_availability=always",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"untyped item": true,
					"category": true,
					"aeon learn abilities": true,
					"auto abilities": true,
					"mixes": true,
				},
				expLengths: map[string]int{
					"monsters":             2,
					"treasures":            0,
					"shops":                0,
					"quests":               1,
					"blitzball prizes":     0,
				},
			},
			expUnique:   newExpUnique(55, "chocobo wing"),
			monsters: []testMonItemAmts{
				{
					index:       0,
					monster:     100,
					stealRare: 	 1,
				},
				{
					index:       1,
					monster:     261,
					stealCommon: 2,
				},
			},
			treasures:          nil,
			shops:              []int32{},
			quests:             map[int32]int32{
				16: 99,
			},
			blitzballPrizes:    nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/55?rel_availability=always&rel_repeatable=false",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"untyped item": true,
					"category": true,
					"aeon learn abilities": true,
					"auto abilities": true,
					"mixes": true,
				},
				expLengths: map[string]int{
					"monsters":             0,
					"treasures":            0,
					"shops":                0,
					"quests":               1,
					"blitzball prizes":     0,
				},
			},
			expUnique:   newExpUnique(55, "chocobo wing"),
			monsters: []testMonItemAmts{},
			treasures:          nil,
			shops:              []int32{},
			quests:             map[int32]int32{
				16: 99,
			},
			blitzballPrizes:    nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/55?rel_availability=always&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"untyped item": true,
					"category": true,
					"aeon learn abilities": true,
					"auto abilities": true,
					"mixes": true,
				},
				expLengths: map[string]int{
					"monsters":             2,
					"treasures":            0,
					"shops":                0,
					"quests":               0,
					"blitzball prizes":     0,
				},
			},
			expUnique:   newExpUnique(55, "chocobo wing"),
			monsters: []testMonItemAmts{
				{
					index:       0,
					monster:     100,
					stealRare: 	 1,
				},
				{
					index:       1,
					monster:     261,
					stealCommon: 2,
				},
			},
			treasures:          nil,
			shops:              []int32{},
			quests:             nil,
			blitzballPrizes:    nil,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/2?rel_availability=always",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"untyped item": true,
					"category": true,
					"aeon learn abilities": true,
					"auto abilities": true,
					"mixes": true,
				},
				expLengths: map[string]int{
					"monsters":             22,
					"treasures":            11,
					"shops":                4,
					"quests":               0,
					"blitzball prizes":     2,
				},
			},
			expUnique:   newExpUnique(2, "hi-potion"),
			monsters: []testMonItemAmts{
				{
					index:       0,
					monster:     39,
					stealRare: 	 1,
					bribe: 		 60,
				},
				{
					index:       8,
					monster:     87,
					stealCommon: 1,
				},
				{
					index:       18,
					monster:     162,
					stealCommon: 1,
					stealRare: 	 2,
				},
			},
			treasures:          map[int32]int32{
				4: 	 1,
				103: 2,
				219: 4,
				224: 4,
				232: 8,
			},
			shops:              []int32{24, 32, 33, 34},
			quests:             nil,
			blitzballPrizes:    map[int32]int32{
				2: 2,
				5: 1,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/2?rel_availability=always&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"untyped item": true,
					"category": true,
					"aeon learn abilities": true,
					"auto abilities": true,
					"mixes": true,
				},
				expLengths: map[string]int{
					"monsters":             22,
					"treasures":            0,
					"shops":                4,
					"quests":               0,
					"blitzball prizes":     2,
				},
			},
			expUnique:   newExpUnique(2, "hi-potion"),
			monsters: []testMonItemAmts{
				{
					index:       0,
					monster:     39,
					stealRare: 	 1,
					bribe: 		 60,
				},
				{
					index:       8,
					monster:     87,
					stealCommon: 1,
				},
				{
					index:       18,
					monster:     162,
					stealCommon: 1,
					stealRare: 	 2,
				},
			},
			treasures:          nil,
			shops:              []int32{24, 32, 33, 34},
			quests:             nil,
			blitzballPrizes:    map[int32]int32{
				2: 2,
				5: 1,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items/2?rel_availability=post",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{
					"untyped item": true,
					"category": true,
					"aeon learn abilities": true,
					"auto abilities": true,
					"mixes": true,
				},
				expLengths: map[string]int{
					"monsters":             2,
					"treasures":            1,
					"shops":                9,
					"quests":               0,
					"blitzball prizes":     0,
				},
			},
			expUnique:   newExpUnique(2, "hi-potion"),
			monsters: []testMonItemAmts{
				{
					index:       0,
					monster:     239,
					stealCommon: 1,
				},
				{
					index:       1,
					monster:     242,
					stealCommon: 2,
				},
			},
			treasures:          map[int32]int32{
				44: 1,
			},
			shops:              []int32{1, 5, 6, 7, 12, 20, 21, 27, 36},
			quests:             nil,
			blitzballPrizes:    nil,
		},
	}

	testSingleResources(t, tests, "GetItem", testCfg.HandleItems, compareItems)
}

func TestRetrieveItems(t *testing.T) {
	t.Parallel()
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?methods=shop,treasures",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'treasures' used for parameter 'methods'. allowed values: 'monster', 'treasure', 'shop', 'quest', 'blitzball'.",
			},
		},
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
				requestURL:     "/api/items?methods=treasure,blitzball&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   50,
			results: []int32{1, 2, 9, 11, 12, 13, 16, 21, 71, 85, 101, 111},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?has_ability=true&related_stat=mp",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{65, 67, 68, 69},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?location=17&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   22,
			results: []int32{1, 7, 21, 34, 76, 102, 111},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?sublocation=12",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{1, 2, 9, 11},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?area=240&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   31,
			results: []int32{4, 23, 60, 89, 110},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=post",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{23, 45, 67, 68, 76, 79, 88, 99},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=post-story&location=18&methods=monster",
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{3, 5, 43, 45, 58, 61, 83, 96},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=post-game&repeatable=false&pre_airship=false",
				expectedStatus: http.StatusOK,
			},
			count:   0,
			results: []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=post-game&repeatable=false&pre_airship=true",
				expectedStatus: http.StatusOK,
			},
			count:   0,
			results: []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=post-game&repeatable=true&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   112,
			results: []int32{1, 17, 19, 38, 54, 67, 88, 100, 112},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=post&repeatable=true",
				expectedStatus: http.StatusOK,
			},
			count:   19,
			results: []int32{23, 45, 67, 68, 76, 88, 90, 95, 105},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=always&repeatable=true&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   93,
			results: []int32{1, 7, 21, 37, 56, 69, 80, 104, 112},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=always&repeatable=false",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{85, 86, 87, 89, 90, 94, 105},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=story&sublocation=25",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{6, 82},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=post-game&area=23",
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{1, 9, 30, 31, 32, 38, 41, 70, 71, 72},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=always&methods=shop",
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{1, 2, 9, 14, 20, 36, 100},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?availability=always&methods=shop&location=12",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{1, 9, 11, 12, 13, 14},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/items?repeatable=true&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   112,
			results: []int32{1, 17, 19, 38, 54, 67, 88, 100, 112},
		},
	}

	testIdList(t, tests, testCfg.e.items.endpoint, "RetrieveItems", testCfg.HandleItems, compareAPIResourceLists[NamedApiResourceList])
}
