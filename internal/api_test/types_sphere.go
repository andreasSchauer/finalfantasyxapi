package apiTest

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
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "item", test.cfg.E.Items.Endpoint, exp.item, got.Item)
	compStructPtrs(test, "created node", exp.createdNode, got.CreatedNode)
	checkTestStructsInSlice(test, "monsters", exp.monsters, got.Monsters, compareMonItemAmts)
	checkResAmtsID(test, "treasures", exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "shops", test.cfg.E.Shops.Endpoint, exp.shops, got.Shops)
	checkResAmtsID(test, "quests", exp.quests, got.Quests)
	checkResAmtsID(test, "blitzball prizes", exp.blitzballPrizes, got.BlitzballPrizes)
}
