package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type expArenaCreation struct {
	testGeneral
	expUnique
	category                  string
	monster                   int32
	parentSubquest            int32
	reward                    testResAmount[apiresource.Typed]
	requiredCatchAmount       int32
	unlockedCreationsCategory *string
	requiredMonsters          []int32
}

func (e expArenaCreation) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareArenaCreations(test test, exp expArenaCreation, got apitypes.ArenaCreation) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compare("category", exp.category, got.Category)
	test.compIDAPIResource("monster", apilookups.E.Monsters.Endpoint, exp.monster, got.Monster)
	test.compIDAPIResource("parent subquest", apilookups.E.Subquests.Endpoint, exp.parentSubquest, got.ParentSubquest)
	compareResAmounts(test, "reward", exp.reward, got.Reward)
	test.compare("required catch amount", exp.requiredCatchAmount, got.RequiredCatchAmount)
	test.compare("unlocked creations category", exp.unlockedCreationsCategory, got.UnlockedCreationsCategory)
	test.checkResIDsInSlice("required monsters", apilookups.E.Monsters.Endpoint, exp.requiredMonsters, got.RequiredMonsters)
}
