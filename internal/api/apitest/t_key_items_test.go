package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetKeyItem(t *testing.T) {
	t.Parallel()
	tests := []KeyItem{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items/61",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "key-item with provided id '61' doesn't exist. max id: 60",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items/39",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"areas":     2,
					"treasures": 2,
					"quests":    0,
				},
			},
			Unique:          newExpUnique(39, "al bhed primer v"),
			UntypedItem:     151,
			Category:        database.KeyItemCategoryPrimer,
			CelestialWeapon: nil,
			Primer:          h.GetInt32Ptr(5),
			Areas:           []int32{64, 168},
			Treasures:       []int32{68, 220},
			Quests:          []int32{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items/23",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"areas":     1,
					"treasures": 1,
					"quests":    0,
				},
			},
			Unique:          newExpUnique(23, "mercury crest"),
			UntypedItem:     135,
			Category:        database.KeyItemCategoryCelestial,
			CelestialWeapon: h.GetInt32Ptr(7),
			Primer:          nil,
			Areas:           []int32{170},
			Treasures:       []int32{231},
			Quests:          []int32{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items/8",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"areas":     1,
					"treasures": 0,
					"quests":    1,
				},
			},
			Unique:          newExpUnique(8, "mark of conquest"),
			UntypedItem:     120,
			Category:        database.KeyItemCategoryOther,
			CelestialWeapon: nil,
			Primer:          nil,
			Areas:           []int32{203},
			Treasures:       []int32{},
			Quests:          []int32{1},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items/39?rel_availability=always",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"areas":     2,
					"treasures": 2,
					"quests":    0,
				},
			},
			Unique:          newExpUnique(39, "al bhed primer v"),
			UntypedItem:     151,
			Category:        database.KeyItemCategoryPrimer,
			CelestialWeapon: nil,
			Primer:          h.GetInt32Ptr(5),
			Areas:           []int32{64, 168},
			Treasures:       []int32{68, 220},
			Quests:          []int32{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items/39?rel_availability=pre-story",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"areas":     0,
					"treasures": 0,
					"quests":    0,
				},
			},
			Unique:          newExpUnique(39, "al bhed primer v"),
			UntypedItem:     151,
			Category:        database.KeyItemCategoryPrimer,
			CelestialWeapon: nil,
			Primer:          h.GetInt32Ptr(5),
			Areas:           []int32{},
			Treasures:       []int32{},
			Quests:          []int32{},
		},
	}

	testSingleResources(t, tests, "GetKeyItem", testSrv.HandleKeyItems)
}

func TestRetrieveKeyItems(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?methods=a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value 'a' used for parameter 'methods'. allowed values: 'treasure', 'quest'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   60,
			Results: []int32{1, 16, 27, 48, 54, 60},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?availability=post&category=2",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{14, 16, 18, 24},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?methods=quest",
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{3, 8, 9, 16, 22, 24},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?location=15",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{4, 18, 25, 32, 49, 50},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?sublocation=14", // stadium (luca)
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{19, 20, 28, 40},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?area=236", // omega ruins bottom
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{60},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?availability=post",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{8, 14, 16, 18, 24, 60},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?availability=post-game&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   52,
			Results: []int32{3, 8, 11, 16, 26, 43, 60},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?availability=story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{1, 2, 5, 6, 53, 54, 55, 56},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?availability=always&location=8",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{19, 20, 28, 40, 41},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?availability=always&sublocation=26", // macalania woods
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{4, 25, 32, 49},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?availability=pre-airship&area=201", // calm lands - central
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{6, 12},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/key-items?availability=pre-airship&methods=quest",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{3, 4, 9, 10, 12, 20, 22},
		},
	}

	testIDList(t, tests, apilookups.E.KeyItems.Endpoint, "RetrieveKeyItems", testSrv.HandleKeyItems, compareAPIResourceLists[apiresource.Named])
}
