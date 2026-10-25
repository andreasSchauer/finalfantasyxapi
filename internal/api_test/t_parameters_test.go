package apiTest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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
				expectedErr:    "parameter 'ass' does not exist for endpoint /areas. use /api/areas/parameters for available parameters.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/parameters?treasures=true",
				expectedStatus: http.StatusBadRequest,
				endpoint:       testCfg.E.Areas.Endpoint,
				handler:        testCfg.HandleAreas,
				expectedErr:    "only the following default parameters are allowed when using /api/areas/parameters: 'limit', 'offset'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums/parameters?ass=max",
				expectedStatus: http.StatusBadRequest,
				endpoint:       testCfg.E.Enums.Endpoint,
				handler:        testCfg.HandleEnums,
				expectedErr:    "parameter 'ass' does not exist for endpoint /enums. use /api/enums/parameters for available parameters.",
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
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLimit, api.QpnOffset, api.QpnItem, api.QpnSaveSphere, api.QpnSublocation, api.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Sublocations.Endpoint,
				handler:        testCfg.HandleSublocations,
			},
			count:   23,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLocation, api.QpnItem, api.QpnMethods, api.QpnAeons, api.QpnShops, api.QpnFMVs, api.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Locations.Endpoint,
				handler:        testCfg.HandleLocations,
			},
			count:   22,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnKeyItem, api.QpnCharacters, api.QpnAeons, api.QpnTreasures, api.QpnSidequests, api.QpnIDs, api.QpnFlip}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.ArenaCreations.Endpoint,
				handler:        testCfg.HandleArenaCreations,
			},
			count:   4,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLimit, api.QpnOffset, api.QpnCategory}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.BlitzballPrizes.Endpoint,
				handler:        testCfg.HandleBlitzballPrizes,
			},
			count:   4,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLimit, api.QpnOffset, api.QpnCategory}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.FMVs.Endpoint,
				handler:        testCfg.HandleFMVs,
			},
			count:   4,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLimit, api.QpnOffset, api.QpnLocation}),
		},

		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Monsters.Endpoint,
				handler:        testCfg.HandleMonsters,
			},
			count:   34,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnKimahriStats, api.QpnAlteredState, api.QpnOmnisElements, api.QpnStatusResists, api.QpnAutoAbility, api.QpnArea, api.QpnDistance, api.QpnUnderwater, api.QpnSpecies, api.QpnIDs, api.QpnYunaStats}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.MonsterFormations.Endpoint,
				handler:        testCfg.HandleMonsterFormations,
			},
			count:   14,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnMonster, api.QpnLocation, api.QpnAmbush, api.QpnCategory, api.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.OverdriveModes.Endpoint,
				handler:        testCfg.HandleOverdriveModes,
			},
			count:   4,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLimit, api.QpnOffset, api.QpnType}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Shops.Endpoint,
				handler:        testCfg.HandleShops,
			},
			count:   13,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLocation, api.QpnAutoAbility, api.QpnItems, api.QpnEquipment, api.QpnAvailability, api.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Songs.Endpoint,
				handler:        testCfg.HandleSongs,
			},
			count:   10,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnArea, api.QpnFMVs, api.QpnComposer, api.QpnSpecialUse}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Sidequests.Endpoint,
				handler:        testCfg.HandleSidequests,
			},
			count:   4,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLimit, api.QpnOffset, api.QpnAvailability}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Subquests.Endpoint,
				handler:        testCfg.HandleSubquests,
			},
			count:   5,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLimit, api.QpnOffset, api.QpnAvailability, api.QpnRepeatable}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Treasures.Endpoint,
				handler:        testCfg.HandleTreasures,
			},
			count:   14,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnSublocation, api.QpnLootType, api.QpnTreasureType, api.QpnAnima, api.QpnAvailability}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums/parameters",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Enums.Endpoint,
				handler:        testCfg.HandleEnums,
			},
			count:   2,
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLimit, api.QpnOffset}),
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
			results: api.QpnsToNamedParams([]api.QueryParamName{api.QpnLimit}),
		},
	}

	testNameList(t, tests, "", "Parameters", nil, compareParameterLists)
}
