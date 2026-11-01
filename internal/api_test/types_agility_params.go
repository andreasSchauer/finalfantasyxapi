package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type testAgilityParams struct {
	agilityTier int32
	tickSpeed   int32
	minICV      *int32
	maxICV      *int32
}

func compareAgilityParams(test test, fieldName string, exp testAgilityParams, got api.AgilityParams) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - agility tier", test.cfg.E.AgilityTiers.Endpoint, exp.agilityTier, got.AgilityTier)
	test.compare(fieldName+" - tick speed", exp.tickSpeed, got.TickSpeed)
	test.compare(fieldName+" - min icv", exp.minICV, got.MinICV)
	test.compare(fieldName+" - max icv", exp.maxICV, got.MaxICV)
}
