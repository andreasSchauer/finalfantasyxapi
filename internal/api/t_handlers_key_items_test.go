package api

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetKeyItem(t *testing.T) {
	t.Parallel()
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
				expLengths: map[string]int{
					"areas":     2,
					"treasures": 2,
					"quests":    0,
				},
			},
			expUnique:       newExpUnique(39, "al bhed primer v"),
			untypedItem:     151,
			category:        database.KeyItemCategoryPrimer,
			celestialWeapon: nil,
			primer:          h.GetInt32Ptr(5),
			areas:           []int32{65, 170},
			treasures:       []int32{68, 220},
			quests:          []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items/23",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"areas":     1,
					"treasures": 1,
					"quests":    0,
				},
			},
			expUnique:       newExpUnique(23, "mercury crest"),
			untypedItem:     135,
			category:        database.KeyItemCategoryCelestial,
			celestialWeapon: h.GetInt32Ptr(7),
			primer:          nil,
			areas:           []int32{172},
			treasures:       []int32{231},
			quests:          []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items/8",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"areas":     1,
					"treasures": 0,
					"quests":    1,
				},
			},
			expUnique:       newExpUnique(8, "mark of conquest"),
			untypedItem:     120,
			category:        database.KeyItemCategoryOther,
			celestialWeapon: nil,
			primer:          nil,
			areas:           []int32{205},
			treasures:       []int32{},
			quests:          []int32{1},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items/39?rel_availability=always",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"areas":     2,
					"treasures": 2,
					"quests":    0,
				},
			},
			expUnique:       newExpUnique(39, "al bhed primer v"),
			untypedItem:     151,
			category:        database.KeyItemCategoryPrimer,
			celestialWeapon: nil,
			primer:          h.GetInt32Ptr(5),
			areas:           []int32{65, 170},
			treasures:       []int32{68, 220},
			quests:          []int32{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items/39?rel_availability=pre-story",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"areas":     0,
					"treasures": 0,
					"quests":    0,
				},
			},
			expUnique:       newExpUnique(39, "al bhed primer v"),
			untypedItem:     151,
			category:        database.KeyItemCategoryPrimer,
			celestialWeapon: nil,
			primer:          h.GetInt32Ptr(5),
			areas:           []int32{},
			treasures:       []int32{},
			quests:          []int32{},
		},
	}

	testSingleResources(t, tests, "GetKeyItem", testCfg.HandleKeyItems, compareKeyItems)
}

func TestRetrieveKeyItems(t *testing.T) {
	t.Parallel()
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?methods=a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'a' used for parameter 'methods'. allowed values: 'treasure', 'quest'.",
			},
		},
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
				requestURL:     "/api/key-items?methods=quest",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{3, 8, 9, 16, 22, 24},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?location=15",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{4, 18, 25, 32, 49, 50},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?sublocation=13",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{19, 20, 28, 40},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?area=239",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{60},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?availability=post",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{8, 14, 16, 18, 24, 60},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?availability=post-game&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   52,
			results: []int32{3, 8, 11, 16, 26, 43, 60},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?availability=story",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{1, 2, 5, 6, 53, 54, 55, 56},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?availability=always&location=8",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{19, 20, 28, 40, 41},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?availability=always&sublocation=25",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{4, 25, 32, 49},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?availability=pre-airship&area=203",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{6, 12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/key-items?availability=pre-airship&methods=quest",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{3, 4, 9, 10, 12, 20, 22},
		},
	}

	testIdList(t, tests, testCfg.e.keyItems.endpoint, "RetrieveKeyItems", testCfg.HandleKeyItems, compareAPIResourceLists[NamedApiResourceList])
}
