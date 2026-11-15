package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func TestParameters(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListNames{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/areas/parameters?ass=max",
				ExpectedStatus: http.StatusBadRequest,
				Endpoint:       apilookups.E.Areas.Endpoint,
				Handler:        testSrv.HandleAreas,
				ExpectedErr:    "parameter 'ass' does not exist for endpoint /areas. use /api/areas/parameters for available parameters",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/areas/parameters?treasures=true",
				ExpectedStatus: http.StatusBadRequest,
				Endpoint:       apilookups.E.Areas.Endpoint,
				Handler:        testSrv.HandleAreas,
				ExpectedErr:    "only the following default parameters are allowed when using /api/areas/parameters: 'limit', 'offset'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enums/parameters?ass=max",
				ExpectedStatus: http.StatusBadRequest,
				Endpoint:       apilookups.E.Enums.Endpoint,
				Handler:        testSrv.HandleEnums,
				ExpectedErr:    "parameter 'ass' does not exist for endpoint /enums. use /api/enums/parameters for available parameters",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/areas/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Areas.Endpoint,
				Handler:        testSrv.HandleAreas,
			},
			Count:   28,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnItem, ts.QpnSaveSphere, ts.QpnSublocation, ts.QpnIDs}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/sublocations/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Sublocations.Endpoint,
				Handler:        testSrv.HandleSublocations,
			},
			Count:   23,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLocation, ts.QpnItem, ts.QpnMethods, ts.QpnAeons, ts.QpnShops, ts.QpnFMVs, ts.QpnIDs}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/locations/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Locations.Endpoint,
				Handler:        testSrv.HandleLocations,
			},
			Count:   22,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnKeyItem, ts.QpnCharacters, ts.QpnAeons, ts.QpnTreasures, ts.QpnSidequests, ts.QpnIDs, ts.QpnFlip}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/arena-creations/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.ArenaCreations.Endpoint,
				Handler:        testSrv.HandleArenaCreations,
			},
			Count:   4,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnCategory}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/blitzball-prizes/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.BlitzballPrizes.Endpoint,
				Handler:        testSrv.HandleBlitzballPrizes,
			},
			Count:   4,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnCategory}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/fmvs/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.FMVs.Endpoint,
				Handler:        testSrv.HandleFMVs,
			},
			Count:   4,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnLocation}),
		},

		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Monsters.Endpoint,
				Handler:        testSrv.HandleMonsters,
			},
			Count:   34,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnKimahriStats, ts.QpnAlteredState, ts.QpnOmnisElements, ts.QpnStatusResists, ts.QpnAutoAbility, ts.QpnArea, ts.QpnDistance, ts.QpnUnderwater, ts.QpnSpecies, ts.QpnIDs, ts.QpnYunaStats}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.MonsterFormations.Endpoint,
				Handler:        testSrv.HandleMonsterFormations,
			},
			Count:   14,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnMonster, ts.QpnLocation, ts.QpnAmbush, ts.QpnCategory, ts.QpnIDs}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/overdrive-modes/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.OverdriveModes.Endpoint,
				Handler:        testSrv.HandleOverdriveModes,
			},
			Count:   4,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnType}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/shops/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Shops.Endpoint,
				Handler:        testSrv.HandleShops,
			},
			Count:   13,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLocation, ts.QpnAutoAbility, ts.QpnItems, ts.QpnEquipment, ts.QpnAvailability, ts.QpnIDs}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/songs/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Songs.Endpoint,
				Handler:        testSrv.HandleSongs,
			},
			Count:   10,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnArea, ts.QpnFMVs, ts.QpnComposer, ts.QpnSpecialUse}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/sidequests/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Sidequests.Endpoint,
				Handler:        testSrv.HandleSidequests,
			},
			Count:   4,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnAvailability}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Subquests.Endpoint,
				Handler:        testSrv.HandleSubquests,
			},
			Count:   5,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset, ts.QpnAvailability, ts.QpnRepeatable}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/parameters?limit=max",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Treasures.Endpoint,
				Handler:        testSrv.HandleTreasures,
			},
			Count:   14,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnSublocation, ts.QpnLootType, ts.QpnTreasureType, ts.QpnAnima, ts.QpnAvailability}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enums/parameters",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Enums.Endpoint,
				Handler:        testSrv.HandleEnums,
			},
			Count:   2,
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit, ts.QpnOffset}),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enums/parameters?limit=1",
				ExpectedStatus: http.StatusOK,
				Endpoint:       apilookups.E.Enums.Endpoint,
				Handler:        testSrv.HandleEnums,
			},
			Count:   2,
			Next:    h.GetStrPtr("/enums/parameters?limit=1&offset=1"),
			Results: ts.QpnsToNamedParams([]ts.QueryParamName{ts.QpnLimit}),
		},
	}

	testNameList(t, tests, "Parameters", compareParameterLists)
}
