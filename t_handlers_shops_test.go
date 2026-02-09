package main

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)



func TestGetShop(t *testing.T) {
	tests := []expShop{
		{
			testGeneral: testGeneral{
				requestURL: "/api/shops/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/shops/0",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{},
			},
			expIdOnly: newExpIdOnly(0),
			area: 0,
			category: "",
			preAirship: &testSubShop{
				items: []testShopItem{
					newTestShopItem(0, 0, 0),
				},
				equipment: []testShopEquipment{
					{
						index: 0,
						equipment: testFoundEquipment{
							index: 				0, // probably not used here
							equipmentName: 		0,
							abilities: 			[]int32{},
							emptySlotsAmount: 	0,
						},
						price: 0,
					},
				},
			},
			postAirship: &testSubShop{
				items: []testShopItem{
					newTestShopItem(0, 0, 0),
				},
				equipment: []testShopEquipment{
					{
						index: 0,
						equipment: testFoundEquipment{
							index: 				0, // probably not used here
							equipmentName: 		0,
							abilities: 			[]int32{},
							emptySlotsAmount: 	0,
						},
						price: 0,
					},
				},
			},
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
			count:   0,
			results: []int32{},
		},
	}

	testIdList(t, tests, testCfg.e.shops.endpoint, "RetrieveShops", testCfg.HandleShops, compareAPIResourceLists[UnnamedApiResourceList])
}
