package api

import (
	"net/http"
	"testing"
)

func TestGetPrimer(t *testing.T) {
	tests := []expPrimer{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers/27",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "primer with provided id '27' doesn't exist. max id: 26.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers/1",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"areas": 		2,
					"treasures": 	2,
				},
			},
			expUnique:  newExpUnique(1, "al bhed primer i"),
			keyItem:	35,
			areas: 		[]int32{15, 169},
			treasures: 	[]int32{13, 215},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers/22",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"areas": 		1,
					"treasures": 	1,
				},
			},
			expUnique:  newExpUnique(22, "al bhed primer xxii"),
			keyItem:	56,
			areas: 		[]int32{197},
			treasures: 	[]int32{252},
		},
	}

	testSingleResources(t, tests, "GetPrimer", testCfg.HandlePrimers, comparePrimers)
}

func TestRetrievePrimers(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   26,
			results: []int32{1, 5, 17, 20, 26},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/primers?availability=story",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{1, 3, 8, 19, 20, 21, 22},
		},
	}

	testIdList(t, tests, testCfg.e.primers.endpoint, "RetrievePrimers", testCfg.HandlePrimers, compareAPIResourceLists[NamedApiResourceList])
}
