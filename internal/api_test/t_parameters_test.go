package apitest

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func TestParameters(t *testing.T) {
	t.Parallel()
	tests := []expListNames{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/parameters?ass=max",
				expectedStatus: http.StatusBadRequest,
				endpoint:       testCfg.E.Areas.Endpoint,
				handler:        testCfg.HandleAreas,
				expectedErr:    "parameter 'ass' does not exist for endpoint /areas. use /api/areas/parameters for available parameters",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/parameters?treasures=true",
				expectedStatus: http.StatusBadRequest,
				endpoint:       testCfg.E.Areas.Endpoint,
				handler:        testCfg.HandleAreas,
				expectedErr:    "only the following default parameters are allowed when using /api/areas/parameters: 'limit', 'offset'",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums/parameters?ass=max",
				expectedStatus: http.StatusBadRequest,
				endpoint:       testCfg.E.Enums.Endpoint,
				handler:        testCfg.HandleEnums,
				expectedErr:    "parameter 'ass' does not exist for endpoint /enums. use /api/enums/parameters for available parameters",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Areas.Endpoint,
				handler:        testCfg.HandleAreas,
			},
			count:   28,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnItem, ts.QpnSaveSphere, ts.QpnSublocation, ts.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Sublocations.Endpoint,
				handler:        testCfg.HandleSublocations,
			},
			count:   23,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLocation, ts.QpnItem, ts.QpnMethods, ts.QpnAeons, ts.QpnShops, ts.QpnFMVs, ts.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Locations.Endpoint,
				handler:        testCfg.HandleLocations,
			},
			count:   22,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnKeyItem, ts.QpnCharacters, ts.QpnAeons, ts.QpnTreasures, ts.QpnSidequests, ts.QpnIDs, ts.QpnFlip}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.ArenaCreations.Endpoint,
				handler:        testCfg.HandleArenaCreations,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnCategory}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.BlitzballPrizes.Endpoint,
				handler:        testCfg.HandleBlitzballPrizes,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnCategory}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.FMVs.Endpoint,
				handler:        testCfg.HandleFMVs,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnLocation}),
		},

		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Monsters.Endpoint,
				handler:        testCfg.HandleMonsters,
			},
			count:   34,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnKimahriStats, ts.QpnAlteredState, ts.QpnOmnisElements, ts.QpnStatusResists, ts.QpnAutoAbility, ts.QpnArea, ts.QpnDistance, ts.QpnUnderwater, ts.QpnSpecies, ts.QpnIDs, ts.QpnYunaStats}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.MonsterFormations.Endpoint,
				handler:        testCfg.HandleMonsterFormations,
			},
			count:   14,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnMonster, ts.QpnLocation, ts.QpnAmbush, ts.QpnCategory, ts.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.OverdriveModes.Endpoint,
				handler:        testCfg.HandleOverdriveModes,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnType}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Shops.Endpoint,
				handler:        testCfg.HandleShops,
			},
			count:   13,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLocation, ts.QpnAutoAbility, ts.QpnItems, ts.QpnEquipment, ts.QpnAvailability, ts.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Songs.Endpoint,
				handler:        testCfg.HandleSongs,
			},
			count:   10,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnArea, ts.QpnFMVs, ts.QpnComposer, ts.QpnSpecialUse}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Sidequests.Endpoint,
				handler:        testCfg.HandleSidequests,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnAvailability}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Subquests.Endpoint,
				handler:        testCfg.HandleSubquests,
			},
			count:   5,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnAvailability, ts.QpnRepeatable}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Treasures.Endpoint,
				handler:        testCfg.HandleTreasures,
			},
			count:   14,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnSublocation, ts.QpnLootType, ts.QpnTreasureType, ts.QpnAnima, ts.QpnAvailability}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums/parameters",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Enums.Endpoint,
				handler:        testCfg.HandleEnums,
			},
			count:   2,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums/parameters?limit=1",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Enums.Endpoint,
				handler:        testCfg.HandleEnums,
			},
			count:   2,
			next:    h.GetStrPtr("/enums/parameters?limit=1&offset=1"),
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit}),
		},
	}

	testNameList(t, tests, "Parameters", compareParameterLists)
}
