package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Monster struct {
	testGeneral
	NameVer
	appliedState     *testAppliedState
	agility          *testAgilityParams
	species          database.MonsterSpecies
	ctbIconType      database.CtbIconType
	distance         int32
	properties       []int32
	autoAbilities    []int32
	ronsoRages       []int32
	areas            []int32
	formations       []int32
	baseStats        map[string]int32
	regenAmounts     []testRegenAmount
	items            *testMonItems
	bribeChances     []apitypes.BribeChance
	equipment        *testMonEquipment
	elemResists      []testElemResist
	statusImmunities []int32
	statusResists    map[string]int32
	defaultState     *testDefaultState
	abilities        []int32
}

func (e Monster) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMonsters(test test, exp Monster, got apitypes.Monster) {
	test.t.Helper()
	test.compareExpNameVer(exp.NameVer, got.ID, got.Name, got.Version)
	test.compare("species", string(exp.species), got.Species)
	test.compare("ctb icon type", string(exp.ctbIconType), got.CTBIconType)
	test.compare("distance", exp.distance, got.Distance)
	test.checkResAmtTypes("base stats", exp.baseStats, got.BaseStats)
	test.checkTestStructsInSlice("regen amounts", exp.regenAmounts, got.RegenAmounts, compareRegenAmounts)
	test.checkResAmtTypes("status resists", exp.statusResists, got.StatusResists)
	test.compTestStructPtrs("agility params", exp.agility, got.AgilityParameters, compareAgilityParams)
	test.compStructSlices("bribe chances", exp.bribeChances, got.BribeChances)
	test.compTestStructSlices("elemental resists", exp.elemResists, got.ElemResists, compareElemResists)
	test.compTestStructPtrs("applied state", exp.appliedState, got.AppliedState, compareMonsterAppliedStates)
	testMonsterDefaultState(test, exp.defaultState, got.AlteredStates)
	test.compTestStructPtrs("items", exp.items, got.Items, compareMonsterItems)
	test.compTestStructPtrs("equipment", exp.equipment, got.Equipment, compareMonsterEquipment)
	test.checkResIDsInSlice("properties", apilookups.E.Properties.Endpoint, exp.properties, got.Properties)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("ronso rages", apilookups.E.RonsoRages.Endpoint, exp.ronsoRages, got.RonsoRages)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, exp.areas, got.Areas)
	test.checkResIDsInSlice("formations", apilookups.E.MonsterFormations.Endpoint, exp.formations, got.Formations)
	test.checkResIDsInSlice("status immunities", apilookups.E.StatusConditions.Endpoint, exp.statusImmunities, got.StatusImmunities)
	test.checkResIDsInSlice("abilities", apilookups.E.Abilities.Endpoint, exp.abilities, got.Abilities)
}

type testMonItems struct {
	itemDropChance int32
	items          map[string]*testResAmount[apiresource.Typed]
	otherItems     []testPossibleItem
}

func compareMonsterItems(test test, fieldName string, exp testMonItems, got apitypes.MonsterItems) {
	test.t.Helper()
	itemMap := exp.items

	test.compare("item drop chance", exp.itemDropChance, got.DropChance)
	test.compTestStructPtrs("steal common", itemMap["steal common"], got.StealCommon, compareResAmounts)
	test.compTestStructPtrs("steal rare", itemMap["steal rare"], got.StealRare, compareResAmounts)
	test.compTestStructPtrs("drop common", itemMap["drop common"], got.DropCommon, compareResAmounts)
	test.compTestStructPtrs("drop rare", itemMap["drop rare"], got.DropRare, compareResAmounts)
	test.compTestStructPtrs("sec drop common", itemMap["sec drop common"], got.SecondaryDropCommon, compareResAmounts)
	test.compTestStructPtrs("sec drop rare", itemMap["sec drop rare"], got.SecondaryDropRare, compareResAmounts)
	test.compTestStructPtrs("bribe", itemMap["bribe"], got.Bribe, compareResAmounts)
	test.compTestStructSlices("other items", exp.otherItems, got.OtherItems, comparePossibleItems)
}

type testMonEquipment struct {
	abilitySlots      apitypes.MonsterEquipmentSlots
	attachedAbilities apitypes.MonsterEquipmentSlots
	weaponAbilities   []int32
	armorAbilities    []int32
}

func compareMonsterEquipment(test test, _ string, exp testMonEquipment, got apitypes.MonsterEquipment) {
	test.t.Helper()
	test.compStructs("ability slots", exp.abilitySlots, got.AbilitySlots)
	test.compStructs("attached abilities", exp.attachedAbilities, got.AttachedAbilities)
	test.checkResIDsInSlice("weapon abilities", apilookups.E.AutoAbilities.Endpoint, exp.weaponAbilities, got.WeaponAbilities)
	test.checkResIDsInSlice("armor abilities", apilookups.E.AutoAbilities.Endpoint, exp.armorAbilities, got.ArmorAbilities)
}
