package apiTest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expArenaCreation struct {
	testGeneral
	expUnique
	category                  string
	monster                   int32
	parentSubquest            int32
	reward                    testResAmount[api.TypedAPIResource]
	requiredCatchAmount       int32
	unlockedCreationsCategory *string
	requiredMonsters          []int32
}

func (e expArenaCreation) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareArenaCreations(test test, exp expArenaCreation, got api.ArenaCreation) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "category", exp.category, got.Category)
	compIdApiResource(test, "monster", test.cfg.E.Monsters.Endpoint, exp.monster, got.Monster)
	compIdApiResource(test, "parent subquest", test.cfg.E.Subquests.Endpoint, exp.parentSubquest, got.ParentSubquest)
	compareResAmounts(test, "reward", exp.reward, got.Reward)
	compare(test, "required catch amount", exp.requiredCatchAmount, got.RequiredCatchAmount)
	compare(test, "unlocked creations category", exp.unlockedCreationsCategory, got.UnlockedCreationsCategory)
	checkResIDsInSlice(test, "required monsters", test.cfg.E.Monsters.Endpoint, exp.requiredMonsters, got.RequiredMonsters)
}
