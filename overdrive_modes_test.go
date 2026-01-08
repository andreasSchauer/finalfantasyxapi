package main

import (
	"encoding/json"
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetOverdriveMode(t *testing.T) {
	tests := []struct {
		testGeneral
		expUnique
		expOverdriveModes
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/ally/2",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    `Wrong format. Usage: /api/overdrive-modes/{name or id}`,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/18",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "provided overdrive-mode ID is out of range. Max ID: 17",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "overdrive-mode not found: a.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/ally/",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"effect": true,
				},
				expLengths: map[string]int{
					"actions": 7,
				},
			},
			expUnique: expUnique{
				id:   14,
				name: "ally",
			},
			expOverdriveModes: expOverdriveModes{
				description: "Charges on character's turn.",
				modeType:    2,
				fillRate:    h.GetFloat32Ptr(0.03),
				actionsAmount: map[string]int32{
					"tidus":   600,
					"yuna":    500,
					"wakka":   350,
					"lulu":    480,
					"kimahri": 300,
					"auron":   450,
					"rikku":   320,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/3",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"description": true,
				},
			},
			expUnique: expUnique{
				id:   3,
				name: "comrade",
			},
			expOverdriveModes: expOverdriveModes{
				effect:   "The gauge fills when an ally takes damage.",
				modeType: 1,
				fillRate: nil,
				actionsAmount: map[string]int32{
					"tidus": 300,
					"wakka": 100,
					"rikku": 100,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/1/",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"description": true,
					"effect":      true,
					"fill rate":   true,
				},
				expLengths: map[string]int{
					"actions": 0,
				},
			},
			expUnique: expUnique{
				id:   1,
				name: "stoic",
			},
			expOverdriveModes: expOverdriveModes{
				modeType: 1,
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "GetOverdriveMode", i+1, testCfg.HandleOverdriveModes)
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

		var got OverdriveMode
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testExpectedUnique(test, tc.expUnique, got.ID, got.Name)
		compare(test, "description", tc.description, got.Description)
		compare(test, "effect", tc.effect, got.Effect)
		compAPIResourcesFromID(test, "type", testCfg.e.overdriveModeType.endpoint, tc.modeType, got.Type)
		compare(test, "fill rate", tc.fillRate, got.FillRate)
		checkResAmtsInSlice(test, "actions", tc.actionsAmount, got.Actions)
	}
}

func TestRetrieveOverdriveModes(t *testing.T) {
	tests := []struct {
		testGeneral
		expList
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes?type=f",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value: 'f', use /api/type to see valid values",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    17,
				previous: nil,
				next:     nil,
				results: []string{
					"/overdrive-modes/1",
					"/overdrive-modes/8",
					"/overdrive-modes/17",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes?type=formula",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"previous": true,
					"next":     true,
				},
			},
			expList: expList{
				count: 4,
				results: []string{
					"/overdrive-modes/1",
					"/overdrive-modes/2",
					"/overdrive-modes/3",
					"/overdrive-modes/4",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes?type=2",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"previous": true,
					"next":     true,
				},
			},
			expList: expList{
				count: 13,
				results: []string{
					"/overdrive-modes/5",
					"/overdrive-modes/12",
					"/overdrive-modes/17",
				},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "RetrieveOverdriveModes", i+1, testCfg.HandleOverdriveModes)
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

		var got NamedApiResourceList
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testAPIResourceList(test, tc.expList, got)
	}
}
