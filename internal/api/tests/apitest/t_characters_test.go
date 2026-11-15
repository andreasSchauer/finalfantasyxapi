package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetCharacter(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Character{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters/9",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "character with provided id '9' doesn't exist. max id: 8",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters/1?osg_stats=standard",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 2,
					"default abilities": 0,
					"std sg abilities":  11,
					"exp sg abilities":  11,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   1,
				Name: "tidus",
			},
			UntypedUnit: 1,
			Area:        1,
			Agility: apitesttypes.TestAgilityParams{
				AgilityTier: 16,
				TickSpeed:   6,
				MinICV:      h.GetInt32Ptr(17),
				MaxICV:      h.GetInt32Ptr(18),
			},
			WeaponType:       "sword",
			CelestialWeapon:  h.GetInt32Ptr(1),
			CelestialFormula: h.GetStrPtr(string(database.CelestialFormulaHpHigh)),
			OverdriveCommand: h.GetInt32Ptr(1),
			CharacterClasses: []int32{1, 5},
			BaseStats: map[string]int32{
				"hp":            3120,
				"mp":            152,
				"strength":      54,
				"defense":       34,
				"magic":         11,
				"magic defense": 7,
				"agility":       44,
				"luck":          19,
				"evasion":       29,
				"accuracy":      29,
			},
			RegenAmounts: []apitesttypes.TestRegenAmount{
				{
					Index:    0,
					Ticks:    0,
					HealedHP: 100,
				},
				{
					Index:    4,
					Ticks:    4,
					HealedHP: 148,
				},
				{
					Index:    5,
					Ticks:    5,
					HealedHP: 160,
				},
				{
					Index:    7,
					Ticks:    15,
					HealedHP: 282,
				},
			},
			DefaultPlayerAbilities: []int32{},
			StdSphereGrid: &apitesttypes.SphereGrid{
				PlayerAbilities: []int32{9, 10, 13, 22, 27, 38, 56, 59},
			},
			ExpSphereGrid: &apitesttypes.SphereGrid{
				PlayerAbilities: []int32{10, 22, 25, 57, 58, 59},
			},
			OverdriveModes: map[string]int32{
				"warrior":   150,
				"victim":    120,
				"coward":    600,
				"daredevil": 170,
				"loner":     60,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters/5",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"character classes": 2,
					"default abilities": 1,
					"std sg abilities":  5,
					"exp sg abilities":  4,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   5,
				Name: "kimahri",
			},
			UntypedUnit: 5,
			Area:        42,
			Agility: apitesttypes.TestAgilityParams{
				AgilityTier: 6,
				TickSpeed:   16,
				MinICV:      h.GetInt32Ptr(46),
				MaxICV:      h.GetInt32Ptr(48),
			},
			WeaponType:       "spear",
			CelestialWeapon:  h.GetInt32Ptr(5),
			CelestialFormula: h.GetStrPtr(string(database.CelestialFormulaHpHigh)),
			OverdriveCommand: h.GetInt32Ptr(5),
			CharacterClasses: []int32{1, 9},
			BaseStats: map[string]int32{
				"hp":            644,
				"mp":            78,
				"strength":      16,
				"defense":       15,
				"magic":         17,
				"magic defense": 5,
				"agility":       6,
				"luck":          18,
				"evasion":       5,
				"accuracy":      5,
			},
			RegenAmounts: []apitesttypes.TestRegenAmount{
				{
					Index:    0,
					Ticks:    0,
					HealedHP: 100,
				},
				{
					Index:    1,
					Ticks:    1,
					HealedHP: 102,
				},
				{
					Index:    3,
					Ticks:    3,
					HealedHP: 107,
				},
				{
					Index:    9,
					Ticks:    30,
					HealedHP: 175,
				},
			},
			DefaultPlayerAbilities: []int32{33},
			StdSphereGrid: &apitesttypes.SphereGrid{
				PlayerAbilities: []int32{12, 32, 33, 52, 87},
			},
			ExpSphereGrid: &apitesttypes.SphereGrid{
				PlayerAbilities: []int32{12, 32, 33, 52},
			},
			OverdriveModes: map[string]int32{
				"comrade":   100,
				"healer":    100,
				"tactician": 60,
				"rook":      120,
				"daredevil": 200,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters/4?osg_stats=expert",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts": true,
				},
				ExpLengths: map[string]int{
					"character classes": 2,
					"default abilities": 4,
					"std sg abilities":  20,
					"exp sg abilities":  19,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   4,
				Name: "lulu",
			},
			UntypedUnit: 4,
			Area:        36,
			Agility: apitesttypes.TestAgilityParams{
				AgilityTier: 12,
				TickSpeed:   10,
				MinICV:      h.GetInt32Ptr(26),
				MaxICV:      h.GetInt32Ptr(30),
			},
			WeaponType:       "doll",
			CelestialWeapon:  h.GetInt32Ptr(4),
			CelestialFormula: h.GetStrPtr(string(database.CelestialFormulaMpHigh)),
			OverdriveCommand: h.GetInt32Ptr(4),
			CharacterClasses: []int32{1, 8},
			BaseStats: map[string]int32{
				"hp":            2580,
				"mp":            332,
				"strength":      15,
				"defense":       38,
				"magic":         54,
				"magic defense": 78,
				"agility":       22,
				"luck":          17,
				"evasion":       83,
				"accuracy":      13,
			},
			DefaultPlayerAbilities: []int32{69, 70, 71, 72},
			StdSphereGrid: &apitesttypes.SphereGrid{
				PlayerAbilities: []int32{21, 29, 43, 69, 74, 78, 83, 86},
			},
			ExpSphereGrid: &apitesttypes.SphereGrid{
				PlayerAbilities: []int32{29, 30, 43, 70, 74, 75, 79, 81, 87},
			},
			OverdriveModes: map[string]int32{
				"comrade":   100,
				"victim":    130,
				"tactician": 75,
				"slayer":    130,
				"hero":      70,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters/8",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts": true,
				},
				ExpLengths: map[string]int{
					"character classes": 1,
					"default abilities": 15,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   8,
				Name: "seymour",
			},
			UntypedUnit: 8,
			Area:        102,
			Agility: apitesttypes.TestAgilityParams{
				AgilityTier: 12,
				TickSpeed:   10,
				MinICV:      h.GetInt32Ptr(28),
				MaxICV:      h.GetInt32Ptr(30),
			},
			WeaponType:       "seymour-staff",
			CelestialWeapon:  nil,
			CelestialFormula: nil,
			OverdriveCommand: nil,
			CharacterClasses: []int32{12},
			BaseStats: map[string]int32{
				"hp":            1200,
				"mp":            999,
				"strength":      20,
				"defense":       25,
				"magic":         35,
				"magic defense": 100,
				"agility":       20,
				"luck":          18,
				"evasion":       10,
				"accuracy":      10,
			},
			DefaultPlayerAbilities: []int32{45, 46, 48, 69, 71, 74, 75, 76},
			StdSphereGrid:          nil,
			ExpSphereGrid:          nil,
			OverdriveModes:         nil,
		},
	}

	testSingleResources(t, tests, "GetCharacter", testSrv.HandleCharacters)
}

func TestRetrieveCharacters(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters",
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters?story_based=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{8},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/characters?underwater=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{1, 3, 7},
		},
	}

	testIDList(t, tests, apilookups.E.Characters.Endpoint, "RetrieveCharacters", testSrv.HandleCharacters, compareAPIResourceLists[apiresource.Named])
}
