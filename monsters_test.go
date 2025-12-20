package main

import (
	"encoding/json"
	"net/http"
	"slices"
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
				requestURL:     "/api/overdrive-modes/ally/2",
				expectedStatus: http.StatusBadRequest,
				expectedErr:    `Wrong format. Usage: /api/overdrive-modes/{name or id}`,
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/18",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "provided overdrive-mode ID is out of range. Max ID: 17",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/a",
				expectedStatus: http.StatusNotFound,
				expectedErr:    "overdrive-mode not found: a.",
			},
		},
		{
			testGeneral: testGeneral{
				requestURL:     "/api/overdrive-modes/ally/",
				expectedStatus: http.StatusOK,
				dontCheck: map[string]bool{
					"effect": true,
				},
				expLengths: map[string]int{
					"actions": 7,
				},
			},
			expNameVer: expNameVer{
				id:      145,
				name:    "north",
				version: h.GetInt32Ptr(1),
			},
			expMonsters: expMonsters{},
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

		compareResources(t, testCfg, testName, "species", tc.species, got.Species, tc.dontCheck)
		compareResources(t, testCfg, testName, "ctb icon type", tc.ctbIconType, got.CTBIconType, tc.dontCheck)
		compare(t, testName, "distance", tc.distance, got.Distance, tc.dontCheck)
		checkResAmtsInSlice(t, testName, "base stats", tc.baseStats, got.BaseStats, tc.expLengths)
		checkResAmtsInSlice(t, testName, "status resists", tc.statusResists, got.StatusResists, tc.expLengths)
		compStructPtrs(t, testName, "agility parameters", tc.agility, got.AgilityParameters, tc.dontCheck)


		if tc.appliedState == nil && got.AppliedState != nil {
			t.Fatalf("%s: applied state to be nil, but got %v", testName, h.DerefOrNil(got.Items))
		}

		if got.AppliedState != nil {
			compare(t, testName, "applied state condition", tc.appliedState.condition, got.AppliedState.Condition, tc.dontCheck)

			compareResourcePtrs(t, testCfg, testName, "applied status", tc.appliedState.appliedStatus, got.AppliedState.AppliedStatus, tc.dontCheck)

			if got.AlteredStates[0].Condition != "default" {
				t.Fatalf("%s: first altered state must be default when another is applied, got: %s", testName, got.AlteredStates[0].Condition)
			}
		}


		checks := []resListTest{
			newResListTest("properties", tc.properties, got.Properties),
			newResListTest("auto abilities", tc.autoAbilities, got.AutoAbilities),
			newResListTest("ronso rages", tc.ronsoRages, got.RonsoRages),
			newResListTest("locations", tc.locations, got.Locations),
			newResListTest("formations", tc.formations, got.Formations),
			newResListTest("status immunities", tc.statusImmunities, got.StatusImmunities),
			newResListTest("abilities", tc.abilities, got.Abilities),
		}

		
		if tc.items == nil && got.Items != nil {
			t.Fatalf("%s: expected items to be nil, but got %v", testName, h.DerefOrNil(got.Items))
		}

		if got.Items != nil {
			expItems := *tc.items
			gotItems := *got.Items
			checks = append(checks, newResListTest("other items", expItems.otherItems, gotItems.OtherItems))

			compare(t, testName, "item drop chance", expItems.itemDropChance, gotItems.DropChance, tc.dontCheck)
			itemMap := expItems.items
			compareResourcePtrs(t, testCfg, testName, "steal common", itemMap["steal common"], gotItems.StealCommon, tc.dontCheck)
			compareResourcePtrs(t, testCfg, testName, "steal rare", itemMap["steal rare"], gotItems.StealRare, tc.dontCheck)
			compareResourcePtrs(t, testCfg, testName, "drop common", itemMap["drop common"], gotItems.DropCommon, tc.dontCheck)
			compareResourcePtrs(t, testCfg, testName, "drop rare", itemMap["drop rare"], gotItems.DropRare, tc.dontCheck)
			compareResourcePtrs(t, testCfg, testName, "sec drop common", itemMap["sec drop common"], gotItems.SecondaryDropCommon, tc.dontCheck)
			compareResourcePtrs(t, testCfg, testName, "sec drop rare", itemMap["sec drop rare"], gotItems.SecondaryDropRare, tc.dontCheck)
			compareResourcePtrs(t, testCfg, testName, "bribe", itemMap["bribe"], gotItems.Bribe, tc.dontCheck)
		}


		if tc.equipment == nil && got.Equipment != nil {
			t.Fatalf("%s: expected equipment to be nil, but got %v", testName, h.DerefOrNil(got.Equipment))
		}

		if got.Equipment != nil {
			expEquipment := *tc.equipment
			gotEquipment := *got.Equipment

			if !tc.dontCheck["ability slots"] {
				compStructs(t, testName, "ability slots", expEquipment.abilitySlots, gotEquipment.AbilitySlots)
			}

			if !tc.dontCheck["attached abilities"] {
				compStructs(t, testName, "attached abilities", expEquipment.attachedAbilities, gotEquipment.AttachedAbilities)
			}

			equipChecks := []resListTest{
				newResListTest("weapon abilities", expEquipment.weaponAbilities, gotEquipment.WeaponAbilities),
				newResListTest("armor abilities", expEquipment.armorAbilities, gotEquipment.ArmorAbilities),
			}

			checks = slices.Concat(checks, equipChecks)
		}

		testResourceLists(t, testCfg, testName, checks, tc.expLengths)
	}
}

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
