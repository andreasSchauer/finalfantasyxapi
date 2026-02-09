package main

import (
	"net/http"
	"testing"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)




func TestGetFMV(t *testing.T) {
	tests := []expFMV{
		{
			testGeneral: testGeneral{
				requestURL: "/api/fmvs/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/fmvs/0",
				expectedStatus: http.StatusOK,
				dontCheck: 		map[string]bool{},
				expLengths: 	map[string]int{},
			},
			expUnique: newExpUnique(0, ""),
			area: 0,
			song: h.GetInt32Ptr(0),
		},
	}

	testSingleResources(t, tests, "GetFMV", testCfg.HandleFMVs, compareFMVs)
}

func TestRetrieveFMVs(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   0,
			results: []int32{},
		},
	}

	testIdList(t, tests, testCfg.e.fmvs.endpoint, "RetrieveFMVs", testCfg.HandleFMVs, compareAPIResourceLists[NamedApiResourceList])
}
