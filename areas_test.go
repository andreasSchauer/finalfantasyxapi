package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetArea(t *testing.T) {
	tests := []struct {
		testInOut
		expectedSingle
		expResAreas
	}{
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/0",
				expectedStatus: http.StatusNotFound,
				expectedErr: 	"Couldn't get Area. Area with this ID doesn't exist.",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/241",
				expectedStatus: http.StatusNotFound,
				expectedErr: 	"Couldn't get Area. Area with this ID doesn't exist.",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/a",
				expectedStatus: http.StatusBadRequest,
				expectedErr: 	"Wrong format.",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/145/",
				expectedStatus: http.StatusOK,
			},
			expectedSingle: expectedSingle{
				id:     	145,
				name:    	"north",
				version: 	h.GetInt32Ptr(1),
				lenMap: 	map[string]int{
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
			expectedSingle: expectedSingle{
				id:      36,
				name:    "besaid village",
				version: nil,
				lenMap: map[string]int{
					"connected areas": 	7,
					"monsters":        	0,
					"characters":      	2,
					"treasures":		6,
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
			expectedSingle: expectedSingle{
				id:      69,
				name:    "main gate",
				version: nil,
				lenMap: map[string]int{
					"connected areas": 	6,
					"shops":        	1,
					"bg music":      	2,
					"cues music":		1,
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
			expectedSingle: expectedSingle{
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
			expectedSingle: expectedSingle{
				id:      42,
				name:    "deck",
				version: nil,
				lenMap: map[string]int{
					"characters":		1,
					"formations":		1,
					"monsters":			2,
					"fmvs music":		1,
					"boss music":		1,
					"fmvs":				5,
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
		testName := getTestName("GetArea", tc.requestURL, i+1)

		req := httptest.NewRequest(http.MethodGet, tc.requestURL, nil)
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(testCfg.HandleAreas)
		handler.ServeHTTP(rr, req)

		if rr.Code != tc.expectedStatus {
			t.Fatalf("%s: expected %d, got %d, body=%s", testName, tc.expectedStatus, rr.Code, rr.Body.String())
		}

		if tc.expectedErr != "" {
			raw := rr.Body.String()
			if !strings.Contains(raw, tc.expectedErr) {
				t.Fatalf("%s: expected error message to contain %s, got %q", testName, tc.expectedErr, raw)
			}
			continue
		}

		var a Area
		if err := json.NewDecoder(rr.Body).Decode(&a); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		if a.ID != tc.id {
			t.Fatalf("%s: expected id %d, got %d", testName, tc.id, a.ID)
		}

		if a.Name != tc.name {
			t.Fatalf("%s: expected name %s, got %s", testName, tc.name, a.Name)
		}

		if h.DerefOrNil(a.Version) != h.DerefOrNil(tc.version) {
			t.Fatalf("%s: expected version %d, got %d", testName, h.DerefOrNil(tc.version), h.DerefOrNil(a.Version))
		}

		if !resourcesMatch(testCfg, a.ParentLocation, tc.parentLocation) {
			t.Fatalf("%s: expected location %s, got %s", testName, tc.parentLocation, a.ParentLocation.URL)
		}

		if !resourcesMatch(testCfg, a.ParentSublocation, tc.parentSublocation) {
			t.Fatalf("%s: expected sublocation %s, got %s", testName, tc.parentSublocation, a.ParentSublocation)
		}

		if !resourcePtrsMatch(testCfg, a.Sidequest, tc.sidequest) {
			t.Fatalf("%s: expected sidequest %s, got %s", testName, h.DerefOrNil(tc.sidequest), derefResourcePtr(a.Sidequest).ToKeyFields())
		}

		checks := []testCheck{
			{name: "connected areas", got: toIface(a.ConnectedAreas), expected: tc.connectedAreas},
			{name: "characters", got: toIface(a.Characters), expected: tc.characters},
			{name: "aeons", got: toIface(a.Aeons), expected: tc.aeons},
			{name: "shops", got: toIface(a.Shops), expected: tc.shops},
			{name: "treasures", got: toIface(a.Treasures), expected: tc.treasures},
			{name: "monsters", got: toIface(a.Monsters), expected: tc.monsters},
			{name: "formations", got: toIface(a.Formations), expected: tc.formations},
			{name: "fmvs", got: toIface(a.FMVs), expected: tc.fmvs},
		}

		if a.Music != nil {
			musicChecks := []testCheck{
				{name: "bg music", got: toIface(a.Music.BackgroundMusic), expected: tc.bgMusic},
				{name: "cues music", got: toIface(a.Music.Cues), expected: tc.cuesMusic},
				{name: "fmvs music", got: toIface(a.Music.FMVs), expected: tc.fmvsMusic},
				{name: "boss music", got: toIface(a.Music.BossFights), expected: tc.bossMusic},
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
				expectedErr: "invalid value. usage: comp-sphere={boolean}",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas?item=113",
				expectedStatus: http.StatusNotFound,
				expectedErr: "provided item ID is out of range. Max ID: 112",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas?key-item=61",
				expectedStatus: http.StatusNotFound,
				expectedErr: "provided key-item ID is out of range. Max ID: 60",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas?location=0",
				expectedStatus: http.StatusNotFound,
				expectedErr: "provided location ID is out of range. Max ID: 26",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/areas/",
				expectedStatus: http.StatusOK,
			},
			expectedList: expectedList{
				count: 		240,
				next: 		h.GetStrPtr("/areas?limit=20&offset=20"),
				results: 	[]string{
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
				count:		240,
				results: 	[]string{
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
				count: 		240,
				next: 		h.GetStrPtr("/areas?limit=30&offset=80"),
				previous: 	h.GetStrPtr("/areas?limit=30&offset=20"),
				results: 	[]string{
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
				count: 		3,
				results: 	[]string{
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
				count: 		5,
				results: 	[]string{
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
				count: 		7,
				results: 	[]string{
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
				count: 		11,
				results: 	[]string{
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
				count: 		2,
				results: 	[]string{
					"/areas/46",
					"/areas/169",
				},
			},
		},
	}

	for i, tc := range tests {
		testName := getTestName("RetrieveAreas", tc.requestURL, i+1)

		req := httptest.NewRequest(http.MethodGet, tc.requestURL, nil)
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(testCfg.HandleAreas)
		handler.ServeHTTP(rr, req)

		if rr.Code != tc.expectedStatus {
			t.Fatalf("%s: expected %d, got %d, body=%s", testName, tc.expectedStatus, rr.Code, rr.Body.String())
		}

		if tc.expectedErr != "" {
			raw := rr.Body.String()
			if !strings.Contains(raw, tc.expectedErr) {
				t.Fatalf("%s: expected error message to contain %s, got %q", testName, tc.expectedErr, raw)
			}
			continue
		}

		var res LocationApiResourceList
		if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		if res.Count != tc.count {
			t.Fatalf("%s: expected count %d, got %d", testName, tc.count, res.Count)
		}

		if !paginationURLsMatch(testCfg, res.Previous, tc.previous) {
			t.Fatalf("%s: expected previous url %s, got %s", testName, h.DerefOrNil(tc.previous), h.DerefOrNil(res.Previous))
		}

		if !paginationURLsMatch(testCfg, res.Next, tc.next) {
			t.Fatalf("%s: expected next url %s, got %s", testName, h.DerefOrNil(tc.next), h.DerefOrNil(res.Next))
		}

		checks := []testCheck{
			{
				name: "results",
				got: toIface(res.Results),
				expected: tc.results,
			},
		}

		testResponseChecks(t, testCfg, testName, checks, nil)
	}
}