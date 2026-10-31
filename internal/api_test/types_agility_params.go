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
	compIDAPIResource(test, fieldName+" - agility tier", test.cfg.E.AgilityTiers.Endpoint, exp.agilityTier, got.AgilityTier)
	compare(test, fieldName+" - tick speed", exp.tickSpeed, got.TickSpeed)
	compare(test, fieldName+" - min icv", exp.minICV, got.MinICV)
	compare(test, fieldName+" - max icv", exp.maxICV, got.MaxICV)
}
