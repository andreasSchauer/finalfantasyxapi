package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func TestGetAllItem(t *testing.T) {
	t.Parallel()
	tests := []MasterItem{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/173",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "all item with provided id '173' doesn't exist. max id: 172",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/4",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
			},
			Unique:    NewExpUnique(4, "mega-potion"),
			TypedItem: "/items/4",
			ItemType:  database.ItemTypeItem,
			ObtainableFrom: apitypes.ObtainableFrom{
				Monsters:        true,
				Treasures:       true,
				Shops:           false,
				Quests:          true,
				BlitzballPrizes: true,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/111",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
			},
			Unique:    NewExpUnique(111, "underdog's secret"),
			TypedItem: "/items/111",
			ItemType:  database.ItemTypeItem,
			ObtainableFrom: apitypes.ObtainableFrom{
				Monsters:        true,
				Treasures:       false,
				Shops:           false,
				Quests:          true,
				BlitzballPrizes: true,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/130",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
			},
			Unique:    NewExpUnique(130, "saturn sigil"),
			TypedItem: "/key-items/18",
			ItemType:  database.ItemTypeKeyItem,
			ObtainableFrom: apitypes.ObtainableFrom{
				Monsters:        false,
				Treasures:       false,
				Shops:           false,
				Quests:          true,
				BlitzballPrizes: false,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/2?rel_availability=always",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
			},
			Unique:    NewExpUnique(2, "hi-potion"),
			TypedItem: "/items/2",
			ItemType:  database.ItemTypeItem,
			ObtainableFrom: apitypes.ObtainableFrom{
				Monsters:        true,
				Treasures:       true,
				Shops:           true,
				Quests:          false,
				BlitzballPrizes: true,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/2?rel_availability=always&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
			},
			Unique:    NewExpUnique(2, "hi-potion"),
			TypedItem: "/items/2",
			ItemType:  database.ItemTypeItem,
			ObtainableFrom: apitypes.ObtainableFrom{
				Monsters:        true,
				Treasures:       false,
				Shops:           true,
				Quests:          false,
				BlitzballPrizes: true,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/2?rel_availability=post",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
			},
			Unique:    NewExpUnique(2, "hi-potion"),
			TypedItem: "/items/2",
			ItemType:  database.ItemTypeItem,
			ObtainableFrom: apitypes.ObtainableFrom{
				Monsters:        true,
				Treasures:       true,
				Shops:           true,
				Quests:          false,
				BlitzballPrizes: false,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/55?rel_availability=post",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
			},
			Unique:    NewExpUnique(55, "chocobo wing"),
			TypedItem: "/items/55",
			ItemType:  database.ItemTypeItem,
			ObtainableFrom: apitypes.ObtainableFrom{
				Monsters:        true,
				Treasures:       false,
				Shops:           false,
				Quests:          false,
				BlitzballPrizes: false,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/55?rel_availability=always",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
			},
			Unique:    NewExpUnique(55, "chocobo wing"),
			TypedItem: "/items/55",
			ItemType:  database.ItemTypeItem,
			ObtainableFrom: apitypes.ObtainableFrom{
				Monsters:        true,
				Treasures:       false,
				Shops:           false,
				Quests:          true,
				BlitzballPrizes: false,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/55?rel_availability=always&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
			},
			Unique:    NewExpUnique(55, "chocobo wing"),
			TypedItem: "/items/55",
			ItemType:  database.ItemTypeItem,
			ObtainableFrom: apitypes.ObtainableFrom{
				Monsters:        true,
				Treasures:       false,
				Shops:           false,
				Quests:          false,
				BlitzballPrizes: false,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items/55?rel_availability=always&rel_repeatable=false",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
			},
			Unique:    NewExpUnique(55, "chocobo wing"),
			TypedItem: "/items/55",
			ItemType:  database.ItemTypeItem,
			ObtainableFrom: apitypes.ObtainableFrom{
				Monsters:        false,
				Treasures:       false,
				Shops:           false,
				Quests:          true,
				BlitzballPrizes: false,
			},
		},
	}

	testSingleResources(t, tests, "GetMasterItem", testSrv.HandleAllItems)
}

func TestRetrieveAllItems(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   172,
			Results: []int32{1, 28, 53, 99, 148, 172},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?type=1&methods=quest&limit=maX",
				ExpectedStatus: http.StatusOK,
			},
			Count:   53,
			Results: []int32{1, 10, 29, 46, 54, 60, 80, 109, 112},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?location=14&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   42,
			Results: []int32{1, 4, 16, 47, 57, 87, 100, 134, 160},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?sublocation=20&limit=max", // moonflow
				ExpectedStatus: http.StatusOK,
			},
			Count:   33,
			Results: []int32{3, 14, 38, 103, 117, 142, 158},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?area=36", // besaid village
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{1, 2, 9, 138},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?availability=story&location=9",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{5, 81},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?availability=post-game&methods=shop&pre_airship=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   17,
			Results: []int32{1, 9, 11, 17, 20, 36, 95, 100},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?availability=post-game&methods=shop&pre_airship=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   17,
			Results: []int32{1, 2, 11, 12, 13, 16, 18, 20, 36, 95, 100},
		},

		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?repeatable=true&methods=monster&area=170", // sanubia desert - west
				ExpectedStatus: http.StatusOK,
			},
			Count:   20,
			Results: []int32{2, 10, 21, 36, 42, 54, 77, 102, 112},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?availability=story&repeatable=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{113, 114, 117, 118, 165, 166, 167, 168},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/all-items?availability=post-game&repeatable=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   52,
			Results: []int32{115, 120, 125, 133, 145, 157, 172},
		},
	}

	testIDList(t, tests, apilookups.E.AllItems.Endpoint, "RetrieveMasterItems", testSrv.HandleAllItems, compareAPIResourceLists[apiresource.Named])
}
