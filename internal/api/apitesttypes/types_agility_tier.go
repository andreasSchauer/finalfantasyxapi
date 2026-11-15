package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type AgilityTier struct {
	TestGeneral
	IDOnly
	FromAgility int32
	ToAgility   int32
	TickSpeed   int32
	MonMaxICV   *int32
	MonMinICV   *int32
	CharMaxICV  *int32
	CharMinICVs []apitypes.AgilitySubtier
}

func (e AgilityTier) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e AgilityTier) Compare(test Test, got apitypes.AgilityTier) {
	test.T.Helper()
	test.compareExpIDOnly(e.IDOnly, got.ID)
	test.compare("from agility", e.FromAgility, got.FromAgility)
	test.compare("to agility", e.ToAgility, got.ToAgility)
	test.compare("tick speed", e.TickSpeed, got.TickSpeed)
	test.compare("mon max icv", e.MonMaxICV, got.MonMaxICV)
	test.compare("mon min icv", e.MonMinICV, got.MonMinICV)
	test.compare("char max icv", e.CharMaxICV, got.CharMaxICV)
	test.compStructSlices("char min icvs", e.CharMinICVs, got.CharMinICVs)
}
