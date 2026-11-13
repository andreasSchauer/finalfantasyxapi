package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetMonsterFormation(t *testing.T) {
	t.Parallel()
	tests := []MonsterFormation{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations/332",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "monster-formation with provided id '332' doesn't exist. max id: 331",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations/27",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         2,
					"areas":            1,
					"trigger commands": 0,
				},
			},
			IDOnly:      newExpIDOnly(27),
			category:       "boss-fight",
			isForcedAmbush: false,
			canEscape:      false,
			ap:             20,
			apOverkill:     30,
			gil:            211,
			bossMusic:      h.GetInt32Ptr(16),
			monsters: map[string]int32{
				"sinspawn echuilles": 1,
				"sinscale - 3":       4,
			},
			areas:           []int32{42},
			triggerCommands: []TestFormationTC{},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations/77",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"trigger commands": true,
				},
				expLengths: map[string]int{
					"monsters":         1,
					"areas":            3,
					"trigger commands": 0,
				},
			},
			IDOnly:      newExpIDOnly(77),
			category:       "random-encounter",
			isForcedAmbush: false,
			canEscape:      true,
			ap:             170,
			apOverkill:     255,
			gil:            140,
			monsters: map[string]int32{
				"garuda - 3": 1,
			},
			areas: []int32{99, 100, 106},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations/137",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         3,
					"areas":            1,
					"trigger commands": 2,
				},
			},
			IDOnly:      newExpIDOnly(137),
			category:       "boss-fight",
			isForcedAmbush: false,
			canEscape:      false,
			ap:             5080,
			apOverkill:     7620,
			gil:            8600,
			bossMusic:      h.GetInt32Ptr(55),
			monsters: map[string]int32{
				"seymour":            1,
				"anima - 1":          1,
				"guado guardian - 1": 2,
			},
			areas: []int32{164},
			triggerCommands: []TestFormationTC{
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
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations/265",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"monsters":         1,
					"areas":            3,
					"trigger commands": 0,
				},
			},
			IDOnly:      newExpIDOnly(265),
			category:       "random-encounter",
			isForcedAmbush: true,
			canEscape:      true,
			ap:             21000,
			apOverkill:     31500,
			gil:            1900,
			monsters: map[string]int32{
				"great malboro": 1,
			},
			areas:           []int32{234, 236, 237},
			triggerCommands: []TestFormationTC{},
		},
	}

	testSingleResources(t, tests, "GetMonsterFormation", testSrv.HandleMonsterFormations, compareMonsterFormations)
}

func TestRetrieveMonsterFormations(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?limit=asd",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'asd' used for parameter 'limit'. usage: '?limit={int|'max'}'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   331,
			results: []int32{1, 175, 238, 307, 331},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?monster=44",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{63, 64, 67, 68, 69},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?location=12",
				expectedStatus: http.StatusOK,
			},
			count:   12,
			results: []int32{90, 93, 98, 102, 105},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?sublocation=7", // besaid
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{9, 15, 19, 24, 25},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?area=232", // sin - city of dying dreams
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{257, 258, 262, 266},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?ambush=true",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{226, 265, 298},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?category=boss-fight&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   52,
			results: []int32{3, 37, 83, 137, 189, 204, 245, 270, 296},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?availability=post&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   70,
			results: []int32{7, 20, 70, 171, 213, 266, 287, 314, 331},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?availability=pre-story&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   86,
			results: []int32{1, 2, 28, 37, 134, 164, 178, 203, 245, 254},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?repeatable=true&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   237,
			results: []int32{9, 35, 58, 81, 103, 148, 165, 193, 239, 280, 319, 331},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?category=boss-fight&repeatable=true",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{226, 227},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?availability=always&location=23",
				expectedStatus: http.StatusOK,
			},
			count:   14,
			results: []int32{230, 233, 234, 237, 240, 244},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?availability=post&location=25",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{256, 260, 265, 266, 269},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?availability=pre-story&sublocation=34", // highbridge
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{174, 190, 191, 192},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?availability=story&area=234", // sin - the nucleus
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{260, 262, 263, 265, 266, 267, 268, 269},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?availability=post&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   70,
			results: []int32{7, 20, 70, 78, 83, 171, 213, 255, 264, 295, 310, 331},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations?escape=false&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   82,
			results: []int32{1, 38, 137, 183, 253, 279, 296},
		},
	}

	testIDList(t, tests, apilookups.E.MonsterFormations.Endpoint, "RetrieveMonsterFormations", testSrv.HandleMonsterFormations, compareAPIResourceLists[apiresource.Unnamed])
}

func TestSubsectionMonsterFormations(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/locations/23/monster-formations",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleLocations,
			},
			count:          18,
			parentResource: h.GetStrPtr("/locations/23"),
			results:        []int32{228, 231, 235, 237, 240, 244, 245},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/locations/4/monster-formations",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleLocations,
			},
			count:          14,
			parentResource: h.GetStrPtr("/locations/4"),
			results:        []int32{9, 13, 15, 18, 21, 23, 25},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/sublocations/38/monster-formations",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleSublocations,
			},
			count:          14,
			parentResource: h.GetStrPtr("/sublocations/38"),
			results:        []int32{198, 202, 216, 217, 221, 224, 227},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/sublocations/18/monster-formations",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleSublocations,
			},
			count:          10,
			parentResource: h.GetStrPtr("/sublocations/18"),
			results:        []int32{86, 89, 90, 93, 95},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/areas/89/monster-formations",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleAreas,
			},
			count:          10,
			parentResource: h.GetStrPtr("/areas/89"),
			results:        []int32{46, 49, 54, 57, 58, 59},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/areas/170/monster-formations",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleAreas,
			},
			count:          12,
			parentResource: h.GetStrPtr("/areas/170"),
			results:        []int32{140, 144, 149, 150, 153, 155},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/200/monster-formations",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleMonsters,
			},
			count:          1,
			parentResource: h.GetStrPtr("/monsters/200"),
			results:        []int32{170},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/155/monster-formations",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleMonsters,
			},
			count:          2,
			parentResource: h.GetStrPtr("/monsters/155"),
			results:        []int32{216, 232},
		},
	}

	testIDList(t, tests, apilookups.E.MonsterFormations.Endpoint, "SubsectionMonsterFormations", nil, compareSimpleResourceLists[apiresource.Unnamed, sectiontypes.MonsterFormationSimple])
}
