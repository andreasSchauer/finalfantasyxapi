package main

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetOverdriveMode(t *testing.T) {
	tests := []expOverdriveMode{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/ally/2",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. usage: '/api/overdrive-modes', '/api/overdrive-modes/{id}', '/api/overdrive-modes/{name}'",
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
			expUnique: newExpUnique(14, "ally"),
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
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/3",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"description": true,
				},
			},
			expUnique: newExpUnique(3, "comrade"),
			effect:   "The gauge fills when an ally takes damage.",
			modeType: 1,
			fillRate: nil,
			actionsAmount: map[string]int32{
				"tidus": 300,
				"wakka": 100,
				"rikku": 100,
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
			expUnique: newExpUnique(1, "stoic"),
			modeType: 1,
		},
	}

	testSingleResources(t, tests, "GetOverdriveMode", testCfg.HandleOverdriveModes, compareOverdriveModes)
}

func TestRetrieveOverdriveModes(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes?type=f",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid enum value 'f' used for parameter 'type'. use /api/overdrive-modes/parameters to see allowed values.",
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
