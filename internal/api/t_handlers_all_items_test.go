package api

import (
	"net/http"
	"testing"
)

func TestGetAllItem(t *testing.T) {
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
			expUnique:    		newExpUnique(4, "mega-potion"),
			typedItem: 			"/items/4",
			itemType: 			1,
			obtainableFrom: 	ObtainableFrom{
				Monsters: 	true,
				Treasures: 	true,
				Shops: 		false,
				Quests: 	true,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/111",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique:    		newExpUnique(111, "underdog's secret"),
			typedItem: 			"/items/111",
			itemType: 			1,
			obtainableFrom: 	ObtainableFrom{
				Monsters: 	true,
				Treasures: 	false,
				Shops: 		false,
				Quests: 	true,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/all-items/130",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
			},
			expUnique:    		newExpUnique(130, "saturn sigil"),
			typedItem: 			"/key-items/18",
			itemType: 			2,
			obtainableFrom: 	ObtainableFrom{
				Monsters: 	false,
				Treasures: 	false,
				Shops: 		false,
				Quests: 	true,
			},
		},
	}

	testSingleResources(t, tests, "GetMasterItem", testCfg.HandleAllItems, compareMasterItems)
}

func TestRetrieveAllItems(t *testing.T) {
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
				requestURL:     "/api/all-items?type=1&method=quest&limit=maX",
				expectedStatus: http.StatusOK,
			},
			count:   53,
			results: []int32{1, 10, 29, 46, 54, 60, 80, 109, 112},
		},
	}

	testIdList(t, tests, testCfg.e.allItems.endpoint, "RetrieveMasterItems", testCfg.HandleAllItems, compareAPIResourceLists[NamedApiResourceList])
}
