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

func (e Monster) Compare(test test, got apitypes.Monster) {
	test.t.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("species", string(e.species), got.Species)
	test.compare("ctb icon type", string(e.ctbIconType), got.CTBIconType)
	test.compare("distance", e.distance, got.Distance)
	test.checkResAmtTypes("base stats", e.baseStats, got.BaseStats)
	test.checkTestStructsInSlice("regen amounts", e.regenAmounts, got.RegenAmounts)
	test.checkResAmtTypes("status resists", e.statusResists, got.StatusResists)
	test.compTestStructPtrs("agility params", e.agility, got.AgilityParameters)
	test.compStructSlices("bribe chances", e.bribeChances, got.BribeChances)
	test.compTestStructSlices("elemental resists", e.elemResists, got.ElemResists)
	test.compTestStructPtrs("applied state", e.appliedState, got.AppliedState)
	testMonsterDefaultState(test, e.defaultState, got.AlteredStates)
	test.compTestStructPtrs("items", e.items, got.Items)
	test.compTestStructPtrs("equipment", e.equipment, got.Equipment)
	test.checkResIDsInSlice("properties", apilookups.E.Properties.Endpoint, e.properties, got.Properties)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("ronso rages", apilookups.E.RonsoRages.Endpoint, e.ronsoRages, got.RonsoRages)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, e.areas, got.Areas)
	test.checkResIDsInSlice("formations", apilookups.E.MonsterFormations.Endpoint, e.formations, got.Formations)
	test.checkResIDsInSlice("status immunities", apilookups.E.StatusConditions.Endpoint, e.statusImmunities, got.StatusImmunities)
	test.checkResIDsInSlice("abilities", apilookups.E.Abilities.Endpoint, e.abilities, got.Abilities)
}

type testMonItems struct {
	itemDropChance int32
	items          map[string]*testResAmount[apiresource.Typed]
	otherItems     []testPossibleItem
}

func (e testMonItems) Compare(test test, fieldName string, got apitypes.MonsterItems) {
	test.t.Helper()
	itemMap := e.items

	test.compare("item drop chance", e.itemDropChance, got.DropChance)
	test.compTestStructPtrs("steal common", itemMap["steal common"], got.StealCommon)
	test.compTestStructPtrs("steal rare", itemMap["steal rare"], got.StealRare)
	test.compTestStructPtrs("drop common", itemMap["drop common"], got.DropCommon)
	test.compTestStructPtrs("drop rare", itemMap["drop rare"], got.DropRare)
	test.compTestStructPtrs("sec drop common", itemMap["sec drop common"], got.SecondaryDropCommon)
	test.compTestStructPtrs("sec drop rare", itemMap["sec drop rare"], got.SecondaryDropRare)
	test.compTestStructPtrs("bribe", itemMap["bribe"], got.Bribe)
	test.compTestStructSlices("other items", e.otherItems, got.OtherItems)
}

type testMonEquipment struct {
	abilitySlots      apitypes.MonsterEquipmentSlots
	attachedAbilities apitypes.MonsterEquipmentSlots
	weaponAbilities   []int32
	armorAbilities    []int32
}

func (e testMonEquipment) Compare(test test, _ string, got apitypes.MonsterEquipment) {
	test.t.Helper()
	test.compStructs("ability slots", e.abilitySlots, got.AbilitySlots)
	test.compStructs("attached abilities", e.attachedAbilities, got.AttachedAbilities)
	test.checkResIDsInSlice("weapon abilities", apilookups.E.AutoAbilities.Endpoint, e.weaponAbilities, got.WeaponAbilities)
	test.checkResIDsInSlice("armor abilities", apilookups.E.AutoAbilities.Endpoint, e.armorAbilities, got.ArmorAbilities)
}
