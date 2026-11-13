package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetMonster(t *testing.T) {
	t.Parallel()
	tests := []Monster{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/308",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "monster with provided id '308' doesn't exist. max id: 307",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "monster not found: 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/a/2",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "monster not found: 'a', version '2'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/1/?ids=1",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'ids'. parameter 'ids' can only be used in the following format: '/api/monsters/simple?ids={id},...'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/a/2/3",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. usage: '/api/monsters', '/api/monsters/{id}', '/api/monsters/{name}', '/api/monsters/{name}/{version}', '/api/monsters/id/{subsection}'. supported subsections: 'abilities', 'areas', 'monster-formations', 'simple'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/211?has_overdrive=true",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'has_overdrive'. parameter 'has_overdrive' can only be used with list-endpoints",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/1?altered_state=1",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'sinscale - 1 (dream zanarkand)' has no altered states",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/105?altered_state=5",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '5' used for parameter 'altered_state' is out of range. max id: 1",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/210?omnis_elements=iifii",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id '210'. parameter 'omnis_elements' can only be used with ids: '211'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/211?omnis_elements=iifii",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input. omnis_elements must contain a combination of exactly four letters. valid letters are 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water)",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/211?omnis_elements=iftw",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid letter 't' for omnis_elements. use any four-letter-combination of 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water)",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/169?kimahri_stats=hp:1000",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id '169'. parameter 'kimahri_stats' can only be used with ids: '167', '168'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/168?kimahri_stats=hp:100000",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "hp in 'kimahri_stats' can't be higher than 99999",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/168?kimahri_stats=hp-3000,strength:255",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input for parameter 'kimahri_stats': 'hp-3000' . usage: '?kimahri_stats={stat}:{int},...'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/167?kimahri_stats=defense:5",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid stat 'defense' in 'kimahri_stats'. 'kimahri_stats' only uses 'hp', 'strength', 'magic', 'agility'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/1?yuna_stats=hp:200",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id '1'. parameter 'yuna_stats' can only be used with ids: '216', '217', '218', '219', '220', '221', '222', '223', '224', '225'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/216?yuna_stats=hp:999999",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "hp in 'yuna_stats' can't be higher than 99999",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/216?yuna_stats=evasion:999999",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "evasion in 'yuna_stats' can't be higher than 255",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/27",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"ronso rages": true,
					"other items": true,
				},
				expLengths: map[string]int{
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
			NameVer: newExpNameVer(27, "yellow element", 0),
			agility: &TestAgilityParams{
				agilityTier: 6,
				tickSpeed:   16,
				minICV:      h.GetInt32Ptr(48),
				maxICV:      h.GetInt32Ptr(53),
			},
			species:       database.MonsterSpeciesElemental,
			ctbIconType:   database.CtbIconTypeMonster,
			distance:      1,
			properties:    []int32{2},
			autoAbilities: []int32{},
			areas:         []int32{53},
			formations:    []int32{30, 31, 32},
			baseStats: map[string]int32{
				"hp":      300,
				"defense": 120,
				"magic":   18,
				"evasion": 0,
			},
			regenAmounts: []TestRegenAmount{
				{
					index:    0,
					ticks:    0,
					healedHP: 100,
				},
				{
					index:    3,
					ticks:    3,
					healedHP: 103,
				},
				{
					index:    7,
					ticks:    15,
					healedHP: 117,
				},
				{
					index:    9,
					ticks:    30,
					healedHP: 135,
				},
			},
			items: &TestMonItems{
				itemDropChance: 255,
				items: map[string]*TestResAmount[apiresource.Typed]{
					"steal common": h.GetStructPtr(newTestResAmount[apiresource.Typed](27, 1)),
					"steal rare":   h.GetStructPtr(newTestResAmount[apiresource.Typed](28, 1)),
					"drop common":  h.GetStructPtr(newTestResAmount[apiresource.Typed](71, 1)),
					"drop rare":    h.GetStructPtr(newTestResAmount[apiresource.Typed](71, 1)),
					"bribe":        h.GetStructPtr(newTestResAmount[apiresource.Typed](28, 6)),
				},
			},
			bribeChances: []apitypes.BribeChance{
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
			equipment: &TestMonEquipment{
				abilitySlots: apitypes.MonsterEquipmentSlots{
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
				attachedAbilities: apitypes.MonsterEquipmentSlots{
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
				weaponAbilities: []int32{2, 6, 26},
				armorAbilities:  []int32{58},
			},
			elemResists: []TestElemResist{
				{element: 1, affinity: database.ElementalAffinityHalved},
				{element: 2, affinity: database.ElementalAffinityAbsorb},
				{element: 3, affinity: database.ElementalAffinityWeak},
				{element: 4, affinity: database.ElementalAffinityHalved},
				{element: 5, affinity: database.ElementalAffinityNeutral},
			},
			statusImmunities: []int32{1, 4, 13},
			statusResists: map[string]int32{
				"silence": 20,
			},
			abilities: []int32{70},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/magIc_urn/1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"other items": 3,
				},
			},
			NameVer:   newExpNameVer(156, "magic urn", 1),
			agility:      nil,
			regenAmounts: nil,
			items: &TestMonItems{
				itemDropChance: 0,
				items: map[string]*TestResAmount[apiresource.Typed]{
					"steal common": h.GetStructPtr(newTestResAmount[apiresource.Typed](1, 1)),
					"steal rare":   h.GetStructPtr(newTestResAmount[apiresource.Typed](1, 1)),
				},
				otherItems: []TestPossibleItem{
					newTestPossibleItem(0, 9, 1, 60),
					newTestPossibleItem(1, 64, 1, 20),
					newTestPossibleItem(2, 7, 1, 20),
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/sphErimorph?altered_state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"regen amounts":  true,
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"other items":    true,
					"auto-abilities": true,
					"ronso rages":    true,
				},
				expLengths: map[string]int{
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
			NameVer: newExpNameVer(86, "spherimorph", 0),
			agility: &TestAgilityParams{
				agilityTier: 10,
				tickSpeed:   12,
				minICV:      h.GetInt32Ptr(36),
				maxICV:      h.GetInt32Ptr(40),
			},
			appliedState: &TestAppliedState{
				condition:   "Fire-elemental.",
				isTemporary: false,
			},
			properties: []int32{1},
			areas:      []int32{148},
			formations: []int32{126},
			items: &TestMonItems{
				itemDropChance: 255,
				items: map[string]*TestResAmount[apiresource.Typed]{
					"steal common": h.GetStructPtr(newTestResAmount[apiresource.Typed](5, 1)),
					"steal rare":   h.GetStructPtr(newTestResAmount[apiresource.Typed](6, 1)),
					"drop common":  h.GetStructPtr(newTestResAmount[apiresource.Typed](82, 1)),
					"drop rare":    h.GetStructPtr(newTestResAmount[apiresource.Typed](82, 1)),
					"bribe":        nil,
				},
			},
			equipment: &TestMonEquipment{
				abilitySlots: apitypes.MonsterEquipmentSlots{
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
				attachedAbilities: apitypes.MonsterEquipmentSlots{
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
				weaponAbilities: []int32{2, 5, 6},
				armorAbilities:  []int32{55, 58, 61, 64},
			},
			elemResists: []TestElemResist{
				{element: 1, affinity: database.ElementalAffinityAbsorb},
				{element: 2, affinity: database.ElementalAffinityAbsorb},
				{element: 3, affinity: database.ElementalAffinityAbsorb},
				{element: 4, affinity: database.ElementalAffinityWeak},
				{element: 5, affinity: database.ElementalAffinityAbsorb},
			},
			statusImmunities: []int32{2, 7, 12, 14, 32, 42, 45},
			statusResists: map[string]int32{
				"poison": 90,
			},
			abilities: []int32{69, 70, 71, 72, 415, 481, 548, 565, 705, 804, 831},
			defaultState: &TestDefaultState{
				IsTemporary: false,
				Change: &TestAltStateChange{
					ElemResists: []TestElemResist{
						{element: 1, affinity: database.ElementalAffinityVaries},
						{element: 2, affinity: database.ElementalAffinityVaries},
						{element: 3, affinity: database.ElementalAffinityVaries},
						{element: 4, affinity: database.ElementalAffinityVaries},
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/105?altered_state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"status immunities": 10,
					"status resists":    2,
					"altered states":    1,
				},
			},
			NameVer: newExpNameVer(105, "sand worm", 0),
			appliedState: &TestAppliedState{
				condition:   "While 'Readying Quake'.",
				isTemporary: true,
			},
			bribeChances:     nil,
			statusImmunities: []int32{1, 2, 5, 9, 13, 32},
			statusResists: map[string]int32{
				"darkness":    50,
				"power break": 50,
			},
			defaultState: &TestDefaultState{
				IsTemporary: false,
				Gain: &TestAltStateGain{
					StatusResists: map[string]int32{
						"sleep": 80,
					},
				},
				Loss: &TestAltStateLoss{
					StatusImmunities: []int32{13, 32},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/neslug?altered_state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"properties":     3,
					"auto-abilities": 1,
					"altered states": 2,
				},
			},
			NameVer: newExpNameVer(287, "neslug", 0),
			appliedState: &TestAppliedState{
				condition:   "While hidden in its shell.",
				isTemporary: true,
			},
			regenAmounts: []TestRegenAmount{
				{
					index:    0,
					ticks:    0,
					healedHP: 100,
				},
				{
					index:    1,
					ticks:    1,
					healedHP: 15725,
				},
				{
					index:    2,
					ticks:    2,
					healedHP: 31350,
				},
				{
					index:    3,
					ticks:    3,
					healedHP: 46975,
				},
			},
			properties:    []int32{6, 8},
			autoAbilities: []int32{102},
			defaultState: &TestDefaultState{
				IsTemporary: false,
				Loss: &TestAltStateLoss{
					Properties:    []int32{6},
					AutoAbilities: []int32{102},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/neslug?altered_state=2",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"properties":     2,
					"auto-abilities": 0,
					"altered states": 2,
				},
			},
			NameVer: newExpNameVer(287, "neslug", 0),
			appliedState: &TestAppliedState{
				condition:   "Without its shell.",
				isTemporary: false,
			},
			agility: &TestAgilityParams{
				agilityTier: 18,
				tickSpeed:   4,
				minICV:      h.GetInt32Ptr(12),
				maxICV:      h.GetInt32Ptr(13),
			},
			properties:    []int32{8},
			autoAbilities: []int32{},
			baseStats: map[string]int32{
				"agility": 120,
			},
			defaultState: &TestDefaultState{
				IsTemporary: false,
				Change: &TestAltStateChange{
					BaseStats: map[string]int32{
						"agility": 43,
					},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/evrae?altered_state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"altered states": 1,
				},
			},
			NameVer: newExpNameVer(114, "evrae", 0),
			appliedState: &TestAppliedState{
				condition:   "When the Airship is far away.",
				isTemporary: false,
			},
			distance: 4,
			defaultState: &TestDefaultState{
				IsTemporary: false,
				Change: &TestAltStateChange{
					Distance: h.GetInt32Ptr(1),
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/penance?altered_state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"altered states": 1,
				},
			},
			NameVer: newExpNameVer(305, "penance", 0),
			appliedState: &TestAppliedState{
				condition:     "When HP falls below 9000000.",
				isTemporary:   false,
				appliedStatus: h.GetInt32Ptr(21),
			},
			agility: &TestAgilityParams{
				agilityTier: 19,
				tickSpeed:   1,
				minICV:      h.GetInt32Ptr(9),
				maxICV:      h.GetInt32Ptr(10),
			},
			defaultState: &TestDefaultState{
				IsTemporary: false,
				Gain: &TestAltStateGain{
					AutoAbilities: []int32{99},
				},
				Loss: &TestAltStateLoss{
					Status: h.GetInt32Ptr(21),
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/211?omnis_elements=iLfw",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{},
			},
			NameVer: newExpNameVer(211, "seymour omnis", 0),
			elemResists: []TestElemResist{
				{
					element:  1,
					affinity: database.ElementalAffinityHalved,
				},
				{
					element:  2,
					affinity: database.ElementalAffinityHalved,
				},
				{
					element:  3,
					affinity: database.ElementalAffinityHalved,
				},
				{
					element:  4,
					affinity: database.ElementalAffinityHalved,
				},
				{
					element:  5,
					affinity: database.ElementalAffinityNeutral,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/211?omnis_elements=iiff",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{},
			},
			NameVer: newExpNameVer(211, "seymour omnis", 0),
			elemResists: []TestElemResist{
				{
					element:  1,
					affinity: database.ElementalAffinityImmune,
				},
				{
					element:  2,
					affinity: database.ElementalAffinityNeutral,
				},
				{
					element:  3,
					affinity: database.ElementalAffinityNeutral,
				},
				{
					element:  4,
					affinity: database.ElementalAffinityImmune,
				},
				{
					element:  5,
					affinity: database.ElementalAffinityNeutral,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/211?omnis_elements=iiii",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{},
			},
			NameVer: newExpNameVer(211, "seymour omnis", 0),
			elemResists: []TestElemResist{
				{
					element:  1,
					affinity: database.ElementalAffinityWeak,
				},
				{
					element:  2,
					affinity: database.ElementalAffinityNeutral,
				},
				{
					element:  3,
					affinity: database.ElementalAffinityNeutral,
				},
				{
					element:  4,
					affinity: database.ElementalAffinityAbsorb,
				},
				{
					element:  5,
					affinity: database.ElementalAffinityNeutral,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/biran_ronso?kimahri_stats=hP:1000,strEngth:255,mAgic:255,agIlity:255",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"ronso rages": 4,
				},
			},
			NameVer: newExpNameVer(167, "biran ronso", 0),
			agility: &TestAgilityParams{
				agilityTier: 19,
				tickSpeed:   3,
				minICV:      h.GetInt32Ptr(9),
				maxICV:      h.GetInt32Ptr(10),
			},
			ronsoRages: []int32{4, 5, 8, 11},
			baseStats: map[string]int32{
				"hp":       3549656,
				"strength": 12,
				"magic":    4,
				"agility":  251,
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/yenke_ronso?kimahri_stats=hp:3500,strength:35,magic:45,agility:28",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"ronso rages": 4,
				},
			},
			NameVer: newExpNameVer(168, "yenke ronso", 0),
			agility: &TestAgilityParams{
				agilityTier: 12,
				tickSpeed:   10,
				minICV:      h.GetInt32Ptr(30),
				maxICV:      h.GetInt32Ptr(33),
			},
			ronsoRages: []int32{2, 6, 7, 9},
			baseStats: map[string]int32{
				"hp":       10896,
				"strength": 13,
				"magic":    22,
				"agility":  22,
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/yenke_ronso?kimahri_stats=hp:1500",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{},
			},
			NameVer: newExpNameVer(168, "yenke ronso", 0),
			agility: &TestAgilityParams{
				agilityTier: 2,
				tickSpeed:   26,
				minICV:      h.GetInt32Ptr(84),
				maxICV:      h.GetInt32Ptr(93),
			},
			baseStats: map[string]int32{
				"hp":       864,
				"strength": 8,
				"magic":    12,
				"agility":  1,
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/vidatu",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{},
			},
			NameVer: newExpNameVer(273, "vidatu", 0),
			regenAmounts: []TestRegenAmount{
				{
					index:    0,
					ticks:    0,
					healedHP: 100,
				},
				{
					index:    5,
					ticks:    5,
					healedHP: 1955,
				},
				{
					index:    6,
					ticks:    10,
					healedHP: 3810,
				},
				{
					index:    8,
					ticks:    20,
					healedHP: 7521,
				},
				{
					index:    10,
					ticks:    50,
					healedHP: 18654,
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/216?yuna_stats=hp:2000,mp:250,strength:10,defense:28,magic:72,magic_defense:50,agility:41,luck:25,evasion:30,accuracy:15",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{},
			},
			NameVer: newExpNameVer(216, "valefor", 1),
			agility: &TestAgilityParams{
				agilityTier: 15,
				tickSpeed:   7,
				minICV:      h.GetInt32Ptr(-1),
				maxICV:      h.GetInt32Ptr(-1),
			},
			baseStats: map[string]int32{
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
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/216?battles=600",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{},
			},
			NameVer: newExpNameVer(216, "valefor", 1),
			agility: &TestAgilityParams{
				agilityTier: 14,
				tickSpeed:   8,
				minICV:      h.GetInt32Ptr(-1),
				maxICV:      h.GetInt32Ptr(-1),
			},
			baseStats: map[string]int32{
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
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/216?yuna_stats=hp:2000,mp:250,strength:10,defense:28,magic:72,magic_defense:50,agility:41,luck:25,evasion:30,accuracy:15&battles=600",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{},
			},
			NameVer: newExpNameVer(216, "valefor", 1),
			agility: &TestAgilityParams{
				agilityTier: 15,
				tickSpeed:   7,
				minICV:      h.GetInt32Ptr(-1),
				maxICV:      h.GetInt32Ptr(-1),
			},
			baseStats: map[string]int32{
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
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/87?rel_availability=pre-story&rel_repeatable=false",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"areas":     1,
					"formation": 1,
				},
			},
			NameVer: newExpNameVer(87, "evil eye", 1),
			areas:      []int32{159},
			formations: []int32{134},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/87?rel_availability=pre-story&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"areas":     1,
					"formation": 2,
				},
			},
			NameVer: newExpNameVer(87, "evil eye", 1),
			areas:      []int32{159},
			formations: []int32{128, 132},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/87?rel_availability=always&rel_repeatable=true",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{
					"areas":     2,
					"formation": 2,
				},
			},
			NameVer: newExpNameVer(87, "evil eye", 1),
			areas:      []int32{154, 157},
			formations: []int32{128, 132},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/seymour",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
				expLengths: map[string]int{},
			},
			NameVer: newExpNameVer(93, "seymour", 0),
			agility: &TestAgilityParams{
				agilityTier: 12,
				tickSpeed:   10,
				minICV:      h.GetInt32Ptr(-1),
				maxICV:      h.GetInt32Ptr(-1),
			},
			autoAbilities: []int32{3},
		},
	}

	testSingleResources(t, tests, "GetMonster", testSrv.HandleMonsters, compareMonsters)
}

func TestGetMultipleMonsters(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/guado_guardian",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   3,
			results: []int32{94, 96, 113},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/yojimbo",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   3,
			results: []int32{165, 222, 235},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/mimic",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   4,
			results: []int32{249, 250, 251, 252},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/%3F%3F%3F",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   4,
			results: []int32{68, 69, 108, 253},
		},
	}

	testIDList(t, tests, apilookups.E.Monsters.Endpoint, "GetMultipleMonsters", testSrv.HandleMonsters, compareAPIResourceLists[apiresource.Named])
}

func TestRetrieveMonsters(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?limit=asd",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'asd' used for parameter 'limit'. usage: '?limit={int|'max'}'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?ids=1,33,22,68,",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'ids'. parameter 'ids' can only be used in the following format: '/api/monsters/simple?ids={id},...'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?elemental_resists=weak",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input for parameter 'elemental_resists': 'weak'. usage: '?elemental_resists={element|id}:{affinity|id},...'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?elemental_resists=fire-weak,water:absorb",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input for parameter 'elemental_resists': 'fire-weak'. usage: '?elemental_resists={element|id}:{affinity|id},...'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/?omnis_elements=ffff",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'omnis_elements'. parameter 'omnis_elements' can only be used with single-resource-endpoints",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?elemental_resists=weak:fire",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "unknown element 'weak' used for parameter 'elemental_resists'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?elemental_resists=fire:weak,fire:neutral",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "duplicate use of id '1' for parameter 'elemental_resists'. each element can only be used once",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?status_resists=4&resistance=350",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value '350' used for parameter 'resistance'. value must be an integer ranging from 1 to 254",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?status_resists=4&resistance=frank",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'frank' used for parameter 'resistance'. usage: '?resistance={int|'immune'}'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?resistance=50",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'resistance'. when using parameter 'resistance', the following parameter(s) must be present: 'status_resists'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?methods=steal",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'methods'. when using parameter 'methods', the following parameter(s) must be present: 'item'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?item=22&methods=steals",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'steals' used for parameter 'methods'. allowed values: 'steal', 'drop', 'bribe', 'other'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?item=asf&methods=drop",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id 'asf' used for parameter 'item'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?distance=5",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value '5' used for parameter 'distance'. value must be an integer ranging from 1 to 4",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?distance=frank",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'frank' used for parameter 'distance'. usage: '?distance={int},...|{int}-{int}'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?ronso_rage=13",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '13' used for parameter 'ronso_rage' is out of range. max id: 12",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?species=wywrm",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid enum value 'wywrm' used for parameter 'species'. use /api/monsters/parameters to see allowed values",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    307,
			previous: nil,
			next:     nil,
			results:  []int32{1, 175, 238, 307},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?elemental_resists=FIre:weAk,water:neutral",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"next": true,
				},
			},
			count:   22,
			results: []int32{11, 23, 64, 148},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?limit=max&status_resists=37",
				expectedStatus: http.StatusOK,
			},
			count:   43,
			results: []int32{32, 127, 211, 233, 295},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?limit=max&status_resists=1,4,10&resistance=50",
				expectedStatus: http.StatusOK,
			},
			count:   150,
			results: []int32{3, 128, 188, 227, 249},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?limit=max&status_resists=4&resistance=immune",
				expectedStatus: http.StatusOK,
			},
			count:   163,
			results: []int32{5, 67, 100, 151, 258},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?limit=max&item=7",
				expectedStatus: http.StatusOK,
			},
			count:   22,
			results: []int32{32, 91, 156, 192, 295, 305},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?item=7&methods=drop",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{32, 91},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?auto_ability=24",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{199, 200, 209, 244, 257, 279, 285},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?auto_ability=2&is_forced=false",
				expectedStatus: http.StatusOK,
			},
			count:   3,
			results: []int32{167, 168, 202},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?empty_slots=3",
				expectedStatus: http.StatusOK,
			},
			count:   8,
			results: []int32{74, 80, 118, 119, 161, 201, 244, 247},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?ronso_rage=12",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{255, 292},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?location=15",
				expectedStatus: http.StatusOK,
			},
			count:   19,
			results: []int32{80, 90, 297},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?sublocation=26", // macalania woods
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{80, 86},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?area=89", // mi'ihen highroad - central
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{38, 45},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?distance=2&availability=pre-story",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{57, 58, 92, 119, 194},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?repeatable=true&capture=false&has_overdrive=true",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{229, 236, 299},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?underwater=true&category=bOss",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{5, 71, 291},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?zombie=true&species=wyRm",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{134},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?creation_area=DJose",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{60, 63, 67},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?availability=post&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   71,
			results: []int32{201, 202, 210, 227, 246, 268, 284, 307},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?availability=post&location=4",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{293},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?availability=post-game&location=4",
				expectedStatus: http.StatusOK,
			},
			count:   4,
			results: []int32{16, 17, 18, 293},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?availability=post&location=23",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{301},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?availability=post&location=25",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{170, 201, 202, 203, 208, 210, 247},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?availability=post-game&area=140", // thunder plains - north
				expectedStatus: http.StatusOK,
			},
			count:   10,
			results: []int32{72, 74, 78, 295, 296},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?availability=story&sublocation=35", // calm lands
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{149, 150},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?availability=story&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   110,
			results: []int32{1, 8, 23, 57, 109, 127, 185, 217, 226},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?availability=story&repeatable=true&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   25,
			results: []int32{7, 31, 111, 118, 130, 209},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters?availability=post-game&repeatable=false&limit=max",
				expectedStatus: http.StatusOK,
			},
			count:   22,
			results: []int32{28, 107, 165, 293, 298, 305, 307},
		},
	}

	testIDList(t, tests, apilookups.E.Monsters.Endpoint, "RetrieveMonsters", testSrv.HandleMonsters, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionMonsters(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/?ids=1&location=15",
				expectedStatus: http.StatusBadRequest,
				handler:        testSrv.HandleMonsters,
				expectedErr:    "invalid usage of parameter 'ids'. parameter 'ids' can only be used in the following format: '/api/monsters/simple?ids={id},...'",
			},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/simple?limit=max",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleMonsters,
				dontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			count:   307,
			results: []int32{1, 22, 33, 68, 123, 199, 278, 307},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/simple?ids=1-21",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleMonsters,
				dontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			count:   21,
			results: []int32{1, 7, 18, 21},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/simple?ids=1,33,22,22,22,22,68,",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleMonsters,
				dontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			count:   4,
			results: []int32{1, 22, 33, 68},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/simple?ids=1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleMonsters,
				dontCheck: map[string]bool{
					"parent resource": true,
				},
			},
			count:   21,
			results: []int32{1, 5, 16, 18, 21},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/areas/89/monsters/",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleAreas,
			},
			count:          6,
			parentResource: h.GetStrPtr("/areas/89"),
			results:        []int32{38, 39, 40, 42, 43, 45},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/areas/23/monsters/",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleAreas,
			},
			count:          4,
			parentResource: h.GetStrPtr("/areas/23"),
			results:        []int32{15, 16, 17, 18},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/areas/236/monsters/",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleAreas,
			},
			count:          21,
			next:           nil,
			parentResource: h.GetStrPtr("/areas/236"),
			results:        []int32{190, 201, 210, 239, 245, 249, 253},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations/193/monsters/",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleMonsterFormations,
			},
			count:          3,
			parentResource: h.GetStrPtr("/monster-formations/193"),
			results:        []int32{143, 145, 148},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monster-formations/140/monsters/",
				expectedStatus: http.StatusOK,
				handler:        testSrv.HandleMonsterFormations,
			},
			count:          3,
			parentResource: h.GetStrPtr("/monster-formations/140"),
			results:        []int32{104, 99, 103},
		},
		{
			TestGeneral: TestGeneral{
				requestURL:     "/api/monsters/simple?ids=1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1",
				expectedStatus: http.StatusBadRequest,
				handler:        testSrv.HandleMonsters,
				expectedErr:    "fetch limit exceeded. the maximum amount of inputs is 1000",
			},
		},
	}

	testIDList(t, tests, apilookups.E.Monsters.Endpoint, "SubsectionMonsters", nil, compareSimpleResourceLists[apiresource.Named, sectiontypes.MonsterSimple])
}
