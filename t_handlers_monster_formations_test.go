package main

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetMonsterFormation(t *testing.T) {
	tests := []expMonsterFormation{
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
				expLengths: 	map[string]int{
					"monsters":         2,
					"areas":            1,
					"trigger commands": 0,
				},
			},
			expIdOnly: newExpIdOnly(27),
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
			expIdOnly: newExpIdOnly(77),
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
			expIdOnly: newExpIdOnly(137),
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
			expIdOnly: newExpIdOnly(265),
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

	testSingleResources(t, tests, "GetMonsterFormation", testCfg.HandleMonsterFormations, compareMonsterFormations)
}

func TestRetrieveMonsterFormations(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations?limit=asd",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'asd' used for parameter 'limit'. usage: '?limit{int|'max'}'.",
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


func TestSubsectionMonsterFormations(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/23/monster-formations",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleLocations,
			},
			count:          18,
			parentResource: h.GetStrPtr("/locations/23"),
			results:        []int32{228, 231, 235, 237, 240, 244, 245},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/4/monster-formations",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleLocations,
			},
			count:          14,
			parentResource: h.GetStrPtr("/locations/4"),
			results:        []int32{9, 13, 15, 18, 21, 23, 25},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/36/monster-formations",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleSublocations,
			},
			count:          14,
			parentResource: h.GetStrPtr("/sublocations/36"),
			results:        []int32{198, 202, 216, 217, 221, 224, 227},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/17/monster-formations",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleSublocations,
			},
			count:          10,
			parentResource: h.GetStrPtr("/sublocations/17"),
			results:        []int32{86, 89, 90, 93, 95},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/89/monster-formations",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleAreas,
			},
			count:          10,
			parentResource: h.GetStrPtr("/areas/89"),
			results:        []int32{46, 49, 54, 57, 58, 59},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/172/monster-formations",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleAreas,
			},
			count:          12,
			parentResource: h.GetStrPtr("/areas/172"),
			results:        []int32{140, 144, 149, 150, 153, 155},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/200/monster-formations",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleMonsters,
			},
			count:          1,
			parentResource: h.GetStrPtr("/monsters/200"),
			results:        []int32{170},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/155/monster-formations",
				expectedStatus: http.StatusOK,
				handler: 		testCfg.HandleMonsters,
			},
			count:          2,
			parentResource: h.GetStrPtr("/monsters/155"),
			results:        []int32{216, 232},
		},
	}

	testIdList(t, tests, testCfg.e.monsterFormations.endpoint, "SubsectionMonsterFormations", nil, compareSubResourceLists[UnnamedAPIResource, MonsterFormationSub])
}