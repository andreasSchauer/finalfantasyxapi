package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetMonsterFormation(t *testing.T) {
	tests := []struct {
		testGeneral
		expIdOnly
		expMonsterFormations
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/332",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "monster formation with provided id '332' doesn't exist. max id: 331.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/27",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         2,
					"areas":            1,
					"trigger commands": 0,
				},
			},
			expIdOnly: expIdOnly{
				id: 27,
			},
			expMonsterFormations: expMonsterFormations{
				category:       "boss-fight",
				isForcedAmbush: false,
				canEscape:      false,
				bossMusic:      h.GetInt32Ptr(16),
				monsters: map[string]int32{
					"sinspawn echuilles": 1,
					"sinscale - 3": 4,
				},
				areas: []int32{47},
				triggerCommands: []testFormationTC{},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/77",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         1,
					"areas":            3,
					"trigger commands": 0,
				},
			},
			expIdOnly: expIdOnly{
				id: 77,
			},
			expMonsterFormations: expMonsterFormations{
				category:       "random-encounter",
				isForcedAmbush: false,
				canEscape:      true,
				monsters: map[string]int32{
					"garuda - 3": 1,
				},
				areas: []int32{100, 101, 107},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/137",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         3,
					"areas":            1,
					"trigger commands": 2,
				},
			},
			expIdOnly: expIdOnly{
				id: 137,
			},
			expMonsterFormations: expMonsterFormations{
				category:       "boss-fight",
				isForcedAmbush: false,
				canEscape:      false,
				bossMusic:      h.GetInt32Ptr(55),
				monsters: map[string]int32{
					"seymour": 1,
					"anima - 1": 1,
					"guado guardian - 1": 2,
				},
				areas: []int32{166},
				triggerCommands: []testFormationTC{
					{
						Ability: 4,
						Users:   []int32{1},
					},
					{
						Ability: 5,
						Users:   []int32{3, 5},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/265",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         1,
					"areas":            3,
					"trigger commands": 0,
				},
			},
			expIdOnly: expIdOnly{
				id: 265,
			},
			expMonsterFormations: expMonsterFormations{
				category:       "random-encounter",
				isForcedAmbush: true,
				canEscape:      true,
				monsters: map[string]int32{
					"great malboro": 1,
				},
				areas: []int32{236, 239, 240},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, err := setupTest(t, tc.testGeneral, "GetMonsterFormation", i+1, testCfg.HandleMonsterFormations)
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

		var got MonsterFormation
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testExpectedIdOnly(test, tc.expIdOnly, got.ID)

		compare(test, "category", tc.category, got.Category)
		compare(test, "is forced ambush", tc.isForcedAmbush, got.IsForcedAmbush)
		compare(test, "can escape", tc.canEscape, got.CanEscape)
		compResPtrsFromID(test, "boss song", testCfg.e.songs.endpoint, tc.bossMusic, got.BossMusic)
		checkResAmtsInSlice(test, "monsters", tc.monsters, got.Monsters)
		compareCustomObjSlices(test, "trigger commands", tc.triggerCommands, got.TriggerCommands, compareFormationTCs)

		checks := []resListTest{
			newResListTestFromIDs("areas", testCfg.e.areas.endpoint, tc.areas, got.Areas),
		}

		testResourceLists(test, checks)
	}
}


func TestRetrieveMonsterFormations(t *testing.T) {
	tests := []struct {
		testGeneral
		expList
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?limit=asd",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'asd' for parameter 'limit'. usage: '?limit{integer or 'max'}'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?limit=max",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    331,
				previous: nil,
				next:     nil,
				results:  []int32{1, 175, 238, 307, 331},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?monster=44",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    5,
				previous: nil,
				next:     nil,
				results:  []int32{63, 64, 67, 68, 69},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?location=12",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    12,
				previous: nil,
				next:     nil,
				results:  []int32{90, 93, 98, 102, 105},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?sublocation=6",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    14,
				previous: nil,
				next:     nil,
				results:  []int32{9, 15, 19, 24, 25},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?area=234",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    8,
				previous: nil,
				next:     nil,
				results:  []int32{257, 258, 262, 266},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?ambush=true",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    3,
				previous: nil,
				next:     nil,
				results:  []int32{226, 265, 298},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?category=boss-fight&limit=max",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    63,
				previous: nil,
				next:     nil,
				results:  []int32{3, 36, 83, 137, 189, 204, 245, 272, 296},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, err := setupTest(t, tc.testGeneral, "RetrieveMonsterFormations", i+1, testCfg.HandleMonsterFormations)
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

		var got UnnamedApiResourceList
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testAPIResourceList(test, testCfg.e.monsterFormations.endpoint, tc.expList, got)
	}
}
