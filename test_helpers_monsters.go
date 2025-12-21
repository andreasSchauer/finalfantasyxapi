package main

import (
	"slices"
	"testing"
)

func testMonsterElemResists(t *testing.T, testCfg *Config, testName string, exp []testElemResist, got []ElementalResist, dontCheck map[string]bool) {
	compare(t, testName, "elemental resists length", 5, len(got), dontCheck)

	for i, resist := range exp {
		compAPIResources(t, testCfg, testName, "elements", resist.element, got[i].Element, dontCheck)
		compAPIResources(t, testCfg, testName, "affinities", resist.affinity, got[i].Affinity, dontCheck)
	}
}


func testMonsterItems(t *testing.T, testCfg *Config, testName string, expItems *testItems, gotItems *MonsterItems, checks *[]resListTest, dontCheck map[string]bool) {
	if !bothPtrsPresent(t, testName, "monster items", expItems, gotItems) {
		return
	}

	exp := *expItems
	got := *gotItems
	*checks = append(*checks, newResListTest("other items", exp.otherItems, got.OtherItems))
	itemMap := exp.items

	compare(t, testName, "item drop chance", exp.itemDropChance, got.DropChance, dontCheck)
	compResourcePtrs(t, testCfg, testName, "steal common", itemMap["steal common"], got.StealCommon, dontCheck)
	compResourcePtrs(t, testCfg, testName, "steal rare", itemMap["steal rare"], got.StealRare, dontCheck)
	compResourcePtrs(t, testCfg, testName, "drop common", itemMap["drop common"], got.DropCommon, dontCheck)
	compResourcePtrs(t, testCfg, testName, "drop rare", itemMap["drop rare"], got.DropRare, dontCheck)
	compResourcePtrs(t, testCfg, testName, "sec drop common", itemMap["sec drop common"], got.SecondaryDropCommon, dontCheck)
	compResourcePtrs(t, testCfg, testName, "sec drop rare", itemMap["sec drop rare"], got.SecondaryDropRare, dontCheck)
	compResourcePtrs(t, testCfg, testName, "bribe", itemMap["bribe"], got.Bribe, dontCheck)
}

func testMonsterEquipment(t *testing.T, testName string, expEquipment *testEquipment, gotEquipment *MonsterEquipment, checks *[]resListTest, dontCheck map[string]bool) {
	if !bothPtrsPresent(t, testName, "monster equipment", expEquipment, gotEquipment) {
		return
	}

	exp := *expEquipment
	got := *gotEquipment

	if !dontCheck["ability slots"] {
		compStructs(t, testName, "ability slots", exp.abilitySlots, got.AbilitySlots)
	}

	if !dontCheck["attached abilities"] {
		compStructs(t, testName, "attached abilities", exp.attachedAbilities, got.AttachedAbilities)
	}

	equipChecks := []resListTest{
		newResListTest("weapon abilities", exp.weaponAbilities, got.WeaponAbilities),
		newResListTest("armor abilities", exp.armorAbilities, got.ArmorAbilities),
	}

	*checks = slices.Concat(*checks, equipChecks)
}


func testMonsterAltStates(t *testing.T, testCfg *Config, testName string, expState *testAppliedState, gotState *AppliedState, gotAltStates []AlteredState) {
	if !bothPtrsPresent(t, testName, "altered states", expState, gotState) {
		return
	}

	compare(t, testName, "applied state condition", expState.condition, gotState.Condition, nil)

	compResourcePtrs(t, testCfg, testName, "applied status", expState.appliedStatus, gotState.AppliedStatus, nil)

	if gotAltStates[0].Condition != "default" {
		t.Fatalf("%s: first altered state must be default when another is applied, got: %s", testName, gotAltStates[0].Condition)
	}

}