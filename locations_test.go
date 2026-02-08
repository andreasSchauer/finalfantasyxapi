package main

import (
	"errors"
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expLocations struct {
	testGeneral
	expUnique
	connectedLocations []int32
	sublocations       []int32
	expLocRel
}

func TestGetLocation(t *testing.T) {
	tests := []expLocations{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "location with provided id '0' doesn't exist. max id: 26.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/27",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "location with provided id '27' doesn't exist. max id: 26.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "location not found: 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/15/",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"connected locations": 4,
					"sublocations":        3,
					"characters":          0,
					"aeons":               1,
					"shops":               6,
					"treasures":           28,
					"monsters":            19,
					"formations":          18,
					"sidequests":          1,
					"bg music":            11,
					"cues music":          3,
					"fmvs music":          1,
					"boss music":          3,
					"fmvs":                2,
				},
			},
			expUnique: expUnique{
				id:   15,
				name: "macalania",
			},
			connectedLocations: []int32{14, 16, 19, 20},
			sublocations:       []int32{25, 26, 27},
			expLocRel: expLocRel{
				aeons:      []int32{4},
				shops:      []int32{22, 25, 36},
				treasures:  []int32{187, 193, 199, 209, 214},
				monsters:   []int32{80, 83, 87, 93, 94, 97, 297},
				formations: []int32{120, 125, 129, 135, 137},
				sidequests: []int32{6},
				bgMusic:    []int32{12, 22, 30, 43, 52, 55, 56},
				cuesMusic:  []int32{4, 57, 59},
				fmvsMusic:  []int32{70},
				bossMusic:  []int32{16, 55, 57},
				fmvs:       []int32{27, 36},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/OmeGa_rUInS",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"connected locations": 0,
					"sublocations":        1,
					"characters":          0,
					"aeons":               0,
					"shops":               0,
					"treasures":           16,
					"monsters":            24,
					"formations":          22,
					"sidequests":          0,
					"bg music":            2,
					"cues music":          0,
					"fmvs music":          0,
					"boss music":          2,
					"fmvs":                0,
				},
			},
			expUnique: expUnique{
				id:   26,
				name: "omega ruins",
			},
			expLocRel: expLocRel{
				treasures:  []int32{328, 332, 337, 343},
				monsters:   []int32{190, 201, 210, 239, 245, 250, 255},
				formations: []int32{258, 262, 265, 283, 289, 296},
				bgMusic:    []int32{81, 82},
				bossMusic:  []int32{16, 80},
			},
		},
	}

	for i, tc := range tests {
		test, got, err := setupTest[Location](t, tc.testGeneral, "GetLocation", i+1, testCfg.HandleLocations)
		if errors.Is(err, errCorrect) {
			continue
		}

		testExpectedUnique(test, tc.expUnique, got.ID, got.Name)

		compareResListTests(test, []resListTest{
			rltIDs("connected locations", testCfg.e.locations.endpoint, tc.connectedLocations, got.ConnectedLocations),
			rltIDs("sublocations", testCfg.e.sublocations.endpoint, tc.sublocations, got.Sublocations),
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

func TestRetrieveLocations(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   26,
			results: []int32{1, 6, 18, 26},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?monsters=false",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{7, 13},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?characters=true",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{1, 4, 5, 8, 10, 12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?aeons=true",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{2, 4, 6, 11, 15, 19, 21, 22},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations?item=45&method=monster",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{18, 20, 25, 26},
		},
	}

	testIdList(t, tests, testCfg.e.locations.endpoint, "RetrieveLocations", testCfg.HandleLocations, compareAPIResourceLists[NamedApiResourceList])
}

func TestLocationsConnected(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/5/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          2,
			parentResource: h.GetStrPtr("/locations/5"),
			results:        []int32{4, 6},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/11/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          2,
			parentResource: h.GetStrPtr("/locations/11"),
			results:        []int32{10, 12},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/20/connected/",
				expectedStatus: http.StatusOK,
			},
			count:          4,
			parentResource: h.GetStrPtr("/locations/20"),
			results:        []int32{15, 21, 22, 23},
		},
	}

	testIdList(t, tests, testCfg.e.locations.endpoint, "SubsectionLocationsConnected", testCfg.HandleLocations, compareSubResourceLists[NamedAPIResource, LocationSub])
}
