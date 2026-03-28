package api

type expSidequest struct {
	testGeneral
	expUnique
	untypedQuest int32
	completion   *testQuestCompletion
	subquests    []int32
}

func (e expSidequest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSidequests(test test, exp expSidequest, got Sidequest) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "untyped quest", test.cfg.e.quests.endpoint, exp.untypedQuest, got.UntypedQuest)
	compTestStructPtrs(test, "completion", exp.completion, got.Completion, compareQuestCompletions)
	checkResIDsInSlice(test, "subquests", test.cfg.e.subquests.endpoint, exp.subquests, got.Subquests)
}

type expSubquest struct {
	testGeneral
	expUnique
	untypedQuest    int32
	parentSidequest int32
	completion      testQuestCompletion
	arenaCreation   *int32
}

func (e expSubquest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSubquests(test test, exp expSubquest, got Subquest) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "untyped quest", test.cfg.e.quests.endpoint, exp.untypedQuest, got.UntypedQuest)
	compIdApiResource(test, "parent sidequest", test.cfg.e.sidequests.endpoint, exp.parentSidequest, got.ParentSidequest)
	compareQuestCompletions(test, "completion", exp.completion, got.Completion)
	compIdApiResourcePtrs(test, "arena creation", test.cfg.e.arenaCreations.endpoint, exp.arenaCreation, got.ArenaCreation)
}

type testQuestCompletion struct {
	areas  []int32
	reward testResAmount[TypedAPIResource]
}

func compareQuestCompletions(test test, fieldName string, exp testQuestCompletion, got QuestCompletion) {
	checkResIDsInSlice(test, fieldName+" - areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas)
	compareResAmounts(test, fieldName+" - rewards", exp.reward, got.Reward)
}
