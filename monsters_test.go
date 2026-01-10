package main

import (
	"encoding/json"
	"net/http"
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func TestGetMonster(t *testing.T) {
	tests := []struct {
		testGeneral
		expNameVer
		expMonsters
	}{
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
				expectedErr:    "wrong format. usage: '/api/monsters/{name or id}', '/api/monsters/{name}/{version}', or  '/api/monsters/{id}/{subsection}'. available subsections: abilities.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?has-overdrive=true",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'has-overdrive'. parameter 'has-overdrive' can only be used with list-endpoints.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/1?altered-state=1",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "monster 'sinscale', version '1' has no altered states",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/105?altered-state=5",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '5' in 'altered-state' is out of range. max id: 1.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/210?omnis-elements=iifii",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id '210'. parameter 'omnis-elements' can only be used with ids: 211.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis-elements=iifii",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input. omnis-elements must contain a combination of exactly four letters. valid letters are 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis-elements=iftw",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid letter 't' for omnis-elements. use any four-letter-combination of 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/169?kimahri-stats=hp-1000",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid id '169'. parameter 'kimahri-stats' can only be used with ids: 167, 168.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/168?kimahri-stats=hp-100000",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "kimahri's hp can't be higher than 99999.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/167?kimahri-stats=defense-5",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid stat 'defense' in 'kimahri-stats'. 'kimahri-stats' only uses 'hp', 'strength', 'magic', 'agility'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/27",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"properties":        1,
					"auto-abilities":	 0,
					"locations":         1,
					"formations":        3,
					"base stats":        10,
					"other items":		 0,
					"weapon abilities":  3,
					"armor abilities":   1,
					"status immunities": 7,
					"status resists":    1,
					"altered states":    0,
					"abilities":         1,
				},
			},
			expNameVer: expNameVer{
				id:      27,
				name:    "yellow element",
				version: nil,
			},
			expMonsters: expMonsters{
				agility: &AgilityParams{
					TickSpeed: 16,
					MinICV:    h.GetInt32Ptr(48),
					MaxICV:    h.GetInt32Ptr(53),
				},
				species:     19,
				ctbIconType: 1,
				distance:    1,
				properties: []int32{2},
				locations: []int32{54},
				formations: []int32{42},
				baseStats: map[string]int32{
					"hp":      300,
					"defense": 120,
					"magic":   18,
					"evasion": 0,
				},
				items: &testItems{
					itemDropChance: 255,
					items: map[string]*int32{
						"steal common": h.GetInt32Ptr(27),
						"steal rare":   h.GetInt32Ptr(28),
						"drop common":  h.GetInt32Ptr(71),
						"drop rare":    h.GetInt32Ptr(71),
						"bribe":        h.GetInt32Ptr(28),
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
				equipment: &testEquipment{
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
					armorAbilities: []int32{58},
				},
				elemResists: []testElemResist{
					{ element:  1, affinity: 3 },
					{ element:  2, affinity: 5 },
					{ element:  3, affinity: 2 },
					{ element:  4, affinity: 3 },
					{ element:  5, affinity: 1 },
				},
				statusImmunities: []int32{1, 4, 14},
				statusResists: map[string]int32{
					"silence": 20,
				},
				abilities: []string{
					"/player-abilities/76",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/magIc-urn/1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"species":       true,
					"ctb icon type": true,
					"distance":      true,
				},
				expLengths: map[string]int{
					"other items": 3,
				},
			},
			expNameVer: expNameVer{
				id:      156,
				name:    "magic urn",
				version: h.GetInt32Ptr(1),
			},
			expMonsters: expMonsters{
				agility: nil,
				items: &testItems{
					itemDropChance: 0,
					items: map[string]*int32{
						"steal common": h.GetInt32Ptr(1),
						"steal rare":   h.GetInt32Ptr(1),
					},
					otherItems: []int32{9, 64, 7},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/sphErimorph?altered-state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"species":       true,
					"ctb icon type": true,
					"distance":      true,
				},
				expLengths: map[string]int{
					"properties":        1,
					"auto-abilities": 	 0,
					"locations":         1,
					"formations":        1,
					"base stats":        10,
					"other items":		 0,
					"weapon abilities":  5,
					"armor abilities":   4,
					"status immunities": 16,
					"status resists":    1,
					"altered states":    4,
					"abilities":         11,
				},
			},
			expNameVer: expNameVer{
				id:      86,
				name:    "spherimorph",
				version: nil,
			},
			expMonsters: expMonsters{
				agility: &AgilityParams{
					TickSpeed: 12,
					MinICV:    h.GetInt32Ptr(36),
					MaxICV:    h.GetInt32Ptr(40),
				},
				appliedState: &testAppliedState{
					condition: "Fire-elemental.",
					isTemporary: false,
				},
				properties: []int32{1},
				locations: []int32{150},
				formations: []int32{220},
				items: &testItems{
					itemDropChance: 255,
					items: map[string]*int32{
						"steal common": h.GetInt32Ptr(5),
						"steal rare":   h.GetInt32Ptr(6),
						"drop common":  h.GetInt32Ptr(82),
						"drop rare":    h.GetInt32Ptr(82),
						"bribe":        nil,
					},
				},
				bribeChances: nil,
				equipment: &testEquipment{
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
					armorAbilities: []int32{55, 58, 61, 64},
				},
				elemResists: []testElemResist{
					{
						element:  1,
						affinity: 5,
					},
					{
						element:  2,
						affinity: 5,
					},
					{
						element:  3,
						affinity: 5,
					},
					{
						element:  4,
						affinity: 2,
					},
					{
						element:  5,
						affinity: 5,
					},
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
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/105?altered-state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params": true,
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"items":		  true,
					"equipment": 	  true,
				},
				expLengths: map[string]int{
					"status immunities": 10,
					"status resists":    2,
					"altered states":    1,
				},
			},
			expNameVer: expNameVer{
				id:      105,
				name:    "sand worm",
				version: nil,
			},
			expMonsters: expMonsters{
				appliedState: &testAppliedState{
					condition: "While 'Readying Quake'.",
					isTemporary: true,
				},
				bribeChances: nil,
				statusImmunities: []int32{1, 2, 5, 10, 14, 33},
				statusResists: map[string]int32{
					"darkness": 50,
					"power break": 50,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/neslug?altered-state=1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params": true,
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"items":		  true,
					"equipment": 	  true,
				},
				expLengths: map[string]int{
					"properties": 		 2,
					"auto-abilities":    1,
					"altered states":    2,
				},
			},
			expNameVer: expNameVer{
				id:      287,
				name:    "neslug",
				version: nil,
			},
			expMonsters: expMonsters{
				appliedState: &testAppliedState{
					condition: "While hidden in its shell.",
					isTemporary: true,
				},
				properties: []int32{6, 8},
				autoAbilities: []int32{102},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/neslug?altered-state=2",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"items":		  true,
					"equipment": 	  true,
				},
				expLengths: map[string]int{
					"properties": 		 1,
					"auto-abilities":    0,
					"altered states":    2,
				},
			},
			expNameVer: expNameVer{
				id:      287,
				name:    "neslug",
				version: nil,
			},
			expMonsters: expMonsters{
				appliedState: &testAppliedState{
					condition: "Without its shell.",
					isTemporary: false,
				},
				agility: &AgilityParams{
					TickSpeed: 4,
					MinICV: h.GetInt32Ptr(12),
					MaxICV: h.GetInt32Ptr(13),
				},
				properties: []int32{8},
				baseStats: map[string]int32{
					"agility": 120,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis-elements=iLfw",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params": true,
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"items": 		  true,
					"equipment": 	  true,
				},
				expLengths: map[string]int{},
			},
			expNameVer: expNameVer{
				id:      211,
				name:    "seymour omnis",
				version: nil,
			},
			expMonsters: expMonsters{
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
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis-elements=iiff",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params": true,
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"items": 		  true,
					"equipment": 	  true,
				},
				expLengths: map[string]int{},
			},
			expNameVer: expNameVer{
				id:      211,
				name:    "seymour omnis",
				version: nil,
			},
			expMonsters: expMonsters{
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
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis-elements=iiii",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"agility params": true,
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"items": 		  true,
					"equipment": 	  true,
				},
				expLengths: map[string]int{},
			},
			expNameVer: expNameVer{
				id:      211,
				name:    "seymour omnis",
				version: nil,
			},
			expMonsters: expMonsters{
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
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/biran-ronso?kimahri-stats=hP-1000,strEngth-255,mAgic-255,agIlity-255",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"items":		  true,
					"equipment": 	  true,
				},
				expLengths: map[string]int{
					"ronso rages": 4,
				},
			},
			expNameVer: expNameVer{
				id:      167,
				name:    "biran ronso",
				version: nil,
			},
			expMonsters: expMonsters{
				agility: &AgilityParams{
					TickSpeed: 3,
					MinICV: h.GetInt32Ptr(9),
					MaxICV: h.GetInt32Ptr(10),
				},
				ronsoRages: []int32{4, 5, 8, 11},
				baseStats: map[string]int32{
					"hp": 3549664,
					"strength": 12,
					"magic": 4,
					"agility": 251,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/yenke-ronso?kimahri-stats=hp-3500,strength-35,magic-45,agility-28",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"items":		  true,
					"equipment": 	  true,
				},
				expLengths: map[string]int{
					"ronso rages": 4,
				},
			},
			expNameVer: expNameVer{
				id:      168,
				name:    "yenke ronso",
				version: nil,
			},
			expMonsters: expMonsters{
				agility: &AgilityParams{
					TickSpeed: 10,
					MinICV: h.GetInt32Ptr(30),
					MaxICV: h.GetInt32Ptr(33),
				},
				ronsoRages: []int32{2, 6, 7, 9},
				baseStats: map[string]int32{
					"hp": 10902,
					"strength": 13,
					"magic": 22,
					"agility": 22,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/yenke-ronso?kimahri-stats=hp-1500",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"items":		  true,
					"equipment": 	  true,
				},
				expLengths: map[string]int{},
			},
			expNameVer: expNameVer{
				id:      168,
				name:    "yenke ronso",
				version: nil,
			},
			expMonsters: expMonsters{
				agility: &AgilityParams{
					TickSpeed: 26,
					MinICV: h.GetInt32Ptr(84),
					MaxICV: h.GetInt32Ptr(93),
				},
				baseStats: map[string]int32{
					"hp": 870,
					"strength": 8,
					"magic": 12,
					"agility": 1,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/seymour",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"species":        true,
					"ctb icon type":  true,
					"distance":       true,
					"items":		  true,
					"equipment": 	  true,
				},
				expLengths: map[string]int{},
			},
			expNameVer: expNameVer{
				id:      93,
				name:    "seymour",
				version: nil,
			},
			expMonsters: expMonsters{
				agility: &AgilityParams{
					TickSpeed: 10,
					MinICV: h.GetInt32Ptr(-1),
					MaxICV: h.GetInt32Ptr(-1),
				},
				autoAbilities: []int32{3},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "GetMonster", i+1, testCfg.HandleMonsters)
		if correctErr {
			continue
		}

		test := test{
			t: t,
			cfg: testCfg,
			name: testName,
			expLengths: tc.expLengths,
			dontCheck: tc.dontCheck,
		}

		var got Monster
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testExpectedNameVer(test, tc.expNameVer, got.ID, got.Name, got.Version)

		compAPIResourcesFromID(test, "species", testCfg.e.monsterSpecies.endpoint, tc.species, got.Species)
		compAPIResourcesFromID(test, "ctb icon type", testCfg.e.ctbIconType.endpoint, tc.ctbIconType, got.CTBIconType)
		compare(test, "distance", tc.distance, got.Distance)
		checkResAmtsInSlice(test, "base stats", tc.baseStats, got.BaseStats)
		checkResAmtsInSlice(test, "status resists", tc.statusResists, got.StatusResists)
		compStructPtrs(test, "agility params", tc.agility, got.AgilityParameters)
		compStructSlices(test, "bribe chances", tc.bribeChances, got.BribeChances)
		testMonsterElemResists(test, tc.elemResists, got.ElemResists)
		testMonsterAltStates(test, tc.appliedState, got.AppliedState, got.AlteredStates)

		checks := []resListTest{
			newResListTestFromIDs("properties", testCfg.e.properties.endpoint, tc.properties, got.Properties),
			newResListTestFromIDs("auto-abilities", testCfg.e.autoAbilities.endpoint, tc.autoAbilities, got.AutoAbilities),
			newResListTestFromIDs("ronso rages", testCfg.e.ronsoRages.endpoint, tc.ronsoRages, got.RonsoRages),
			newResListTestFromIDs("locations", testCfg.e.areas.endpoint, tc.locations, got.Locations),
			newResListTestFromIDs("formations", testCfg.e.monsterFormations.endpoint, tc.formations, got.Formations),
			newResListTestFromIDs("status immunities", testCfg.e.statusConditions.endpoint, tc.statusImmunities, got.StatusImmunities),
			newResListTest("abilities", tc.abilities, got.Abilities),
		}

		testMonsterItems(test, tc.items, got.Items, &checks)
		testMonsterEquipment(test, tc.equipment, got.Equipment, &checks)
		testResourceLists(test, checks)
	}
}



func TestGetMultipleMonsters(t *testing.T) {
	tests := []struct {
		testGeneral
		expList
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/guado-guardian",
				expectedStatus: http.StatusMultipleChoices,
			},
			expList: expList{
				count:    3,
				results: []int32{94, 96, 113},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/yojimbo",
				expectedStatus: http.StatusMultipleChoices,
			},
			expList: expList{
				count:    3,
				results: []int32{165, 222, 234},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/mimic",
				expectedStatus: http.StatusMultipleChoices,
			},
			expList: expList{
				count:    4,
				results: []int32{249, 250, 251, 252},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/%3F%3F%3F",
				expectedStatus: http.StatusMultipleChoices,
			},
			expList: expList{
				count:    4,
				results: []int32{68, 69, 108, 253},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "GetMultipleMonsters", i+1, testCfg.HandleMonsters)
		if correctErr {
			continue
		}

		test := test{
			t: t,
			cfg: testCfg,
			name: testName,
			expLengths: tc.expLengths,
			dontCheck: tc.dontCheck,
		}

		var got NamedApiResourceList
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testAPIResourceList(test, testCfg.e.monsters.endpoint, tc.expList, got)
	}
}



func TestRetrieveMonsters(t *testing.T) {
	tests := []struct {
		testGeneral
		expList
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=asd",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'asd' for parameter 'limit'. usage: '?limit{integer or 'max'}'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental-affinities=weak",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input. usage: '?elemental-affinities={element_name/id}-{affinity_name/id},{element_name/id}-{affinity_name/id},...'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/?omnis-elements=ffff",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'omnis-elements'. parameter 'omnis-elements' can only be used with single-resource-endpoints.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental-affinities=weak-fire",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "unknown element 'weak' in 'elemental-affinities'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental-affinities=fire-weak,fire-neutral",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "duplicate use of id '1' in 'elemental-affinities'. each element can only be used once.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?status-resists=4&resistance=350",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value '350'. 'resistance' must be an integer ranging from 1 to 254.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?status-resists=4&resistance=frank",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'frank' for parameter 'resistance'. usage: 'status-resists={status_condition_id},{status_condition_id},...&resistance={1-254 or 'immune'}'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?resistance=50",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'resistance'. parameter 'resistance' can only be used in combination with parameter(s): status-resists.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?method=steal",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage of parameter 'method'. parameter 'method' can only be used in combination with parameter(s): item.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?item=22&method=steals",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid method value: 'steals'. allowed values: 'steal', 'drop', 'bribe', 'other'.",
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
				requestURL:     "/api/monsters?auto-abilities=1,4,4,1,3,3,4",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "duplicate use of id '4' in 'auto-abilities'. each id can only be used once.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?distance=5",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value '5'. 'distance' must be an integer ranging from 1 to 4.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?distance=frank",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value 'frank' for parameter 'distance'. usage: '?distance={value}'.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?ronso-rage=13",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided id '13' in 'ronso-rage' is out of range. max id: 12.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?species=wywrm",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid enum value: 'wywrm'. use /api/species to see valid values.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    307,
				previous: nil,
				next:     nil,
				results: []int32{1, 175, 238, 307},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental-affinities=FIre-weAk,water-neutral",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"next": true,
				},
			},
			expList: expList{
				count:    22,
				results: []int32{11, 23, 64, 148},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max&status-resists=38",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    43,
				results: []int32{32, 127, 211, 233, 295},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max&status-resists=1,4,11&resistance=50",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    150,
				results: []int32{3, 128, 188, 227, 249},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max&status-resists=4&resistance=immune",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    163,
				results: []int32{5, 67, 100, 151, 258},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=max&item=7",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    22,
				results: []int32{32, 91, 156, 192, 295, 305},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?item=7&method=drop",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    2,
				results: []int32{32, 91},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?auto-abilities=96,101",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    5,
				results: []int32{97, 146, 172, 211, 304},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?ronso-rage=12",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    2,
				results: []int32{255, 292},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?location=15",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    19,
				results: []int32{80, 90, 297},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?sublocation=25",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    7,
				results: []int32{80, 86},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?area=90",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    6,
				results: []int32{38, 45},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?distance=2&story-based=false",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    2,
				results: []int32{191, 289},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?repeatable=true&capture=false&has-overdrive=true",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    11,
				results: []int32{229, 236, 299},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?underwater=true&type=bOss",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    6,
				results: []int32{5, 71, 291},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?zombie=true&species=wyRm",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    1,
				results: []int32{134},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?creation-area=DJose",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    7,
				results: []int32{60, 63, 67},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "RetrieveMonsters", i+1, testCfg.HandleMonsters)
		if correctErr {
			continue
		}

		test := test{
			t: t,
			cfg: testCfg,
			name: testName,
			expLengths: tc.expLengths,
			dontCheck: tc.dontCheck,
		}

		var got NamedApiResourceList
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testAPIResourceList(test, testCfg.e.monsters.endpoint, tc.expList, got)
	}
}

