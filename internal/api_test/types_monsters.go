package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type expMonster struct {
	testGeneral
	expNameVer
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
	bribeChances     []api.BribeChance
	equipment        *testMonEquipment
	elemResists      []testElemResist
	statusImmunities []int32
	statusResists    map[string]int32
	defaultState     *testDefaultState
	abilities        []int32
}

func (e expMonster) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMonsters(test test, exp expMonster, got api.Monster) {
	test.t.Helper()
	test.compareExpNameVer(exp.expNameVer, got.ID, got.Name, got.Version)
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
	test.checkResIDsInSlice("properties", api.E.Properties.Endpoint, exp.properties, got.Properties)
	test.checkResIDsInSlice("auto-abilities", api.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("ronso rages", api.E.RonsoRages.Endpoint, exp.ronsoRages, got.RonsoRages)
	test.checkResIDsInSlice("areas", api.E.Areas.Endpoint, exp.areas, got.Areas)
	test.checkResIDsInSlice("formations", api.E.MonsterFormations.Endpoint, exp.formations, got.Formations)
	test.checkResIDsInSlice("status immunities", api.E.StatusConditions.Endpoint, exp.statusImmunities, got.StatusImmunities)
	test.checkResIDsInSlice("abilities", api.E.Abilities.Endpoint, exp.abilities, got.Abilities)
}

type testMonItems struct {
	itemDropChance int32
	items          map[string]*testResAmount[api.TypedAPIResource]
	otherItems     []testPossibleItem
}

func compareMonsterItems(test test, fieldName string, exp testMonItems, got api.MonsterItems) {
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
	abilitySlots      api.MonsterEquipmentSlots
	attachedAbilities api.MonsterEquipmentSlots
	weaponAbilities   []int32
	armorAbilities    []int32
}

func compareMonsterEquipment(test test, _ string, exp testMonEquipment, got api.MonsterEquipment) {
	test.t.Helper()
	test.compStructs("ability slots", exp.abilitySlots, got.AbilitySlots)
	test.compStructs("attached abilities", exp.attachedAbilities, got.AttachedAbilities)
	test.checkResIDsInSlice("weapon abilities", api.E.AutoAbilities.Endpoint, exp.weaponAbilities, got.WeaponAbilities)
	test.checkResIDsInSlice("armor abilities", api.E.AutoAbilities.Endpoint, exp.armorAbilities, got.ArmorAbilities)
}
