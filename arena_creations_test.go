package main

import (
	//"net/http"
	"testing"
	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expArenaCreation struct {
	testGeneral
	expUnique
	category			string
	monster				int32
	parentSubquest		int32
	reward				testItemAmount
	requiredCatchAmount	int32
	requiredMonsters	[]int32
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
	compareResListTest(test, rltIDs("required monsters", test.cfg.e.monsters.endpoint, exp.requiredMonsters, got.RequiredMonsters))
}

func TestGetArenaCreation(t *testing.T) {
	tests := []expArenaCreation{}

	testSingleResources(t, tests, "GetArenaCreation", testCfg.HandleArenaCreations, compareArenaCreations)
}

func TestRetrieveArenaCreations(t *testing.T) {
	tests := []expListIDs{}

	testIdList(t, tests, testCfg.e.arenaCreations.endpoint, "RetrieveArenaCreations", testCfg.HandleArenaCreations, compareAPIResourceLists[NamedApiResourceList])
}
