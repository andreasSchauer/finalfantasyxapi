package main

import (
	"net/http"
	"testing"
)


func TestGetTreasure(t *testing.T) {
	tests := []expTreasure{
		{
			testGeneral: testGeneral{
				requestURL: "/api/treasures/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/treasures/0",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{},
			},
			expIdOnly: newExpIdOnly(0),
			area: 				0,
			isPostAirship: 		false,
			isAnimaTreasure: 	false,
			treasureType: 		0,
			lootType: 			0,
			gilAmount: 			nil,
			items: []testItemAmount{
				newTestItemAmount("/items/0", 0),
			},
			equipment: &testFoundEquipment{
				index: 				0,
				equipmentName: 		0,
				abilities: 			[]int32{},
				emptySlotsAmount: 	0,
			},
		},
	}

	testSingleResources(t, tests, "GetTreasure", testCfg.HandleTreasures, compareTreasures)
}

func TestRetrieveTreasures(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   0,
			results: []int32{},
		},
	}

	testIdList(t, tests, testCfg.e.treasures.endpoint, "RetrieveTreasures", testCfg.HandleTreasures, compareAPIResourceLists[UnnamedApiResourceList])
}
