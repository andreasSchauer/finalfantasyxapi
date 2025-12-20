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
		testInOut
		expectedNameVer
		expResAreas
	}{
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "Couldn't get Area. Area with this ID doesn't exist.",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/241",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "Couldn't get Area. Area with this ID doesn't exist.",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Wrong format.",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/145/",
				expectedStatus: http.StatusOK,
			},
			expectedNameVer: expectedNameVer{
				id:      145,
				name:    "north",
				version: h.GetInt32Ptr(1),
				lenMap: map[string]int{
					"connected areas": 2,
					"monsters":        6,
					"formations":      6,
				},
			},
			expResAreas: expResAreas{
				parentLocation:    "/locations/15",
				parentSublocation: "/sublocations/25",
				locBasedExpect: locBasedExpect{
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
			testInOut: testInOut{
				requestURL:     "/api/areas/36",
				expectedStatus: http.StatusOK,
			},
			expectedNameVer: expectedNameVer{
				id:      36,
				name:    "besaid village",
				version: nil,
				lenMap: map[string]int{
					"connected areas": 7,
					"monsters":        0,
					"characters":      2,
					"treasures":       6,
				},
			},
			expResAreas: expResAreas{
				parentLocation:    "/locations/4",
				parentSublocation: "/sublocations/7",
				locBasedExpect: locBasedExpect{
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
			testInOut: testInOut{
				requestURL:     "/api/areas/69",
				expectedStatus: http.StatusOK,
			},
			expectedNameVer: expectedNameVer{
				id:      69,
				name:    "main gate",
				version: nil,
				lenMap: map[string]int{
					"connected areas": 6,
					"shops":           1,
					"bg music":        2,
					"cues music":      1,
				},
			},
			expResAreas: expResAreas{
				parentLocation:    "/locations/8",
				parentSublocation: "/sublocations/13",
				locBasedExpect: locBasedExpect{
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
			testInOut: testInOut{
				requestURL:     "/api/areas/140",
				expectedStatus: http.StatusOK,
			},
			expectedNameVer: expectedNameVer{
				id:      140,
				name:    "agency front",
				version: nil,
			},
			expResAreas: expResAreas{
				parentLocation:    "/locations/14",
				parentSublocation: "/sublocations/24",
				locBasedExpect: locBasedExpect{
					sidequest: h.GetStrPtr("/sidequests/7"),
				},
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/42",
				expectedStatus: http.StatusOK,
			},
			expectedNameVer: expectedNameVer{
				id:      42,
				name:    "deck",
				version: nil,
				lenMap: map[string]int{
					"characters": 1,
					"formations": 1,
					"monsters":   2,
					"fmvs music": 1,
					"boss music": 1,
					"fmvs":       5,
				},
			},
			expResAreas: expResAreas{
				parentLocation:    "/locations/5",
				parentSublocation: "/sublocations/8",
				locBasedExpect: locBasedExpect{
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
		rr, testName, correctErr := setupTest(t, tc.testInOut, "GetArea", i+1, testCfg.HandleAreas)
		if correctErr {
			continue
		}

		var a Area
		if err := json.NewDecoder(rr.Body).Decode(&a); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testExpectedNameVer(t, testName, tc.expectedNameVer, a.ID, a.Name, a.Version)

		testResourceMatch(t, testCfg, testName, "location", tc.parentLocation, a.ParentLocation)
		testResourceMatch(t, testCfg, testName, "sublocation", tc.parentSublocation, a.ParentSublocation)
		testResourcePtrMatch(t, testCfg, testName, "sidequest", tc.sidequest, a.Sidequest)

		checks := []testCheck{
			{name: "connected areas", got: toHasAPIResSlice(a.ConnectedAreas), expected: tc.connectedAreas},
			{name: "characters", got: toHasAPIResSlice(a.Characters), expected: tc.characters},
			{name: "aeons", got: toHasAPIResSlice(a.Aeons), expected: tc.aeons},
			{name: "shops", got: toHasAPIResSlice(a.Shops), expected: tc.shops},
			{name: "treasures", got: toHasAPIResSlice(a.Treasures), expected: tc.treasures},
			{name: "monsters", got: toHasAPIResSlice(a.Monsters), expected: tc.monsters},
			{name: "formations", got: toHasAPIResSlice(a.Formations), expected: tc.formations},
			{name: "fmvs", got: toHasAPIResSlice(a.FMVs), expected: tc.fmvs},
		}

		if a.Music != nil {
			musicChecks := []testCheck{
				{name: "bg music", got: toHasAPIResSlice(a.Music.BackgroundMusic), expected: tc.bgMusic},
				{name: "cues music", got: toHasAPIResSlice(a.Music.Cues), expected: tc.cuesMusic},
				{name: "fmvs music", got: toHasAPIResSlice(a.Music.FMVs), expected: tc.fmvsMusic},
				{name: "boss music", got: toHasAPIResSlice(a.Music.BossFights), expected: tc.bossMusic},
			}

			checks = slices.Concat(checks, musicChecks)
		}

		testResponseChecks(t, testCfg, testName, checks, tc.lenMap)
	}
}

func TestRetrieveAreas(t *testing.T) {
	tests := []struct {
		testInOut
		expectedList
	}{
		{
			testInOut: testInOut{
				requestURL:     "/api/areas?comp-sphere=fa",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value. usage: comp-sphere={boolean}",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas?item=113",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "provided item ID is out of range. Max ID: 112",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas?key-item=61",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "provided key-item ID is out of range. Max ID: 60",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas?location=0",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "provided location ID is out of range. Max ID: 26",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/",
				expectedStatus: http.StatusOK,
			},
			expectedList: expectedList{
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
			testInOut: testInOut{
				requestURL:     "/api/areas?limit=240",
				expectedStatus: http.StatusOK,
			},
			expectedList: expectedList{
				count: 240,
				results: []string{
					"/areas/1",
					"/areas/50",
					"/areas/240",
				},
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas?offset=50&limit=30",
				expectedStatus: http.StatusOK,
			},
			expectedList: expectedList{
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
			testInOut: testInOut{
				requestURL:     "/api/areas?monsters=true&chocobo=true&save-sphere=true",
				expectedStatus: http.StatusOK,
			},
			expectedList: expectedList{
				count: 3,
				results: []string{
					"/areas/88",
					"/areas/97",
					"/areas/203",
				},
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas?item=elixir&story-based=false&monsters=false",
				expectedStatus: http.StatusOK,
			},
			expectedList: expectedList{
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
			testInOut: testInOut{
				requestURL:     "/api/areas?characters=true",
				expectedStatus: http.StatusOK,
			},
			expectedList: expectedList{
				count: 7,
				results: []string{
					"/areas/1",
					"/areas/20",
					"/areas/103",
				},
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas?sidequests=true",
				expectedStatus: http.StatusOK,
			},
			expectedList: expectedList{
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
			testInOut: testInOut{
				requestURL:     "/api/areas?key-item=37",
				expectedStatus: http.StatusOK,
			},
			expectedList: expectedList{
				count: 2,
				results: []string{
					"/areas/46",
					"/areas/169",
				},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testInOut, "RetrieveAreas", i+1, testCfg.HandleAreas)
		if correctErr {
			continue
		}

		var res LocationApiResourceList
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testAPIResourceList(t, testCfg, testName, res, tc.expectedList)
	}
}
