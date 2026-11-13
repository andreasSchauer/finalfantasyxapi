package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareArenaCreations(test Test, exp ArenaCreation, got apitypes.ArenaCreation) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("category", exp.category, got.Category)
	test.compIDAPIResource("monster", apilookups.E.Monsters.Endpoint, exp.monster, got.Monster)
	test.compIDAPIResource("parent subquest", apilookups.E.Subquests.Endpoint, exp.parentSubquest, got.ParentSubquest)
	compareResAmounts(test, "reward", exp.reward, got.Reward)
	test.compare("required catch amount", exp.requiredCatchAmount, got.RequiredCatchAmount)
	test.compare("unlocked creations category", exp.unlockedCreationsCategory, got.UnlockedCreationsCategory)
	test.checkResIDsInSlice("required monsters", apilookups.E.Monsters.Endpoint, exp.requiredMonsters, got.RequiredMonsters)
}
