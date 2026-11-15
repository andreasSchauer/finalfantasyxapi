package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetItem(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Item{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/113",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "item with provided id '113' doesn't exist. max id: 112",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/69",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
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
			Unique:      apitesttypes.NewExpUnique(69, "three stars"),
			UntypedItem: 69,
			Category:    database.ItemCategorySupport,
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:   0,
					Monster: 203,
					Bribe:   14,
				},
				{
					Index:      2,
					Monster:    266,
					DropCommon: 1,
				},
				{
					Index:       4,
					Monster:     301,
					StealCommon: 2,
				},
			},
			Treasures: nil,
			Shops:     []int32{},
			Quests: map[int32]int32{
				38: 60,
				68: 60,
			},
			BlitzballPrizes: map[int32]int32{
				1: 1,
			},
			AeonLearnAbilities: map[int32]int32{
				43: 5,
			},
			AutoAbilities: map[string]int32{
				"one mp cost":    20,
				"break mp limit": 30,
			},
			Mixes: []int32{6, 12, 26, 49, 62, 64},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/69?rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"untyped item":         true,
					"category":             true,
					"treasures":            true,
					"shops":                true,
					"blitzball prizes":     true,
					"auto abilities":       true,
					"aeon learn abilities": true,
					"mixes":                true,
				},
				ExpLengths: map[string]int{
					"monsters": 4,
					"quests":   0,
				},
			},
			Unique: apitesttypes.NewExpUnique(69, "three stars"),
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:   0,
					Monster: 203,
					Bribe:   14,
				},
				{
					Index:      2,
					Monster:    266,
					DropCommon: 1,
				},
				{
					Index:     3,
					Monster:   291,
					StealRare: 1,
				},
			},
			Quests: nil,
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/26",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
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
			Unique:      apitesttypes.NewExpUnique(26, "fire gem"),
			UntypedItem: 26,
			Category:    database.ItemCategoryOffensive,
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:   0,
					Monster: 109,
					Bribe:   14,
				},
				{
					Index:       1,
					Monster:     141,
					StealCommon: 1,
					StealRare:   2,
					Bribe:       10,
				},
			},
			Treasures:       nil,
			Shops:           []int32{},
			Quests:          nil,
			BlitzballPrizes: nil,
			AeonLearnAbilities: map[int32]int32{
				77: 4,
			},
			AutoAbilities: map[string]int32{
				"fire eater": 20,
			},
			Mixes: []int32{1, 6, 15, 31, 55, 63},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/2",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
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
			Unique:      apitesttypes.NewExpUnique(2, "hi-potion"),
			UntypedItem: 2,
			Category:    database.ItemCategoryHealing,
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:     1,
					Monster:   39,
					StealRare: 1,
					Bribe:     60,
				},
				{
					Index:       36,
					Monster:     242,
					StealCommon: 2,
				},
			},
			Treasures: map[int32]int32{
				4:   1,
				33:  1,
				73:  2,
				142: 2,
				219: 4,
				232: 8,
			},
			Shops:  []int32{1, 7, 12, 20, 25, 28, 32, 36},
			Quests: nil,
			BlitzballPrizes: map[int32]int32{
				2: 2,
				5: 1,
			},
			AeonLearnAbilities: map[int32]int32{
				45: 99,
			},
			AutoAbilities: nil,
			Mixes:         []int32{34, 40, 51, 56, 61},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/64",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
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
			Unique:      apitesttypes.NewExpUnique(64, "stamina tablet"),
			UntypedItem: 64,
			Category:    database.ItemCategorySupport,
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:   0,
					Monster: 69,
					Other:   1,
				},
				{
					Index:   2,
					Monster: 156,
					Other:   1,
				},
			},
			Treasures: nil,
			Shops:     []int32{},
			Quests: map[int32]int32{
				29: 60,
			},
			BlitzballPrizes:    nil,
			AeonLearnAbilities: nil,
			AutoAbilities: map[string]int32{
				"auto-potion": 4,
				"hp stroll":   2,
			},
			Mixes: []int32{42, 56, 60, 63, 64},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/64?rel_availability=pre-story",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
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
			Unique:      apitesttypes.NewExpUnique(64, "stamina tablet"),
			UntypedItem: 64,
			Category:    database.ItemCategorySupport,
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:       0,
					Monster:     195,
					StealCommon: 1,
				},
			},
			Treasures:          nil,
			Shops:              []int32{},
			Quests:             nil,
			BlitzballPrizes:    nil,
			AeonLearnAbilities: nil,
			AutoAbilities: map[string]int32{
				"auto-potion": 4,
				"hp stroll":   2,
			},
			Mixes: []int32{42, 56, 60, 63, 64},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/55?rel_availability=post",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"untyped item":         true,
					"category":             true,
					"aeon learn abilities": true,
					"auto abilities":       true,
					"mixes":                true,
				},
				ExpLengths: map[string]int{
					"monsters":         3,
					"treasures":        0,
					"shops":            0,
					"quests":           0,
					"blitzball prizes": 0,
				},
			},
			Unique: apitesttypes.NewExpUnique(55, "chocobo wing"),
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:   0,
					Monster: 242,
					Bribe:   60,
				},
				{
					Index:     1,
					Monster:   269,
					StealRare: 1,
				},
				{
					Index:     2,
					Monster:   270,
					StealRare: 1,
				},
			},
			Treasures:       nil,
			Shops:           []int32{},
			Quests:          nil,
			BlitzballPrizes: nil,
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/55?rel_availability=always",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"untyped item":         true,
					"category":             true,
					"aeon learn abilities": true,
					"auto abilities":       true,
					"mixes":                true,
				},
				ExpLengths: map[string]int{
					"monsters":         2,
					"treasures":        0,
					"shops":            0,
					"quests":           1,
					"blitzball prizes": 0,
				},
			},
			Unique: apitesttypes.NewExpUnique(55, "chocobo wing"),
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:     0,
					Monster:   100,
					StealRare: 1,
				},
				{
					Index:       1,
					Monster:     261,
					StealCommon: 2,
				},
			},
			Treasures: nil,
			Shops:     []int32{},
			Quests: map[int32]int32{
				16: 99,
			},
			BlitzballPrizes: nil,
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/55?rel_availability=always&rel_repeatable=false",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"untyped item":         true,
					"category":             true,
					"aeon learn abilities": true,
					"auto abilities":       true,
					"mixes":                true,
				},
				ExpLengths: map[string]int{
					"monsters":         0,
					"treasures":        0,
					"shops":            0,
					"quests":           1,
					"blitzball prizes": 0,
				},
			},
			Unique:    apitesttypes.NewExpUnique(55, "chocobo wing"),
			Monsters:  []apitesttypes.TestMonItemAmts{},
			Treasures: nil,
			Shops:     []int32{},
			Quests: map[int32]int32{
				16: 99,
			},
			BlitzballPrizes: nil,
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/55?rel_availability=always&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"untyped item":         true,
					"category":             true,
					"aeon learn abilities": true,
					"auto abilities":       true,
					"mixes":                true,
				},
				ExpLengths: map[string]int{
					"monsters":         2,
					"treasures":        0,
					"shops":            0,
					"quests":           0,
					"blitzball prizes": 0,
				},
			},
			Unique: apitesttypes.NewExpUnique(55, "chocobo wing"),
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:     0,
					Monster:   100,
					StealRare: 1,
				},
				{
					Index:       1,
					Monster:     261,
					StealCommon: 2,
				},
			},
			Treasures:       nil,
			Shops:           []int32{},
			Quests:          nil,
			BlitzballPrizes: nil,
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/2?rel_availability=always",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"untyped item":         true,
					"category":             true,
					"aeon learn abilities": true,
					"auto abilities":       true,
					"mixes":                true,
				},
				ExpLengths: map[string]int{
					"monsters":         22,
					"treasures":        11,
					"shops":            4,
					"quests":           0,
					"blitzball prizes": 2,
				},
			},
			Unique: apitesttypes.NewExpUnique(2, "hi-potion"),
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:     0,
					Monster:   39,
					StealRare: 1,
					Bribe:     60,
				},
				{
					Index:       8,
					Monster:     87,
					StealCommon: 1,
				},
				{
					Index:       18,
					Monster:     162,
					StealCommon: 1,
					StealRare:   2,
				},
			},
			Treasures: map[int32]int32{
				4:   1,
				103: 2,
				219: 4,
				224: 4,
				232: 8,
			},
			Shops:  []int32{24, 32, 33, 34},
			Quests: nil,
			BlitzballPrizes: map[int32]int32{
				2: 2,
				5: 1,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/2?rel_availability=always&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"untyped item":         true,
					"category":             true,
					"aeon learn abilities": true,
					"auto abilities":       true,
					"mixes":                true,
				},
				ExpLengths: map[string]int{
					"monsters":         22,
					"treasures":        0,
					"shops":            4,
					"quests":           0,
					"blitzball prizes": 2,
				},
			},
			Unique: apitesttypes.NewExpUnique(2, "hi-potion"),
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:     0,
					Monster:   39,
					StealRare: 1,
					Bribe:     60,
				},
				{
					Index:       8,
					Monster:     87,
					StealCommon: 1,
				},
				{
					Index:       18,
					Monster:     162,
					StealCommon: 1,
					StealRare:   2,
				},
			},
			Treasures: nil,
			Shops:     []int32{24, 32, 33, 34},
			Quests:    nil,
			BlitzballPrizes: map[int32]int32{
				2: 2,
				5: 1,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items/2?rel_availability=post",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"untyped item":         true,
					"category":             true,
					"aeon learn abilities": true,
					"auto abilities":       true,
					"mixes":                true,
				},
				ExpLengths: map[string]int{
					"monsters":         2,
					"treasures":        1,
					"shops":            9,
					"quests":           0,
					"blitzball prizes": 0,
				},
			},
			Unique: apitesttypes.NewExpUnique(2, "hi-potion"),
			Monsters: []apitesttypes.TestMonItemAmts{
				{
					Index:       0,
					Monster:     239,
					StealCommon: 1,
				},
				{
					Index:       1,
					Monster:     242,
					StealCommon: 2,
				},
			},
			Treasures: map[int32]int32{
				44: 1,
			},
			Shops:           []int32{1, 5, 6, 7, 12, 20, 21, 27, 36},
			Quests:          nil,
			BlitzballPrizes: nil,
		},
	}

	testSingleResources(t, tests, "GetItem", testSrv.HandleItems)
}

func TestRetrieveItems(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?methods=shop,treasures",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value 'treasures' used for parameter 'methods'. allowed values: 'monster', 'treasure', 'shop', 'quest', 'blitzball'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   112,
			Results: []int32{1, 42, 67, 78, 90, 110, 112},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?category=healing,support&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   31,
			Results: []int32{1, 5, 11, 21, 23, 54, 59, 64, 69},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?category=healing,support&limit=max&flip=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   81,
			Results: []int32{17, 29, 41, 49, 73, 86, 106, 112},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?methods=treasure,blitzball&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   50,
			Results: []int32{1, 2, 9, 11, 12, 13, 16, 21, 71, 85, 101, 111},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?has_ability=true&related_stat=mp",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{65, 67, 68, 69},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?location=17&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   22,
			Results: []int32{1, 7, 21, 34, 76, 102, 111},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?sublocation=13", // ss winno
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{1, 2, 9, 11},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?area=237&limit=max", // omega ruins top
				ExpectedStatus: http.StatusOK,
			},
			Count:   31,
			Results: []int32{4, 23, 60, 89, 110},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=post",
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{23, 45, 67, 68, 76, 79, 88, 99},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=post-story&location=18&methods=monster",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{3, 5, 43, 45, 58, 61, 83, 96},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=post-game&repeatable=false&pre_airship=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   0,
			Results: []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=post-game&repeatable=false&pre_airship=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   0,
			Results: []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=post-game&repeatable=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   112,
			Results: []int32{1, 17, 19, 38, 54, 67, 88, 100, 112},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=post&repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   19,
			Results: []int32{23, 45, 67, 68, 76, 88, 90, 95, 105},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=always&repeatable=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   93,
			Results: []int32{1, 7, 21, 37, 56, 69, 80, 104, 112},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=always&repeatable=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{85, 86, 87, 89, 90, 94, 105},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=story&sublocation=26", // macalania woods
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{6, 82},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=post-game&area=23", // besaid - waterfall way (1)
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{1, 9, 30, 31, 32, 38, 41, 70, 71, 72},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=always&methods=shop",
				ExpectedStatus: http.StatusOK,
			},
			Count:   14,
			Results: []int32{1, 2, 9, 14, 20, 36, 100},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?availability=always&methods=shop&location=12",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{1, 9, 11, 12, 13, 14},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/items?repeatable=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   112,
			Results: []int32{1, 17, 19, 38, 54, 67, 88, 100, 112},
		},
	}

	testIDList(t, tests, apilookups.E.Items.Endpoint, "RetrieveItems", testSrv.HandleItems, compareAPIResourceLists[apiresource.Named])
}
