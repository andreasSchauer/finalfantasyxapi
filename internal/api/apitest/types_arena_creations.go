package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type ArenaCreation struct {
	testGeneral
	Unique
	category                  string
	monster                   int32
	parentSubquest            int32
	reward                    testResAmount[apiresource.Typed]
	requiredCatchAmount       int32
	unlockedCreationsCategory *string
	requiredMonsters          []int32
}

func (e ArenaCreation) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e ArenaCreation) Compare(test test, got apitypes.ArenaCreation) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("category", e.category, got.Category)
	test.compIDAPIResource("monster", apilookups.E.Monsters.Endpoint, e.monster, got.Monster)
	test.compIDAPIResource("parent subquest", apilookups.E.Subquests.Endpoint, e.parentSubquest, got.ParentSubquest)
	e.reward.Compare(test, "reward", got.Reward)
	test.compare("required catch amount", e.requiredCatchAmount, got.RequiredCatchAmount)
	test.compare("unlocked creations category", e.unlockedCreationsCategory, got.UnlockedCreationsCategory)
	test.checkResIDsInSlice("required monsters", apilookups.E.Monsters.Endpoint, e.requiredMonsters, got.RequiredMonsters)
}
