package api

import (
	"net/http"
	"testing"
)

func TestSections(t *testing.T) {
	tests := []expListNames{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.areas.endpoint,
				handler:        testCfg.HandleAreas,
			},
			count: 6,
			results: []string{
				"simple",
				"connected",
				"monsters",
				"monster-formations",
				"songs",
				"treasures",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.sublocations.endpoint,
				handler:        testCfg.HandleSublocations,
			},
			count: 8,
			results: []string{
				"simple",
				"areas",
				"connected",
				"monsters",
				"monster-formations",
				"shops",
				"songs",
				"treasures",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.locations.endpoint,
				handler:        testCfg.HandleLocations,
			},
			count: 9,
			results: []string{
				"simple",
				"areas",
				"connected",
				"monsters",
				"monster-formations",
				"shops",
				"songs",
				"sublocations",
				"treasures",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.arenaCreations.endpoint,
				handler:        testCfg.HandleArenaCreations,
			},
			count: 0,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.blitzballPrizes.endpoint,
				handler:        testCfg.HandleBlitzballPrizes,
			},
			count: 0,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.fmvs.endpoint,
				handler:        testCfg.HandleFMVs,
			},
			count: 0,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.monsters.endpoint,
				handler:        testCfg.HandleMonsters,
			},
			count: 3,
			results: []string{
				"simple",
				"areas",
				"monster-formations",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.monsterFormations.endpoint,
				handler:        testCfg.HandleMonsterFormations,
			},
			count: 2,
			results: []string{
				"simple",
				"monsters",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.overdriveModes.endpoint,
				handler:        testCfg.HandleOverdriveModes,
			},
			count: 0,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.shops.endpoint,
				handler:        testCfg.HandleShops,
			},
			count: 1,
			results: []string{
				"simple",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.songs.endpoint,
				handler:        testCfg.HandleSongs,
			},
			count: 0,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.sidequests.endpoint,
				handler:        testCfg.HandleSidequests,
			},
			count: 1,
			results: []string{
				"subquests",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.subquests.endpoint,
				handler:        testCfg.HandleSubquests,
			},
			count: 0,
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/sections",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.treasures.endpoint,
				handler:        testCfg.HandleTreasures,
			},
			count: 0,
		},
	}

	testNameList(t, tests, "", "Sections", nil, compareSectionLists)
}
