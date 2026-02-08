package main

import (
	"errors"
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expOverdriveModes struct {
	testGeneral
	expUnique
	description   string
	effect        string
	modeType      int32
	fillRate      *float32
	actionsAmount map[string]int32
}

func TestGetOverdriveMode(t *testing.T) {
	tests := []expOverdriveModes{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/ally/2",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. usage: '/api/overdrive-modes/{name or id}'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/18",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "overdrive mode with provided id '18' doesn't exist. max id: 17.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "overdrive mode not found: 'a'.",
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
			description: "Charges on character's turn.",
			modeType:    2,
			fillRate:    h.GetFloat32Ptr(0.03),
			actionsAmount: map[string]int32{
				"/characters/1":   	600,
				"/characters/2":    500,
				"/characters/3":   	350,
				"/characters/4":    480,
				"/characters/5": 	300,
				"/characters/6":   	450,
				"/characters/7":   	320,
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
			effect:   "The gauge fills when an ally takes damage.",
			modeType: 1,
			fillRate: nil,
			actionsAmount: map[string]int32{
				"/characters/1": 300,
				"/characters/3": 100,
				"/characters/7": 100,
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
			modeType: 1,
		},
	}

	for i, tc := range tests {
		test, got, err := setupTest[OverdriveMode](t, tc.testGeneral, "GetOverdriveMode", i+1, testCfg.HandleOverdriveModes)
		if errors.Is(err, errCorrect) {
			continue
		}

		testExpectedUnique(test, tc.expUnique, got.ID, got.Name)
		compare(test, "description", tc.description, got.Description)
		compare(test, "effect", tc.effect, got.Effect)
		compIdApiResource(test, "type", testCfg.e.overdriveModeType.endpoint, tc.modeType, got.Type)
		compare(test, "fill rate", tc.fillRate, got.FillRate)
		checkResAmtsInSlice(test, "actions", tc.actionsAmount, got.Actions)
	}
}

func TestRetrieveOverdriveModes(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes?type=f",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid enum value: 'f'. use /api/type to see valid values.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/",
				expectedStatus: http.StatusOK,
			},
			count:    17,
			previous: nil,
			next:     nil,
			results:  []int32{1, 8, 17},
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
			count:   4,
			results: []int32{1, 2, 3, 4},
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
			count:   13,
			results: []int32{5, 12, 17},
		},
	}

	testIdList(t, tests, testCfg.e.overdriveModes.endpoint, "RetrieveOverdriveModes", testCfg.HandleOverdriveModes, compareAPIResourceLists[NamedApiResourceList])
}

func TestOverdriveModesSections(t *testing.T) {
	tests := []expListNames{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/sections",
				expectedStatus: http.StatusOK,
			},
			count: 0,
		},
	}

	testNameList(t, tests, testCfg.e.overdriveModes.endpoint, "OverdriveModeSections", testCfg.HandleOverdriveModes, compareSectionLists)
}
