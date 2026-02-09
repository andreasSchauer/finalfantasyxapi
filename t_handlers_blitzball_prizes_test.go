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
				requestURL: "/api/blitzball-prizes/9",
				expectedStatus: http.StatusNotFound,
				expectedErr: "blitzball prize table with provided id '9' doesn't exist. max id: 8.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: "/api/blitzball-prizes/9/2",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "endpoint 'blitzball-prizes' doesn't have any subsections.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: "/api/blitzball-prizes/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "invalid id 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/blitzball-prizes/1",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{
					"items": 8,
				},
			},
			expIdOnly: 	newExpIdOnly(1),
			category: 	"league",
			slot: 		"1st",
			items: []testPossibleItem{
				newTestPossibleItem(1, "/items/69", 1, 15),
				newTestPossibleItem(4, "/items/53", 1, 10),
				newTestPossibleItem(6, "/items/96", 1, 10),
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/blitzball-prizes/8",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{
					"items": 3,
				},
			},
			expIdOnly: 	newExpIdOnly(8),
			category: 	"tournament",
			slot: 		"top-scorer",
			items: []testPossibleItem{
				newTestPossibleItem(0, "/items/5", 1, 20),
				newTestPossibleItem(2, "/items/101", 1, 20),
			},
		},
	}

	testSingleResources(t, tests, "GetBlitzballPrize", testCfg.HandleBlitzballPrizes, compareBlitzballPrizes)
}

func TestRetrieveBlitzballPrizes(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes?category=asd",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "invalid enum value: 'asd'. use /api/blitzball-tournament-category to see valid values.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes?category=league",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{1, 2, 3, 4},
		},
	}

	testIdList(t, tests, testCfg.e.blitzballPrizes.endpoint, "RetrieveBlitzballPositions", testCfg.HandleBlitzballPrizes, compareAPIResourceLists[UnnamedApiResourceList])
}
