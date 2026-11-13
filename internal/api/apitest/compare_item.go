package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareItems(test Test, exp Item, got apitypes.Item) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
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

func compareMonItemAmts(test Test, fieldName string, exp TestMonItemAmts, got apitypes.MonItemAmts) {
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
