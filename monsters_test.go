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
				requestURL:     "/api/monsters/27",
				expectedStatus: http.StatusOK,
				dontCheck:      map[string]bool{},
				expLengths: map[string]int{
					"properties":        1,
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
				bribeChances: nil,
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
		compStructPtrs(t, testName, "agility parameters", tc.agility, got.AgilityParameters, tc.dontCheck)
		compStructSlices(t, testName, "bribe chances", tc.bribeChances, got.BribeChances, tc.expLengths)
		testMonsterElemResists(t, testCfg, testName, tc.elemResists, got.ElemResists, tc.dontCheck)
		testMonsterAltStates(t, testCfg, testName, tc.appliedState, got.AppliedState, got.AlteredStates)

		checks := []resListTest{
			newResListTest("properties", tc.properties, got.Properties),
			newResListTest("auto abilities", tc.autoAbilities, got.AutoAbilities),
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




// TestGetMonsterMultiple for 300 responses

/*
func TestRetrieveMonsters(t *testing.T) {
	tests := []struct {
		testGeneral
		expList
	}{
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes?type=f",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    "invalid value: f, use /api/overdrive-mode-type to see valid values.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/",
				expectedStatus: http.StatusOK,
			},
			expList: expList{
				count:    17,
				previous: nil,
				next:     nil,
				results: []string{
					"/overdrive-modes/1",
					"/overdrive-modes/8",
					"/overdrive-modes/17",
				},
			},
		},
	}

	for i, tc := range tests {
		rr, testName, correctErr := setupTest(t, tc.testGeneral, "RetrieveOverdriveModes", i+1, testCfg.HandleOverdriveModes)
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
*/
