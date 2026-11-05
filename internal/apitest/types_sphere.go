package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expSphere struct {
	testGeneral
	expUnique
	item            int32
	createdNode     *api.CreatedNode
	monsters        []testMonItemAmts
	treasures       map[int32]int32
	shops           []int32
	quests          map[int32]int32
	blitzballPrizes map[int32]int32
}

func (e expSphere) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSpheres(test test, exp expSphere, got api.Sphere) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("item", api.E.Items.Endpoint, exp.item, got.Item)
	test.compStructPtrs("created node", exp.createdNode, got.CreatedNode)
	test.checkTestStructsInSlice("monsters", exp.monsters, got.Monsters, compareMonItemAmts)
	test.checkResAmtsID("treasures", exp.treasures, got.Treasures)
	test.checkResIDsInSlice("shops", api.E.Shops.Endpoint, exp.shops, got.Shops)
	test.checkResAmtsID("quests", exp.quests, got.Quests)
	test.checkResAmtsID("blitzball prizes", exp.blitzballPrizes, got.BlitzballPrizes)
}
