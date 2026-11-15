package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func TestGetMix(t *testing.T) {
	t.Parallel()
	tests := []Mix{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/mixes/65",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "mix with provided id '65' doesn't exist. max id: 64",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/mixes/23",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"combinations": 62,
				},
			},
			Unique:    newExpUnique(23, "tidal wave"),
			category:  database.MixCategoryWaterElemental,
			overdrive: 74,
			combinations: []testMixCombination{
				{
					index:      7,
					firstItem:  30,
					secondItem: 89,
				},
				{
					index:      39,
					firstItem:  31,
					secondItem: 99,
				},
				{
					index:      60,
					firstItem:  32,
					secondItem: 98,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/mixes/23?best=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"combinations": 2,
				},
			},
			Unique:    newExpUnique(23, "tidal wave"),
			category:  database.MixCategoryWaterElemental,
			overdrive: 74,
			combinations: []testMixCombination{
				{
					index:      0,
					firstItem:  31,
					secondItem: 82,
				},
				{
					index:      1,
					firstItem:  32,
					secondItem: 81,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/mixes/34?contains_item=2",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"combinations": 15,
				},
			},
			Unique:    newExpUnique(34, "ultra potion"),
			category:  database.MixCategoryRecovery,
			overdrive: 85,
			combinations: []testMixCombination{
				{
					index:      0,
					firstItem:  1,
					secondItem: 2,
				},
				{
					index:      2,
					firstItem:  2,
					secondItem: 17,
				},
				{
					index:      14,
					firstItem:  2,
					secondItem: 101,
				},
			},
		},
	}

	testSingleResources(t, tests, "GetMix", testSrv.HandleMixes)
}

func TestRetrieveMixes(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/mixes?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   64,
			Results: []int32{1, 14, 38, 44, 55, 64},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/mixes?category=fire-elemental",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{11, 12, 13, 14, 15},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/mixes?req_item=45",
				ExpectedStatus: http.StatusOK,
			},
			Count:   14,
			Results: []int32{5, 9, 33, 58, 63},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/mixes?req_item=45&second_item=98",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{63},
		},
	}

	testIDList(t, tests, apilookups.E.Mixes.Endpoint, "RetrieveMixes", testSrv.HandleMixes, compareAPIResourceLists[apiresource.Named])
}
