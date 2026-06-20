package api

import (
	"net/http"
	"testing"
)

func TestGetAllItem(t *testing.T) {
	t.Parallel()
	tests := []expMasterItem{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/173",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "all item with provided id '173' doesn't exist. max id: 172.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/4",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique: newExpUnique(4, "mega-potion"),
			typedItem: "/items/4",
			itemType:  1,
			obtainableFrom: ObtainableFrom{
				Monsters:  		 true,
				Treasures: 		 true,
				Shops:     		 false,
				Quests:    		 true,
				BlitzballPrizes: true,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/111",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique: newExpUnique(111, "underdog's secret"),
			typedItem: "/items/111",
			itemType:  1,
			obtainableFrom: ObtainableFrom{
				Monsters:  		 true,
				Treasures: 		 false,
				Shops:     		 false,
				Quests:    		 true,
				BlitzballPrizes: true,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/130",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique: newExpUnique(130, "saturn sigil"),
			typedItem: "/key-items/18",
			itemType:  2,
			obtainableFrom: ObtainableFrom{
				Monsters:  		 false,
				Treasures: 		 false,
				Shops:     		 false,
				Quests:    		 true,
				BlitzballPrizes: false,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/2?rel_availability=always",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique: newExpUnique(2, "hi-potion"),
			typedItem: "/items/2",
			itemType:  1,
			obtainableFrom: ObtainableFrom{
				Monsters:  		 true,
				Treasures: 		 true,
				Shops:     		 true,
				Quests:    		 false,
				BlitzballPrizes: true,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/2?rel_availability=always&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique: newExpUnique(2, "hi-potion"),
			typedItem: "/items/2",
			itemType:  1,
			obtainableFrom: ObtainableFrom{
				Monsters:  		 true,
				Treasures: 		 false,
				Shops:     		 true,
				Quests:    		 false,
				BlitzballPrizes: true,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/2?rel_availability=post",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique: newExpUnique(2, "hi-potion"),
			typedItem: "/items/2",
			itemType:  1,
			obtainableFrom: ObtainableFrom{
				Monsters:  		 true,
				Treasures: 		 true,
				Shops:     		 true,
				Quests:    		 false,
				BlitzballPrizes: false,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/55?rel_availability=post",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique: newExpUnique(55, "chocobo wing"),
			typedItem: "/items/55",
			itemType:  1,
			obtainableFrom: ObtainableFrom{
				Monsters:  		 true,
				Treasures: 		 false,
				Shops:     		 false,
				Quests:    		 false,
				BlitzballPrizes: false,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/55?rel_availability=always",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique: newExpUnique(55, "chocobo wing"),
			typedItem: "/items/55",
			itemType:  1,
			obtainableFrom: ObtainableFrom{
				Monsters:  		 true,
				Treasures: 		 false,
				Shops:     		 false,
				Quests:    		 true,
				BlitzballPrizes: false,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/55?rel_availability=always&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique: newExpUnique(55, "chocobo wing"),
			typedItem: "/items/55",
			itemType:  1,
			obtainableFrom: ObtainableFrom{
				Monsters:  		 true,
				Treasures: 		 false,
				Shops:     		 false,
				Quests:    		 false,
				BlitzballPrizes: false,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/55?rel_availability=always&rel_repeatable=false",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique: newExpUnique(55, "chocobo wing"),
			typedItem: "/items/55",
			itemType:  1,
			obtainableFrom: ObtainableFrom{
				Monsters:  		 false,
				Treasures: 		 false,
				Shops:     		 false,
				Quests:    		 true,
				BlitzballPrizes: false,
			},
		},
	}

	testSingleResources(t, tests, "GetMasterItem", testCfg.HandleAllItems, compareMasterItems)
}

func TestRetrieveAllItems(t *testing.T) {
	t.Parallel()
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   172,
			results: []int32{1, 28, 53, 99, 148, 172},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?type=1&methods=quest&limit=maX",
				expectedStatus: http.StatusOK,
			},
			count:   53,
			results: []int32{1, 10, 29, 46, 54, 60, 80, 109, 112},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?location=14&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   42,
			results: []int32{1, 4, 16, 47, 57, 87, 100, 134, 160},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?sublocation=19&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   33,
			results: []int32{3, 14, 38, 103, 117, 142, 158},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?area=36",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{1, 2, 9, 138},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?availability=story&location=9",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{5, 81},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?availability=post-game&methods=shop&pre_airship=false",
				expectedStatus: http.StatusOK,
			},
			count:   17,
			results: []int32{1, 9, 11, 17, 20, 36, 95, 100},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?availability=post-game&methods=shop&pre_airship=true",
				expectedStatus: http.StatusOK,
			},
			count:   17,
			results: []int32{1, 2, 11, 12, 13, 16, 18, 20, 36, 95, 100},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?repeatable=true&methods=monster&area=172",
				expectedStatus: http.StatusOK,
			},
			count:   20,
			results: []int32{2, 10, 21, 36, 42, 54, 77, 102, 112},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?availability=story&repeatable=false",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{113, 114, 117, 118, 165, 166, 167, 168},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items?availability=post-game&repeatable=false&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   52,
			results: []int32{115, 120, 125, 133, 145, 157, 172},
		},
	}

	testIdList(t, tests, testCfg.e.allItems.endpoint, "RetrieveMasterItems", testCfg.HandleAllItems, compareAPIResourceLists[NamedApiResourceList])
}
