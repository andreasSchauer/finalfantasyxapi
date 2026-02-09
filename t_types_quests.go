package main

type expSidequest struct {
	testGeneral
	expUnique
	completion		*testQuestCompletion
	subquests		[]int32
}

func (e expSidequest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSidequests(test test, exp expSidequest, got Sidequest) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compTestStructPtrs(test, "completion", exp.completion, got.Completion, compareQuestCompletions)
	checkResIDsInSlice(test, "subquests", test.cfg.e.subquests.endpoint, exp.subquests, got.Subquests)
}

type expSubquest struct {
	testGeneral
	expUnique
	parentSidequest	int32
	completions		[]testQuestCompletion
}

func (e expSubquest) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSubquests(test test, exp expSubquest, got Subquest) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "parent sidequest", test.cfg.e.sidequests.endpoint, exp.parentSidequest, got.ParentSidequest)
	checkTestStructsInSlice(test, "completions", exp.completions, got.Completions, compareQuestCompletions)
}

type testQuestCompletion struct {
	index	int
	areas	[]int32
	reward	testItemAmount
}

func (t testQuestCompletion) GetIndex() int {
	return t.index
}

func compareQuestCompletions(test test, exp testQuestCompletion, got QuestCompletion) {
	checkResIDsInSlice(test, "quest completion - areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas)
	compareItemAmounts(test, exp.reward, got.Reward)
}