package main

import (
	"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)



func TestGetArenaCreation(t *testing.T) {
	tests := []expArenaCreation{
		{
			testGeneral: testGeneral{
				requestURL: "/api/arena-creations/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/arena-creations/0",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{},
			},
			expUnique: newExpUnique(0, ""),
			category: 				"",
			monster: 				0,
			parentSubquest: 		0,
			reward: 				newTestItemAmount("/items/0", 0),
			requiredCatchAmount: 	0,
			requiredMonsters: 		[]int32{},
		},
	}

	testSingleResources(t, tests, "GetArenaCreation", testCfg.HandleArenaCreations, compareArenaCreations)
}

func TestRetrieveArenaCreations(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   0,
			results: []int32{},
		},
	}

	testIdList(t, tests, testCfg.e.arenaCreations.endpoint, "RetrieveArenaCreations", testCfg.HandleArenaCreations, compareAPIResourceLists[NamedApiResourceList])
}
