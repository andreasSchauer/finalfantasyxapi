package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type ArenaCreation struct {
	TestGeneral
	Unique
	Category                  string
	Monster                   int32
	ParentSubquest            int32
	Reward                    testResAmount[apiresource.Typed]
	RequiredCatchAmount       int32
	UnlockedCreationsCategory *string
	RequiredMonsters          []int32
}

func (e ArenaCreation) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e ArenaCreation) Compare(test Test, got apitypes.ArenaCreation) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("category", e.Category, got.Category)
	test.compIDAPIResource("monster", apilookups.E.Monsters.Endpoint, e.Monster, got.Monster)
	test.compIDAPIResource("parent subquest", apilookups.E.Subquests.Endpoint, e.ParentSubquest, got.ParentSubquest)
	e.Reward.Compare(test, "reward", got.Reward)
	test.compare("required catch amount", e.RequiredCatchAmount, got.RequiredCatchAmount)
	test.compare("unlocked creations category", e.UnlockedCreationsCategory, got.UnlockedCreationsCategory)
	test.checkResIDsInSlice("required monsters", apilookups.E.Monsters.Endpoint, e.RequiredMonsters, got.RequiredMonsters)
}
