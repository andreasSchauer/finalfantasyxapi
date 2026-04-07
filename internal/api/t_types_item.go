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