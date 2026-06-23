package api

import (
	"net/http"
	"testing"
)

func TestParameters(t *testing.T) {
	t.Parallel()
	tests := []expListNames{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.areas.endpoint,
				handler:        testCfg.HandleAreas,
			},
			count:   28,
			results: qpnsToNamedParams([]QueryParamName{qpnLimit, qpnOffset, qpnItem, qpnSaveSphere, qpnSublocation, qpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.sublocations.endpoint,
				handler:        testCfg.HandleSublocations,
			},
			count:   23,
			results: qpnsToNamedParams([]QueryParamName{qpnLocation, qpnItem, qpnMethods, qpnAeons, qpnShops, qpnFMVs, qpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.locations.endpoint,
				handler:        testCfg.HandleLocations,
			},
			count:   22,
			results: qpnsToNamedParams([]QueryParamName{qpnKeyItem, qpnCharacters, qpnAeons, qpnTreasures, qpnSidequests, qpnIDs, qpnFlip}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.arenaCreations.endpoint,
				handler:        testCfg.HandleArenaCreations,
			},
			count:   4,
			results: qpnsToNamedParams([]QueryParamName{qpnLimit, qpnOffset, qpnCategory}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.blitzballPrizes.endpoint,
				handler:        testCfg.HandleBlitzballPrizes,
			},
			count:   4,
			results: qpnsToNamedParams([]QueryParamName{qpnLimit, qpnOffset, qpnCategory}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.fmvs.endpoint,
				handler:        testCfg.HandleFMVs,
			},
			count:   4,
			results: qpnsToNamedParams([]QueryParamName{qpnLimit, qpnOffset, qpnLocation}),
		},

		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.monsters.endpoint,
				handler:        testCfg.HandleMonsters,
			},
			count:   33,
			results: qpnsToNamedParams([]QueryParamName{qpnKimahriStats, qpnAeonStats, qpnAlteredState, qpnOmnisElements, qpnStatusResists, qpnAutoAbility, qpnArea, qpnDistance, qpnUnderwater, qpnSpecies, qpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.monsterFormations.endpoint,
				handler:        testCfg.HandleMonsterFormations,
			},
			count:   13,
			results: qpnsToNamedParams([]QueryParamName{qpnMonster, qpnLocation, qpnAmbush, qpnCategory, qpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.overdriveModes.endpoint,
				handler:        testCfg.HandleOverdriveModes,
			},
			count:   4,
			results: qpnsToNamedParams([]QueryParamName{qpnLimit, qpnOffset, qpnType}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.shops.endpoint,
				handler:        testCfg.HandleShops,
			},
			count:   13,
			results: qpnsToNamedParams([]QueryParamName{qpnLocation, qpnAutoAbility, qpnItems, qpnEquipment, qpnAvailability, qpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.songs.endpoint,
				handler:        testCfg.HandleSongs,
			},
			count:   10,
			results: qpnsToNamedParams([]QueryParamName{qpnArea, qpnFMVs, qpnComposer, qpnSpecialUse}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.sidequests.endpoint,
				handler:        testCfg.HandleSidequests,
			},
			count:   4,
			results: qpnsToNamedParams([]QueryParamName{qpnLimit, qpnOffset, qpnAvailability}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.subquests.endpoint,
				handler:        testCfg.HandleSubquests,
			},
			count:   5,
			results: qpnsToNamedParams([]QueryParamName{qpnLimit, qpnOffset, qpnAvailability, qpnRepeatable}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.e.treasures.endpoint,
				handler:        testCfg.HandleTreasures,
			},
			count:   14,
			results: qpnsToNamedParams([]QueryParamName{qpnSublocation, qpnLootType, qpnTreasureType, qpnAnima, qpnAvailability}),
		},
	}

	testNameList(t, tests, "", "Parameters", nil, compareParameterLists)
}
