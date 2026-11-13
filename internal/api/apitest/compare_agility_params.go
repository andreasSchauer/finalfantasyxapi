package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareAgilityParams(test Test, fieldName string, exp TestAgilityParams, got apitypes.AgilityParams) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - agility tier", apilookups.E.AgilityTiers.Endpoint, exp.agilityTier, got.AgilityTier)
	test.compare(fieldName+" - tick speed", exp.tickSpeed, got.TickSpeed)
	test.compare(fieldName+" - min icv", exp.minICV, got.MinICV)
	test.compare(fieldName+" - max icv", exp.maxICV, got.MaxICV)
}
