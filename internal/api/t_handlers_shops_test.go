package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetShop(t *testing.T) {
	tests := []expShop{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/37",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "shop with provided id '37' doesn't exist. max id: 36.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/34/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "subsection 'a' does not exist for endpoint /shops. supported subsections: 'simple'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
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
			expIdOnly: newExpIdOnly(2),
			area:      48,
			category:  "standard",
			preAirship: &testSubShop{
				items: []testShopItem{},
				equipment: []testShopEquipment{
					{
						index: 1,
						equipment: testFoundEquipment{
							equipmentName:    332,
							abilities:        []int32{1},
							emptySlotsAmount: 0,
						},
						price: 250,
					},
					{
						index: 4,
						equipment: testFoundEquipment{
							equipmentName:    335,
							abilities:        []int32{2, 1},
							emptySlotsAmount: 0,
						},
						price: 375,
					},
					{
						index: 8,
						equipment: testFoundEquipment{
							equipmentName:    1034,
							abilities:        []int32{111},
							emptySlotsAmount: 0,
						},
						price: 150,
					},
				},
			},
			postAirship: &testSubShop{
				items: []testShopItem{},
				equipment: []testShopEquipment{
					{
						index: 5,
						equipment: testFoundEquipment{
							equipmentName:    336,
							abilities:        []int32{2, 1},
							emptySlotsAmount: 2,
						},
						price: 1875,
					},
					{
						index: 6,
						equipment: testFoundEquipment{
							equipmentName:    337,
							abilities:        []int32{2, 1},
							emptySlotsAmount: 2,
						},
						price: 1875,
					},
					{
						index: 11,
						equipment: testFoundEquipment{
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
			testGeneral: testGeneral{
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
			expIdOnly: newExpIdOnly(16),
			area:      122,
			category:  "standard",
			preAirship: &testSubShop{
				items: []testShopItem{
					newTestShopItem(0, 1, 100),
					newTestShopItem(1, 9, 200),
					newTestShopItem(3, 13, 100),
					newTestShopItem(4, 14, 100),
				},
				equipment: []testShopEquipment{
					{
						index: 0,
						equipment: testFoundEquipment{
							equipmentName:    331,
							abilities:        []int32{30, 1},
							emptySlotsAmount: 0,
						},
						price: 2250,
					},
					{
						index: 3,
						equipment: testFoundEquipment{
							equipmentName:    272,
							abilities:        []int32{2, 5, 8},
							emptySlotsAmount: 0,
						},
						price: 8700,
					},
				},
			},
			postAirship: &testSubShop{
				items: []testShopItem{
					newTestShopItem(1, 9, 115),
					newTestShopItem(4, 14, 57),
					newTestShopItem(5, 12, 57),
				},
				equipment: []testShopEquipment{
					{
						index: 2,
						equipment: testFoundEquipment{
							equipmentName:    347,
							abilities:        []int32{30, 8},
							emptySlotsAmount: 0,
						},
						price: 2156,
					},
					{
						index: 4,
						equipment: testFoundEquipment{
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
			testGeneral: testGeneral{
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
			expIdOnly: newExpIdOnly(31),
			area:      202,
			category:  "travel-agency",
			preAirship: &testSubShop{
				items: []testShopItem{
					newTestShopItem(1, 2, 500),
					newTestShopItem(4, 13, 50),
					newTestShopItem(9, 19, 100),
				},
				equipment: []testShopEquipment{
					{
						index: 0,
						equipment: testFoundEquipment{
							equipmentName:    395,
							abilities:        []int32{35, 33},
							emptySlotsAmount: 2,
						},
						price: 38625,
					},
					{
						index: 4,
						equipment: testFoundEquipment{
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

	testSingleResources(t, tests, "GetShop", testCfg.HandleShops, compareShops)
}

func TestRetrieveShops(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   36,
			results: []int32{1, 7, 8, 16, 23, 28, 35, 36},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops?auto_ability=103",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{5, 6, 9, 12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops?equipment=false",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{3, 8, 11, 18, 26, 29, 30},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops?items=false",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{2},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops?post_airship=true",
				expectedStatus: http.StatusOK,
			},
			count:   18,
			results: []int32{1, 6, 7, 12, 17, 20, 21, 27, 32, 33, 36},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops?location=12",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{13, 14, 15, 16, 17, 18},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops?sublocation=25",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{22, 36},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops?category=oaka",
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{3, 4, 10, 13, 22, 26, 30},
		},
	}

	testIdList(t, tests, testCfg.e.shops.endpoint, "RetrieveShops", testCfg.HandleShops, compareAPIResourceLists[UnnamedApiResourceList])
}

func TestSubsectionShops(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/8/shops/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleLocations,
			},
			count:          3,
			parentResource: h.GetStrPtr("/locations/8"),
			results:        []int32{4, 5, 6},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/25/shops/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleSublocations,
			},
			count:          2,
			parentResource: h.GetStrPtr("/sublocations/25"),
			results:        []int32{22, 36},
		},
	}

	testIdList(t, tests, testCfg.e.shops.endpoint, "SubsectionShops", nil, compareSubResourceLists[UnnamedAPIResource, ShopSub])
}
