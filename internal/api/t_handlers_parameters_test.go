package api

import (
	"net/http"
	"testing"
)

func TestParameters(t *testing.T) {
	tests := []expListNames{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.areas.endpoint,
				handler:        testCfg.HandleAreas,
			},
			count:   20,
			results: []string{"limit", "offset", "item", "save_sphere", "sublocation"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.sublocations.endpoint,
				handler:        testCfg.HandleSublocations,
			},
			count:   14,
			results: []string{"location", "item", "method", "aeons", "shops", "fmvs"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.locations.endpoint,
				handler:        testCfg.HandleLocations,
			},
			count:   13,
			results: []string{"key_item", "characters", "aeons", "treasures", "sidequests"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.arenaCreations.endpoint,
				handler:        testCfg.HandleArenaCreations,
			},
			count:   3,
			results: []string{"limit", "offset", "category"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.blitzballPrizes.endpoint,
				handler:        testCfg.HandleBlitzballPrizes,
			},
			count:   3,
			results: []string{"limit", "offset", "category"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.fmvs.endpoint,
				handler:        testCfg.HandleFMVs,
			},
			count:   3,
			results: []string{"limit", "offset", "location"},
		},

		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.monsters.endpoint,
				handler:        testCfg.HandleMonsters,
			},
			count:   27,
			results: []string{"kimahri_stats", "aeon_stats", "altered_state", "omnis_elements", "status_resists", "auto_ability", "area", "distance", "underwater", "species"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.monsterFormations.endpoint,
				handler:        testCfg.HandleMonsterFormations,
			},
			count:   8,
			results: []string{"monster", "location", "ambush", "category"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.overdriveModes.endpoint,
				handler:        testCfg.HandleOverdriveModes,
			},
			count:   3,
			results: []string{"limit", "offset", "type"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.shops.endpoint,
				handler:        testCfg.HandleShops,
			},
			count:   10,
			results: []string{"location", "auto_ability", "items", "equipment", "pre_airship"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.songs.endpoint,
				handler:        testCfg.HandleSongs,
			},
			count:   9,
			results: []string{"area", "fmvs", "composer", "special_use"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.sidequests.endpoint,
				handler:        testCfg.HandleSidequests,
			},
			count:   2,
			results: []string{"limit", "offset"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.subquests.endpoint,
				handler:        testCfg.HandleSubquests,
			},
			count:   2,
			results: []string{"limit", "offset"},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.treasures.endpoint,
				handler:        testCfg.HandleTreasures,
			},
			count:   9,
			results: []string{"sublocation", "loot_type", "treasure_type", "anima", "airship"},
		},
	}

	testNameList(t, tests, "", "Parameters", nil, compareParameterLists)
}
