package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func TestParameters(t *testing.T) {
	t.Parallel()
	tests := []ListNames{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/areas/parameters?ass=max",
				expectedStatus: http.StatusBadRequest,
				endpoint:       apilookups.E.Areas.Endpoint,
				handler:        testSrv.HandleAreas,
				expectedErr:    "parameter 'ass' does not exist for endpoint /areas. use /api/areas/parameters for available parameters",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/areas/parameters?treasures=true",
				expectedStatus: http.StatusBadRequest,
				endpoint:       apilookups.E.Areas.Endpoint,
				handler:        testSrv.HandleAreas,
				expectedErr:    "only the following default parameters are allowed when using /api/areas/parameters: 'limit', 'offset'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/enums/parameters?ass=max",
				expectedStatus: http.StatusBadRequest,
				endpoint:       apilookups.E.Enums.Endpoint,
				handler:        testSrv.HandleEnums,
				expectedErr:    "parameter 'ass' does not exist for endpoint /enums. use /api/enums/parameters for available parameters",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/areas/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Areas.Endpoint,
				handler:        testSrv.HandleAreas,
			},
			count:   28,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnItem, ts.QpnSaveSphere, ts.QpnSublocation, ts.QpnIDs}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/sublocations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Sublocations.Endpoint,
				handler:        testSrv.HandleSublocations,
			},
			count:   23,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLocation, ts.QpnItem, ts.QpnMethods, ts.QpnAeons, ts.QpnShops, ts.QpnFMVs, ts.QpnIDs}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/locations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Locations.Endpoint,
				handler:        testSrv.HandleLocations,
			},
			count:   22,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnKeyItem, ts.QpnCharacters, ts.QpnAeons, ts.QpnTreasures, ts.QpnSidequests, ts.QpnIDs, ts.QpnFlip}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/arena-creations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.ArenaCreations.Endpoint,
				handler:        testSrv.HandleArenaCreations,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnCategory}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/blitzball-prizes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.BlitzballPrizes.Endpoint,
				handler:        testSrv.HandleBlitzballPrizes,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnCategory}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/fmvs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.FMVs.Endpoint,
				handler:        testSrv.HandleFMVs,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnLocation}),
		},

		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Monsters.Endpoint,
				handler:        testSrv.HandleMonsters,
			},
			count:   34,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnKimahriStats, ts.QpnAlteredState, ts.QpnOmnisElements, ts.QpnStatusResists, ts.QpnAutoAbility, ts.QpnArea, ts.QpnDistance, ts.QpnUnderwater, ts.QpnSpecies, ts.QpnIDs, ts.QpnYunaStats}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.MonsterFormations.Endpoint,
				handler:        testSrv.HandleMonsterFormations,
			},
			count:   14,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnMonster, ts.QpnLocation, ts.QpnAmbush, ts.QpnCategory, ts.QpnIDs}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/overdrive-modes/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.OverdriveModes.Endpoint,
				handler:        testSrv.HandleOverdriveModes,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnType}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/shops/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Shops.Endpoint,
				handler:        testSrv.HandleShops,
			},
			count:   13,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLocation, ts.QpnAutoAbility, ts.QpnItems, ts.QpnEquipment, ts.QpnAvailability, ts.QpnIDs}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/songs/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Songs.Endpoint,
				handler:        testSrv.HandleSongs,
			},
			count:   10,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnArea, ts.QpnFMVs, ts.QpnComposer, ts.QpnSpecialUse}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/sidequests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Sidequests.Endpoint,
				handler:        testSrv.HandleSidequests,
			},
			count:   4,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnAvailability}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/subquests/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Subquests.Endpoint,
				handler:        testSrv.HandleSubquests,
			},
			count:   5,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnAvailability, ts.QpnRepeatable}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/treasures/parameters?limit=max",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Treasures.Endpoint,
				handler:        testSrv.HandleTreasures,
			},
			count:   14,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnSublocation, ts.QpnLootType, ts.QpnTreasureType, ts.QpnAnima, ts.QpnAvailability}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/enums/parameters",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Enums.Endpoint,
				handler:        testSrv.HandleEnums,
			},
			count:   2,
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset}),
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/enums/parameters?limit=1",
				expectedStatus: http.StatusOK,
				endpoint:       apilookups.E.Enums.Endpoint,
				handler:        testSrv.HandleEnums,
			},
			count:   2,
			next:    h.GetStrPtr("/enums/parameters?limit=1&offset=1"),
			results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit}),
		},
	}

	testNameList(t, tests, "Parameters", compareParameterLists)
}
