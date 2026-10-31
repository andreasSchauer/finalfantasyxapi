package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareStats(test test, exp expStat, got api.Stat) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "spheres", test.cfg.E.Spheres.Endpoint, exp.spheres, got.Spheres)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "player abilities", test.cfg.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, "overdrive abilities", test.cfg.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, "item abilities", test.cfg.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, "trigger commands", test.cfg.E.TriggerCommands.Endpoint, exp.triggerCommands, got.TriggerCommands)
	checkResIDsInSlice(test, "status conditions", test.cfg.E.StatusConditions.Endpoint, exp.statusConditions, got.StatusConditions)
	checkResIDsInSlice(test, "properties", test.cfg.E.Properties.Endpoint, exp.properties, got.Properties)
}
