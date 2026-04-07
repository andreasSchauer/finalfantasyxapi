package api

type expItem struct {
	testGeneral
	expUnique
	untypedItem        int32
	category           int32
	monsters           []testMonItemAmts
	treasures          map[int32]int32
	shops              []int32
	quests             map[int32]int32
	blitzballPrizes    map[int32]int32
	aeonLearnAbilities map[int32]int32
	autoAbilities      map[string]int32
	mixes              []int32
}

func (i expItem) GetTestGeneral() testGeneral {
	return i.testGeneral
}

func compareItems(test test, exp expItem, got Item) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "untyped item", test.cfg.e.allItems.endpoint, exp.untypedItem, got.UntypedItem)
	compIdApiResource(test, "category", test.cfg.e.itemCategory.endpoint, exp.category, got.Category)
	checkTestStructsInSlice(test, "monsters", exp.monsters, got.Monsters, compareMonItemAmts)
	checkResAmtsID(test, "treasures", exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "shops", test.cfg.e.shops.endpoint, exp.shops, got.Shops)
	checkResAmtsID(test, "quests", exp.quests, got.Quests)
	checkResAmtsID(test, "blitzball prizes", exp.blitzballPrizes, got.BlitzballPrizes)
	checkResAmtsID(test, "aeon learn abilities", exp.aeonLearnAbilities, got.AeonLearnAbilities)
	checkResAmts(test, "auto abilities", exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "mixes", test.cfg.e.mixes.endpoint, exp.mixes, got.Mixes)
}

type testMonItemAmts struct {
	index		  int
	monster       int32
	steal         *CommonRareAmount
	drop          *CommonRareAmount
	secondaryDrop *CommonRareAmount
	bribe         int32
	other         int32
}

func (t testMonItemAmts) GetIndex() int {
	return t.index
}

func compareMonItemAmts(test test, fieldName string, exp testMonItemAmts, got MonItemAmts) {
	compIdApiResource(test, fieldName+" - monster", test.cfg.e.monsters.endpoint, exp.monster, got.Monster)
	compStructPtrs(test, fieldName+" - steal", exp.steal, got.Steal)
	compStructPtrs(test, fieldName+" - drop", exp.drop, got.Drop)
	compStructPtrs(test, fieldName+" - secondary drop", exp.secondaryDrop, got.SecondaryDrop)
	compare(test, fieldName+" - bribe", exp.bribe, got.Bribe)
	compare(test, fieldName+" - other", exp.other, got.Other)
}


type expKeyItem struct {
	testGeneral
	expUnique
	untypedItem        	int32
	category           	int32
	celestialWeapon		*int32	
	primer				*int32
	areas				[]int32
	treasures			[]int32
	quests				[]int32
}

func (i expKeyItem) GetTestGeneral() testGeneral {
	return i.testGeneral
}

func compareKeyItems(test test, exp expKeyItem, got KeyItem) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "untyped item", test.cfg.e.allItems.endpoint, exp.untypedItem, got.UntypedItem)
	compIdApiResource(test, "category", test.cfg.e.keyItemCategory.endpoint, exp.category, got.Category)
	compIdApiResourcePtrs(test, "celestial weapon", test.cfg.e.celestialWeapons.endpoint, exp.celestialWeapon, got.CelestialWeapon)
	compIdApiResourcePtrs(test, "primer", test.cfg.e.primers.endpoint, exp.primer, got.Primer)
	checkResIDsInSlice(test, "areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas)
	checkResIDsInSlice(test, "treasures", test.cfg.e.treasures.endpoint, exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "quests", test.cfg.e.quests.endpoint, exp.quests, got.Quests)
}


type expMasterItem struct {
	testGeneral
	expUnique
	itemType			int32
	typedItem			string
	obtainableFrom		ObtainableFrom
}

func (i expMasterItem) GetTestGeneral() testGeneral {
	return i.testGeneral
}

func compareMasterItems(test test, exp expMasterItem, got MasterItem) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "item type", test.cfg.e.itemType.endpoint, exp.itemType, got.Type)
	compPathApiResource(test, "typed item", exp.typedItem, got.TypedItem)
	compStructs(test, "obtainable from", exp.obtainableFrom, got.ObtainableFrom)
}


type expPrimer struct {
	testGeneral
	expUnique
	keyItem		int32
	areas		[]int32
	treasures	[]int32
}

func (p expPrimer) GetTestGeneral() testGeneral {
	return p.testGeneral
}

func comparePrimers(test test, exp expPrimer, got Primer) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "key item", test.cfg.e.keyItemCategory.endpoint, exp.keyItem, got.KeyItem)
	checkResIDsInSlice(test, "areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas)
	checkResIDsInSlice(test, "treasures", test.cfg.e.treasures.endpoint, exp.treasures, got.Treasures)
}


type expMix struct {
	testGeneral
	expUnique
	category		int32
	overdrive		int32
	combinations	[]testMixCombination
}

func (m expMix) GetTestGeneral() testGeneral {
	return m.testGeneral
}

func compareMixes(test test, exp expMix, got Mix) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "category", test.cfg.e.mixCategory.endpoint, exp.category, got.Category)
	compIdApiResource(test, "overdrive", test.cfg.e.overdrives.endpoint, exp.overdrive, got.Overdrive)
	compTestStructSlices(test, "combinations", exp.combinations, got.Combinations, compareMixCombinations)
}

type testMixCombination struct {
	firstItem	int32
	secondItem	int32
}

func compareMixCombinations(test test, fieldName string, exp testMixCombination, got MixCombination) {
	compIdApiResource(test, fieldName+" - first item", test.cfg.e.items.endpoint, exp.firstItem, got.FirstItem)
	compIdApiResource(test, fieldName+" - second item", test.cfg.e.items.endpoint, exp.secondItem, got.SecondItem)
}