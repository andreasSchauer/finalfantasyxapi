package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareSpheres(test Test, exp Sphere, got apitypes.Sphere) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("item", apilookups.E.Items.Endpoint, exp.item, got.Item)
	test.compStructPtrs("created node", exp.createdNode, got.CreatedNode)
	test.checkTestStructsInSlice("monsters", exp.monsters, got.Monsters, compareMonItemAmts)
	test.checkResAmtsID("treasures", exp.treasures, got.Treasures)
	test.checkResIDsInSlice("shops", apilookups.E.Shops.Endpoint, exp.shops, got.Shops)
	test.checkResAmtsID("quests", exp.quests, got.Quests)
	test.checkResAmtsID("blitzball prizes", exp.blitzballPrizes, got.BlitzballPrizes)
}
