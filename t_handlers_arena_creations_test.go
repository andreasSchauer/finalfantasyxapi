package main

import (
	"net/http"
	"testing"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)



func TestGetArenaCreation(t *testing.T) {
	tests := []expArenaCreation{
		{
			testGeneral: testGeneral{
				requestURL: "/api/arena-creations/36",
				expectedStatus: http.StatusNotFound,
				expectedErr: "arena creation with provided id '36' doesn't exist. max id: 35.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: "/api/arena-creations/a",
				expectedStatus: http.StatusNotFound,
				expectedErr: "arena creation not found: 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: "/api/arena-creations/a/2",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "wrong format. usage: '/api/arena-creations', '/api/arena-creations/{id}', '/api/arena-creations/{name}'",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/arena-creations/13",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{
					"required monsters": 10,
				},
			},
			expUnique: 				newExpUnique(13, "vorban"),
			category: 				"area",
			monster: 				268,
			parentSubquest: 		13,
			reward: 				newTestItemAmount("/items/107", 60),
			requiredCatchAmount: 	1,
			requiredMonsters: 		[]int32{239, 243, 245, 248},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/arena-creations/one-eye",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{
					"required monsters": 5,
				},
			},
			expUnique: 				newExpUnique(19, "one-eye"),
			category: 				"species",
			monster: 				274,
			parentSubquest: 		19,
			reward: 				newTestItemAmount("/items/64", 60),
			requiredCatchAmount: 	4,
			requiredMonsters: 		[]int32{40, 73, 87, 170, 240},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/arena-creations/31",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{
					"required monsters": 0,
				},
			},
			expUnique: 					newExpUnique(31, "th'uban"),
			category: 					"original",
			monster: 					286,
			parentSubquest: 			31,
			reward: 					newTestItemAmount("/items/110", 99),
			requiredCatchAmount: 		6,
			unlockedCreationsCategory: 	h.GetStrPtr("species"),
			requiredMonsters: 			nil,
		},
	}

	testSingleResources(t, tests, "GetArenaCreation", testCfg.HandleArenaCreations, compareArenaCreations)
}

func TestRetrieveArenaCreations(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations?category=asd",
				expectedStatus: http.StatusBadRequest,
				expectedErr: 	"invalid enum value 'asd' used for parameter 'category'. use /api/arena-creations/parameters to see allowed values.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   35,
			results: []int32{1, 11, 14, 20, 27, 31, 35},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations?category=species&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{14, 15, 18, 19, 23, 27},
		},
	}

	testIdList(t, tests, testCfg.e.arenaCreations.endpoint, "RetrieveArenaCreations", testCfg.HandleArenaCreations, compareAPIResourceLists[NamedApiResourceList])
}
