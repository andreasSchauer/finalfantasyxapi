package main

import (
	"fmt"
	"slices"
)

type testAppliedState struct {
	condition     string
	isTemporary   bool
	appliedStatus *int32
}

type testDefaultState struct {
	IsTemporary bool                 `json:"is_temporary"`
	Changes     []testAltStateChange `json:"changes"`
}

type testAltStateChange struct {
	AlterationType   string
	Distance         *int32
	Properties       []int32
	AutoAbilities    []int32
	BaseStats        map[string]int32
	ElemResists      []testElemResist
	StatusImmunities []int32
	StatusResists    map[string]int32
	AddedStatus      *InflictedStatus
	RemovedStatus    *int32
}

type testElemResist struct {
	element  int32
	affinity int32
}

func compareMonsterElemResists(test test, exp testElemResist, got ElementalResist) {
	elemEndpoint := test.cfg.e.elements.endpoint
	affinityEndpoint := test.cfg.e.affinities.endpoint

	compAPIResourcesFromID(test, "elements", elemEndpoint, exp.element, got.Element)
	compAPIResourcesFromID(test, "affinities", affinityEndpoint, exp.affinity, got.Affinity)
}

type testMonItems struct {
	itemDropChance int32
	items          map[string]*int32
	otherItems     []int32
}

type testMonEquipment struct {
	abilitySlots      MonsterEquipmentSlots
	attachedAbilities MonsterEquipmentSlots
	weaponAbilities   []int32
	armorAbilities    []int32
}

func testMonsterItems(test test, expItems *testMonItems, gotItems *MonsterItems, checks *[]resListTest) {
	if test.dontCheck != nil && test.dontCheck["items"] {
		return
	}

	if !bothPtrsPresent(test, "monster items", expItems, gotItems) {
		return
	}

	exp := *expItems
	got := *gotItems
	endpoint := test.cfg.e.items.endpoint
	*checks = append(*checks, rltIDs("other items", endpoint, exp.otherItems, got.OtherItems))
	itemMap := exp.items

	compare(test, "item drop chance", exp.itemDropChance, got.DropChance)
	compResPtrsFromID(test, "steal common", endpoint, itemMap["steal common"], got.StealCommon)
	compResPtrsFromID(test, "steal rare", endpoint, itemMap["steal rare"], got.StealRare)
	compResPtrsFromID(test, "drop common", endpoint, itemMap["drop common"], got.DropCommon)
	compResPtrsFromID(test, "drop rare", endpoint, itemMap["drop rare"], got.DropRare)
	compResPtrsFromID(test, "sec drop common", endpoint, itemMap["sec drop common"], got.SecondaryDropCommon)
	compResPtrsFromID(test, "sec drop rare", endpoint, itemMap["sec drop rare"], got.SecondaryDropRare)
	compResPtrsFromID(test, "bribe", endpoint, itemMap["bribe"], got.Bribe)
}

func testMonsterEquipment(test test, expEquipment *testMonEquipment, gotEquipment *MonsterEquipment, checks *[]resListTest) {
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
		rltIDs("weapon abilities", test.cfg.e.autoAbilities.endpoint, exp.weaponAbilities, got.WeaponAbilities),
		rltIDs("armor abilities", test.cfg.e.autoAbilities.endpoint, exp.armorAbilities, got.ArmorAbilities),
	}

	*checks = slices.Concat(*checks, equipChecks)
}

func testMonsterAppliedState(test test, exp *testAppliedState, got *AppliedState) {
	if !bothPtrsPresent(test, "applied state", exp, got) {
		return
	}

	compare(test, "applied state condition", exp.condition, got.Condition)
	compare(test, "applied state isTemporary", exp.isTemporary, got.IsTemporary)
	compResPtrsFromID(test, "applied status", test.cfg.e.statusConditions.endpoint, exp.appliedStatus, got.AppliedStatus)
}

func testMonsterDefaultState(test test, exp *testDefaultState, gotStates []AlteredState) {
	if test.dontCheck != nil && test.dontCheck["default state"] {
		return
	}

	if exp == nil && len(gotStates) == 0 {
		return
	}
	if exp == nil && len(gotStates) != 0 {
		test.t.Fatalf("%s: expected default state to be nil, but got alt states", test.name)
	}
	if exp != nil && len(gotStates) == 0 {
		test.t.Fatalf("%s: expected default state to be not nil, but got no alt states", test.name)
	}

	compare(test, "altered states length", test.expLengths["altered states"], len(gotStates))

	got := gotStates[0]
	if got.Condition != "default" {
		test.t.Fatalf("%s: first altered state must be default when another is applied, got: %s", test.name, got.Condition)
	}

	compare(test, "default state isTemporary", exp.IsTemporary, got.IsTemporary)
	compare(test, "def state changes length", len(exp.Changes), len(got.Changes))

	for i, expChange := range exp.Changes {
		desc := fmt.Sprintf("def state change: %s ", expChange.AlterationType)
		gotChange := got.Changes[i]

		compare(test, desc+"type", expChange.AlterationType, string(gotChange.AlterationType))
		compare(test, desc+"distance", expChange.Distance, gotChange.Distance)
		checkResAmtsInSlice(test, desc+"base stats", expChange.BaseStats, gotChange.BaseStats)
		checkResAmtsInSlice(test, desc+"status resists", expChange.StatusResists, gotChange.StatusResists)
		compStructPtrs(test, desc+"added status", expChange.AddedStatus, gotChange.AddedStatus)
		compResPtrsFromID(test, desc+"removed status", test.cfg.e.statusConditions.endpoint, expChange.RemovedStatus, gotChange.RemovedStatus)
		compareCustomObjSlices(test, desc+"elemental resists", expChange.ElemResists, gotChange.ElemResists, compareMonsterElemResists)

		checks := []resListTest{
			rltIDs(desc+"properties", test.cfg.e.properties.endpoint, expChange.Properties, gotChange.Properties),
			rltIDs(desc+"auto-abilities", test.cfg.e.autoAbilities.endpoint, expChange.AutoAbilities, gotChange.AutoAbilities),
			rltIDs(desc+"status immunities", test.cfg.e.statusConditions.endpoint, expChange.StatusImmunities, gotChange.StatusImmunities),
		}

		compareResListTests(test, checks)
	}
}
