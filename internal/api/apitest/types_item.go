package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Item struct {
	testGeneral
	Unique
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

func (e Item) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Item) Compare(test test, got apitypes.Item) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped item", apilookups.E.AllItems.Endpoint, e.untypedItem, got.UntypedItem)
	test.compare("category", string(e.category), got.Category)
	test.checkTestStructsInSlice("monsters", e.monsters, got.Monsters)
	test.checkResAmtsID("treasures", e.treasures, got.Treasures)
	test.checkResIDsInSlice("shops", apilookups.E.Shops.Endpoint, e.shops, got.Shops)
	test.checkResAmtsID("quests", e.quests, got.Quests)
	test.checkResAmtsID("blitzball prizes", e.blitzballPrizes, got.BlitzballPrizes)
	test.checkResAmtsID("aeon learn abilities", e.aeonLearnAbilities, got.AeonLearnAbilities)
	test.checkResAmts("auto abilities", e.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("mixes", apilookups.E.Mixes.Endpoint, e.mixes, got.Mixes)
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

func (e testMonItemAmts) GetIndex() int {
	return e.index
}

func (e testMonItemAmts) Compare(test test, fieldName string, got apitypes.MonItemAmts) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - monster", apilookups.E.Monsters.Endpoint, e.monster, got.Monster)
	test.compare(fieldName+" - steal common", e.stealCommon, got.StealCommon)
	test.compare(fieldName+" - steal rare", e.stealRare, got.StealRare)
	test.compare(fieldName+" - drop common", e.dropCommon, got.DropCommon)
	test.compare(fieldName+" - drop rare", e.dropRare, got.DropRare)
	test.compare(fieldName+" - secondary drop common", e.secondaryDropCommon, got.SecondaryDropCommon)
	test.compare(fieldName+" - secondary drop rare", e.secondaryDropRare, got.SecondaryDropRare)
	test.compare(fieldName+" - bribe", e.bribe, got.Bribe)
	test.compare(fieldName+" - other", e.other, got.Other)
}
