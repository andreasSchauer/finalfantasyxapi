package main

import (
	"slices"
)

func testMonsterElemResists(test test, exp []testElemResist, got []ElementalResist) {
	compare(test, "elemental resists length", 5, len(got))

	for i, resist := range exp {
		elemEndpoint := test.cfg.e.elements.endpoint
		affinityEndpoint := test.cfg.e.affinities.endpoint

		compAPIResourcesFromID(test, "elements", elemEndpoint, resist.element, got[i].Element)
		compAPIResourcesFromID(test, "affinities", affinityEndpoint, resist.affinity, got[i].Affinity)
	}
}

func testMonsterItems(test test, expItems *testItems, gotItems *MonsterItems, checks *[]resListTest) {
	if test.dontCheck != nil && test.dontCheck["items"] {
		return
	}

	if !bothPtrsPresent(test, "monster items", expItems, gotItems) {
		return
	}

	exp := *expItems
	got := *gotItems
	endpoint := test.cfg.e.items.endpoint
	*checks = append(*checks, newResListTestFromIDs("other items", endpoint, exp.otherItems, got.OtherItems))
	itemMap := exp.items

	compare(test, "item drop chance", exp.itemDropChance, got.DropChance)
	compResPtrsFromID(test, endpoint, "steal common", itemMap["steal common"], got.StealCommon)
	compResPtrsFromID(test, endpoint, "steal rare", itemMap["steal rare"], got.StealRare)
	compResPtrsFromID(test, endpoint, "drop common", itemMap["drop common"], got.DropCommon)
	compResPtrsFromID(test, endpoint, "drop rare", itemMap["drop rare"], got.DropRare)
	compResPtrsFromID(test, endpoint, "sec drop common", itemMap["sec drop common"], got.SecondaryDropCommon)
	compResPtrsFromID(test, endpoint, "sec drop rare", itemMap["sec drop rare"], got.SecondaryDropRare)
	compResPtrsFromID(test, endpoint, "bribe", itemMap["bribe"], got.Bribe)
}

func testMonsterEquipment(test test, expEquipment *testEquipment, gotEquipment *MonsterEquipment, checks *[]resListTest) {
	if test.dontCheck != nil && test.dontCheck["equipment"] {
		return
	}

	if !bothPtrsPresent(test, "monster equipment", expEquipment, gotEquipment) {
		return
	}

	exp := *expEquipment
	got := *gotEquipment

	if !test.dontCheck["ability slots"] {
		compStructs(test, "ability slots", exp.abilitySlots, got.AbilitySlots)
	}

	if !test.dontCheck["attached abilities"] {
		compStructs(test, "attached abilities", exp.attachedAbilities, got.AttachedAbilities)
	}

	equipChecks := []resListTest{
		newResListTestFromIDs("weapon abilities", test.cfg.e.autoAbilities.endpoint, exp.weaponAbilities, got.WeaponAbilities),
		newResListTestFromIDs("armor abilities", test.cfg.e.autoAbilities.endpoint, exp.armorAbilities, got.ArmorAbilities),
	}

	*checks = slices.Concat(*checks, equipChecks)
}

func testMonsterAltStates(test test, expState *testAppliedState, gotState *AppliedState, gotAltStates []AlteredState) {
	if !bothPtrsPresent(test, "altered states", expState, gotState) {
		return
	}

	compare(test, "applied state condition", expState.condition, gotState.Condition)

	compResPtrsFromID(test, test.cfg.e.statusConditions.endpoint, "applied status", expState.appliedStatus, gotState.AppliedStatus)

	if gotAltStates[0].Condition != "default" {
		test.t.Fatalf("%s: first altered state must be default when another is applied, got: %s", test.name, gotAltStates[0].Condition)
	}

}
