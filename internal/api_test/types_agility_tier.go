package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expAgilityTier struct {
	testGeneral
	expIDOnly
	fromAgility int32
	toAgility   int32
	tickSpeed   int32
	monMaxICV   *int32
	monMinICV   *int32
	charMaxICV  *int32
	charMinICVs []api.AgilitySubtier
}

func (e expAgilityTier) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAgilityTiers(test test, exp expAgilityTier, got api.AgilityTier) {
	test.t.Helper()
	compareExpIDOnly(test, exp.expIDOnly, got.ID)
	compare(test, "from agility", exp.fromAgility, got.FromAgility)
	compare(test, "to agility", exp.toAgility, got.ToAgility)
	compare(test, "tick speed", exp.tickSpeed, got.TickSpeed)
	compare(test, "mon max icv", exp.monMaxICV, got.MonMaxICV)
	compare(test, "mon min icv", exp.monMinICV, got.MonMinICV)
	compare(test, "char max icv", exp.charMaxICV, got.CharMaxICV)
	compStructSlices(test, "char min icvs", exp.charMinICVs, got.CharMinICVs)
}
