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
				expectedStatus: http.StatusNotFound,
				expectedErr: "fmv not found: 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: "/api/fmvs/53",
				expectedStatus: http.StatusNotFound,
				expectedErr: "fmv with provided id '53' doesn't exist. max id: 52.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: "/api/fmvs/2/2",
				expectedStatus: http.StatusBadRequest,
				expectedErr: "invalid subsection '2'. subsection can't be an integer. use /api/fmvs/sections for valid subsections.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/fmvs/9",
				expectedStatus: http.StatusOK,
			},
			expUnique: newExpUnique(9, "fear on the sea"),
			area: 42,
			song: h.GetInt32Ptr(16),
		},
		{
			testGeneral: testGeneral{
				requestURL: 	"/api/fmvs/38",
				expectedStatus: http.StatusOK,
			},
			expUnique: newExpUnique(38, "the last chapter"),
			area: 220,
			song: h.GetInt32Ptr(77),
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
			count:   52,
			results: []int32{1, 13, 27, 34, 45, 46, 52},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs?location=15",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{27, 36},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs?location=5",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{9, 10, 11, 12, 13, 14},
		},
	}

	testIdList(t, tests, testCfg.e.fmvs.endpoint, "RetrieveFMVs", testCfg.HandleFMVs, compareAPIResourceLists[NamedApiResourceList])
}
