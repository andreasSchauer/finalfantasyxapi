package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

type expMonsterFormation struct {
	testGeneral
	expIDOnly
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

func (e expMonsterFormation) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMonsterFormations(test test, exp expMonsterFormation, got apitypes.MonsterFormation) {
	test.t.Helper()
	test.compareExpIDOnly(exp.expIDOnly, got.ID)
	test.compare("category", exp.category, got.Category)
	test.compare("is forced ambush", exp.isForcedAmbush, got.IsForcedAmbush)
	test.compare("can escape", exp.canEscape, got.CanEscape)
	test.compare("ap", exp.ap, got.AP)
	test.compare("ap overkill", exp.apOverkill, got.ApOverkill)
	test.compare("gil", exp.gil, got.Gil)
	test.compIDAPIResourcePtrs("boss song", api.E.Songs.Endpoint, exp.bossMusic, got.BossMusic)
	test.checkResAmts("monsters", exp.monsters, got.Monsters)
	test.compTestStructSlices("trigger commands", exp.triggerCommands, got.TriggerCommands, compareFormationTCs)
	test.checkResIDsInSlice("areas", api.E.Areas.Endpoint, exp.areas, got.Areas)
}

type testFormationTC struct {
	Ability int32
	Users   []int32
}

func compareFormationTCs(test test, fieldName string, exp testFormationTC, got apitypes.FormationTriggerCommand) {
	test.t.Helper()
	tcEndpoint := api.E.TriggerCommands.Endpoint
	charClassesEndpoint := api.E.CharacterClasses.Endpoint

	test.compIDAPIResource(fieldName+" - tc ability", tcEndpoint, exp.Ability, got.Ability)
	test.checkResIDsInSlice(fieldName+" - tc users", charClassesEndpoint, exp.Users, got.Users)
}
