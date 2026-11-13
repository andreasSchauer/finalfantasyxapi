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
				requestURL:     "/api/shops/37",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "shop with provided id '37' doesn't exist. max id: 36",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops/34/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "subsection 'a' does not exist for endpoint /shops. supported subsections: 'simple'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops/2",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"pre airship - items":      0,
					"pre airship - equipment":  10,
					"post airship - items":     0,
					"post airship - equipment": 14,
				},
			},
			IDOnly: newExpIDOnly(2),
			area:      47,
			category:  database.ShopCategoryStandard,
			preAirship: &TestSubShop{
				items: []TestShopItem{},
				equipment: []TestShopEquipment{
					{
						index: 1,
						equipment: TestFoundEquipment{
							equipmentName:    332,
							abilities:        []int32{1},
							emptySlotsAmount: 0,
						},
						price: 250,
					},
					{
						index: 4,
						equipment: TestFoundEquipment{
							equipmentName:    335,
							abilities:        []int32{2, 1},
							emptySlotsAmount: 0,
						},
						price: 375,
					},
					{
						index: 8,
						equipment: TestFoundEquipment{
							equipmentName:    1034,
							abilities:        []int32{111},
							emptySlotsAmount: 0,
						},
						price: 150,
					},
				},
			},
			postAirship: &TestSubShop{
				items: []TestShopItem{},
				equipment: []TestShopEquipment{
					{
						index: 5,
						equipment: TestFoundEquipment{
							equipmentName:    336,
							abilities:        []int32{2, 1},
							emptySlotsAmount: 2,
						},
						price: 1875,
					},
					{
						index: 6,
						equipment: TestFoundEquipment{
							equipmentName:    337,
							abilities:        []int32{2, 1},
							emptySlotsAmount: 2,
						},
						price: 1875,
					},
					{
						index: 11,
						equipment: TestFoundEquipment{
							equipmentName:    940,
							abilities:        []int32{111},
							emptySlotsAmount: 3,
						},
						price: 2250,
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops/16",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"pre airship - items":      6,
					"pre airship - equipment":  5,
					"post airship - items":     6,
					"post airship - equipment": 5,
				},
			},
			IDOnly: newExpIDOnly(16),
			area:      121,
			category:  database.ShopCategoryStandard,
			preAirship: &TestSubShop{
				items: []TestShopItem{
					newTestShopItem(0, 1, 100),
					newTestShopItem(1, 9, 200),
					newTestShopItem(3, 13, 100),
					newTestShopItem(4, 14, 100),
				},
				equipment: []TestShopEquipment{
					{
						index: 0,
						equipment: TestFoundEquipment{
							equipmentName:    331,
							abilities:        []int32{30, 1},
							emptySlotsAmount: 0,
						},
						price: 2250,
					},
					{
						index: 3,
						equipment: TestFoundEquipment{
							equipmentName:    272,
							abilities:        []int32{2, 5, 8},
							emptySlotsAmount: 0,
						},
						price: 8700,
					},
				},
			},
			postAirship: &TestSubShop{
				items: []TestShopItem{
					newTestShopItem(1, 9, 115),
					newTestShopItem(4, 14, 57),
					newTestShopItem(5, 12, 57),
				},
				equipment: []TestShopEquipment{
					{
						index: 2,
						equipment: TestFoundEquipment{
							equipmentName:    347,
							abilities:        []int32{30, 8},
							emptySlotsAmount: 0,
						},
						price: 2156,
					},
					{
						index: 4,
						equipment: TestFoundEquipment{
							equipmentName:    329,
							abilities:        []int32{2, 9},
							emptySlotsAmount: 0,
						},
						price: 3536,
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops/31",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"pre airship - items":      11,
					"pre airship - equipment":  7,
					"post airship - items":     0,
					"post airship - equipment": 0,
				},
			},
			IDOnly: newExpIDOnly(31),
			area:      200,
			category:  database.ShopCategoryTravelAgency,
			preAirship: &TestSubShop{
				items: []TestShopItem{
					newTestShopItem(1, 2, 500),
					newTestShopItem(4, 13, 50),
					newTestShopItem(9, 19, 100),
				},
				equipment: []TestShopEquipment{
					{
						index: 0,
						equipment: TestFoundEquipment{
							equipmentName:    395,
							abilities:        []int32{35, 33},
							emptySlotsAmount: 2,
						},
						price: 38625,
					},
					{
						index: 4,
						equipment: TestFoundEquipment{
							equipmentName:    765,
							abilities:        []int32{76, 70},
							emptySlotsAmount: 0,
						},
						price: 6825,
					},
				},
			},
			postAirship: nil,
		},
	}

	testSingleResources(t, tests, "GetShop", testSrv.HandleShops, compareShops)
}

func TestRetrieveShops(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   36,
			results: []int32{1, 7, 8, 16, 23, 28, 35, 36},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?auto_ability=103",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{5, 6, 9},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?equipment=false",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{3, 8, 11, 18, 26, 29, 30},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?items=false",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{2},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?availability=always",
				expectedStatus: http.StatusOK,
			},
			count:   16,
			results: []int32{1, 6, 7, 12, 17, 20, 21, 32, 33, 34},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?location=12",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{13, 14, 15, 16, 17, 18},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?sublocation=26", // macalania woods
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{22, 36},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?category=oaka",
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{3, 4, 10, 13, 22, 26, 30},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?auto_ability=111&character=3&empty_slots=0",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{1, 2, 9, 21},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?empty_slots=2-4&availability=post",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{1, 2, 12, 24, 27, 32, 36},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?availability=post&location=15",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{36},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?availability=pre-airship&sublocation=16", // highroad (mi'ihen)
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{7, 8},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?availability=pre-story&empty_slots=2&character=7&auto_ability=41",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{35},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?availability=pre-story&equipment=true",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{4, 10, 19, 22, 25, 31, 35},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?availability=always&empty_slots=2&character=3",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{34},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops?availability=post&auto_ability=4",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{12},
		},
	}

	testIDList(t, tests, apilookups.E.Shops.Endpoint, "RetrieveShops", testSrv.HandleShops, compareAPIResourceLists[apiresource.Unnamed])
}

func TestSubsectionShops(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/locations/8/shops/",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleLocations,
			},
			count:          3,
			parentResource: h.GetStrPtr("/locations/8"),
			results:        []int32{4, 5, 6},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/sublocations/26/shops/",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleSublocations,
			},
			count:          2,
			parentResource: h.GetStrPtr("/sublocations/26"),
			results:        []int32{22, 36},
		},
	}

	testIDList(t, tests, apilookups.E.Shops.Endpoint, "SubsectionShops", nil, compareSimpleResourceLists[apiresource.Unnamed, sectiontypes.ShopSimple])
}
