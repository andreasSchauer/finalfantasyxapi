package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Item struct {
	TestGeneral
	Unique
	UntypedItem        int32
	Category           database.ItemCategory
	Monsters           []TestMonItemAmts
	Treasures          map[int32]int32
	Shops              []int32
	Quests             map[int32]int32
	BlitzballPrizes    map[int32]int32
	AeonLearnAbilities map[int32]int32
	AutoAbilities      map[string]int32
	Mixes              []int32
}

func (e Item) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Item) Compare(test Test, got apitypes.Item) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("untyped item", apilookups.E.AllItems.Endpoint, e.UntypedItem, got.UntypedItem)
	test.compare("category", string(e.Category), got.Category)
	test.checkTestStructsInSlice("monsters", e.Monsters, got.Monsters)
	test.checkResAmtsID("treasures", e.Treasures, got.Treasures)
	test.CheckResIDsInSlice("shops", apilookups.E.Shops.Endpoint, e.Shops, got.Shops)
	test.checkResAmtsID("quests", e.Quests, got.Quests)
	test.checkResAmtsID("blitzball prizes", e.BlitzballPrizes, got.BlitzballPrizes)
	test.checkResAmtsID("aeon learn abilities", e.AeonLearnAbilities, got.AeonLearnAbilities)
	test.checkResAmts("auto abilities", e.AutoAbilities, got.AutoAbilities)
	test.CheckResIDsInSlice("mixes", apilookups.E.Mixes.Endpoint, e.Mixes, got.Mixes)
}

type TestMonItemAmts struct {
	Index               int
	Monster             int32
	StealCommon         int32
	StealRare           int32
	DropCommon          int32
	DropRare            int32
	SecondaryDropCommon int32
	SecondaryDropRare   int32
	Bribe               int32
	Other               int32
}

func (e TestMonItemAmts) GetIndex() int {
	return e.Index
}

func (e TestMonItemAmts) Compare(test Test, fn string, got apitypes.MonItemAmts) {
	test.T.Helper()
	test.compIDAPIResource(childFn(fn, "monster"), apilookups.E.Monsters.Endpoint, e.Monster, got.Monster)
	test.compare(childFn(fn, "steal common"), e.StealCommon, got.StealCommon)
	test.compare(childFn(fn, "steal rare"), e.StealRare, got.StealRare)
	test.compare(childFn(fn, "drop common"), e.DropCommon, got.DropCommon)
	test.compare(childFn(fn, "drop rare"), e.DropRare, got.DropRare)
	test.compare(childFn(fn, "secondary drop common"), e.SecondaryDropCommon, got.SecondaryDropCommon)
	test.compare(childFn(fn, "secondary drop rare"), e.SecondaryDropRare, got.SecondaryDropRare)
	test.compare(childFn(fn, "bribe"), e.Bribe, got.Bribe)
	test.compare(childFn(fn, "other"), e.Other, got.Other)
}
