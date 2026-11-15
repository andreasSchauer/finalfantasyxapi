package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetShop(t *testing.T) {
	t.Parallel()
	tests := []Shop{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops/37",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "shop with provided id '37' doesn't exist. max id: 36",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops/34/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "subsection 'a' does not exist for endpoint /shops. supported subsections: 'simple'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid id 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops/2",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"pre airship - items":      0,
					"pre airship - equipment":  10,
					"post airship - items":     0,
					"post airship - equipment": 14,
				},
			},
			IDOnly:   newExpIDOnly(2),
			Area:     47,
			Category: database.ShopCategoryStandard,
			PreAirship: &TestSubShop{
				Items: []TestShopItem{},
				Equipment: []TestShopEquipment{
					{
						Index: 1,
						Equipment: TestFoundEquipment{
							EquipmentName:    332,
							Abilities:        []int32{1},
							EmptySlotsAmount: 0,
						},
						Price: 250,
					},
					{
						Index: 4,
						Equipment: TestFoundEquipment{
							EquipmentName:    335,
							Abilities:        []int32{2, 1},
							EmptySlotsAmount: 0,
						},
						Price: 375,
					},
					{
						Index: 8,
						Equipment: TestFoundEquipment{
							EquipmentName:    1034,
							Abilities:        []int32{111},
							EmptySlotsAmount: 0,
						},
						Price: 150,
					},
				},
			},
			PostAirship: &TestSubShop{
				Items: []TestShopItem{},
				Equipment: []TestShopEquipment{
					{
						Index: 5,
						Equipment: TestFoundEquipment{
							EquipmentName:    336,
							Abilities:        []int32{2, 1},
							EmptySlotsAmount: 2,
						},
						Price: 1875,
					},
					{
						Index: 6,
						Equipment: TestFoundEquipment{
							EquipmentName:    337,
							Abilities:        []int32{2, 1},
							EmptySlotsAmount: 2,
						},
						Price: 1875,
					},
					{
						Index: 11,
						Equipment: TestFoundEquipment{
							EquipmentName:    940,
							Abilities:        []int32{111},
							EmptySlotsAmount: 3,
						},
						Price: 2250,
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops/16",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"pre airship - items":      6,
					"pre airship - equipment":  5,
					"post airship - items":     6,
					"post airship - equipment": 5,
				},
			},
			IDOnly:   newExpIDOnly(16),
			Area:     121,
			Category: database.ShopCategoryStandard,
			PreAirship: &TestSubShop{
				Items: []TestShopItem{
					newTestShopItem(0, 1, 100),
					newTestShopItem(1, 9, 200),
					newTestShopItem(3, 13, 100),
					newTestShopItem(4, 14, 100),
				},
				Equipment: []TestShopEquipment{
					{
						Index: 0,
						Equipment: TestFoundEquipment{
							EquipmentName:    331,
							Abilities:        []int32{30, 1},
							EmptySlotsAmount: 0,
						},
						Price: 2250,
					},
					{
						Index: 3,
						Equipment: TestFoundEquipment{
							EquipmentName:    272,
							Abilities:        []int32{2, 5, 8},
							EmptySlotsAmount: 0,
						},
						Price: 8700,
					},
				},
			},
			PostAirship: &TestSubShop{
				Items: []TestShopItem{
					newTestShopItem(1, 9, 115),
					newTestShopItem(4, 14, 57),
					newTestShopItem(5, 12, 57),
				},
				Equipment: []TestShopEquipment{
					{
						Index: 2,
						Equipment: TestFoundEquipment{
							EquipmentName:    347,
							Abilities:        []int32{30, 8},
							EmptySlotsAmount: 0,
						},
						Price: 2156,
					},
					{
						Index: 4,
						Equipment: TestFoundEquipment{
							EquipmentName:    329,
							Abilities:        []int32{2, 9},
							EmptySlotsAmount: 0,
						},
						Price: 3536,
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops/31",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"pre airship - items":      11,
					"pre airship - equipment":  7,
					"post airship - items":     0,
					"post airship - equipment": 0,
				},
			},
			IDOnly:   newExpIDOnly(31),
			Area:     200,
			Category: database.ShopCategoryTravelAgency,
			PreAirship: &TestSubShop{
				Items: []TestShopItem{
					newTestShopItem(1, 2, 500),
					newTestShopItem(4, 13, 50),
					newTestShopItem(9, 19, 100),
				},
				Equipment: []TestShopEquipment{
					{
						Index: 0,
						Equipment: TestFoundEquipment{
							EquipmentName:    395,
							Abilities:        []int32{35, 33},
							EmptySlotsAmount: 2,
						},
						Price: 38625,
					},
					{
						Index: 4,
						Equipment: TestFoundEquipment{
							EquipmentName:    765,
							Abilities:        []int32{76, 70},
							EmptySlotsAmount: 0,
						},
						Price: 6825,
					},
				},
			},
			PostAirship: nil,
		},
	}

	testSingleResources(t, tests, "GetShop", testSrv.HandleShops)
}

func TestRetrieveShops(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   36,
			Results: []int32{1, 7, 8, 16, 23, 28, 35, 36},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?auto_ability=103",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{5, 6, 9},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?equipment=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{3, 8, 11, 18, 26, 29, 30},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?items=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{2},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?availability=always",
				ExpectedStatus: http.StatusOK,
			},
			Count:   16,
			Results: []int32{1, 6, 7, 12, 17, 20, 21, 32, 33, 34},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?location=12",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{13, 14, 15, 16, 17, 18},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?sublocation=26", // macalania woods
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{22, 36},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?category=oaka",
				ExpectedStatus: http.StatusOK,
			},
			Count:   14,
			Results: []int32{3, 4, 10, 13, 22, 26, 30},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?auto_ability=111&character=3&empty_slots=0",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{1, 2, 9, 21},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?empty_slots=2-4&availability=post",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{1, 2, 12, 24, 27, 32, 36},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?availability=post&location=15",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{36},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?availability=pre-airship&sublocation=16", // highroad (mi'ihen)
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{7, 8},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?availability=pre-story&empty_slots=2&character=7&auto_ability=41",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{35},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?availability=pre-story&equipment=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{4, 10, 19, 22, 25, 31, 35},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?availability=always&empty_slots=2&character=3",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{34},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/shops?availability=post&auto_ability=4",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{12},
		},
	}

	testIDList(t, tests, apilookups.E.Shops.Endpoint, "RetrieveShops", testSrv.HandleShops, compareAPIResourceLists[apiresource.Unnamed])
}

func TestSubsectionShops(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/8/shops/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleLocations,
			},
			Count:          3,
			ParentResource: h.GetStrPtr("/locations/8"),
			Results:        []int32{4, 5, 6},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/26/shops/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          2,
			ParentResource: h.GetStrPtr("/sublocations/26"),
			Results:        []int32{22, 36},
		},
	}

	testIDList(t, tests, apilookups.E.Shops.Endpoint, "SubsectionShops", nil, compareSimpleResourceLists[apiresource.Unnamed, sectiontypes.ShopSimple])
}
