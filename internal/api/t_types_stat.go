package api

type expStat struct {
	testGeneral
	expUnique
	spheres            []int32
	autoAbilities      []int32
	playerAbilities    []int32
	overdriveAbilities []int32
	itemAbilities      []int32
	triggerCommands    []int32
	statusConditions   []int32
	properties         []int32
}

func (e expStat) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareStats(test test, exp expStat, got Stat) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "spheres", test.cfg.e.spheres.endpoint, exp.spheres, got.Spheres)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "player abilities", test.cfg.e.playerAbilities.endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, "overdrive abilities", test.cfg.e.overdriveAbilities.endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, "item abilities", test.cfg.e.itemAbilities.endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, "trigger commands", test.cfg.e.triggerCommands.endpoint, exp.triggerCommands, got.TriggerCommands)
	checkResIDsInSlice(test, "status conditions", test.cfg.e.statusConditions.endpoint, exp.statusConditions, got.StatusConditions)
	checkResIDsInSlice(test, "properties", test.cfg.e.properties.endpoint, exp.properties, got.Properties)
}