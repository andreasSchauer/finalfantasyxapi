package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type MonsterFormation struct {
	testGeneral
	IDOnly
	category        string
	isForcedAmbush  bool
	canEscape       bool
	bossMusic       *int32
	ap              int32
	apOverkill      int32
	gil             int32
	monsters        map[string]int32
	areas           []int32
	triggerCommands []testFormationTC
}

func (e MonsterFormation) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMonsterFormations(test test, exp MonsterFormation, got apitypes.MonsterFormation) {
	test.t.Helper()
	test.compareExpIDOnly(exp.IDOnly, got.ID)
	test.compare("category", exp.category, got.Category)
	test.compare("is forced ambush", exp.isForcedAmbush, got.IsForcedAmbush)
	test.compare("can escape", exp.canEscape, got.CanEscape)
	test.compare("ap", exp.ap, got.AP)
	test.compare("ap overkill", exp.apOverkill, got.ApOverkill)
	test.compare("gil", exp.gil, got.Gil)
	test.compIDAPIResourcePtrs("boss song", apilookups.E.Songs.Endpoint, exp.bossMusic, got.BossMusic)
	test.checkResAmts("monsters", exp.monsters, got.Monsters)
	test.compTestStructSlices("trigger commands", exp.triggerCommands, got.TriggerCommands, compareFormationTCs)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, exp.areas, got.Areas)
}

type testFormationTC struct {
	Ability int32
	Users   []int32
}

func compareFormationTCs(test test, fieldName string, exp testFormationTC, got apitypes.FormationTriggerCommand) {
	test.t.Helper()
	tcEndpoint := apilookups.E.TriggerCommands.Endpoint
	charClassesEndpoint := apilookups.E.CharacterClasses.Endpoint

	test.compIDAPIResource(fieldName+" - tc ability", tcEndpoint, exp.Ability, got.Ability)
	test.checkResIDsInSlice(fieldName+" - tc users", charClassesEndpoint, exp.Users, got.Users)
}
