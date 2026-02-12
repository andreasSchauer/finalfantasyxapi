package api

import (
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetMonster(t *testing.T) {
	tests := []expMonster{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/308",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "monster with provided id '308' doesn't exist. max id: 307.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "monster not found: 'a'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/a/2",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "monster not found: 'a', version '2'",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/a/2/3",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "wrong format. usage: '/api/monsters', '/api/monsters/{id}', '/api/monsters/{name}', '/api/monsters/{name}/{version}', '/api/monsters/id/{subsection}'. supported subsections: 'areas', 'monster-formations'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?has_overdrive=true",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'has_overdrive'. parameter 'has_overdrive' can only be used with list-endpoints.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/1?altered_state=1",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'sinscale', version '1' has no altered states.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/105?altered_state=5",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '5' used for parameter 'altered_state' is out of range. max id: 1.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/210?omnis_elements=iifii",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id '210'. parameter 'omnis_elements' can only be used with ids: '211'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis_elements=iifii",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input. omnis_elements must contain a combination of exactly four letters. valid letters are 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis_elements=iftw",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid letter 't' for omnis_elements. use any four-letter-combination of 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/169?kimahri_stats=hp-1000",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id '169'. parameter 'kimahri_stats' can only be used with ids: '167', '168'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/168?kimahri_stats=hp-100000",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "hp in 'kimahri_stats' can't be higher than 99999.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/167?kimahri_stats=defense-5",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid stat 'defense' in 'kimahri_stats'. 'kimahri_stats' only uses 'hp', 'strength', 'magic', 'agility'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/1?aeon_stats=hp-200",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id '1'. parameter 'aeon_stats' can only be used with ids: '216', '217', '218', '219', '220', '221', '222', '223', '224', '225'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/216?aeon_stats=mp-300",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid stat 'mp' in 'aeon_stats'. 'aeon_stats' only uses 'hp', 'strength', 'defense', 'magic', 'magic defense', 'agility', 'evasion', 'accuracy'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/216?aeon_stats=hp-999999",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "hp in 'aeon_stats' can't be higher than 99999.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/216?aeon_stats=evasion-999999",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "evasion in 'aeon_stats' can't be higher than 255.",
			},
		},
		{
			testGeneral: testGeneral{
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
			expNameVer: newExpNameVer(27, "yellow element", 0),
			agility: &AgilityParams{
				TickSpeed: 16,
				MinICV:    h.GetInt32Ptr(48),
				MaxICV:    h.GetInt32Ptr(53),
			},
			species:       19,
			ctbIconType:   1,
			distance:      1,
			properties:    []int32{2},
			autoAbilities: []int32{},
			areas:         []int32{54},
			formations:    []int32{30, 31, 32},
			baseStats: map[string]int32{
				"hp":      300,
				"defense": 120,
				"magic":   18,
				"evasion": 0,
			},
			items: &testMonItems{
				itemDropChance: 255,
				items: map[string]*testItemAmount{
					"steal common": h.GetStructPtr(newTestItemAmount("/items/27", 1)),
					"steal rare":   h.GetStructPtr(newTestItemAmount("/items/28", 1)),
					"drop common":  h.GetStructPtr(newTestItemAmount("/items/71", 1)),
					"drop rare":    h.GetStructPtr(newTestItemAmount("/items/71", 1)),
					"bribe":        h.GetStructPtr(newTestItemAmount("/items/28", 6)),
				},
			},
			bribeChances: []BribeChance{
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
			equipment: &testMonEquipment{
				abilitySlots: MonsterEquipmentSlots{
					MinAmount: 1,
					MaxAmount: 2,
					Chances: []EquipmentSlotsChance{
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
				attachedAbilities: MonsterEquipmentSlots{
					MinAmount: 0,
					MaxAmount: 2,
					Chances: []EquipmentSlotsChance{
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
			elemResists: []testElemResist{
				{element: 1, affinity: 3},
				{element: 2, affinity: 5},
				{element: 3, affinity: 2},
				{element: 4, affinity: 3},
				{element: 5, affinity: 1},
			},
			statusImmunities: []int32{1, 4, 14},
			statusResists: map[string]int32{
				"silence": 20,
			},
			abilities: []string{
				"/player-abilities/76",
			},
		},
		{
			testGeneral: testGeneral{
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
			expNameVer: newExpNameVer(156, "magic urn", 1),
			agility:    nil,
			items: &testMonItems{
				itemDropChance: 0,
				items: map[string]*testItemAmount{
					"steal common": h.GetStructPtr(newTestItemAmount("/items/1", 1)),
					"steal rare":   h.GetStructPtr(newTestItemAmount("/items/1", 1)),
				},
				otherItems: []testPossibleItem{
					newTestPossibleItem(0, "/items/9", 1, 60),
					newTestPossibleItem(1, "/items/64", 1, 20),
					newTestPossibleItem(2, "/items/7", 1, 20),
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/sphErimorph?altered_state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"other items":    true,
					"auto-abilities": true,
					"ronso rages":    true,
				},
				expLengths: map[string]int{
					"properties":        1,
					"auto-abilities":    0,
					"areas":             1,
					"formations":        1,
					"base stats":        10,
					"other items":       0,
					"weapon abilities":  5,
					"armor abilities":   4,
					"status immunities": 16,
					"status resists":    1,
					"altered states":    4,
					"abilities":         11,
				},
			},
			expNameVer: newExpNameVer(86, "spherimorph", 0),
			agility: &AgilityParams{
				TickSpeed: 12,
				MinICV:    h.GetInt32Ptr(36),
				MaxICV:    h.GetInt32Ptr(40),
			},
			appliedState: &testAppliedState{
				condition:   "Fire-elemental.",
				isTemporary: false,
			},
			properties: []int32{1},
			areas:      []int32{150},
			formations: []int32{126},
			items: &testMonItems{
				itemDropChance: 255,
				items: map[string]*testItemAmount{
					"steal common": h.GetStructPtr(newTestItemAmount("/items/5", 1)),
					"steal rare":   h.GetStructPtr(newTestItemAmount("/items/6", 1)),
					"drop common":  h.GetStructPtr(newTestItemAmount("/items/82", 1)),
					"drop rare":    h.GetStructPtr(newTestItemAmount("/items/82", 1)),
					"bribe":        nil,
				},
			},
			equipment: &testMonEquipment{
				abilitySlots: MonsterEquipmentSlots{
					MinAmount: 2,
					MaxAmount: 3,
					Chances: []EquipmentSlotsChance{
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
				attachedAbilities: MonsterEquipmentSlots{
					MinAmount: 1,
					MaxAmount: 3,
					Chances: []EquipmentSlotsChance{
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
			elemResists: []testElemResist{
				{element: 1, affinity: 5},
				{element: 2, affinity: 5},
				{element: 3, affinity: 5},
				{element: 4, affinity: 2},
				{element: 5, affinity: 5},
			},
			statusImmunities: []int32{2, 6, 8, 13, 15, 33, 43, 46},
			statusResists: map[string]int32{
				"poison": 90,
			},
			abilities: []string{
				"/player-abilities/75",
				"/player-abilities/76",
				"/player-abilities/77",
				"/player-abilities/78",
				"/enemy-abilities/2",
				"/enemy-abilities/210",
				"/enemy-abilities/211",
				"/enemy-abilities/212",
				"/enemy-abilities/213",
				"/enemy-abilities/214",
				"/enemy-abilities/215",
			},
			defaultState: &testDefaultState{
				IsTemporary: false,
				Changes: []testAltStateChange{
					{
						AlterationType: "change",
						ElemResists: []testElemResist{
							{element: 1, affinity: 6},
							{element: 2, affinity: 6},
							{element: 3, affinity: 6},
							{element: 4, affinity: 6},
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/105?altered_state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params":    true,
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
			expNameVer: newExpNameVer(105, "sand worm", 0),
			appliedState: &testAppliedState{
				condition:   "While 'Readying Quake'.",
				isTemporary: true,
			},
			bribeChances:     nil,
			statusImmunities: []int32{1, 2, 5, 10, 14, 33},
			statusResists: map[string]int32{
				"darkness":    50,
				"power break": 50,
			},
			defaultState: &testDefaultState{
				IsTemporary: false,
				Changes: []testAltStateChange{
					{
						AlterationType:   "loss",
						StatusImmunities: []int32{14, 33},
					},
					{
						AlterationType: "gain",
						StatusResists: map[string]int32{
							"sleep": 80,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
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
					"properties":     2,
					"auto-abilities": 1,
					"altered states": 2,
				},
			},
			expNameVer: newExpNameVer(287, "neslug", 0),
			appliedState: &testAppliedState{
				condition:   "While hidden in its shell.",
				isTemporary: true,
			},
			properties:    []int32{6, 8},
			autoAbilities: []int32{102},
			defaultState: &testDefaultState{
				IsTemporary: false,
				Changes: []testAltStateChange{
					{
						AlterationType: "loss",
						Properties:     []int32{6},
						AutoAbilities:  []int32{102},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/neslug?altered_state=2",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
					"properties":     1,
					"auto-abilities": 0,
					"altered states": 2,
				},
			},
			expNameVer: newExpNameVer(287, "neslug", 0),
			appliedState: &testAppliedState{
				condition:   "Without its shell.",
				isTemporary: false,
			},
			agility: &AgilityParams{
				TickSpeed: 4,
				MinICV:    h.GetInt32Ptr(12),
				MaxICV:    h.GetInt32Ptr(13),
			},
			properties:    []int32{8},
			autoAbilities: []int32{},
			baseStats: map[string]int32{
				"agility": 120,
			},
			defaultState: &testDefaultState{
				IsTemporary: false,
				Changes: []testAltStateChange{
					{
						AlterationType: "change",
						BaseStats: map[string]int32{
							"agility": 43,
						},
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/evrae?altered_state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params":    true,
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
			expNameVer: newExpNameVer(114, "evrae", 0),
			appliedState: &testAppliedState{
				condition:   "When the Airship is far away.",
				isTemporary: false,
			},
			distance: 4,
			defaultState: &testDefaultState{
				IsTemporary: false,
				Changes: []testAltStateChange{
					{
						AlterationType: "change",
						Distance:       h.GetInt32Ptr(1),
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/penance?altered_state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params":    true,
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
					"status immunities": true,
					"abilities":         true,
				},
				expLengths: map[string]int{
					"altered states": 1,
				},
			},
			expNameVer: newExpNameVer(305, "penance", 0),
			appliedState: &testAppliedState{
				condition:     "When HP falls below 9000000.",
				isTemporary:   false,
				appliedStatus: h.GetInt32Ptr(22),
			},
			defaultState: &testDefaultState{
				IsTemporary: false,
				Changes: []testAltStateChange{
					{
						AlterationType: "gain",
						AutoAbilities:  []int32{99},
					},
					{
						AlterationType: "loss",
						RemovedStatus:  h.GetInt32Ptr(22),
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis_elements=iLfw",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params":    true,
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
			expNameVer: newExpNameVer(211, "seymour omnis", 0),
			elemResists: []testElemResist{
				{
					element:  1,
					affinity: 3,
				},
				{
					element:  2,
					affinity: 3,
				},
				{
					element:  3,
					affinity: 3,
				},
				{
					element:  4,
					affinity: 3,
				},
				{
					element:  5,
					affinity: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis_elements=iiff",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params":    true,
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
			expNameVer: newExpNameVer(211, "seymour omnis", 0),
			elemResists: []testElemResist{
				{
					element:  1,
					affinity: 4,
				},
				{
					element:  2,
					affinity: 1,
				},
				{
					element:  3,
					affinity: 1,
				},
				{
					element:  4,
					affinity: 4,
				},
				{
					element:  5,
					affinity: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis_elements=iiii",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params":    true,
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
			expNameVer: newExpNameVer(211, "seymour omnis", 0),
			elemResists: []testElemResist{
				{
					element:  1,
					affinity: 2,
				},
				{
					element:  2,
					affinity: 1,
				},
				{
					element:  3,
					affinity: 1,
				},
				{
					element:  4,
					affinity: 5,
				},
				{
					element:  5,
					affinity: 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/biran_ronso?kimahri_stats=hP-1000,strEngth-255,mAgic-255,agIlity-255",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
			expNameVer: newExpNameVer(167, "biran ronso", 0),
			agility: &AgilityParams{
				TickSpeed: 3,
				MinICV:    h.GetInt32Ptr(9),
				MaxICV:    h.GetInt32Ptr(10),
			},
			ronsoRages: []int32{4, 5, 8, 11},
			baseStats: map[string]int32{
				"hp":       3549664,
				"strength": 12,
				"magic":    4,
				"agility":  251,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/yenke_ronso?kimahri_stats=hp-3500,strength-35,magic-45,agility-28",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
			expNameVer: newExpNameVer(168, "yenke ronso", 0),
			agility: &AgilityParams{
				TickSpeed: 10,
				MinICV:    h.GetInt32Ptr(30),
				MaxICV:    h.GetInt32Ptr(33),
			},
			ronsoRages: []int32{2, 6, 7, 9},
			baseStats: map[string]int32{
				"hp":       10902,
				"strength": 13,
				"magic":    22,
				"agility":  22,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/yenke_ronso?kimahri_stats=hp-1500",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
			expNameVer: newExpNameVer(168, "yenke ronso", 0),
			agility: &AgilityParams{
				TickSpeed: 26,
				MinICV:    h.GetInt32Ptr(84),
				MaxICV:    h.GetInt32Ptr(93),
			},
			baseStats: map[string]int32{
				"hp":       870,
				"strength": 8,
				"magic":    12,
				"agility":  1,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/216?aeon_stats=hp-200,strength-235,agility-68,evasion-2,accuracy-150,defense-46,magic-188,magic_defense-2",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
			expNameVer: newExpNameVer(216, "valefor", 1),
			agility: &AgilityParams{
				TickSpeed: 5,
				MinICV:    h.GetInt32Ptr(-1),
				MaxICV:    h.GetInt32Ptr(-1),
			},
			baseStats: map[string]int32{
				"hp":            725,
				"mp":            1,
				"strength":      235,
				"defense":       46,
				"magic":         188,
				"magic defense": 23,
				"agility":       68,
				"luck":          1,
				"evasion":       19,
				"accuracy":      150,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/seymour",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
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
			expNameVer: newExpNameVer(93, "seymour", 0),
			agility: &AgilityParams{
				TickSpeed: 10,
				MinICV:    h.GetInt32Ptr(-1),
				MaxICV:    h.GetInt32Ptr(-1),
			},
			autoAbilities: []int32{3},
		},
	}

	testSingleResources(t, tests, "GetMonster", testCfg.HandleMonsters, compareMonsters)
}

func TestGetMultipleMonsters(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/guado_guardian",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   3,
			results: []int32{94, 96, 113},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/yojimbo",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   3,
			results: []int32{165, 222, 234},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/mimic",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   4,
			results: []int32{249, 250, 251, 252},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/%3F%3F%3F",
				expectedStatus: http.StatusMultipleChoices,
			},
			count:   4,
			results: []int32{68, 69, 108, 253},
		},
	}

	testIdList(t, tests, testCfg.e.monsters.endpoint, "GetMultipleMonsters", testCfg.HandleMonsters, compareAPIResourceLists[NamedApiResourceList])
}

func TestRetrieveMonsters(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=asd",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'asd' used for parameter 'limit'. usage: '?limit{int|'max'}'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental_resists=weak",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input for parameter 'elemental_resists'. usage: '?elemental_resists={element|id}-{affinity|id},...'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/?omnis_elements=ffff",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'omnis_elements'. parameter 'omnis_elements' can only be used with single-resource-endpoints.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental_resists=weak-fire",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "unknown element 'weak' used for parameter 'elemental_resists'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental_resists=fire-weak,fire-neutral",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "duplicate use of id '1' for parameter 'elemental_resists'. each element can only be used once.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?status_resists=4&resistance=350",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value '350' used for parameter 'resistance'. value must be an integer ranging from 1 to 254.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?status_resists=4&resistance=frank",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'frank' used for parameter 'resistance'. usage: 'status_resists={id},...&resistance={int|'immune'}'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?resistance=50",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'resistance'. parameter 'resistance' can only be used in combination with parameter(s): 'status_resists'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?method=steal",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'method'. parameter 'method' can only be used in combination with parameter(s): 'item'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?item=22&method=steals",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'steals' used for 'method'. allowed values: 'steal', 'drop', 'bribe', 'other'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?item=asf&method=drop",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id 'asf' used for parameter 'item'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?auto_abilities=1,4,4,1,3,3,4",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "duplicate use of id '4' for parameter 'auto_abilities'. each id can only be used once.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?distance=5",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value '5' used for parameter 'distance'. value must be an integer ranging from 1 to 4.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?distance=frank",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'frank' used for parameter 'distance'. usage: '?distance={int}'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?ronso_rage=13",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '13' used for parameter 'ronso_rage' is out of range. max id: 12.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?species=wywrm",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid enum value 'wywrm' used for parameter 'species'. use /api/monsters/parameters to see allowed values.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max",
				expectedStatus: http.StatusOK,
			},
			count:    307,
			previous: nil,
			next:     nil,
			results:  []int32{1, 175, 238, 307},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental_resists=FIre-weAk,water-neutral",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"next": true,
				},
			},
			count:   22,
			results: []int32{11, 23, 64, 148},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max&status_resists=38",
				expectedStatus: http.StatusOK,
			},
			count:   43,
			results: []int32{32, 127, 211, 233, 295},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max&status_resists=1,4,11&resistance=50",
				expectedStatus: http.StatusOK,
			},
			count:   150,
			results: []int32{3, 128, 188, 227, 249},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max&status_resists=4&resistance=immune",
				expectedStatus: http.StatusOK,
			},
			count:   163,
			results: []int32{5, 67, 100, 151, 258},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max&item=7",
				expectedStatus: http.StatusOK,
			},
			count:   22,
			results: []int32{32, 91, 156, 192, 295, 305},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?item=7&method=drop",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{32, 91},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?auto_abilities=96,101",
				expectedStatus: http.StatusOK,
			},
			count:   5,
			results: []int32{97, 146, 172, 211, 304},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?ronso_rage=12",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{255, 292},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?location=15",
				expectedStatus: http.StatusOK,
			},
			count:   19,
			results: []int32{80, 90, 297},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?sublocation=25",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{80, 86},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?area=90",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{38, 45},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?distance=2&story_based=false",
				expectedStatus: http.StatusOK,
			},
			count:   2,
			results: []int32{191, 289},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?repeatable=true&capture=false&has_overdrive=true",
				expectedStatus: http.StatusOK,
			},
			count:   11,
			results: []int32{229, 236, 299},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?underwater=true&type=bOss",
				expectedStatus: http.StatusOK,
			},
			count:   6,
			results: []int32{5, 71, 291},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?zombie=true&species=wyRm",
				expectedStatus: http.StatusOK,
			},
			count:   1,
			results: []int32{134},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?creation_area=DJose",
				expectedStatus: http.StatusOK,
			},
			count:   7,
			results: []int32{60, 63, 67},
		},
	}

	testIdList(t, tests, testCfg.e.monsters.endpoint, "RetrieveMonsters", testCfg.HandleMonsters, compareAPIResourceLists[NamedApiResourceList])
}

func TestSubsectionMonsters(t *testing.T) {
	tests := []expListIDs{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/90/monsters/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          6,
			parentResource: h.GetStrPtr("/areas/90"),
			results:        []int32{38, 39, 40, 42, 43, 45},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/23/monsters/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          4,
			parentResource: h.GetStrPtr("/areas/23"),
			results:        []int32{15, 16, 17, 18},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/areas/239/monsters/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleAreas,
			},
			count:          21,
			next:           h.GetStrPtr("/areas/239/monsters?limit=20&offset=20"),
			parentResource: h.GetStrPtr("/areas/239"),
			results:        []int32{190, 201, 210, 239, 245, 249, 253},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/193/monsters/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleMonsterFormations,
			},
			count:          3,
			parentResource: h.GetStrPtr("/monster-formations/193"),
			results:        []int32{143, 145, 148},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monster-formations/140/monsters/",
				expectedStatus: http.StatusOK,
				handler:        testCfg.HandleMonsterFormations,
			},
			count:          3,
			parentResource: h.GetStrPtr("/monster-formations/140"),
			results:        []int32{104, 99, 103},
		},
	}

	testIdList(t, tests, testCfg.e.monsters.endpoint, "SubsectionMonsters", nil, compareSubResourceLists[NamedAPIResource, MonsterSub])
}
