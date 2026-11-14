package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type AgilityTier struct {
	testGeneral
	IDOnly
	fromAgility int32
	toAgility   int32
	tickSpeed   int32
	monMaxICV   *int32
	monMinICV   *int32
	charMaxICV  *int32
	charMinICVs []apitypes.AgilitySubtier
}

func (e AgilityTier) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e AgilityTier) Compare(test test, got apitypes.AgilityTier) {
	test.t.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compare("from agility", e.fromAgility, got.FromAgility)
	test.compare("to agility", e.toAgility, got.ToAgility)
	test.compare("tick speed", e.tickSpeed, got.TickSpeed)
	test.compare("mon max icv", e.monMaxICV, got.MonMaxICV)
	test.compare("mon min icv", e.monMinICV, got.MonMinICV)
	test.compare("char max icv", e.charMaxICV, got.CharMaxICV)
	test.compStructSlices("char min icvs", e.charMinICVs, got.CharMinICVs)
}
