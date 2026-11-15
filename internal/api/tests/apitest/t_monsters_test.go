package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetMonster(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Monster{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/308",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "monster with provided id '308' doesn't exist. max id: 307",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/a",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "monster not found: 'a'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/a/2",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "monster not found: 'a', version '2'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/1/?ids=1",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid usage of parameter 'ids'. parameter 'ids' can only be used in the following format: '/api/monsters/simple?ids={id},...'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/a/2/3",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. usage: '/api/monsters', '/api/monsters/{id}', '/api/monsters/{name}', '/api/monsters/{name}/{version}', '/api/monsters/id/{subsection}'. supported subsections: 'abilities', 'areas', 'monster-formations', 'simple'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/211?has_overdrive=true",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid usage of parameter 'has_overdrive'. parameter 'has_overdrive' can only be used with list-endpoints",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/1?altered_state=1",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "monster 'sinscale - 1 (dream zanarkand)' has no altered states",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/105?altered_state=5",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "provided id '5' used for parameter 'altered_state' is out of range. max id: 1",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/210?omnis_elements=iifii",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid id '210'. parameter 'omnis_elements' can only be used with ids: '211'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/211?omnis_elements=iifii",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid input. omnis_elements must contain a combination of exactly four letters. valid letters are 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water)",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/211?omnis_elements=iftw",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid letter 't' for omnis_elements. use any four-letter-combination of 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water)",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/169?kimahri_stats=hp:1000",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid id '169'. parameter 'kimahri_stats' can only be used with ids: '167', '168'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/168?kimahri_stats=hp:100000",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "hp in 'kimahri_stats' can't be higher than 99999",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/168?kimahri_stats=hp-3000,strength:255",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid input for parameter 'kimahri_stats': 'hp-3000' . usage: '?kimahri_stats={stat}:{int},...'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/167?kimahri_stats=defense:5",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid stat 'defense' in 'kimahri_stats'. 'kimahri_stats' only uses 'hp', 'strength', 'magic', 'agility'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/1?yuna_stats=hp:200",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid id '1'. parameter 'yuna_stats' can only be used with ids: '216', '217', '218', '219', '220', '221', '222', '223', '224', '225'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/216?yuna_stats=hp:999999",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "hp in 'yuna_stats' can't be higher than 99999",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/216?yuna_stats=evasion:999999",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "evasion in 'yuna_stats' can't be higher than 255",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/27",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"ronso rages": true,
					"other items": true,
				},
				ExpLengths: map[string]int{
					"properties":        1,
					"auto-abilities":    0,
					"areas":             1,
					"formations":        3,
					"base stats":        10,
					"other items":       0,
					"weapon abilities":  3,
					"armor abilities":   1,
					"status immunities": 7,
					"status resists":    1,
					"altered states":    0,
					"abilities":         1,
				},
			},
			NameVer: apitesttypes.NewExpNameVer(27, "yellow element", 0),
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 6,
				TickSpeed:   16,
				MinICV:      h.GetInt32Ptr(48),
				MaxICV:      h.GetInt32Ptr(53),
			},
			Species:       database.MonsterSpeciesElemental,
			CtbIconType:   database.CtbIconTypeMonster,
			Distance:      1,
			Properties:    []int32{2},
			AutoAbilities: []int32{},
			Areas:         []int32{53},
			Formations:    []int32{30, 31, 32},
			BaseStats: map[string]int32{
				"hp":      300,
				"defense": 120,
				"magic":   18,
				"evasion": 0,
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
					HealedHP: 103,
				},
				{
					Index:    7,
					Ticks:    15,
					HealedHP: 117,
				},
				{
					Index:    9,
					Ticks:    30,
					HealedHP: 135,
				},
			},
			Items: &apitesttypes.TestMonItems{
				ItemDropChance: 255,
				Items: map[string]*apitesttypes.TestResAmount[apiresource.Typed]{
					"steal common": h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](27, 1)),
					"steal rare":   h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](28, 1)),
					"drop common":  h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](71, 1)),
					"drop rare":    h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](71, 1)),
					"bribe":        h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](28, 6)),
				},
			},
			BribeChances: []apitypes.BribeChance{
				{
					Gil:    3000,
					Chance: 25,
				},
				{
					Gil:    4500,
					Chance: 50,
				},
				{
					Gil:    6000,
					Chance: 75,
				},
				{
					Gil:    7500,
					Chance: 100,
				},
			},
			Equipment: &apitesttypes.TestMonEquipment{
				AbilitySlots: apitypes.MonsterEquipmentSlots{
					MinAmount: 1,
					MaxAmount: 2,
					Chances: []apitypes.EquipmentSlotsChance{
						{
							Amount: 1,
							Chance: 75,
						},
						{
							Amount: 2,
							Chance: 25,
						},
					},
				},
				AttachedAbilities: apitypes.MonsterEquipmentSlots{
					MinAmount: 0,
					MaxAmount: 2,
					Chances: []apitypes.EquipmentSlotsChance{
						{
							Amount: 0,
							Chance: 50,
						},
						{
							Amount: 1,
							Chance: 50,
						},
					},
				},
				WeaponAbilities: []int32{2, 6, 26},
				ArmorAbilities:  []int32{58},
			},
			ElemResists: []apitesttypes.TestElemResist{
				{Element: 1, Affinity: database.ElementalAffinityHalved},
				{Element: 2, Affinity: database.ElementalAffinityAbsorb},
				{Element: 3, Affinity: database.ElementalAffinityWeak},
				{Element: 4, Affinity: database.ElementalAffinityHalved},
				{Element: 5, Affinity: database.ElementalAffinityNeutral},
			},
			StatusImmunities: []int32{1, 4, 13},
			StatusResists: map[string]int32{
				"silence": 20,
			},
			Abilities: []int32{70},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/magIc_urn/1",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
					"elemental resists": true,
				},
				ExpLengths: map[string]int{
					"other items": 3,
				},
			},
			NameVer:      apitesttypes.NewExpNameVer(156, "magic urn", 1),
			Agility:      nil,
			RegenAmounts: nil,
			Items: &apitesttypes.TestMonItems{
				ItemDropChance: 0,
				Items: map[string]*apitesttypes.TestResAmount[apiresource.Typed]{
					"steal common": h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](1, 1)),
					"steal rare":   h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](1, 1)),
				},
				OtherItems: []apitesttypes.TestPossibleItem{
					apitesttypes.NewTestPossibleItem(0, 9, 1, 60),
					apitesttypes.NewTestPossibleItem(1, 64, 1, 20),
					apitesttypes.NewTestPossibleItem(2, 7, 1, 20),
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/sphErimorph?altered_state=1",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":  true,
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"other items":    true,
					"auto-abilities": true,
					"ronso rages":    true,
				},
				ExpLengths: map[string]int{
					"properties":        2,
					"auto-abilities":    0,
					"areas":             1,
					"formations":        1,
					"base stats":        10,
					"other items":       0,
					"weapon abilities":  5,
					"armor abilities":   4,
					"status immunities": 15,
					"status resists":    1,
					"altered states":    4,
					"abilities":         11,
				},
			},
			NameVer: apitesttypes.NewExpNameVer(86, "spherimorph", 0),
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 10,
				TickSpeed:   12,
				MinICV:      h.GetInt32Ptr(36),
				MaxICV:      h.GetInt32Ptr(40),
			},
			AppliedState: &apitesttypes.TestAppliedState{
				Condition:   "Fire-elemental.",
				IsTemporary: false,
			},
			Properties: []int32{1},
			Areas:      []int32{148},
			Formations: []int32{126},
			Items: &apitesttypes.TestMonItems{
				ItemDropChance: 255,
				Items: map[string]*apitesttypes.TestResAmount[apiresource.Typed]{
					"steal common": h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](5, 1)),
					"steal rare":   h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](6, 1)),
					"drop common":  h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](82, 1)),
					"drop rare":    h.GetStructPtr(apitesttypes.NewTestResAmount[apiresource.Typed](82, 1)),
					"bribe":        nil,
				},
			},
			Equipment: &apitesttypes.TestMonEquipment{
				AbilitySlots: apitypes.MonsterEquipmentSlots{
					MinAmount: 2,
					MaxAmount: 3,
					Chances: []apitypes.EquipmentSlotsChance{
						{
							Amount: 2,
							Chance: 50,
						},
						{
							Amount: 3,
							Chance: 50,
						},
					},
				},
				AttachedAbilities: apitypes.MonsterEquipmentSlots{
					MinAmount: 1,
					MaxAmount: 3,
					Chances: []apitypes.EquipmentSlotsChance{
						{
							Amount: 1,
							Chance: 50,
						},
						{
							Amount: 2,
							Chance: 50,
						},
					},
				},
				WeaponAbilities: []int32{2, 5, 6},
				ArmorAbilities:  []int32{55, 58, 61, 64},
			},
			ElemResists: []apitesttypes.TestElemResist{
				{Element: 1, Affinity: database.ElementalAffinityAbsorb},
				{Element: 2, Affinity: database.ElementalAffinityAbsorb},
				{Element: 3, Affinity: database.ElementalAffinityAbsorb},
				{Element: 4, Affinity: database.ElementalAffinityWeak},
				{Element: 5, Affinity: database.ElementalAffinityAbsorb},
			},
			StatusImmunities: []int32{2, 7, 12, 14, 32, 42, 45},
			StatusResists: map[string]int32{
				"poison": 90,
			},
			Abilities: []int32{69, 70, 71, 72, 415, 481, 548, 565, 705, 804, 831},
			DefaultState: &apitesttypes.TestDefaultState{
				IsTemporary: false,
				Change: &apitesttypes.TestAltStateChange{
					ElemResists: []apitesttypes.TestElemResist{
						{Element: 1, Affinity: database.ElementalAffinityVaries},
						{Element: 2, Affinity: database.ElementalAffinityVaries},
						{Element: 3, Affinity: database.ElementalAffinityVaries},
						{Element: 4, Affinity: database.ElementalAffinityVaries},
					},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/105?altered_state=1",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"agility params":    true,
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{
					"status immunities": 10,
					"status resists":    2,
					"altered states":    1,
				},
			},
			NameVer: apitesttypes.NewExpNameVer(105, "sand worm", 0),
			AppliedState: &apitesttypes.TestAppliedState{
				Condition:   "While 'Readying Quake'.",
				IsTemporary: true,
			},
			BribeChances:     nil,
			StatusImmunities: []int32{1, 2, 5, 9, 13, 32},
			StatusResists: map[string]int32{
				"darkness":    50,
				"power break": 50,
			},
			DefaultState: &apitesttypes.TestDefaultState{
				IsTemporary: false,
				Gain: &apitesttypes.TestAltStateGain{
					StatusResists: map[string]int32{
						"sleep": 80,
					},
				},
				Loss: &apitesttypes.TestAltStateLoss{
					StatusImmunities: []int32{13, 32},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/neslug?altered_state=1",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"agility params":    true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"equipment":         true,
					"elemental resists": true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
					"other items":       true,
				},
				ExpLengths: map[string]int{
					"properties":     3,
					"auto-abilities": 1,
					"altered states": 2,
				},
			},
			NameVer: apitesttypes.NewExpNameVer(287, "neslug", 0),
			AppliedState: &apitesttypes.TestAppliedState{
				Condition:   "While hidden in its shell.",
				IsTemporary: true,
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
					HealedHP: 15725,
				},
				{
					Index:    2,
					Ticks:    2,
					HealedHP: 31350,
				},
				{
					Index:    3,
					Ticks:    3,
					HealedHP: 46975,
				},
			},
			Properties:    []int32{6, 8},
			AutoAbilities: []int32{102},
			DefaultState: &apitesttypes.TestDefaultState{
				IsTemporary: false,
				Loss: &apitesttypes.TestAltStateLoss{
					Properties:    []int32{6},
					AutoAbilities: []int32{102},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/neslug?altered_state=2",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"elemental resists": true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{
					"properties":     2,
					"auto-abilities": 0,
					"altered states": 2,
				},
			},
			NameVer: apitesttypes.NewExpNameVer(287, "neslug", 0),
			AppliedState: &apitesttypes.TestAppliedState{
				Condition:   "Without its shell.",
				IsTemporary: false,
			},
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 18,
				TickSpeed:   4,
				MinICV:      h.GetInt32Ptr(12),
				MaxICV:      h.GetInt32Ptr(13),
			},
			Properties:    []int32{8},
			AutoAbilities: []int32{},
			BaseStats: map[string]int32{
				"agility": 120,
			},
			DefaultState: &apitesttypes.TestDefaultState{
				IsTemporary: false,
				Change: &apitesttypes.TestAltStateChange{
					BaseStats: map[string]int32{
						"agility": 43,
					},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/evrae?altered_state=1",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"agility params":    true,
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{
					"altered states": 1,
				},
			},
			NameVer: apitesttypes.NewExpNameVer(114, "evrae", 0),
			AppliedState: &apitesttypes.TestAppliedState{
				Condition:   "When the Airship is far away.",
				IsTemporary: false,
			},
			Distance: 4,
			DefaultState: &apitesttypes.TestDefaultState{
				IsTemporary: false,
				Change: &apitesttypes.TestAltStateChange{
					Distance: h.GetInt32Ptr(1),
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/penance?altered_state=1",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"species":           true,
					"regen amounts":     true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{
					"altered states": 1,
				},
			},
			NameVer: apitesttypes.NewExpNameVer(305, "penance", 0),
			AppliedState: &apitesttypes.TestAppliedState{
				Condition:     "When HP falls below 9000000.",
				IsTemporary:   false,
				AppliedStatus: h.GetInt32Ptr(21),
			},
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 19,
				TickSpeed:   1,
				MinICV:      h.GetInt32Ptr(9),
				MaxICV:      h.GetInt32Ptr(10),
			},
			DefaultState: &apitesttypes.TestDefaultState{
				IsTemporary: false,
				Gain: &apitesttypes.TestAltStateGain{
					AutoAbilities: []int32{99},
				},
				Loss: &apitesttypes.TestAltStateLoss{
					Status: h.GetInt32Ptr(21),
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/211?omnis_elements=iLfw",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"agility params":    true,
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{},
			},
			NameVer: apitesttypes.NewExpNameVer(211, "seymour omnis", 0),
			ElemResists: []apitesttypes.TestElemResist{
				{
					Element:  1,
					Affinity: database.ElementalAffinityHalved,
				},
				{
					Element:  2,
					Affinity: database.ElementalAffinityHalved,
				},
				{
					Element:  3,
					Affinity: database.ElementalAffinityHalved,
				},
				{
					Element:  4,
					Affinity: database.ElementalAffinityHalved,
				},
				{
					Element:  5,
					Affinity: database.ElementalAffinityNeutral,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/211?omnis_elements=iiff",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"agility params":    true,
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{},
			},
			NameVer: apitesttypes.NewExpNameVer(211, "seymour omnis", 0),
			ElemResists: []apitesttypes.TestElemResist{
				{
					Element:  1,
					Affinity: database.ElementalAffinityImmune,
				},
				{
					Element:  2,
					Affinity: database.ElementalAffinityNeutral,
				},
				{
					Element:  3,
					Affinity: database.ElementalAffinityNeutral,
				},
				{
					Element:  4,
					Affinity: database.ElementalAffinityImmune,
				},
				{
					Element:  5,
					Affinity: database.ElementalAffinityNeutral,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/211?omnis_elements=iiii",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"agility params":    true,
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{},
			},
			NameVer: apitesttypes.NewExpNameVer(211, "seymour omnis", 0),
			ElemResists: []apitesttypes.TestElemResist{
				{
					Element:  1,
					Affinity: database.ElementalAffinityWeak,
				},
				{
					Element:  2,
					Affinity: database.ElementalAffinityNeutral,
				},
				{
					Element:  3,
					Affinity: database.ElementalAffinityNeutral,
				},
				{
					Element:  4,
					Affinity: database.ElementalAffinityAbsorb,
				},
				{
					Element:  5,
					Affinity: database.ElementalAffinityNeutral,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/biran_ronso?kimahri_stats=hP:1000,strEngth:255,mAgic:255,agIlity:255",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{
					"ronso rages": 4,
				},
			},
			NameVer: apitesttypes.NewExpNameVer(167, "biran ronso", 0),
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 19,
				TickSpeed:   3,
				MinICV:      h.GetInt32Ptr(9),
				MaxICV:      h.GetInt32Ptr(10),
			},
			RonsoRages: []int32{4, 5, 8, 11},
			BaseStats: map[string]int32{
				"hp":       3549656,
				"strength": 12,
				"magic":    4,
				"agility":  251,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/yenke_ronso?kimahri_stats=hp:3500,strength:35,magic:45,agility:28",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{
					"ronso rages": 4,
				},
			},
			NameVer: apitesttypes.NewExpNameVer(168, "yenke ronso", 0),
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 12,
				TickSpeed:   10,
				MinICV:      h.GetInt32Ptr(30),
				MaxICV:      h.GetInt32Ptr(33),
			},
			RonsoRages: []int32{2, 6, 7, 9},
			BaseStats: map[string]int32{
				"hp":       10896,
				"strength": 13,
				"magic":    22,
				"agility":  22,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/yenke_ronso?kimahri_stats=hp:1500",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{},
			},
			NameVer: apitesttypes.NewExpNameVer(168, "yenke ronso", 0),
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 2,
				TickSpeed:   26,
				MinICV:      h.GetInt32Ptr(84),
				MaxICV:      h.GetInt32Ptr(93),
			},
			BaseStats: map[string]int32{
				"hp":       864,
				"strength": 8,
				"magic":    12,
				"agility":  1,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/vidatu",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"agility params":    true,
					"base stats":        true,
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{},
			},
			NameVer: apitesttypes.NewExpNameVer(273, "vidatu", 0),
			RegenAmounts: []apitesttypes.TestRegenAmount{
				{
					Index:    0,
					Ticks:    0,
					HealedHP: 100,
				},
				{
					Index:    5,
					Ticks:    5,
					HealedHP: 1955,
				},
				{
					Index:    6,
					Ticks:    10,
					HealedHP: 3810,
				},
				{
					Index:    8,
					Ticks:    20,
					HealedHP: 7521,
				},
				{
					Index:    10,
					Ticks:    50,
					HealedHP: 18654,
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/216?yuna_stats=hp:2000,mp:250,strength:10,defense:28,magic:72,magic_defense:50,agility:41,luck:25,evasion:30,accuracy:15",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{},
			},
			NameVer: apitesttypes.NewExpNameVer(216, "valefor", 1),
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 15,
				TickSpeed:   7,
				MinICV:      h.GetInt32Ptr(-1),
				MaxICV:      h.GetInt32Ptr(-1),
			},
			BaseStats: map[string]int32{
				"hp":            2146,
				"mp":            1,
				"strength":      47,
				"defense":       72,
				"magic":         76,
				"magic defense": 59,
				"agility":       35,
				"luck":          1,
				"evasion":       27,
				"accuracy":      44,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/216?battles=600",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{},
			},
			NameVer: apitesttypes.NewExpNameVer(216, "valefor", 1),
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 14,
				TickSpeed:   8,
				MinICV:      h.GetInt32Ptr(-1),
				MaxICV:      h.GetInt32Ptr(-1),
			},
			BaseStats: map[string]int32{
				"hp":            2225,
				"mp":            1,
				"strength":      45,
				"defense":       66,
				"magic":         64,
				"magic defense": 69,
				"agility":       34,
				"luck":          1,
				"evasion":       42,
				"accuracy":      20,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/216?yuna_stats=hp:2000,mp:250,strength:10,defense:28,magic:72,magic_defense:50,agility:41,luck:25,evasion:30,accuracy:15&battles=600",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{},
			},
			NameVer: apitesttypes.NewExpNameVer(216, "valefor", 1),
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 15,
				TickSpeed:   7,
				MinICV:      h.GetInt32Ptr(-1),
				MaxICV:      h.GetInt32Ptr(-1),
			},
			BaseStats: map[string]int32{
				"hp":            2225,
				"mp":            1,
				"strength":      47,
				"defense":       72,
				"magic":         76,
				"magic defense": 69,
				"agility":       35,
				"luck":          1,
				"evasion":       42,
				"accuracy":      44,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/87?rel_availability=pre-story&rel_repeatable=false",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"bribe chances":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"status immunities": true,
					"abilities":         true,
					"base stats":        true,
					"agility params":    true,
				},
				ExpLengths: map[string]int{
					"areas":     1,
					"formation": 1,
				},
			},
			NameVer:    apitesttypes.NewExpNameVer(87, "evil eye", 1),
			Areas:      []int32{159},
			Formations: []int32{134},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/87?rel_availability=pre-story&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"bribe chances":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"status immunities": true,
					"abilities":         true,
					"base stats":        true,
					"agility params":    true,
				},
				ExpLengths: map[string]int{
					"areas":     1,
					"formation": 2,
				},
			},
			NameVer:    apitesttypes.NewExpNameVer(87, "evil eye", 1),
			Areas:      []int32{159},
			Formations: []int32{128, 132},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/87?rel_availability=always&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"bribe chances":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"status immunities": true,
					"abilities":         true,
					"base stats":        true,
					"agility params":    true,
				},
				ExpLengths: map[string]int{
					"areas":     2,
					"formation": 2,
				},
			},
			NameVer:    apitesttypes.NewExpNameVer(87, "evil eye", 1),
			Areas:      []int32{154, 157},
			Formations: []int32{128, 132},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/seymour",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"regen amounts":     true,
					"species":           true,
					"ctb icon type":     true,
					"distance":          true,
					"items":             true,
					"other items":       true,
					"equipment":         true,
					"default state":     true,
					"elemental resists": true,
					"properties":        true,
					"auto-abilities":    true,
					"ronso rages":       true,
					"areas":             true,
					"formations":        true,
					"status immunities": true,
					"abilities":         true,
				},
				ExpLengths: map[string]int{},
			},
			NameVer: apitesttypes.NewExpNameVer(93, "seymour", 0),
			Agility: &apitesttypes.TestAgilityParams{
				AgilityTier: 12,
				TickSpeed:   10,
				MinICV:      h.GetInt32Ptr(-1),
				MaxICV:      h.GetInt32Ptr(-1),
			},
			AutoAbilities: []int32{3},
		},
	}

	testSingleResources(t, tests, "GetMonster", testSrv.HandleMonsters)
}

func TestGetMultipleMonsters(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/guado_guardian",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   3,
			Results: []int32{94, 96, 113},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/yojimbo",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   3,
			Results: []int32{165, 222, 235},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/mimic",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   4,
			Results: []int32{249, 250, 251, 252},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/%3F%3F%3F",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   4,
			Results: []int32{68, 69, 108, 253},
		},
	}

	testIDList(t, tests, apilookups.E.Monsters.Endpoint, "GetMultipleMonsters", testSrv.HandleMonsters, compareAPIResourceLists[apiresource.Named])
}

func TestRetrieveMonsters(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?limit=asd",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value 'asd' used for parameter 'limit'. usage: '?limit={int|'max'}'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?ids=1,33,22,68,",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid usage of parameter 'ids'. parameter 'ids' can only be used in the following format: '/api/monsters/simple?ids={id},...'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?elemental_resists=weak",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid input for parameter 'elemental_resists': 'weak'. usage: '?elemental_resists={element|id}:{affinity|id},...'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?elemental_resists=fire-weak,water:absorb",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid input for parameter 'elemental_resists': 'fire-weak'. usage: '?elemental_resists={element|id}:{affinity|id},...'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/?omnis_elements=ffff",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid usage of parameter 'omnis_elements'. parameter 'omnis_elements' can only be used with single-resource-endpoints",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?elemental_resists=weak:fire",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "unknown element 'weak' used for parameter 'elemental_resists'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?elemental_resists=fire:weak,fire:neutral",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "duplicate use of id '1' for parameter 'elemental_resists'. each element can only be used once",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?status_resists=4&resistance=350",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value '350' used for parameter 'resistance'. value must be an integer ranging from 1 to 254",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?status_resists=4&resistance=frank",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value 'frank' used for parameter 'resistance'. usage: '?resistance={int|'immune'}'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?resistance=50",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid usage of parameter 'resistance'. when using parameter 'resistance', the following parameter(s) must be present: 'status_resists'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?methods=steal",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid usage of parameter 'methods'. when using parameter 'methods', the following parameter(s) must be present: 'item'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?item=22&methods=steals",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value 'steals' used for parameter 'methods'. allowed values: 'steal', 'drop', 'bribe', 'other'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?item=asf&methods=drop",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid id 'asf' used for parameter 'item'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?distance=5",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value '5' used for parameter 'distance'. value must be an integer ranging from 1 to 4",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?distance=frank",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid value 'frank' used for parameter 'distance'. usage: '?distance={int},...|{int}-{int}'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?ronso_rage=13",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "provided id '13' used for parameter 'ronso_rage' is out of range. max id: 12",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?species=wywrm",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid enum value 'wywrm' used for parameter 'species'. use /api/monsters/parameters to see allowed values",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    307,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 175, 238, 307},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?elemental_resists=FIre:weAk,water:neutral",
				ExpectedStatus: http.StatusOK,
				DontCheck: map[string]bool{
					"next": true,
				},
			},
			Count:   22,
			Results: []int32{11, 23, 64, 148},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?limit=max&status_resists=37",
				ExpectedStatus: http.StatusOK,
			},
			Count:   43,
			Results: []int32{32, 127, 211, 233, 295},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?limit=max&status_resists=1,4,10&resistance=50",
				ExpectedStatus: http.StatusOK,
			},
			Count:   150,
			Results: []int32{3, 128, 188, 227, 249},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?limit=max&status_resists=4&resistance=immune",
				ExpectedStatus: http.StatusOK,
			},
			Count:   163,
			Results: []int32{5, 67, 100, 151, 258},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?limit=max&item=7",
				ExpectedStatus: http.StatusOK,
			},
			Count:   22,
			Results: []int32{32, 91, 156, 192, 295, 305},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?item=7&methods=drop",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{32, 91},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?auto_ability=24",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{199, 200, 209, 244, 257, 279, 285},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?auto_ability=2&is_forced=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{167, 168, 202},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?empty_slots=3",
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{74, 80, 118, 119, 161, 201, 244, 247},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?ronso_rage=12",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{255, 292},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?location=15",
				ExpectedStatus: http.StatusOK,
			},
			Count:   19,
			Results: []int32{80, 90, 297},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?sublocation=26", // macalania woods
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{80, 86},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?area=89", // mi'ihen highroad - central
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{38, 45},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?distance=2&availability=pre-story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   5,
			Results: []int32{57, 58, 92, 119, 194},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?repeatable=true&capture=false&has_overdrive=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{229, 236, 299},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?underwater=true&category=bOss",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{5, 71, 291},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?zombie=true&species=wyRm",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{134},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?creation_area=DJose",
				ExpectedStatus: http.StatusOK,
			},
			Count:   7,
			Results: []int32{60, 63, 67},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?availability=post&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   71,
			Results: []int32{201, 202, 210, 227, 246, 268, 284, 307},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?availability=post&location=4",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{293},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?availability=post-game&location=4",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{16, 17, 18, 293},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?availability=post&location=23",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{301},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?availability=post&location=25",
				ExpectedStatus: http.StatusOK,
			},
			Count:   11,
			Results: []int32{170, 201, 202, 203, 208, 210, 247},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?availability=post-game&area=140", // thunder plains - north
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{72, 74, 78, 295, 296},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?availability=story&sublocation=35", // calm lands
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{149, 150},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?availability=story&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   110,
			Results: []int32{1, 8, 23, 57, 109, 127, 185, 217, 226},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?availability=story&repeatable=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   25,
			Results: []int32{7, 31, 111, 118, 130, 209},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters?availability=post-game&repeatable=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   22,
			Results: []int32{28, 107, 165, 293, 298, 305, 307},
		},
	}

	testIDList(t, tests, apilookups.E.Monsters.Endpoint, "RetrieveMonsters", testSrv.HandleMonsters, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionMonsters(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/?ids=1&location=15",
				ExpectedStatus: http.StatusBadRequest,
				Handler:        testSrv.HandleMonsters,
				ExpectedErr:    "invalid usage of parameter 'ids'. parameter 'ids' can only be used in the following format: '/api/monsters/simple?ids={id},...'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/simple?limit=max",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsters,
				DontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			Count:   307,
			Results: []int32{1, 22, 33, 68, 123, 199, 278, 307},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/simple?ids=1-21",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsters,
				DontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			Count:   21,
			Results: []int32{1, 7, 18, 21},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/simple?ids=1,33,22,22,22,22,68,",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsters,
				DontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			Count:   4,
			Results: []int32{1, 22, 33, 68},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/simple?ids=1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsters,
				DontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			Count:   21,
			Results: []int32{1, 5, 16, 18, 21},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/areas/89/monsters/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          6,
			ParentResource: h.GetStrPtr("/areas/89"),
			Results:        []int32{38, 39, 40, 42, 43, 45},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/areas/23/monsters/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          4,
			ParentResource: h.GetStrPtr("/areas/23"),
			Results:        []int32{15, 16, 17, 18},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/areas/236/monsters/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          21,
			Next:           nil,
			ParentResource: h.GetStrPtr("/areas/236"),
			Results:        []int32{190, 201, 210, 239, 245, 249, 253},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations/193/monsters/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsterFormations,
			},
			Count:          3,
			ParentResource: h.GetStrPtr("/monster-formations/193"),
			Results:        []int32{143, 145, 148},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monster-formations/140/monsters/",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleMonsterFormations,
			},
			Count:          3,
			ParentResource: h.GetStrPtr("/monster-formations/140"),
			Results:        []int32{104, 99, 103},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/monsters/simple?ids=1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1",
				ExpectedStatus: http.StatusBadRequest,
				Handler:        testSrv.HandleMonsters,
				ExpectedErr:    "fetch limit exceeded. the maximum amount of inputs is 1000",
			},
		},
	}

	testIDList(t, tests, apilookups.E.Monsters.Endpoint, "SubsectionMonsters", nil, compareSimpleResourceLists[apiresource.Named, sectiontypes.MonsterSimple])
}
