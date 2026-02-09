package main

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetBlitzballPrize(t *testing.T) {
	tests := []expBlitzballPrize{
		{
			testGeneral: testGeneral{
				requestURL: "/api/blitzball-prizes/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/blitzball-prizes/0",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{},
			},
			expIdOnly: newExpIdOnly(0),
			category: 	"",
			slot: 		"",
			items: []testPossibleItem{
				newTestPossibleItem(0, "/items/0", 0, 0),
			},
		},
	}

	testSingleResources(t, tests, "GetBlitzballPrize", testCfg.HandleBlitzballPrizes, compareBlitzballPrizes)
}

func TestRetrieveBlitzballPrizes(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   0,
			results: []int32{},
		},
	}

	testIdList(t, tests, testCfg.e.blitzballPrizes.endpoint, "RetrieveBlitzballPositions", testCfg.HandleBlitzballPrizes, compareAPIResourceLists[UnnamedApiResourceList])
}
