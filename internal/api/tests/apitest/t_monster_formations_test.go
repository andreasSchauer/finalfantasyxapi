package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetMonsterFormation(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.MonsterFormation{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations/332",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "monster-formation with provided id '332' doesn't exist. max id: 331",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations/27",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":         2,
					"areas":            1,
					"trigger commands": 0,
				},
			},
			IDOnly:         apitesttypes.NewExpIDOnly(27),
			Category:       "boss-fight",
			IsForcedAmbush: false,
			CanEscape:      false,
			Ap:             20,
			ApOverkill:     30,
			Gil:            211,
			BossMusic:      h.GetInt32Ptr(16),
			Monsters: map[string]int32{
				"sinspawn echuilles": 1,
				"sinscale - 3":       4,
			},
			Areas:           []int32{42},
			TriggerCommands: []apitesttypes.TestFormationTC{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations/77",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"trigger commands": true,
				},
				ExpLengths: map[string]int{
					"monsters":         1,
					"areas":            3,
					"trigger commands": 0,
				},
			},
			IDOnly:         apitesttypes.NewExpIDOnly(77),
			Category:       "random-encounter",
			IsForcedAmbush: false,
			CanEscape:      true,
			Ap:             170,
			ApOverkill:     255,
			Gil:            140,
			Monsters: map[string]int32{
				"garuda - 3": 1,
			},
			Areas: []int32{99, 100, 106},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations/137",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":         3,
					"areas":            1,
					"trigger commands": 2,
				},
			},
			IDOnly:         apitesttypes.NewExpIDOnly(137),
			Category:       "boss-fight",
			IsForcedAmbush: false,
			CanEscape:      false,
			Ap:             5080,
			ApOverkill:     7620,
			Gil:            8600,
			BossMusic:      h.GetInt32Ptr(55),
			Monsters: map[string]int32{
				"seymour":            1,
				"anima - 1":          1,
				"guado guardian - 1": 2,
			},
			Areas: []int32{164},
			TriggerCommands: []apitesttypes.TestFormationTC{
				{
					Ability: 4,
					Users:   []int32{5},
				},
				{
					Ability: 5,
					Users:   []int32{6, 7},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations/265",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":         1,
					"areas":            3,
					"trigger commands": 0,
				},
			},
			IDOnly:         apitesttypes.NewExpIDOnly(265),
			Category:       "random-encounter",
			IsForcedAmbush: true,
			CanEscape:      true,
			Ap:             21000,
			ApOverkill:     31500,
			Gil:            1900,
			Monsters: map[string]int32{
				"great malboro": 1,
			},
			Areas:           []int32{234, 236, 237},
			TriggerCommands: []apitesttypes.TestFormationTC{},
		},
	}

	testSingleResources(t, tests, "GetMonsterFormation", testSrv.HandleMonsterFormations)
}

func TestRetrieveMonsterFormations(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?limit=asd",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value 'asd' used for parameter 'limit'. usage: '?limit={int|'max'}'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   331,
			Results: []int32{1, 175, 238, 307, 331},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?monster=44",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{63, 64, 67, 68, 69},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?location=12",
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{90, 93, 98, 102, 105},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?sublocation=7", // besaid
				ExpectedStatus: http.StatusOK,
			},
			Count:   14,
			Results: []int32{9, 15, 19, 24, 25},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?area=232", // sin - city of dying dreams
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{257, 258, 262, 266},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?ambush=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{226, 265, 298},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?category=boss-fight&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   52,
			Results: []int32{3, 37, 83, 137, 189, 204, 245, 270, 296},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?availability=post&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   70,
			Results: []int32{7, 20, 70, 171, 213, 266, 287, 314, 331},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?availability=pre-story&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   86,
			Results: []int32{1, 2, 28, 37, 134, 164, 178, 203, 245, 254},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?repeatable=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   237,
			Results: []int32{9, 35, 58, 81, 103, 148, 165, 193, 239, 280, 319, 331},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?category=boss-fight&repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{226, 227},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?availability=always&location=23",
				ExpectedStatus: http.StatusOK,
			},
			Count:   14,
			Results: []int32{230, 233, 234, 237, 240, 244},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?availability=post&location=25",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{256, 260, 265, 266, 269},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?availability=pre-story&sublocation=34", // highbridge
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{174, 190, 191, 192},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?availability=story&area=234", // sin - the nucleus
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{260, 262, 263, 265, 266, 267, 268, 269},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?availability=post&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   70,
			Results: []int32{7, 20, 70, 78, 83, 171, 213, 255, 264, 295, 310, 331},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations?escape=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   82,
			Results: []int32{1, 38, 137, 183, 253, 279, 296},
		},
	}

	testIDList(t, tests, apilookups.E.MonsterFormations.Endpoint, "RetrieveMonsterFormations", testSrv.HandleMonsterFormations, compareAPIResourceLists[apiresource.Unnamed])
}

func TestSubsectionMonsterFormations(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/locations/23/monster-formations",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleLocations,
			},
			Count:          18,
			ParentResource: h.GetStrPtr("/locations/23"),
			Results:        []int32{228, 231, 235, 237, 240, 244, 245},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/locations/4/monster-formations",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleLocations,
			},
			Count:          14,
			ParentResource: h.GetStrPtr("/locations/4"),
			Results:        []int32{9, 13, 15, 18, 21, 23, 25},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/sublocations/38/monster-formations",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          14,
			ParentResource: h.GetStrPtr("/sublocations/38"),
			Results:        []int32{198, 202, 216, 217, 221, 224, 227},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/sublocations/18/monster-formations",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          10,
			ParentResource: h.GetStrPtr("/sublocations/18"),
			Results:        []int32{86, 89, 90, 93, 95},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/areas/89/monster-formations",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          10,
			ParentResource: h.GetStrPtr("/areas/89"),
			Results:        []int32{46, 49, 54, 57, 58, 59},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/areas/170/monster-formations",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          12,
			ParentResource: h.GetStrPtr("/areas/170"),
			Results:        []int32{140, 144, 149, 150, 153, 155},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/200/monster-formations",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsters,
			},
			Count:          1,
			ParentResource: h.GetStrPtr("/monsters/200"),
			Results:        []int32{170},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/155/monster-formations",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsters,
			},
			Count:          2,
			ParentResource: h.GetStrPtr("/monsters/155"),
			Results:        []int32{216, 232},
		},
	}

	testIDList(t, tests, apilookups.E.MonsterFormations.Endpoint, "SubsectionMonsterFormations", nil, compareSimpleResourceLists[apiresource.Unnamed, sectiontypes.MonsterFormationSimple])
}
