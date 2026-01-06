package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetArea(t *testing.T) {
	tests := []struct {
		testGeneral
		expNameVer
		expAreas
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "Area with ID 0 doesn't exist. Max ID: 240",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/241",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "Area with ID 241 doesn't exist. Max ID: 240",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Wrong format.",
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
				expLocBased: expLocBased{
					sidequest: h.GetInt32Ptr(6),
					connectedAreas: []int32{144, 149},
					monsters: []int32{81, 84, 85},
					formations: []int32{204, 208},
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
				expLocBased: expLocBased{
					connectedAreas: []int32{26, 37, 41},
					characters: []int32{2, 4},
					treasures: []int32{33, 37},
					bgMusic: []int32{19},
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
				expLocBased: expLocBased{
					shops: []int32{5},
					cuesMusic: []int32{35},
					bgMusic: []int32{32, 34},
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
				expLocBased: expLocBased{
					sidequest: h.GetInt32Ptr(7),
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
				expLocBased: expLocBased{
					characters: []int32{5},
					monsters: []int32{19},
					formations: []int32{37},
					fmvsMusic: []int32{16},
					bossMusic: []int32{16},
					fmvs: []int32{9, 13},
				},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "GetArea", i+1, testCfg.HandleAreas)
		if correctErr {
			continue
		}

		test := test{
			t: t,
			cfg: testCfg,
			name: testName,
			expLengths: tc.expLengths,
			dontCheck: tc.dontCheck,
		}

		var got Area
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testExpectedNameVer(test, tc.expNameVer, got.ID, got.Name, got.Version)
		compAPIResourcesFromID(test, "location", testCfg.e.locations.endpoint, tc.parentLocation, got.ParentLocation)
		compAPIResourcesFromID(test, "sublocation", testCfg.e.sublocations.endpoint, tc.parentSublocation, got.ParentSublocation)
		compResPtrsFromID(test, testCfg.e.sidequests.endpoint, "sidequest", tc.sidequest, got.Sidequest)

		checks := []resListTest{
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
				newResListTestFromIDs("boss music", testCfg.e.songs.endpoint, tc.bossMusic, got.Music.BossFights),
			}

			checks = slices.Concat(checks, musicChecks)
		}

		testResourceLists(test, checks)
	}
}

func TestRetrieveAreas(t *testing.T) {
	tests := []struct {
		testGeneral
		expList
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?comp-sphere=fa",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value. usage: ?comp-sphere={boolean}",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?item=113",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided item ID 113 is out of range in item. Max ID: 112",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?key-item=61",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided key-item ID 61 is out of range in key-item. Max ID: 60",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?location=0",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided location ID 0 is out of range in location. Max ID: 26",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count: 240,
				next:  h.GetStrPtr("/areas?limit=20&offset=20"),
				results: []string{
					"/areas/1",
					"/areas/5",
					"/areas/20",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?limit=240",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count: 240,
				results: []string{
					"/areas/1",
					"/areas/50",
					"/areas/240",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?offset=50&limit=30",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    240,
				next:     h.GetStrPtr("/areas?limit=30&offset=80"),
				previous: h.GetStrPtr("/areas?limit=30&offset=20"),
				results: []string{
					"/areas/51",
					"/areas/80",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?monsters=true&chocobo=true&save-sphere=true",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count: 3,
				results: []string{
					"/areas/88",
					"/areas/97",
					"/areas/203",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?item=elixir&story-based=false&monsters=false",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count: 5,
				results: []string{
					"/areas/35",
					"/areas/129",
					"/areas/140",
					"/areas/163",
					"/areas/208",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?characters=true",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count: 7,
				results: []string{
					"/areas/1",
					"/areas/20",
					"/areas/103",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?sidequests=true",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count: 11,
				results: []string{
					"/areas/75",
					"/areas/140",
					"/areas/144",
					"/areas/145",
					"/areas/182",
					"/areas/185",
					"/areas/203",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?key-item=37",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count: 2,
				results: []string{
					"/areas/46",
					"/areas/169",
				},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "RetrieveAreas", i+1, testCfg.HandleAreas)
		if correctErr {
			continue
		}

		test := test{
			t: t,
			cfg: testCfg,
			name: testName,
			expLengths: tc.expLengths,
			dontCheck: tc.dontCheck,
		}

		var got LocationApiResourceList
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testAPIResourceList(test, tc.expList, got)
	}
}
