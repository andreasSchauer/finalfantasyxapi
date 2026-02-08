package main

import (
	"errors"
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expSublocations struct {
	testGeneral
	expUnique
	parentLocation        int32
	connectedSublocations []int32
	areas                 []int32
	expLocRel
}

func TestGetSublocation(t *testing.T) {
	tests := []expSublocations{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sublocation with provided id '0' doesn't exist. max id: 41.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/42",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sublocation with provided id '42' doesn't exist. max id: 41.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "sublocation not found: 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/34/",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"connected sublocations": 4,
					"areas":                  6,
					"characters":             0,
					"aeons":                  0,
					"shops":                  3,
					"treasures":              9,
					"monsters":               51,
					"formations":             50,
					"sidequests":             2,
					"bg music":               2,
					"cues music":             0,
					"fmvs music":             0,
					"boss music":             1,
					"fmvs":                   0,
				},
			},
			expUnique: expUnique{
				id:   34,
				name: "calm lands",
			},
			parentLocation:        20,
			connectedSublocations: []int32{25, 35, 36, 37},
			areas:                 []int32{202, 203, 204, 205, 206, 207},
			expLocRel: expLocRel{
				shops:      []int32{31, 32, 33},
				treasures:  []int32{265, 268, 270, 273},
				monsters:   []int32{138, 142, 149, 259, 270, 282, 292},
				formations: []int32{193, 198, 205, 306, 312, 320, 331},
				sidequests: []int32{1, 4},
				bgMusic:    []int32{71, 73},
				bossMusic:  []int32{16},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/beSaiD",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"connected sublocations": 2,
					"areas":                  16,
					"characters":             1,
					"aeons":                  1,
					"shops":                  0,
					"treasures":              23,
					"monsters":               12,
					"formations":             14,
					"sidequests":             0,
					"bg music":               5,
					"cues music":             1,
					"fmvs music":             3,
					"boss music":             2,
					"fmvs":                   3,
				},
			},
			expUnique: expUnique{
				id:   6,
				name: "besaid",
			},
			parentLocation:        4,
			connectedSublocations: []int32{7, 8},
			areas:                 []int32{20, 24, 29, 31, 35},
			expLocRel: expLocRel{
				characters: []int32{3},
				aeons:      []int32{1},
				treasures:  []int32{15, 18, 24, 40, 44},
				monsters:   []int32{7, 11, 14, 18, 293},
				formations: []int32{9, 13, 17, 22, 25},
				bgMusic:    []int32{18, 20, 22, 23, 27},
				cuesMusic:  []int32{17},
				fmvsMusic:  []int32{17, 23, 34},
				bossMusic:  []int32{9, 34},
				fmvs:       []int32{5, 6, 8},
			},
		},
	}

	for i, tc := range tests {
		test, got, err := setupTest[Sublocation](t, tc.testGeneral, "GetSublocation", i+1, testCfg.HandleSublocations)
		if errors.Is(err, errCorrect) {
			continue
		}

		testExpectedUnique(test, tc.expUnique, got.ID, got.Name)
		compIdApiResource(test, "location", testCfg.e.locations.endpoint, tc.parentLocation, got.ParentLocation)

		compareResListTests(test, []resListTest{
			rltIDs("connected sublocations", testCfg.e.sublocations.endpoint, tc.connectedSublocations, got.ConnectedSublocations),
			rltIDs("areas", testCfg.e.areas.endpoint, tc.areas, got.Areas),
			rltIDs("characters", testCfg.e.characters.endpoint, tc.characters, got.Characters),
			rltIDs("aeons", testCfg.e.aeons.endpoint, tc.aeons, got.Aeons),
			rltIDs("shops", testCfg.e.shops.endpoint, tc.shops, got.Shops),
			rltIDs("treasures", testCfg.e.treasures.endpoint, tc.treasures, got.Treasures),
			rltIDs("monsters", testCfg.e.monsters.endpoint, tc.monsters, got.Monsters),
			rltIDs("formations", testCfg.e.monsterFormations.endpoint, tc.formations, got.Formations),
			rltIDs("sidequests", testCfg.e.sidequests.endpoint, tc.sidequests, got.Sidequests),
			rltIDs("fmvs", testCfg.e.fmvs.endpoint, tc.fmvs, got.FMVs),
		})

		if got.Music != nil {
			compareResListTests(test, []resListTest{
				rltIDs("bg music", testCfg.e.songs.endpoint, tc.bgMusic, got.Music.BackgroundMusic),
				rltIDs("cues music", testCfg.e.songs.endpoint, tc.cuesMusic, got.Music.Cues),
				rltIDs("fmvs music", testCfg.e.songs.endpoint, tc.fmvsMusic, got.Music.FMVs),
				rltIDs("boss music", testCfg.e.songs.endpoint, tc.bossMusic, got.Music.BossMusic),
			})
		}
	}
}

func TestRetrieveSublocations(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   41,
			results: []int32{1, 25, 41},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?monsters=false",
				expectedStatus: http.StatusOK,
			},
			count:   9,
			results: []int32{2, 9, 18, 21, 23},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?characters=true",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{1, 6, 7, 8, 13, 16, 19},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?aeons=true",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{3, 6, 11, 18, 27, 32, 35, 36},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations?key_item=36",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{7},
		},
	}

	testIdList(t, tests, testCfg.e.sublocations.endpoint, "RetrieveSublocations", testCfg.HandleSublocations, compareAPIResourceLists[NamedApiResourceList])
}

func TestSublocationsConnected(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/4/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          2,
			parentResource: h.GetStrPtr("/sublocations/4"),
			results:        []int32{5, 6},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/11/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          1,
			parentResource: h.GetStrPtr("/sublocations/11"),
			results:        []int32{10},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/26/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          3,
			parentResource: h.GetStrPtr("/sublocations/26"),
			results:        []int32{25, 27, 28},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/34/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          4,
			parentResource: h.GetStrPtr("/sublocations/34"),
			results:        []int32{25, 35, 36, 37},
		},
	}

	testIdList(t, tests, testCfg.e.sublocations.endpoint, "SubsectionSublocationsConnected", testCfg.HandleSublocations, compareSubResourceLists[NamedAPIResource, SublocationSub])
}
