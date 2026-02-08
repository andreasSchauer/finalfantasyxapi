package main

import (
	"fmt"
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

	compIdApiResource(test, "elements", elemEndpoint, exp.element, got.Element)
	compIdApiResource(test, "affinities", affinityEndpoint, exp.affinity, got.Affinity)
}

type testMonItems struct {
	itemDropChance int32
	items          map[string]*testItemAmount
	otherItems     []testPossibleItem
}

type testItemAmount struct {
	item   string
	amount int32
}

func compareItemAmounts(test test, exp testItemAmount, got ItemAmount) {
	compPathApiResource(test, "item amount - item", exp.item, got)
	compare(test, "item amount - amount", exp.amount, got.Amount)
}

type testPossibleItem struct {
	testItemAmount
	chance int32
}

func comparePossibleItems(test test, exp testPossibleItem, got PossibleItem) {
	compareItemAmounts(test, exp.testItemAmount, got.ItemAmount)
	compare(test, "possible item - chance", exp.chance, got.Chance)
}

type testMonEquipment struct {
	abilitySlots      MonsterEquipmentSlots
	attachedAbilities MonsterEquipmentSlots
	weaponAbilities   []int32
	armorAbilities    []int32
}

func compareMonsterItems(test test, expItems *testMonItems, gotItems *MonsterItems) {
	if test.dontCheck != nil && test.dontCheck["items"] {
		return
	}

	if !bothPtrsPresent(test, "monster items", expItems, gotItems) {
		return
	}

	exp := *expItems
	got := *gotItems
	itemMap := exp.items

	compare(test, "item drop chance", exp.itemDropChance, got.DropChance)
	compTestStructPtrs(test, "steal common", itemMap["steal common"], got.StealCommon, compareItemAmounts)
	compTestStructPtrs(test, "steal rare", itemMap["steal rare"], got.StealRare, compareItemAmounts)
	compTestStructPtrs(test, "drop common", itemMap["drop common"], got.DropCommon, compareItemAmounts)
	compTestStructPtrs(test, "drop rare", itemMap["drop rare"], got.DropRare, compareItemAmounts)
	compTestStructPtrs(test, "sec drop common", itemMap["sec drop common"], got.SecondaryDropCommon, compareItemAmounts)
	compTestStructPtrs(test, "sec drop rare", itemMap["sec drop rare"], got.SecondaryDropRare, compareItemAmounts)
	compTestStructPtrs(test, "bribe", itemMap["bribe"], got.Bribe, compareItemAmounts)
	compTestStructSlices(test, "other items", exp.otherItems, got.OtherItems, comparePossibleItems)
}

func compareMonsterEquipment(test test, expEquipment *testMonEquipment, gotEquipment *MonsterEquipment) {
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

	compareResListTests(test, []resListTest{
		rltIDs("weapon abilities", test.cfg.e.autoAbilities.endpoint, exp.weaponAbilities, got.WeaponAbilities),
		rltIDs("armor abilities", test.cfg.e.autoAbilities.endpoint, exp.armorAbilities, got.ArmorAbilities),
	})
}

func compareMonsterAppliedStates(test test, exp *testAppliedState, got *AppliedState) {
	if !bothPtrsPresent(test, "applied state", exp, got) {
		return
	}

	compare(test, "applied state condition", exp.condition, got.Condition)
	compare(test, "applied state isTemporary", exp.isTemporary, got.IsTemporary)
	compIdApiResourcePtrs(test, "applied status", test.cfg.e.statusConditions.endpoint, exp.appliedStatus, got.AppliedStatus)
}

func testMonsterDefaultState(test test, exp *testDefaultState, gotStates []AlteredState) {
	if test.dontCheck != nil && test.dontCheck["default state"] {
		return
	}

	if !defaultAndAltStatesPresent(test, exp, gotStates) {
		return
	}

	compLength(test, "altered states", len(gotStates))

	got := gotStates[0]
	if got.Condition != "default" {
		test.t.Fatalf("%s: first altered state must be default when another is applied, got: %s", test.name, got.Condition)
	}

	compare(test, "default state isTemporary", exp.IsTemporary, got.IsTemporary)
	compare(test, "def state changes length", len(exp.Changes), len(got.Changes))
	compTestStructSlices(test, "alt state changes", exp.Changes, got.Changes, compareMonsterAltStateChanges)
}

func compareMonsterAltStateChanges(test test, exp testAltStateChange, got AltStateChange) {
	desc := fmt.Sprintf("def state change: %s ", exp.AlterationType)

	compare(test, desc+"type", exp.AlterationType, string(got.AlterationType))
	compare(test, desc+"distance", exp.Distance, got.Distance)
	checkResAmtsNameVals(test, desc+"base stats", exp.BaseStats, got.BaseStats)
	checkResAmtsNameVals(test, desc+"status resists", exp.StatusResists, got.StatusResists)
	compStructPtrs(test, desc+"added status", exp.AddedStatus, got.AddedStatus)
	compIdApiResourcePtrs(test, desc+"removed status", test.cfg.e.statusConditions.endpoint, exp.RemovedStatus, got.RemovedStatus)
	compTestStructSlices(test, desc+"elemental resists", exp.ElemResists, got.ElemResists, compareMonsterElemResists)

	compareResListTests(test, []resListTest{
		rltIDs(desc+"properties", test.cfg.e.properties.endpoint, exp.Properties, got.Properties),
		rltIDs(desc+"auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.AutoAbilities, got.AutoAbilities),
		rltIDs(desc+"status immunities", test.cfg.e.statusConditions.endpoint, exp.StatusImmunities, got.StatusImmunities),
	})
}
