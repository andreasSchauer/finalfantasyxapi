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
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "species", string(exp.species), got.Species)
	compare(test, "ctb icon type", string(exp.ctbIconType), got.CTBIconType)
	compare(test, "distance", exp.distance, got.Distance)
	checkResAmtTypes(test, "base stats", exp.baseStats, got.BaseStats)
	checkTestStructsInSlice(test, "regen amounts", exp.regenAmounts, got.RegenAmounts, compareRegenAmounts)
	checkResAmtTypes(test, "status resists", exp.statusResists, got.StatusResists)
	compTestStructPtrs(test, "agility params", exp.agility, got.AgilityParameters, compareAgilityParams)
	compStructSlices(test, "bribe chances", exp.bribeChances, got.BribeChances)
	compTestStructSlices(test, "elemental resists", exp.elemResists, got.ElemResists, compareElemResists)
	compTestStructPtrs(test, "applied state", exp.appliedState, got.AppliedState, compareMonsterAppliedStates)
	testMonsterDefaultState(test, exp.defaultState, got.AlteredStates)
	compTestStructPtrs(test, "items", exp.items, got.Items, compareMonsterItems)
	compTestStructPtrs(test, "equipment", exp.equipment, got.Equipment, compareMonsterEquipment)
	checkResIDsInSlice(test, "properties", test.cfg.E.Properties.Endpoint, exp.properties, got.Properties)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "ronso rages", test.cfg.E.RonsoRages.Endpoint, exp.ronsoRages, got.RonsoRages)
	checkResIDsInSlice(test, "areas", test.cfg.E.Areas.Endpoint, exp.areas, got.Areas)
	checkResIDsInSlice(test, "formations", test.cfg.E.MonsterFormations.Endpoint, exp.formations, got.Formations)
	checkResIDsInSlice(test, "status immunities", test.cfg.E.StatusConditions.Endpoint, exp.statusImmunities, got.StatusImmunities)
	checkResIDsInSlice(test, "abilities", test.cfg.E.Abilities.Endpoint, exp.abilities, got.Abilities)
}

type testMonItems struct {
	itemDropChance int32
	items          map[string]*testResAmount[api.TypedAPIResource]
	otherItems     []testPossibleItem
}

func compareMonsterItems(test test, fieldName string, exp testMonItems, got api.MonsterItems) {
	test.t.Helper()
	itemMap := exp.items

	compare(test, "item drop chance", exp.itemDropChance, got.DropChance)
	compTestStructPtrs(test, "steal common", itemMap["steal common"], got.StealCommon, compareResAmounts)
	compTestStructPtrs(test, "steal rare", itemMap["steal rare"], got.StealRare, compareResAmounts)
	compTestStructPtrs(test, "drop common", itemMap["drop common"], got.DropCommon, compareResAmounts)
	compTestStructPtrs(test, "drop rare", itemMap["drop rare"], got.DropRare, compareResAmounts)
	compTestStructPtrs(test, "sec drop common", itemMap["sec drop common"], got.SecondaryDropCommon, compareResAmounts)
	compTestStructPtrs(test, "sec drop rare", itemMap["sec drop rare"], got.SecondaryDropRare, compareResAmounts)
	compTestStructPtrs(test, "bribe", itemMap["bribe"], got.Bribe, compareResAmounts)
	compTestStructSlices(test, "other items", exp.otherItems, got.OtherItems, comparePossibleItems)
}

type testMonEquipment struct {
	abilitySlots      api.MonsterEquipmentSlots
	attachedAbilities api.MonsterEquipmentSlots
	weaponAbilities   []int32
	armorAbilities    []int32
}

func compareMonsterEquipment(test test, _ string, exp testMonEquipment, got api.MonsterEquipment) {
	test.t.Helper()
	compStructs(test, "ability slots", exp.abilitySlots, got.AbilitySlots)
	compStructs(test, "attached abilities", exp.attachedAbilities, got.AttachedAbilities)
	checkResIDsInSlice(test, "weapon abilities", test.cfg.E.AutoAbilities.Endpoint, exp.weaponAbilities, got.WeaponAbilities)
	checkResIDsInSlice(test, "armor abilities", test.cfg.E.AutoAbilities.Endpoint, exp.armorAbilities, got.ArmorAbilities)
}
