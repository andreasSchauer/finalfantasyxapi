package api

type expModifier struct {
	testGeneral
	expUnique
	autoAbilities      []int32
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	triggerCommands    []int32
	enemyAbilities     []int32
	statusConditions   []int32
	properties         []int32
}

func (e expModifier) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareModifiers(test test, exp expModifier, got Modifier) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "player abilities", test.cfg.e.playerAbilities.endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, "overdrive abilities", test.cfg.e.overdriveAbilities.endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, "item abilities", test.cfg.e.itemAbilities.endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, "trigger commands", test.cfg.e.triggerCommands.endpoint, exp.triggerCommands, got.TriggerCommands)
	checkResIDsInSlice(test, "enemy abilities", test.cfg.e.enemyAbilities.endpoint, exp.enemyAbilities, got.EnemyAbilities)
	checkResIDsInSlice(test, "status conditions", test.cfg.e.statusConditions.endpoint, exp.statusConditions, got.StatusConditions)
	checkResIDsInSlice(test, "properties", test.cfg.e.properties.endpoint, exp.properties, got.Properties)
}