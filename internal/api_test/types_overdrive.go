package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expOverdrive struct {
	testGeneral
	expUnique
	rank               *int32
	countdownInSec     *int32
	user               int32
	overdriveCommand   *int32
	overdriveAbilities []int32
}

func (e expOverdrive) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareOverdrives(test test, exp expOverdrive, got api.Overdrive) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "countdown in sec", exp.countdownInSec, got.CountdownInSec)
	compIDAPIResource(test, "user", test.cfg.E.CharacterClasses.Endpoint, exp.user, got.User)
	compIDAPIResourcePtrs(test, "overdrive command", test.cfg.E.OverdriveCommands.Endpoint, exp.overdriveCommand, got.OverdriveCommand)
	checkResIDsInSlice(test, "overdrive abilities", test.cfg.E.OverdriveAbilities.Endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
}
