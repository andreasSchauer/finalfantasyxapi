package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareMonsterFormations(test test, exp expMonsterFormation, got api.MonsterFormation) {
	test.t.Helper()
	compareExpIDOnly(test, exp.expIDOnly, got.ID)
	compare(test, "category", exp.category, got.Category)
	compare(test, "is forced ambush", exp.isForcedAmbush, got.IsForcedAmbush)
	compare(test, "can escape", exp.canEscape, got.CanEscape)
	compare(test, "ap", exp.ap, got.AP)
	compare(test, "ap overkill", exp.apOverkill, got.ApOverkill)
	compare(test, "gil", exp.gil, got.Gil)
	compIDAPIResourcePtrs(test, "boss song", test.cfg.E.Songs.Endpoint, exp.bossMusic, got.BossMusic)
	checkResAmts(test, "monsters", exp.monsters, got.Monsters)
	compTestStructSlices(test, "trigger commands", exp.triggerCommands, got.TriggerCommands, compareFormationTCs)
	checkResIDsInSlice(test, "areas", test.cfg.E.Areas.Endpoint, exp.areas, got.Areas)
}

type testFormationTC struct {
	Ability int32
	Users   []int32
}

func compareFormationTCs(test test, fieldName string, exp testFormationTC, got api.FormationTriggerCommand) {
	test.t.Helper()
	tcEndpoint := test.cfg.E.TriggerCommands.Endpoint
	charClassesEndpoint := test.cfg.E.CharacterClasses.Endpoint

	compIDAPIResource(test, fieldName+" - tc ability", tcEndpoint, exp.Ability, got.Ability)
	checkResIDsInSlice(test, fieldName+" - tc users", charClassesEndpoint, exp.Users, got.Users)
}
