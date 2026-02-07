package main

import (
	"errors"
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expMonsterFormations struct {
	testGeneral
	expIdOnly
	category        string
	isForcedAmbush  bool
	canEscape       bool
	bossMusic       *int32
	monsters        map[string]int32
	areas           []int32
	triggerCommands []testFormationTC
}

type testFormationTC struct {
	Ability int32
	Users   []int32
}

func compareFormationTCs(test test, exp testFormationTC, got FormationTriggerCommand) {
	tcEndpoint := test.cfg.e.triggerCommands.endpoint
	charClassesEndpoint := test.cfg.e.characterClasses.endpoint

	compIdApiResource(test, "tc ability", tcEndpoint, exp.Ability, got.Ability)
	compareResListTest(test, rltIDs("tc users", charClassesEndpoint, exp.Users, got.Users))
}

func TestGetMonsterFormation(t *testing.T) {
	tests := []expMonsterFormations{
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
			category:       "boss-fight",
			isForcedAmbush: false,
			canEscape:      false,
			bossMusic:      h.GetInt32Ptr(16),
			monsters: map[string]int32{
				"sinspawn echuilles": 1,
				"sinscale - 3":       4,
			},
			areas:           []int32{47},
			triggerCommands: []testFormationTC{},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/77",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"trigger commands": true,
				},
				expLengths: map[string]int{
					"monsters":         1,
					"areas":            3,
					"trigger commands": 0,
				},
			},
			expIdOnly: expIdOnly{
				id: 77,
			},
			category:       "random-encounter",
			isForcedAmbush: false,
			canEscape:      true,
			monsters: map[string]int32{
				"garuda - 3": 1,
			},
			areas: []int32{100, 101, 107},
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
			category:       "boss-fight",
			isForcedAmbush: false,
			canEscape:      false,
			bossMusic:      h.GetInt32Ptr(55),
			monsters: map[string]int32{
				"seymour":            1,
				"anima - 1":          1,
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
			category:       "random-encounter",
			isForcedAmbush: true,
			canEscape:      true,
			monsters: map[string]int32{
				"great malboro": 1,
			},
			areas:           []int32{236, 239, 240},
			triggerCommands: []testFormationTC{},
		},
	}

	for i, tc := range tests {
		test, got, err := setupTest[MonsterFormation](t, tc.testGeneral, "GetMonsterFormation", i+1, testCfg.HandleMonsterFormations)
		if errors.Is(err, errCorrect) {
			continue
		}

		testExpectedIdOnly(test, tc.expIdOnly, got.ID)

		compare(test, "category", tc.category, got.Category)
		compare(test, "is forced ambush", tc.isForcedAmbush, got.IsForcedAmbush)
		compare(test, "can escape", tc.canEscape, got.CanEscape)
		compIdApiResourcePtrs(test, "boss song", testCfg.e.songs.endpoint, tc.bossMusic, got.BossMusic)
		checkResAmtsInSlice(test, "monsters", tc.monsters, got.Monsters)
		compTestStructSlices(test, "trigger commands", tc.triggerCommands, got.TriggerCommands, compareFormationTCs)

		checks := []resListTest{
			rltIDs("areas", testCfg.e.areas.endpoint, tc.areas, got.Areas),
		}

		compareResListTests(test, checks)
	}
}

func TestRetrieveMonsterFormations(t *testing.T) {
	tests := []expListIDs{
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
			count:   331,
			results: []int32{1, 175, 238, 307, 331},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?monster=44",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{63, 64, 67, 68, 69},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?location=12",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{90, 93, 98, 102, 105},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?sublocation=6",
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{9, 15, 19, 24, 25},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?area=234",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{257, 258, 262, 266},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?ambush=true",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{226, 265, 298},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?category=boss-fight&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   63,
			results: []int32{3, 36, 83, 137, 189, 204, 245, 272, 296},
		},
	}

	testIdList(t, tests, testCfg.e.monsterFormations.endpoint, "RetrieveMonsterFormations", testCfg.HandleMonsterFormations, compareAPIResourceLists[UnnamedApiResourceList])
}
