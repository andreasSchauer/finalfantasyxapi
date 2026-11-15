package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetAeon(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Aeon{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons/11",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "aeon with provided id '11' doesn't exist. max id: 10",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons/3",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts": true,
				},
				ExpLengths: map[string]int{
					"character classes": 3,
					"aeon commands":     0,
					"default abilities": 5,
					"overdrives":        1,
					"weapon abilities":  3,
					"armor abilities":   4,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   3,
				Name: "ixion",
			},
			UntypedUnit:    11,
			Area:           114,
			BattlesToRegen: 20,
			Agility: apitesttypes.TestAgilityParams{
				AgilityTier: 7,
				TickSpeed:   15,
				MinICV:      h.GetInt32Ptr(43),
				MaxICV:      h.GetInt32Ptr(45),
			},
			CelestialWeapon:  h.GetInt32Ptr(5),
			CharacterClasses: []int32{2, 3, 15},
			BaseStats: map[string]int32{
				"hp":            891,
				"mp":            25,
				"strength":      20,
				"defense":       26,
				"magic":         20,
				"magic defense": 29,
				"agility":       8,
				"luck":          17,
				"evasion":       11,
				"accuracy":      12,
			},
			AeonCommands:           []int32{},
			DefaultPlayerAbilities: []int32{49, 50, 70, 74, 90},
			Overdrives:             []int32{120},
			WeaponAbilities: []apitesttypes.AeonEquipment{
				{
					AutoAbility:     54,
					CelestialWeapon: false,
				},
				{
					AutoAbility:     2,
					CelestialWeapon: false,
				},
				{
					AutoAbility:     51,
					CelestialWeapon: true,
				},
			},
			ArmorAbilities: []apitesttypes.AeonEquipment{
				{
					AutoAbility:     1,
					CelestialWeapon: false,
				},
				{
					AutoAbility:     128,
					CelestialWeapon: false,
				},
				{
					AutoAbility:     127,
					CelestialWeapon: false,
				},
				{
					AutoAbility:     60,
					CelestialWeapon: false,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons/9?battles=347&yuna_stats=strength:24",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"weapon abilities": true,
					"armor abilities":  true,
				},
				ExpLengths: map[string]int{
					"character classes": 3,
					"aeon commands":     5,
					"default abilities": 10,
					"overdrives":        1,
					"weapon abilities":  3,
					"armor abilities":   3,
				},
			},
			UntypedUnit: 17,
			Unique: apitesttypes.Unique{
				ID:   9,
				Name: "sandy",
			},
			Area:           208,
			BattlesToRegen: 30,
			Agility: apitesttypes.TestAgilityParams{
				AgilityTier: 12,
				TickSpeed:   10,
				MinICV:      h.GetInt32Ptr(26),
				MaxICV:      h.GetInt32Ptr(30),
			},
			CelestialWeapon:  nil,
			CharacterClasses: []int32{2, 4, 21},
			BaseStats: map[string]int32{
				"hp":            4590,
				"mp":            71,
				"strength":      149,
				"defense":       53,
				"magic":         53,
				"magic defense": 57,
				"agility":       22,
				"luck":          17,
				"evasion":       27,
				"accuracy":      18,
			},
			RegenAmounts: []apitesttypes.TestRegenAmount{
				{
					Index:    0,
					Ticks:    0,
					HealedHP: 100,
				},
				{
					Index:    3,
					Ticks:    3,
					HealedHP: 153,
				},
				{
					Index:    7,
					Ticks:    15,
					HealedHP: 368,
				},
				{
					Index:    8,
					Ticks:    20,
					HealedHP: 458,
				},
			},
			AeonCommands:           []int32{2, 3, 4, 7, 9},
			DefaultPlayerAbilities: []int32{45, 46, 47, 56, 60, 61, 62, 68, 100, 102},
			Overdrives:             []int32{124},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons/1?yuna_stats=hp:2000,mp:250,strength:10,defense:28,magic:72,magic_defense:50,agility:41,luck:25,evasion:30,accuracy:15",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":    true,
					"weapon abilities": true,
					"armor abilities":  true,
				},
				ExpLengths: map[string]int{
					"character classes": 3,
					"aeon commands":     0,
					"default abilities": 5,
					"overdrives":        2,
				},
			},
			Unique: apitesttypes.Unique{
				ID:   1,
				Name: "valefor",
			},
			UntypedUnit:    9,
			Area:           33,
			BattlesToRegen: 8,
			Agility: apitesttypes.TestAgilityParams{
				AgilityTier: 15,
				TickSpeed:   7,
				MinICV:      h.GetInt32Ptr(20),
				MaxICV:      h.GetInt32Ptr(21),
			},
			CelestialWeapon:  h.GetInt32Ptr(2),
			CharacterClasses: []int32{2, 3, 13},
			BaseStats: map[string]int32{
				"hp":            2146,
				"mp":            68,
				"strength":      47,
				"defense":       72,
				"magic":         76,
				"magic defense": 59,
				"agility":       35,
				"luck":          25,
				"evasion":       27,
				"accuracy":      44,
			},
			AeonCommands:           []int32{},
			DefaultPlayerAbilities: []int32{69, 70, 71, 72, 88},
			Overdrives:             []int32{117, 118},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons/1?yuna_stats=hp:2000,mp:250,strength:10,defense:28,magic:72,magic_defense:50,agility:41,luck:25,evasion:30,accuracy:15&battles=600",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"area":                  true,
					"battles to regenerate": true,
					"regen amounts":         true,
					"celestial weapon":      true,
					"character classes":     true,
					"aeon commands":         true,
					"default abilities":     true,
					"overdrives":            true,
					"weapon abilities":      true,
					"armor abilities":       true,
				},
				ExpLengths: map[string]int{},
			},
			Unique: apitesttypes.Unique{
				ID:   1,
				Name: "valefor",
			},
			UntypedUnit: 9,
			Agility: apitesttypes.TestAgilityParams{
				AgilityTier: 15,
				TickSpeed:   7,
				MinICV:      h.GetInt32Ptr(20),
				MaxICV:      h.GetInt32Ptr(21),
			},
			BaseStats: map[string]int32{
				"hp":            2225,
				"mp":            71,
				"strength":      47,
				"defense":       72,
				"magic":         76,
				"magic defense": 69,
				"agility":       35,
				"luck":          25,
				"evasion":       42,
				"accuracy":      44,
			},
		},
	}

	testSingleResources(t, tests, "GetAeon", testSrv.HandleAeons)
}

func TestRetrieveAeons(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons?optional=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{6, 7, 8, 9, 10},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/aeons?optional=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{1, 2, 3, 4, 5},
		},
	}

	testIDList(t, tests, apilookups.E.Aeons.Endpoint, "RetrieveAeons", testSrv.HandleAeons, compareAPIResourceLists[apiresource.Named])
}
