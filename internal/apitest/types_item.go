package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type expItem struct {
	testGeneral
	expUnique
	untypedItem        int32
	category           database.ItemCategory
	monsters           []testMonItemAmts
	treasures          map[int32]int32
	shops              []int32
	quests             map[int32]int32
	blitzballPrizes    map[int32]int32
	aeonLearnAbilities map[int32]int32
	autoAbilities      map[string]int32
	mixes              []int32
}

func (e expItem) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareItems(test test, exp expItem, got apitypes.Item) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("untyped item", apilookups.E.AllItems.Endpoint, exp.untypedItem, got.UntypedItem)
	test.compare("category", string(exp.category), got.Category)
	test.checkTestStructsInSlice("monsters", exp.monsters, got.Monsters, compareMonItemAmts)
	test.checkResAmtsID("treasures", exp.treasures, got.Treasures)
	test.checkResIDsInSlice("shops", apilookups.E.Shops.Endpoint, exp.shops, got.Shops)
	test.checkResAmtsID("quests", exp.quests, got.Quests)
	test.checkResAmtsID("blitzball prizes", exp.blitzballPrizes, got.BlitzballPrizes)
	test.checkResAmtsID("aeon learn abilities", exp.aeonLearnAbilities, got.AeonLearnAbilities)
	test.checkResAmts("auto abilities", exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("mixes", apilookups.E.Mixes.Endpoint, exp.mixes, got.Mixes)
}

type testMonItemAmts struct {
	index               int
	monster             int32
	stealCommon         int32
	stealRare           int32
	dropCommon          int32
	dropRare            int32
	secondaryDropCommon int32
	secondaryDropRare   int32
	bribe               int32
	other               int32
}

func (t testMonItemAmts) GetIndex() int {
	return t.index
}

func compareMonItemAmts(test test, fieldName string, exp testMonItemAmts, got apitypes.MonItemAmts) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - monster", apilookups.E.Monsters.Endpoint, exp.monster, got.Monster)
	test.compare(fieldName+" - steal common", exp.stealCommon, got.StealCommon)
	test.compare(fieldName+" - steal rare", exp.stealRare, got.StealRare)
	test.compare(fieldName+" - drop common", exp.dropCommon, got.DropCommon)
	test.compare(fieldName+" - drop rare", exp.dropRare, got.DropRare)
	test.compare(fieldName+" - secondary drop common", exp.secondaryDropCommon, got.SecondaryDropCommon)
	test.compare(fieldName+" - secondary drop rare", exp.secondaryDropRare, got.SecondaryDropRare)
	test.compare(fieldName+" - bribe", exp.bribe, got.Bribe)
	test.compare(fieldName+" - other", exp.other, got.Other)
}
