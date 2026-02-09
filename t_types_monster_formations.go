package main

type expMonsterFormation struct {
	testGeneral
	expIdOnly
	category        string
	isForcedAmbush  bool
	canEscape       bool
	bossMusic       *int32
	monsters        map[string]int32
	areas           []int32
	triggerCommands []testFormationTC
}

func (e expMonsterFormation) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMonsterFormations(test test, exp expMonsterFormation, got MonsterFormation) {
	compareExpIdOnly(test, exp.expIdOnly, got.ID)
	compare(test, "category", exp.category, got.Category)
	compare(test, "is forced ambush", exp.isForcedAmbush, got.IsForcedAmbush)
	compare(test, "can escape", exp.canEscape, got.CanEscape)
	compIdApiResourcePtrs(test, "boss song", test.cfg.e.songs.endpoint, exp.bossMusic, got.BossMusic)
	checkResAmtsNameVals(test, "monsters", exp.monsters, got.Monsters)
	compTestStructSlices(test, "trigger commands", exp.triggerCommands, got.TriggerCommands, compareFormationTCs)
	checkResIDsInSlice(test, "areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas)
}

type testFormationTC struct {
	Ability int32
	Users   []int32
}

func compareFormationTCs(test test, exp testFormationTC, got FormationTriggerCommand) {
	tcEndpoint := test.cfg.e.triggerCommands.endpoint
	charClassesEndpoint := test.cfg.e.characterClasses.endpoint

	compIdApiResource(test, "tc ability", tcEndpoint, exp.Ability, got.Ability)
	checkResIDsInSlice(test, "tc users", charClassesEndpoint, exp.Users, got.Users)
}