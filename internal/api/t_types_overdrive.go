package api

type expOverdrive struct {
	testGeneral
	expUnique
	rank				*int32
	countdownInSec		*int32
	user				int32
	overdriveCommand	*int32
	overdriveAbilities	[]int32
}

func (e expOverdrive) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareOverdrives(test test, exp expOverdrive, got Overdrive) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "countdown in sec", exp.countdownInSec, got.CountdownInSec)
	compIdApiResource(test, "user", test.cfg.e.characterClasses.endpoint, exp.user, got.User)
	compIdApiResourcePtrs(test, "overdrive command", test.cfg.e.overdriveCommands.endpoint, exp.overdriveCommand, got.OverdriveCommand)
	checkResIDsInSlice(test, "overdrive abilities", test.cfg.e.overdriveAbilities.endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
}