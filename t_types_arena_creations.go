package main

type expArenaCreation struct {
	testGeneral
	expUnique
	category					string
	monster						int32
	parentSubquest				int32
	reward						testItemAmount
	requiredCatchAmount			int32
	unlockedCreationsCategory 	*string
	requiredMonsters			[]int32
}

func (e expArenaCreation) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareArenaCreations(test test, exp expArenaCreation, got ArenaCreation) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "category", exp.category, got.Category)
	compIdApiResource(test, "monster", test.cfg.e.monsters.endpoint, exp.monster, got.Monster)
	compIdApiResource(test, "parent subquest", test.cfg.e.subquests.endpoint, exp.parentSubquest, got.ParentSubquest)
	compareItemAmounts(test, exp.reward, got.Reward)
	compare(test, "required catch amount", exp.requiredCatchAmount, got.RequiredCatchAmount)
	compare(test, "unlocked creations category", exp.unlockedCreationsCategory, got.UnlockedCreationsCategory)
	checkResIDsInSlice(test, "required monsters", test.cfg.e.monsters.endpoint, exp.requiredMonsters, got.RequiredMonsters)
}