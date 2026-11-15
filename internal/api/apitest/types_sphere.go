package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Sphere struct {
	TestGeneral
	Unique
	Item            int32
	CreatedNode     *apitypes.CreatedNode
	Monsters        []TestMonItemAmts
	Treasures       map[int32]int32
	Shops           []int32
	Quests          map[int32]int32
	BlitzballPrizes map[int32]int32
}

func (e Sphere) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Sphere) Compare(test Test, got apitypes.Sphere) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("item", apilookups.E.Items.Endpoint, e.Item, got.Item)
	test.compStructPtrs("created node", e.CreatedNode, got.CreatedNode)
	test.checkTestStructsInSlice("monsters", e.Monsters, got.Monsters)
	test.checkResAmtsID("treasures", e.Treasures, got.Treasures)
	test.checkResIDsInSlice("shops", apilookups.E.Shops.Endpoint, e.Shops, got.Shops)
	test.checkResAmtsID("quests", e.Quests, got.Quests)
	test.checkResAmtsID("blitzball prizes", e.BlitzballPrizes, got.BlitzballPrizes)
}
