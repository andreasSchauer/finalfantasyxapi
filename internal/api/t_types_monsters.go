package api

type expMonster struct {
	testGeneral
	expNameVer
	appliedState     *testAppliedState
	agility          *AgilityParams
	species          int32
	ctbIconType      string
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
	compare(test, "ctb icon type", exp.ctbIconType, got.CTBIconType)
	compare(test, "distance", exp.distance, got.Distance)
	checkResAmtsNameVals(test, "base stats", exp.baseStats, got.BaseStats)
	checkResAmtsNameVals(test, "status resists", exp.statusResists, got.StatusResists)
	compStructPtrs(test, "agility params", exp.agility, got.AgilityParameters)
	compStructSlices(test, "bribe chances", exp.bribeChances, got.BribeChances)
	compTestStructSlices(test, "elemental resists", exp.elemResists, got.ElemResists, compareElemResists)
	compTestStructPtrs(test, "applied state", exp.appliedState, got.AppliedState, compareMonsterAppliedStates)
	testMonsterDefaultState(test, exp.defaultState, got.AlteredStates)
	compTestStructPtrs(test, "items", exp.items, got.Items, compareMonsterItems)
	compTestStructPtrs(test, "equipment", exp.equipment, got.Equipment, compareMonsterEquipment)

	checkResIDsInSlice(test, "properties", test.cfg.e.properties.endpoint, exp.properties, got.Properties)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "ronso rages", test.cfg.e.ronsoRages.endpoint, exp.ronsoRages, got.RonsoRages)
	checkResIDsInSlice(test, "areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas)
	checkResIDsInSlice(test, "formations", test.cfg.e.monsterFormations.endpoint, exp.formations, got.Formations)
	checkResIDsInSlice(test, "status immunities", test.cfg.e.statusConditions.endpoint, exp.statusImmunities, got.StatusImmunities)
	checkResPathsInSlice(test, "abilities", exp.abilities, got.Abilities)
}

type testMonItems struct {
	itemDropChance int32
	items          map[string]*testItemAmount
	otherItems     []testPossibleItem
}

func compareMonsterItems(test test, fieldName string, exp testMonItems, got MonsterItems) {
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

type testMonEquipment struct {
	abilitySlots      MonsterEquipmentSlots
	attachedAbilities MonsterEquipmentSlots
	weaponAbilities   []int32
	armorAbilities    []int32
}

func compareMonsterEquipment(test test, _ string, exp testMonEquipment, got MonsterEquipment) {
	compStructs(test, "ability slots", exp.abilitySlots, got.AbilitySlots)
	compStructs(test, "attached abilities", exp.attachedAbilities, got.AttachedAbilities)
	checkResIDsInSlice(test, "weapon abilities", test.cfg.e.autoAbilities.endpoint, exp.weaponAbilities, got.WeaponAbilities)
	checkResIDsInSlice(test, "armor abilities", test.cfg.e.autoAbilities.endpoint, exp.armorAbilities, got.ArmorAbilities)
}
