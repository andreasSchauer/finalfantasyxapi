package main

import (
	"encoding/json"
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetOverdriveMode(t *testing.T) {
	tests := []struct {
		testInOut
		expectedUnique
		expResOverdriveModes
	}{
		{
			testInOut: testInOut{
				requestURL:     "/api/overdrive-modes/ally/2",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    `Wrong format. Usage: /api/overdrive-modes/{name or id}`,
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/overdrive-modes/18",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "provided overdrive-mode ID is out of range. Max ID: 17",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/overdrive-modes/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "overdrive-mode not found: a.",
			},
		},
		{
			testInOut: testInOut{
				requestURL:     "/api/overdrive-modes/ally/",
				expectedStatus: http.StatusOK,
			},
			expectedUnique: expectedUnique{
				id:   14,
				name: "ally",
				lenMap: map[string]int{
					"actions": 7,
				},
			},
			expResOverdriveModes: expResOverdriveModes{
				description: "Charges on character's turn.",
				effect:      "The gauge fills at the start of the character's turn.",
				modeType:    "per-action",
				fillRate:    h.GetFloat32Ptr(0.03),
				actionsAmount: map[string]int32{
					"tidus": 600,
					"yuna": 500,
					"wakka": 350,
					"lulu": 480,
					"kimahri": 300,
					"auron": 450,
					"rikku": 320,
				},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testInOut, "GetOverdriveMode", i+1, testCfg.HandleOverdriveModes)
		if correctErr {
			continue
		}

		var o OverdriveMode
		if err := json.NewDecoder(rr.Body).Decode(&o); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testExpectedUnique(t, testName, tc.expectedUnique, o.ID, o.Name)

		testString(t, testName, "description", tc.description, o.Description)
		testString(t, testName, "effect", tc.effect, o.Effect)
		testString(t, testName, "type", tc.modeType, o.Type)
		testFloat32Ptr(t, testName, "fill rate", tc.fillRate, o.FillRate)

		testResourceAmount(t, testName, o.Name, o.Actions, tc.actionsAmount)
	}
}

/*
func TestRetrieveOverdriveModes(t *testing.T) {
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
		rr, testName, correctErr := setupTest(t, tc.testInOut, "RetrieveOverdriveModes", i+1, testCfg.HandleOverdriveModes)
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
*/
