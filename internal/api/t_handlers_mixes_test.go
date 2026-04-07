package api

import (
	"net/http"
	"testing"
)

func TestGetMix(t *testing.T) {
	tests := []expMix{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/mixes/65",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "mix with provided id '65' doesn't exist. max id: 64.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/mixes/23",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"combinations": 62,
				},
			},
			expUnique:  newExpUnique(23, "tidal wave"),
			category:	9,
			overdrive:	74,
			combinations: []testMixCombination{
				{
					index: 		7,
					firstItem: 	30,
					secondItem: 89,
				},
				{
					index: 		39,
					firstItem: 	31,
					secondItem: 99,
				},
				{
					index: 		60,
					firstItem: 	32,
					secondItem: 98,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/mixes/23?best=true",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"combinations": 2,
				},
			},
			expUnique:  newExpUnique(23, "tidal wave"),
			category:	9,
			overdrive:	74,
			combinations: []testMixCombination{
				{
					index: 		0,
					firstItem: 	31,
					secondItem: 82,
				},
				{
					index: 		1,
					firstItem: 	32,
					secondItem: 81,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/mixes/34?contains_item=2",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: 	map[string]int{
					"combinations": 15,
				},
			},
			expUnique:  newExpUnique(34, "ultra potion"),
			category:	1,
			overdrive:	85,
			combinations: []testMixCombination{
				{
					index: 		0,
					firstItem: 	1,
					secondItem: 2,
				},
				{
					index: 		2,
					firstItem: 	2,
					secondItem: 17,
				},
				{
					index: 		14,
					firstItem: 	2,
					secondItem: 101,
				},
			},
		},
	}

	testSingleResources(t, tests, "GetMix", testCfg.HandleMixes, compareMixes)
}

func TestRetrieveMixes(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/mixes?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   64,
			results: []int32{1, 14, 38, 44, 55, 64},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/mixes?category=fire-elemental",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{11, 12, 13, 14, 15},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/mixes?req_item=45",
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{5, 9, 33, 58, 63},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/mixes?req_item=45&second_item=98",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{63},
		},
	}

	testIdList(t, tests, testCfg.e.mixes.endpoint, "RetrieveMixes", testCfg.HandleMixes, compareAPIResourceLists[NamedApiResourceList])
}
