package main

import (
	"encoding/json"
	"fmt"
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
				expectedErr:    "Couldn't get Area. Area with this ID doesn't exist.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/241",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "Couldn't get Area. Area with this ID doesn't exist.",
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
				parentLocation:    "/locations/15",
				parentSublocation: "/sublocations/25",
				expLocBased: expLocBased{
					sidequest: h.GetStrPtr("/sidequests/6"),
					connectedAreas: []string{
						"/areas/144",
						"/areas/149",
					},
					monsters: []string{
						"/monsters/81",
						"/monsters/84",
						"/monsters/85",
					},
					formations: []string{
						"/monster-formations/204",
						"/monster-formations/208",
					},
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
				parentLocation:    "/locations/4",
				parentSublocation: "/sublocations/7",
				expLocBased: expLocBased{
					connectedAreas: []string{
						"/areas/26",
						"/areas/37",
						"/areas/41",
					},
					characters: []string{
						"/characters/2",
						"/characters/4",
					},
					treasures: []string{
						"/treasures/33",
						"/treasures/37",
					},
					bgMusic: []string{
						"/songs/19",
					},
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
				parentLocation:    "/locations/8",
				parentSublocation: "/sublocations/13",
				expLocBased: expLocBased{
					shops: []string{
						"/shops/5",
					},
					cuesMusic: []string{
						"/songs/35",
					},
					bgMusic: []string{
						"/songs/32",
						"/songs/34",
					},
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
				parentLocation:    "/locations/14",
				parentSublocation: "/sublocations/24",
				expLocBased: expLocBased{
					sidequest: h.GetStrPtr("/sidequests/7"),
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
				parentLocation:    "/locations/5",
				parentSublocation: "/sublocations/8",
				expLocBased: expLocBased{
					characters: []string{
						"/characters/5",
					},
					monsters: []string{
						"/monsters/19",
					},
					formations: []string{
						"/monster-formations/37",
					},
					fmvsMusic: []string{
						"/songs/16",
					},
					bossMusic: []string{
						"/songs/16",
					},
					fmvs: []string{
						"/fmvs/9",
						"/fmvs/13",
					},
				},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "GetArea", i+1, testCfg.HandleAreas)
		if correctErr {
			continue
		}
		fmt.Println(testName)

		var got Area
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testExpectedNameVer(t, testName, tc.expNameVer, got.ID, got.Name, got.Version)

		compAPIResources(t, testCfg, testName, "location", tc.parentLocation, got.ParentLocation, tc.dontCheck)
		compAPIResources(t, testCfg, testName, "sublocation", tc.parentSublocation, got.ParentSublocation, tc.dontCheck)
		compResourcePtrs(t, testCfg, testName, "sidequest", tc.sidequest, got.Sidequest, tc.dontCheck)

		checks := []resListTest{
			newResListTest("connected areas", tc.connectedAreas, got.ConnectedAreas),
			newResListTest("characters", tc.characters, got.Characters),
			newResListTest("aeons", tc.aeons, got.Aeons),
			newResListTest("shops", tc.shops, got.Shops),
			newResListTest("treasures", tc.treasures, got.Treasures),
			newResListTest("monsters", tc.monsters, got.Monsters),
			newResListTest("formations", tc.formations, got.Formations),
			newResListTest("fmvs", tc.fmvs, got.FMVs),
		}

		if got.Music != nil {
			musicChecks := []resListTest{
				newResListTest("bg music", tc.bgMusic, got.Music.BackgroundMusic),
				newResListTest("cues music", tc.cuesMusic, got.Music.Cues),
				newResListTest("fmvs music", tc.fmvsMusic, got.Music.FMVs),
				newResListTest("boss music", tc.bossMusic, got.Music.BossFights),
			}

			checks = slices.Concat(checks, musicChecks)
		}

		testResourceLists(t, testCfg, testName, checks, tc.expLengths)
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
				expectedErr:    "invalid value. usage: comp-sphere={boolean}",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?item=113",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "provided item ID is out of range. Max ID: 112",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?key-item=61",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "provided key-item ID is out of range. Max ID: 60",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas?location=0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "provided location ID is out of range. Max ID: 26",
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

		var got LocationApiResourceList
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testAPIResourceList(t, testCfg, testName, tc.expList, got, tc.dontCheck)
	}
}
