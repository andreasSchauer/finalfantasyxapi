package main

type expMonster struct {
	testGeneral
	expNameVer
	appliedState     *testAppliedState
	agility          *AgilityParams
	species          int32
	ctbIconType      int32
	distance         int32
	properties       []int32
	autoAbilities    []int32
	ronsoRages       []int32
	areas            []int32
	formations       []int32
	baseStats        map[string]int32
	items            *testMonItems
	bribeChances     []BribeChance
	equipment        *testMonEquipment
	elemResists      []testElemResist
	statusImmunities []int32
	statusResists    map[string]int32
	defaultState     *testDefaultState
	abilities        []string
}

func (e expMonster) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMonsters(test test, exp expMonster, got Monster) {
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compIdApiResource(test, "species", test.cfg.e.monsterSpecies.endpoint, exp.species, got.Species)
	compIdApiResource(test, "ctb icon type", test.cfg.e.ctbIconType.endpoint, exp.ctbIconType, got.CTBIconType)
	compare(test, "distance", exp.distance, got.Distance)
	checkResAmtsNameVals(test, "base stats", exp.baseStats, got.BaseStats)
	checkResAmtsNameVals(test, "status resists", exp.statusResists, got.StatusResists)
	compStructPtrs(test, "agility params", exp.agility, got.AgilityParameters)
	compStructSlices(test, "bribe chances", exp.bribeChances, got.BribeChances)
	compTestStructSlices(test, "elemental resists", exp.elemResists, got.ElemResists, compareMonsterElemResists)
	compareMonsterAppliedStates(test, exp.appliedState, got.AppliedState)
	testMonsterDefaultState(test, exp.defaultState, got.AlteredStates)
	compareMonsterItems(test, exp.items, got.Items)
	compareMonsterEquipment(test, exp.equipment, got.Equipment)

	checkResIDsInSlice(test, "properties", test.cfg.e.properties.endpoint, exp.properties, got.Properties)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "ronso rages", test.cfg.e.ronsoRages.endpoint, exp.ronsoRages, got.RonsoRages)
	checkResIDsInSlice(test, "areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas)
	checkResIDsInSlice(test, "formations", test.cfg.e.monsterFormations.endpoint, exp.formations, got.Formations)
	checkResIDsInSlice(test, "status immunities", test.cfg.e.statusConditions.endpoint, exp.statusImmunities, got.StatusImmunities)
	checkResPathsInSlice(test, "abilities", exp.abilities, got.Abilities)
}