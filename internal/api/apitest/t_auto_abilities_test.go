package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
)

func TestGetAutoAbility(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.AutoAbility{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/131",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "auto-ability with provided id '131' doesn't exist. max id: 130",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/50",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters drop":      1,
					"monsters items":     4,
					"shops pre airship":  0,
					"shops post airship": 0,
					"treasures":          0,
					"equipment tables":   2,
				},
			},
			Unique:       apitesttypes.NewExpUnique(50, "gillionaire"),
			MonstersDrop: []int32{302},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:   0,
					Monster: 190,
					Bribe:   5,
				},
				{
					Index:     1,
					Monster:   261,
					StealRare: 1,
				},
				{
					Index:     2,
					Monster:   265,
					StealRare: 1,
				},
				{
					Index:      3,
					Monster:    285,
					DropCommon: 1,
				},
			},
			ShopsPreAirship:  []int32{},
			ShopsPostAirship: []int32{},
			Treasures:        []int32{},
			EquipmentTables:  []int32{7, 30},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/2?rel_availability=post",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters drop":      22,
					"monsters items":     2,
					"shops pre airship":  0,
					"shops post airship": 1,
					"treasures":          1,
					"equipment tables":   2,
				},
			},
			Unique:       apitesttypes.NewExpUnique(2, "piercing"),
			MonstersDrop: []int32{201, 206, 210, 243, 248, 287},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:    0,
					Monster:  247,
					DropRare: 1,
				},
				{
					Index:     1,
					Monster:   285,
					StealRare: 1,
				},
			},
			ShopsPreAirship:  []int32{},
			ShopsPostAirship: []int32{27},
			Treasures:        []int32{31},
			EquipmentTables:  []int32{76, 78},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/2?rel_availability=always",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters drop":      100,
					"monsters items":     3,
					"shops pre airship":  5,
					"shops post airship": 5,
					"treasures":          3,
					"equipment tables":   2,
				},
			},
			Unique:       apitesttypes.NewExpUnique(2, "piercing"),
			MonstersDrop: []int32{16, 25, 41, 54, 78, 139, 163, 183, 192},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:    0,
					Monster:  161,
					DropRare: 1,
				},
				{
					Index:   1,
					Monster: 173,
					Bribe:   30,
				},
				{
					Index:      2,
					Monster:    190,
					DropCommon: 1,
				},
			},
			ShopsPreAirship:  []int32{2, 12, 16, 17, 32},
			ShopsPostAirship: []int32{2, 12, 16, 17, 32},
			Treasures:        []int32{76, 101, 105},
			EquipmentTables:  []int32{76, 78},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/45?rel_availability=always",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters drop":      0,
					"monsters items":     2,
					"shops pre airship":  0,
					"shops post airship": 0,
					"treasures":          0,
					"equipment tables":   2,
				},
			},
			Unique:       apitesttypes.NewExpUnique(45, "triple ap"),
			MonstersDrop: []int32{},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:   0,
					Monster: 142,
					Bribe:   4,
				},
				{
					Index:      1,
					Monster:    291,
					DropCommon: 1,
				},
			},
			ShopsPreAirship:  []int32{},
			ShopsPostAirship: []int32{},
			Treasures:        []int32{},
			EquipmentTables:  []int32{15, 20},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/55?rel_availability=story&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters drop":      3,
					"monsters items":     0,
					"shops pre airship":  2,
					"shops post airship": 0,
					"treasures":          0,
					"equipment tables":   1,
				},
			},
			Unique:           apitesttypes.NewExpUnique(55, "fire ward"),
			MonstersDrop:     []int32{109, 118, 119},
			MonstersItems:    []apitesttypes.TestMonItemAmts{},
			ShopsPreAirship:  []int32{4, 10},
			ShopsPostAirship: []int32{},
			Treasures:        []int32{},
			EquipmentTables:  []int32{137},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/55?rel_availability=post-game&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters drop":      8,
					"monsters items":     2,
					"shops pre airship":  0,
					"shops post airship": 0,
					"treasures":          0,
					"equipment tables":   1,
				},
			},
			Unique:       apitesttypes.NewExpUnique(55, "fire ward"),
			MonstersDrop: []int32{38, 53, 141, 151, 176, 191, 192, 239},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:       0,
					Monster:     38,
					StealCommon: 2,
					StealRare:   3,
				},
				{
					Index:       1,
					Monster:     53,
					StealCommon: 1,
					StealRare:   2,
				},
			},
			ShopsPreAirship:  []int32{},
			ShopsPostAirship: []int32{},
			Treasures:        []int32{},
			EquipmentTables:  []int32{137},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/10",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters drop":      9,
					"monsters items":     9,
					"shops pre airship":  0,
					"shops post airship": 1,
					"treasures":          0,
					"equipment tables":   3,
				},
			},
			Unique:       apitesttypes.NewExpUnique(10, "darkstrike"),
			MonstersDrop: []int32{185, 200, 257, 279, 285},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:     0,
					Monster:   11,
					StealRare: 1,
				},
				{
					Index:       5,
					Monster:     99,
					StealCommon: 1,
					StealRare:   2,
				},
				{
					Index:       8,
					Monster:     271,
					StealCommon: 4,
				},
			},
			ShopsPreAirship:  []int32{},
			ShopsPostAirship: []int32{7},
			Treasures:        []int32{},
			EquipmentTables:  []int32{25, 32, 46},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/89?rel_availability=pre-airship",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters drop":      9,
					"monsters items":     4,
					"shops pre airship":  1,
					"shops post airship": 1,
					"treasures":          2,
					"equipment tables":   2,
				},
			},
			Unique:       apitesttypes.NewExpUnique(89, "sos nulblaze"),
			MonstersDrop: []int32{31, 71, 102, 143, 158, 181},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:   0,
					Monster: 38,
					Bribe:   16,
				},
				{
					Index:   1,
					Monster: 53,
					Bribe:   8,
				},
				{
					Index:   2,
					Monster: 68,
					Other:   2,
				},
				{
					Index:       3,
					Monster:     109,
					StealCommon: 2,
					StealRare:   3,
				},
			},
			ShopsPreAirship:  []int32{14},
			ShopsPostAirship: []int32{14},
			Treasures:        []int32{61, 291},
			EquipmentTables:  []int32{102, 142},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/1?rel_availability=always",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"equipment tables": true,
				},
				ExpLengths: map[string]int{
					"monsters drop":      3,
					"monsters items":     12,
					"shops pre airship":  2,
					"shops post airship": 2,
					"treasures":          4,
				},
			},
			Unique:       apitesttypes.NewExpUnique(1, "sensor"),
			MonstersDrop: []int32{16, 64, 99},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:      0,
					Monster:    39,
					DropCommon: 1,
					DropRare:   1,
				},
				{
					Index:      5,
					Monster:    105,
					DropCommon: 2,
				},
				{
					Index:      11,
					Monster:    189,
					DropCommon: 1,
					DropRare:   1,
				},
			},
			ShopsPreAirship:  []int32{2, 16},
			ShopsPostAirship: []int32{2, 16},
			Treasures:        []int32{39, 54, 104, 174},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/1?rel_availability=always&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"equipment tables": true,
				},
				ExpLengths: map[string]int{
					"monsters drop":      3,
					"monsters items":     12,
					"shops pre airship":  2,
					"shops post airship": 2,
					"treasures":          0,
				},
			},
			Unique:       apitesttypes.NewExpUnique(1, "sensor"),
			MonstersDrop: []int32{16, 64, 99},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:      0,
					Monster:    39,
					DropCommon: 1,
					DropRare:   1,
				},
				{
					Index:      6,
					Monster:    139,
					DropCommon: 1,
					DropRare:   2,
				},
				{
					Index:      9,
					Monster:    177,
					DropCommon: 1,
					DropRare:   2,
				},
			},
			ShopsPreAirship:  []int32{2, 16},
			ShopsPostAirship: []int32{2, 16},
			Treasures:        []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/1?rel_availability=pre-story",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"equipment tables": true,
				},
				ExpLengths: map[string]int{
					"monsters drop":      7,
					"monsters items":     16,
					"shops pre airship":  1,
					"shops post airship": 0,
					"treasures":          1,
				},
			},
			Unique:       apitesttypes.NewExpUnique(1, "sensor"),
			MonstersDrop: []int32{11, 47, 94, 113},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:      0,
					Monster:    6,
					DropCommon: 2,
					DropRare:   2,
				},
				{
					Index:      4,
					Monster:    14,
					DropCommon: 2,
					DropRare:   2,
				},
				{
					Index:      15,
					Monster:    123,
					DropCommon: 1,
					DropRare:   1,
				},
			},
			ShopsPreAirship:  []int32{9},
			ShopsPostAirship: []int32{},
			Treasures:        []int32{84},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/51?rel_availability=always&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"equipment tables": true,
				},
				ExpLengths: map[string]int{
					"monsters drop":      1,
					"monsters items":     8,
					"shops pre airship":  0,
					"shops post airship": 0,
					"treasures":          0,
				},
			},
			Unique:       apitesttypes.NewExpUnique(51, "break damage limit"),
			MonstersDrop: []int32{299},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:    0,
					Monster:  261,
					DropRare: 1,
				},
				{
					Index:    3,
					Monster:  265,
					DropRare: 1,
				},
				{
					Index:      7,
					Monster:    299,
					DropCommon: 1,
				},
			},
			ShopsPreAirship:  []int32{},
			ShopsPostAirship: []int32{},
			Treasures:        []int32{},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities/51?rel_availability=post&rel_repeatable=false",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"equipment tables": true,
				},
				ExpLengths: map[string]int{
					"monsters drop":      8,
					"monsters items":     9,
					"shops pre airship":  0,
					"shops post airship": 0,
					"treasures":          0,
				},
			},
			Unique:       apitesttypes.NewExpUnique(51, "break damage limit"),
			MonstersDrop: []int32{293, 298, 302, 305},
			MonstersItems: []apitesttypes.TestMonItemAmts{
				{
					Index:      0,
					Monster:    293,
					DropCommon: 1,
				},
				{
					Index:      4,
					Monster:    302,
					DropCommon: 1,
				},
				{
					Index:      8,
					Monster:    307,
					DropCommon: 1,
				},
			},
			ShopsPreAirship:  []int32{},
			ShopsPostAirship: []int32{},
			Treasures:        []int32{},
		},
	}

	testSingleResources(t, tests, "GetAutoAbility", testSrv.HandleAutoAbilities)
}

func TestRetrieveAutoAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   130,
			Results: []int32{1, 33, 67, 89, 101, 111, 121, 130},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?category=1&type=armor",
				ExpectedStatus: http.StatusOK,
			},
			Count:   16,
			Results: []int32{103, 106, 110, 112, 114, 118},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?monster=294",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{5, 51, 57, 128, 129},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?monster_items=294",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{51, 113, 117, 121, 129},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?availability=post&customize=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{32, 35, 36, 42, 88, 105, 118},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?availability=post&customize=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   22,
			Results: []int32{3, 38, 42, 50, 84, 98, 110, 121, 127},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?availability=always&monster=85&character=1",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{5, 6, 7, 8, 27, 73, 116},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?availability=pre-story&character=5&area=213", // gagazet - prominence
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{2, 10, 93},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?availability=always&sublocation=26&character=2&methods=treasure", // macalania woods
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{69, 83},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?availability=post&location=4",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{3, 38, 43, 51, 59, 102, 128, 129},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/auto-abilities?availability=post&area=236&repeatable=false", // omega ruins bottom
				ExpectedStatus: http.StatusOK,
			},
			Count:   9,
			Results: []int32{37, 39, 60, 74, 127},
		},
	}

	testIDList(t, tests, apilookups.E.AutoAbilities.Endpoint, "RetrieveAutoAbilities", testSrv.HandleAutoAbilities, compareAPIResourceLists[apiresource.Named])
}
