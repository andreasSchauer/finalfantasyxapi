package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Overdrive struct {
	testGeneral
	Unique
	rank               *int32
	countdownInSec     *int32
	user               int32
	overdriveCommand   *int32
	overdriveAbilities []int32
}

func (e Overdrive) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareOverdrives(test test, exp Overdrive, got apitypes.Overdrive) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("countdown in sec", exp.countdownInSec, got.CountdownInSec)
	test.compIDAPIResource("user", apilookups.E.CharacterClasses.Endpoint, exp.user, got.User)
	test.compIDAPIResourcePtrs("overdrive command", apilookups.E.OverdriveCommands.Endpoint, exp.overdriveCommand, got.OverdriveCommand)
	test.checkResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
}
