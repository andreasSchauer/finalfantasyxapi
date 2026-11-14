package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type testAgilityParams struct {
	agilityTier int32
	tickSpeed   int32
	minICV      *int32
	maxICV      *int32
}

func (e testAgilityParams) Compare(test test, fieldName string, got apitypes.AgilityParams) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - agility tier", apilookups.E.AgilityTiers.Endpoint, e.agilityTier, got.AgilityTier)
	test.compare(fieldName+" - tick speed", e.tickSpeed, got.TickSpeed)
	test.compare(fieldName+" - min icv", e.minICV, got.MinICV)
	test.compare(fieldName+" - max icv", e.maxICV, got.MaxICV)
}
