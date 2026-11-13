package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareElements(test Test, exp Element, got apitypes.Element) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResourcePtrs("status protection", apilookups.E.StatusConditions.Endpoint, exp.statusProtection, got.StatusProtection)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("player abilities", apilookups.E.PlayerAbilities.Endpoint, exp.playerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice("item abilities", apilookups.E.ItemAbilities.Endpoint, exp.itemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice("enemy abilities", apilookups.E.EnemyAbilities.Endpoint, exp.enemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice("monsters weak", apilookups.E.Monsters.Endpoint, exp.monstersWeak, got.MonstersWeak)
	test.checkResIDsInSlice("monsters halved", apilookups.E.Monsters.Endpoint, exp.monstersHalved, got.MonstersHalved)
	test.checkResIDsInSlice("monsters immune", apilookups.E.Monsters.Endpoint, exp.monstersImmune, got.MonstersImmune)
	test.checkResIDsInSlice("monsters absorb", apilookups.E.Monsters.Endpoint, exp.monstersAbsorb, got.MonstersAbsorb)
}
