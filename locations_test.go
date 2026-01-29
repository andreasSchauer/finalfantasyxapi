package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetLocation(t *testing.T) {
	tests := []struct {
		testGeneral
		expNameVer
		expAreas
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "area with provided id '0' doesn't exist. max id: 240.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/241",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "area with provided id '241' doesn't exist. max id: 240.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/145/",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"connected areas": 2,
					"monsters":        6,
					"formations":      6,
				},
			},
			expNameVer: expNameVer{
				id:      145,
				name:    "north",
				version: h.GetInt32Ptr(1),
			},
			expAreas: expAreas{
				parentLocation:    15,
				parentSublocation: 25,
				expLocRel: expLocRel{
					sidequests: []int32{6},
					monsters:   []int32{81, 84, 85},
					formations: []int32{214, 219},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/36",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"sidequests": true,
				},
				expLengths: map[string]int{
					"connected areas": 7,
					"monsters":        0,
					"characters":      2,
					"treasures":       6,
				},
			},
			expNameVer: expNameVer{
				id:      36,
				name:    "besaid village",
				version: nil,
			},
			expAreas: expAreas{
				parentLocation:    4,
				parentSublocation: 7,
				expLocRel: expLocRel{
					characters: []int32{2, 4},
					treasures:  []int32{33, 37},
					bgMusic:    []int32{19},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/69",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"sidequests": true,
				},
				expLengths: map[string]int{
					"connected areas": 6,
					"shops":           1,
					"bg music":        2,
					"cues music":      1,
				},
			},
			expNameVer: expNameVer{
				id:      69,
				name:    "main gate",
				version: nil,
			},
			expAreas: expAreas{
				parentLocation:    8,
				parentSublocation: 13,
				expLocRel: expLocRel{
					shops:     []int32{5},
					cuesMusic: []int32{35},
					bgMusic:   []int32{32, 34},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/140",
				expectedStatus: http.StatusOK,
			},
			expNameVer: expNameVer{
				id:      140,
				name:    "agency front",
				version: nil,
			},
			expAreas: expAreas{
				parentLocation:    14,
				parentSublocation: 24,
				expLocRel: expLocRel{
					sidequests: []int32{7},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/42",
				expectedStatus: http.StatusOK,
				expLengths: map[string]int{
					"characters": 1,
					"formations": 1,
					"monsters":   2,
					"fmvs music": 1,
					"boss music": 1,
					"fmvs":       5,
				},
			},
			expNameVer: expNameVer{
				id:      42,
				name:    "deck",
				version: nil,
			},
			expAreas: expAreas{
				parentLocation:    5,
				parentSublocation: 8,
				expLocRel: expLocRel{
					characters: []int32{5},
					monsters:   []int32{19},
					formations: []int32{36},
					fmvsMusic:  []int32{16},
					bossMusic:  []int32{16},
					fmvs:       []int32{9, 13},
				},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, err := setupTest(t, tc.testGeneral, "GetArea", i+1, testCfg.HandleAreas)
		if errors.Is(err, errCorrect) {
			continue
		}

		test := test{
			t:          t,
			cfg:        testCfg,
			name:       testName,
			expLengths: tc.expLengths,
			dontCheck:  tc.dontCheck,
		}

		var got Area
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testExpectedNameVer(test, tc.expNameVer, got.ID, got.Name, got.Version)
		compAPIResourcesFromID(test, "location", testCfg.e.locations.endpoint, tc.parentLocation, got.ParentLocation)
		compAPIResourcesFromID(test, "sublocation", testCfg.e.sublocations.endpoint, tc.parentSublocation, got.ParentSublocation)

		checks := []resListTest{
			newResListTestFromIDs("sidequests", testCfg.e.sidequests.endpoint, tc.sidequests, got.Sidequests),
			newResListTestFromIDs("connected areas", testCfg.e.areas.endpoint, tc.connectedAreas, got.ConnectedAreas),
			newResListTestFromIDs("characters", testCfg.e.characters.endpoint, tc.characters, got.Characters),
			newResListTestFromIDs("aeons", testCfg.e.aeons.endpoint, tc.aeons, got.Aeons),
			newResListTestFromIDs("shops", testCfg.e.shops.endpoint, tc.shops, got.Shops),
			newResListTestFromIDs("treasures", testCfg.e.treasures.endpoint, tc.treasures, got.Treasures),
			newResListTestFromIDs("monsters", testCfg.e.monsters.endpoint, tc.monsters, got.Monsters),
			newResListTestFromIDs("formations", testCfg.e.monsterFormations.endpoint, tc.formations, got.Formations),
			newResListTestFromIDs("fmvs", testCfg.e.fmvs.endpoint, tc.fmvs, got.FMVs),
		}

		if got.Music != nil {
			musicChecks := []resListTest{
				newResListTestFromIDs("bg music", testCfg.e.songs.endpoint, tc.bgMusic, got.Music.BackgroundMusic),
				newResListTestFromIDs("cues music", testCfg.e.songs.endpoint, tc.cuesMusic, got.Music.Cues),
				newResListTestFromIDs("fmvs music", testCfg.e.songs.endpoint, tc.fmvsMusic, got.Music.FMVs),
				newResListTestFromIDs("boss music", testCfg.e.songs.endpoint, tc.bossMusic, got.Music.BossMusic),
			}

			checks = slices.Concat(checks, musicChecks)
		}

		testResourceLists(test, checks)
	}
}
