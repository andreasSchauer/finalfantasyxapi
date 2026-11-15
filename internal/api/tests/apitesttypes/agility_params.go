package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type TestAgilityParams struct {
	AgilityTier int32
	TickSpeed   int32
	MinICV      *int32
	MaxICV      *int32
}

func (e TestAgilityParams) Compare(test Test, fn string, got apitypes.AgilityParams) {
	test.T.Helper()
	test.compIDAPIResource(childFn(fn, "agility tier"), apilookups.E.AgilityTiers.Endpoint, e.AgilityTier, got.AgilityTier)
	test.compare(childFn(fn, "tick speed"), e.TickSpeed, got.TickSpeed)
	test.compare(childFn(fn, "min icv"), e.MinICV, got.MinICV)
	test.compare(childFn(fn, "max icv"), e.MaxICV, got.MaxICV)
}
