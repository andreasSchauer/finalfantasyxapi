package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Sphere struct {
	testGeneral
	Unique
	item            int32
	createdNode     *apitypes.CreatedNode
	monsters        []testMonItemAmts
	treasures       map[int32]int32
	shops           []int32
	quests          map[int32]int32
	blitzballPrizes map[int32]int32
}

func (e Sphere) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Sphere) Compare(test test, got apitypes.Sphere) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("item", apilookups.E.Items.Endpoint, e.item, got.Item)
	test.compStructPtrs("created node", e.createdNode, got.CreatedNode)
	test.checkTestStructsInSlice("monsters", e.monsters, got.Monsters)
	test.checkResAmtsID("treasures", e.treasures, got.Treasures)
	test.checkResIDsInSlice("shops", apilookups.E.Shops.Endpoint, e.shops, got.Shops)
	test.checkResAmtsID("quests", e.quests, got.Quests)
	test.checkResAmtsID("blitzball prizes", e.blitzballPrizes, got.BlitzballPrizes)
}
