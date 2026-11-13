package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareModifiers(test Test, exp Modifier, got apitypes.Modifier) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("player abilities", apilookups.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice("item abilities", apilookups.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice("trigger commands", apilookups.E.TriggerCommands.Endpoint, exp.triggerCommands, got.TriggerCommands)
	test.checkResIDsInSlice("enemy abilities", apilookups.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice("status conditions", apilookups.E.StatusConditions.Endpoint, exp.statusConditions, got.StatusConditions)
	test.checkResIDsInSlice("properties", apilookups.E.Properties.Endpoint, exp.properties, got.Properties)
}
