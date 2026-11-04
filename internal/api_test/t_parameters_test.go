package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
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
				endpoint:       api.E.Areas.Endpoint,
				handler:        testSrv.Cfg.HandleAreas,
				expectedErr:    "parameter 'ass' does not exist for endpoint /areas. use /api/areas/parameters for available parameters",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/parameters?treasures=true",
				expectedStatus: http.StatusBadRequest,
				endpoint:       api.E.Areas.Endpoint,
				handler:        testSrv.Cfg.HandleAreas,
				expectedErr:    "only the following default parameters are allowed when using /api/areas/parameters: 'limit', 'offset'",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums/parameters?ass=max",
				expectedStatus: http.StatusBadRequest,
				endpoint:       api.E.Enums.Endpoint,
				handler:        testSrv.Cfg.HandleEnums,
				expectedErr:    "parameter 'ass' does not exist for endpoint /enums. use /api/enums/parameters for available parameters",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Areas.Endpoint,
				handler:        testSrv.Cfg.HandleAreas,
			},
			count:   28,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnItem, ts.QpnSaveSphere, ts.QpnSublocation, ts.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Sublocations.Endpoint,
				handler:        testSrv.Cfg.HandleSublocations,
			},
			count:   23,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLocation, ts.QpnItem, ts.QpnMethods, ts.QpnAeons, ts.QpnShops, ts.QpnFMVs, ts.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Locations.Endpoint,
				handler:        testSrv.Cfg.HandleLocations,
			},
			count:   22,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnKeyItem, ts.QpnCharacters, ts.QpnAeons, ts.QpnTreasures, ts.QpnSidequests, ts.QpnIDs, ts.QpnFlip}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.ArenaCreations.Endpoint,
				handler:        testSrv.Cfg.HandleArenaCreations,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnCategory}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.BlitzballPrizes.Endpoint,
				handler:        testSrv.Cfg.HandleBlitzballPrizes,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnCategory}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.FMVs.Endpoint,
				handler:        testSrv.Cfg.HandleFMVs,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnLocation}),
		},

		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Monsters.Endpoint,
				handler:        testSrv.Cfg.HandleMonsters,
			},
			count:   34,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnKimahriStats, ts.QpnAlteredState, ts.QpnOmnisElements, ts.QpnStatusResists, ts.QpnAutoAbility, ts.QpnArea, ts.QpnDistance, ts.QpnUnderwater, ts.QpnSpecies, ts.QpnIDs, ts.QpnYunaStats}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.MonsterFormations.Endpoint,
				handler:        testSrv.Cfg.HandleMonsterFormations,
			},
			count:   14,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnMonster, ts.QpnLocation, ts.QpnAmbush, ts.QpnCategory, ts.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.OverdriveModes.Endpoint,
				handler:        testSrv.Cfg.HandleOverdriveModes,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnType}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Shops.Endpoint,
				handler:        testSrv.Cfg.HandleShops,
			},
			count:   13,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLocation, ts.QpnAutoAbility, ts.QpnItems, ts.QpnEquipment, ts.QpnAvailability, ts.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Songs.Endpoint,
				handler:        testSrv.Cfg.HandleSongs,
			},
			count:   10,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnArea, ts.QpnFMVs, ts.QpnComposer, ts.QpnSpecialUse}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Sidequests.Endpoint,
				handler:        testSrv.Cfg.HandleSidequests,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnAvailability}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Subquests.Endpoint,
				handler:        testSrv.Cfg.HandleSubquests,
			},
			count:   5,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnAvailability, ts.QpnRepeatable}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Treasures.Endpoint,
				handler:        testSrv.Cfg.HandleTreasures,
			},
			count:   14,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnSublocation, ts.QpnLootType, ts.QpnTreasureType, ts.QpnAnima, ts.QpnAvailability}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums/parameters",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Enums.Endpoint,
				handler:        testSrv.Cfg.HandleEnums,
			},
			count:   2,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums/parameters?limit=1",
				expectedStatus: http.StatusOK,
				endpoint:       api.E.Enums.Endpoint,
				handler:        testSrv.Cfg.HandleEnums,
			},
			count:   2,
			next:    h.GetStrPtr("/enums/parameters?limit=1&offset=1"),
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit}),
		},
	}

	testNameList(t, tests, "Parameters", compareParameterLists)
}
