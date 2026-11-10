package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type expAgilityTier struct {
	testGeneral
	expIDOnly
	fromAgility int32
	toAgility   int32
	tickSpeed   int32
	monMaxICV   *int32
	monMinICV   *int32
	charMaxICV  *int32
	charMinICVs []apitypes.AgilitySubtier
}

func (e expAgilityTier) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAgilityTiers(test test, exp expAgilityTier, got apitypes.AgilityTier) {
	test.t.Helper()
	test.compareExpIDOnly(exp.expIDOnly, got.ID)
	test.compare("from agility", exp.fromAgility, got.FromAgility)
	test.compare("to agility", exp.toAgility, got.ToAgility)
	test.compare("tick speed", exp.tickSpeed, got.TickSpeed)
	test.compare("mon max icv", exp.monMaxICV, got.MonMaxICV)
	test.compare("mon min icv", exp.monMinICV, got.MonMinICV)
	test.compare("char max icv", exp.charMaxICV, got.CharMaxICV)
	test.compStructSlices("char min icvs", exp.charMinICVs, got.CharMinICVs)
}
