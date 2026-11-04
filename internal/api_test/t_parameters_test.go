package apitest

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
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
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLimit, typestrings.QpnOffset, typestrings.QpnItem, typestrings.QpnSaveSphere, typestrings.QpnSublocation, typestrings.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sublocations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Sublocations.Endpoint,
				handler:        testCfg.HandleSublocations,
			},
			count:   23,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLocation, typestrings.QpnItem, typestrings.QpnMethods, typestrings.QpnAeons, typestrings.QpnShops, typestrings.QpnFMVs, typestrings.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/locations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Locations.Endpoint,
				handler:        testCfg.HandleLocations,
			},
			count:   22,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnKeyItem, typestrings.QpnCharacters, typestrings.QpnAeons, typestrings.QpnTreasures, typestrings.QpnSidequests, typestrings.QpnIDs, typestrings.QpnFlip}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/arena-creations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.ArenaCreations.Endpoint,
				handler:        testCfg.HandleArenaCreations,
			},
			count:   4,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLimit, typestrings.QpnOffset, typestrings.QpnCategory}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/blitzball-prizes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.BlitzballPrizes.Endpoint,
				handler:        testCfg.HandleBlitzballPrizes,
			},
			count:   4,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLimit, typestrings.QpnOffset, typestrings.QpnCategory}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/fmvs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.FMVs.Endpoint,
				handler:        testCfg.HandleFMVs,
			},
			count:   4,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLimit, typestrings.QpnOffset, typestrings.QpnLocation}),
		},

		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Monsters.Endpoint,
				handler:        testCfg.HandleMonsters,
			},
			count:   34,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnKimahriStats, typestrings.QpnAlteredState, typestrings.QpnOmnisElements, typestrings.QpnStatusResists, typestrings.QpnAutoAbility, typestrings.QpnArea, typestrings.QpnDistance, typestrings.QpnUnderwater, typestrings.QpnSpecies, typestrings.QpnIDs, typestrings.QpnYunaStats}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.MonsterFormations.Endpoint,
				handler:        testCfg.HandleMonsterFormations,
			},
			count:   14,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnMonster, typestrings.QpnLocation, typestrings.QpnAmbush, typestrings.QpnCategory, typestrings.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.OverdriveModes.Endpoint,
				handler:        testCfg.HandleOverdriveModes,
			},
			count:   4,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLimit, typestrings.QpnOffset, typestrings.QpnType}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/shops/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Shops.Endpoint,
				handler:        testCfg.HandleShops,
			},
			count:   13,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLocation, typestrings.QpnAutoAbility, typestrings.QpnItems, typestrings.QpnEquipment, typestrings.QpnAvailability, typestrings.QpnIDs}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/songs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Songs.Endpoint,
				handler:        testCfg.HandleSongs,
			},
			count:   10,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnArea, typestrings.QpnFMVs, typestrings.QpnComposer, typestrings.QpnSpecialUse}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/sidequests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Sidequests.Endpoint,
				handler:        testCfg.HandleSidequests,
			},
			count:   4,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLimit, typestrings.QpnOffset, typestrings.QpnAvailability}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/subquests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Subquests.Endpoint,
				handler:        testCfg.HandleSubquests,
			},
			count:   5,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLimit, typestrings.QpnOffset, typestrings.QpnAvailability, typestrings.QpnRepeatable}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/treasures/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Treasures.Endpoint,
				handler:        testCfg.HandleTreasures,
			},
			count:   14,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnLootType, typestrings.QpnTreasureType, typestrings.QpnAnima, typestrings.QpnAvailability}),
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/enums/parameters",
				expectedStatus: http.StatusOK,
				endpoint:       testCfg.E.Enums.Endpoint,
				handler:        testCfg.HandleEnums,
			},
			count:   2,
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLimit, typestrings.QpnOffset}),
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
			results: typestrings.QpnsToNamedParams([]typestrings.QueryParamName{typestrings.QpnLimit}),
		},
	}

	testNameList(t, tests, "Parameters", compareParameterLists)
}
