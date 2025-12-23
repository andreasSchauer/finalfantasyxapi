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
				expectedErr:    "provided monster ID is out of range. Max ID: 307",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "monster not found: a.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/a/2",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "monster not found: a, version 2",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/a/2/3",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Wrong format. Usage: /api/monsters/{name or id}, or /api/monsters/{name}/{version}",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/1?altered-state=1",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Monster sinscale, Version 1 has no altered states",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/210?omnis-elements=iifii",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Invalid usage. omnis-elements can only be used on Seymour Omnis",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis-elements=iifii",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Invalid input. omnis-elements must contain a combination of exactly four letters. Valid letters are 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis-elements=iftw",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "Invalid letter t for omnis-elements. Use any four letter combination of 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/169?kimahri-stats=hp-1000",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid usage. kimahri-stats can only be used on biran ronso (id: 167), or yenke ronso (id: 168)",
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
				species:     "/monster-species/19",
				ctbIconType: "/ctb-icon-type/1",
				distance:    1,
				properties: []string{
					"/properties/2",
				},
				locations: []string{
					"/areas/54",
				},
				formations: []string{
					"/monster-formations/42",
				},
				baseStats: map[string]int32{
					"hp":      300,
					"defense": 120,
					"magic":   18,
					"evasion": 0,
				},
				items: &testItems{
					itemDropChance: 255,
					items: map[string]*string{
						"steal common": h.GetStrPtr("/items/27"),
						"steal rare":   h.GetStrPtr("/items/28"),
						"drop common":  h.GetStrPtr("/items/71"),
						"drop rare":    h.GetStrPtr("/items/71"),
						"bribe":        h.GetStrPtr("/items/28"),
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
					weaponAbilities: []string{
						"/auto-abilities/2",
						"/auto-abilities/6",
						"/auto-abilities/26",
					},
					armorAbilities: []string{
						"/auto-abilities/58",
					},
				},
				elemResists: []testElemResist{
					{
						element:  "/elements/1",
						affinity: "/affinities/3",
					},
					{
						element:  "/elements/2",
						affinity: "/affinities/5",
					},
					{
						element:  "/elements/3",
						affinity: "/affinities/2",
					},
					{
						element:  "/elements/4",
						affinity: "/affinities/3",
					},
					{
						element:  "/elements/5",
						affinity: "/affinities/1",
					},
				},
				statusImmunities: []string{
					"/status-conditions/1",
					"/status-conditions/4",
					"/status-conditions/14",
				},
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
				requestURL:     "/api/monsters/magic-urn/1",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"species":       true,
					"ctb icon type": true,
					"distance":      true,
				},
				expLengths: map[string]int{},
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
					items: map[string]*string{
						"steal common": h.GetStrPtr("/items/1"),
						"steal rare":   h.GetStrPtr("/items/1"),
					},
					otherItems: []string{
						"/items/9",
						"/items/64",
						"/items/7",
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/spherimorph?altered-state=1",
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
				properties: []string{
					"/properties/1",
				},
				locations: []string{
					"/areas/150",
				},
				formations: []string{
					"/monster-formations/221",
				},
				items: &testItems{
					itemDropChance: 255,
					items: map[string]*string{
						"steal common": h.GetStrPtr("/items/5"),
						"steal rare":   h.GetStrPtr("/items/6"),
						"drop common":  h.GetStrPtr("/items/82"),
						"drop rare":    h.GetStrPtr("/items/82"),
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
					weaponAbilities: []string{
						"/auto-abilities/2",
						"/auto-abilities/5",
						"/auto-abilities/6",
					},
					armorAbilities: []string{
						"/auto-abilities/55",
						"/auto-abilities/58",
						"/auto-abilities/61",
						"/auto-abilities/64",
					},
				},
				elemResists: []testElemResist{
					{
						element:  "/elements/1",
						affinity: "/affinities/5",
					},
					{
						element:  "/elements/2",
						affinity: "/affinities/5",
					},
					{
						element:  "/elements/3",
						affinity: "/affinities/5",
					},
					{
						element:  "/elements/4",
						affinity: "/affinities/2",
					},
					{
						element:  "/elements/5",
						affinity: "/affinities/5",
					},
				},
				statusImmunities: []string{
					"/status-conditions/2",
					"/status-conditions/6",
					"/status-conditions/8",
					"/status-conditions/13",
					"/status-conditions/15",
					"/status-conditions/33",
					"/status-conditions/43",
					"/status-conditions/46",
				},
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
				statusImmunities: []string{
					"/status-conditions/1",
					"/status-conditions/2",
					"/status-conditions/5",
					"/status-conditions/10",
					"/status-conditions/14",
					"/status-conditions/33",
				},
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
				properties: []string{
					"/properties/6",
					"/properties/8",
				},
				autoAbilities: []string{
					"/auto-abilities/102",
				},
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
				properties: []string{
					"/properties/8",
				},
				baseStats: map[string]int32{
					"agility": 120,
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/211?omnis-elements=ilfw",
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
						element:  "/elements/1",
						affinity: "/affinities/3",
					},
					{
						element:  "/elements/2",
						affinity: "/affinities/3",
					},
					{
						element:  "/elements/3",
						affinity: "/affinities/3",
					},
					{
						element:  "/elements/4",
						affinity: "/affinities/3",
					},
					{
						element:  "/elements/5",
						affinity: "/affinities/1",
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
						element:  "/elements/1",
						affinity: "/affinities/4",
					},
					{
						element:  "/elements/2",
						affinity: "/affinities/1",
					},
					{
						element:  "/elements/3",
						affinity: "/affinities/1",
					},
					{
						element:  "/elements/4",
						affinity: "/affinities/4",
					},
					{
						element:  "/elements/5",
						affinity: "/affinities/1",
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
						element:  "/elements/1",
						affinity: "/affinities/2",
					},
					{
						element:  "/elements/2",
						affinity: "/affinities/1",
					},
					{
						element:  "/elements/3",
						affinity: "/affinities/1",
					},
					{
						element:  "/elements/4",
						affinity: "/affinities/5",
					},
					{
						element:  "/elements/5",
						affinity: "/affinities/1",
					},
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/biran-ronso?kimahri-stats=hp-1000,strength-255,magic-255,agility-255",
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
				ronsoRages: []string{
					"/ronso-rages/4",
					"/ronso-rages/5",
					"/ronso-rages/8",
					"/ronso-rages/11",
				},
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
				ronsoRages: []string{
					"/ronso-rages/2",
					"/ronso-rages/6",
					"/ronso-rages/7",
					"/ronso-rages/9",
				},
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
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "GetMonster", i+1, testCfg.HandleMonsters)
		if correctErr {
			continue
		}

		var got Monster
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testExpectedNameVer(t, testName, tc.expNameVer, got.ID, got.Name, got.Version)

		compAPIResources(t, testCfg, testName, "species", tc.species, got.Species, tc.dontCheck)
		compAPIResources(t, testCfg, testName, "ctb icon type", tc.ctbIconType, got.CTBIconType, tc.dontCheck)
		compare(t, testName, "distance", tc.distance, got.Distance, tc.dontCheck)
		checkResAmtsInSlice(t, testName, "base stats", tc.baseStats, got.BaseStats, tc.expLengths)
		checkResAmtsInSlice(t, testName, "status resists", tc.statusResists, got.StatusResists, tc.expLengths)
		compStructPtrs(t, testName, "agility params", tc.agility, got.AgilityParameters, tc.dontCheck)
		compStructSlices(t, testName, "bribe chances", tc.bribeChances, got.BribeChances, tc.expLengths)
		testMonsterElemResists(t, testCfg, testName, tc.elemResists, got.ElemResists, tc.dontCheck)
		testMonsterAltStates(t, testCfg, testName, tc.appliedState, got.AppliedState, got.AlteredStates)

		checks := []resListTest{
			newResListTest("properties", tc.properties, got.Properties),
			newResListTest("auto-abilities", tc.autoAbilities, got.AutoAbilities),
			newResListTest("ronso rages", tc.ronsoRages, got.RonsoRages),
			newResListTest("locations", tc.locations, got.Locations),
			newResListTest("formations", tc.formations, got.Formations),
			newResListTest("status immunities", tc.statusImmunities, got.StatusImmunities),
			newResListTest("abilities", tc.abilities, got.Abilities),
		}

		testMonsterItems(t, testCfg, testName, tc.items, got.Items, &checks, tc.dontCheck)
		testMonsterEquipment(t, testName, tc.equipment, got.Equipment, &checks, tc.dontCheck)
		testResourceLists(t, testCfg, testName, checks, tc.expLengths)
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
				results: []string{
					"/monsters/94",
					"/monsters/96",
					"/monsters/113",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/yojimbo",
				expectedStatus: http.StatusMultipleChoices,
			},
			expList: expList{
				count:    3,
				results: []string{
					"/monsters/165",
					"/monsters/222",
					"/monsters/234",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/mimic",
				expectedStatus: http.StatusMultipleChoices,
			},
			expList: expList{
				count:    4,
				results: []string{
					"/monsters/249",
					"/monsters/250",
					"/monsters/251",
					"/monsters/252",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters/%3F%3F%3F",
				expectedStatus: http.StatusMultipleChoices,
			},
			expList: expList{
				count:    4,
				results: []string{
					"/monsters/68",
					"/monsters/69",
					"/monsters/108",
					"/monsters/253",
				},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "GetMultipleMonsters", i+1, testCfg.HandleMonsters)
		if correctErr {
			continue
		}

		var got NamedApiResourceList
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testAPIResourceList(t, testCfg, testName, tc.expList, got, tc.dontCheck)
	}
}



func TestRetrieveMonsters(t *testing.T) {
	tests := []struct {
		testGeneral
		expList
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental-affinities=weak",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input. usage: elemental-affinities={element}-{affinity},{element}-{affinity}",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental-affinities=weak-fire",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "unknown element 'weak' in elemental-affinities.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?resistance=50",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input. resistance parameter must be paired with status-resists parameter. usage: status-resists={status},{status},...&resistance={1-254 or immune}",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?method=steal",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid input. method parameter must be paired with item parameter. usage: item={item}&method={steal/drop/bribe/other}",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?item=asf&method=drop",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "unknown item 'asf' in item.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?ronso-rage=13",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "provided ronso rage ID 13 in ronso-rage is out of range. Max ID: 12",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?species=wywrm",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value: 'wywrm', use /api/species to see valid values",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=307",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    307,
				previous: nil,
				next:     nil,
				results: []string{
					"/monsters/1",
					"/monsters/175",
					"/monsters/238",
					"/monsters/307",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?elemental-affinities=fire-weak,water-neutral",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"next": true,
				},
			},
			expList: expList{
				count:    22,
				results: []string{
					"/monsters/11",
					"/monsters/23",
					"/monsters/64",
					"/monsters/148",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?limit=307&status-resists=darkness,poison,berserk&resistance=50",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    150,
				results: []string{
					"/monsters/3",
					"/monsters/128",
					"/monsters/188",
					"/monsters/227",
					"/monsters/249",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?item=elixir&method=drop",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    2,
				results: []string{
					"/monsters/32",
					"/monsters/91",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?auto-abilities=sos-haste,auto-haste",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    5,
				results: []string{
					"/monsters/97",
					"/monsters/146",
					"/monsters/172",
					"/monsters/211",
					"/monsters/304",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?ronso-rage=nova",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    2,
				results: []string{
					"/monsters/255",
					"/monsters/292",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?location=macalania",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    19,
				results: []string{
					"/monsters/80",
					"/monsters/90",
					"/monsters/297",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?sublocation=macalania-woods",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    7,
				results: []string{
					"/monsters/80",
					"/monsters/86",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?area=90",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    6,
				results: []string{
					"/monsters/38",
					"/monsters/45",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?distance=2&story-based=false",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    2,
				results: []string{
					"/monsters/191",
					"/monsters/289",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?repeatable=true&capture=false&has-overdrive=true",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    11,
				results: []string{
					"/monsters/229",
					"/monsters/236",
					"/monsters/299",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?underwater=true&type=boss",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    6,
				results: []string{
					"/monsters/5",
					"/monsters/71",
					"/monsters/291",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?zombie=true&species=wyrm",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    1,
				results: []string{
					"/monsters/134",
				},
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/monsters?creation-area=djose",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    7,
				results: []string{
					"/monsters/60",
					"/monsters/63",
					"/monsters/67",
				},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "RetrieveMonsters", i+1, testCfg.HandleMonsters)
		if correctErr {
			continue
		}

		var got NamedApiResourceList
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Fatalf("%s: failed to decode: %v", testName, err)
		}

		testAPIResourceList(t, testCfg, testName, tc.expList, got, tc.dontCheck)
	}
}

