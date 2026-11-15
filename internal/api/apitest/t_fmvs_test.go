package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetFMV(t *testing.T) {
	t.Parallel()
	tests := []FMV{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/fmvs/a",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "fmv not found: 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/fmvs/53",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "fmv with provided id '53' doesn't exist. max id: 52",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/fmvs/2/2",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid subsection '2'. subsection can't be an integer. use /api/fmvs/sections for valid subsections",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/fmvs/9",
				ExpectedStatus: http.StatusOK,
			},
			Unique: newExpUnique(9, "fear on the sea"),
			Area:   42,
			Song:   h.GetInt32Ptr(16),
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/fmvs/38",
				ExpectedStatus: http.StatusOK,
			},
			Unique: newExpUnique(38, "the last chapter"),
			Area:   218,
			Song:   h.GetInt32Ptr(77),
		},
	}

	testSingleResources(t, tests, "GetFMV", testSrv.HandleFMVs)
}

func TestRetrieveFMVs(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/fmvs?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   52,
			Results: []int32{1, 13, 27, 34, 45, 46, 52},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/fmvs?location=15",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{27, 36},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/fmvs?location=5",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{9, 10, 11, 12, 13, 14},
		},
	}

	testIDList(t, tests, apilookups.E.FMVs.Endpoint, "RetrieveFMVs", testSrv.HandleFMVs, compareAPIResourceLists[apiresource.Named])
}
