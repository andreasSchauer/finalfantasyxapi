package api


type testAgilityParams struct {
	agilityTier int32
	tickSpeed   int32
	minICV      *int32
	maxICV      *int32
}

func compareAgilityParams(test test, fieldName string, exp testAgilityParams, got AgilityParams) {
	test.t.Helper()
	compIdApiResource(test, fieldName+" - agility tier", test.cfg.e.agilityTiers.endpoint, exp.agilityTier, got.AgilityTier)
	compare(test, fieldName+" - tick speed", exp.tickSpeed, got.TickSpeed)
	compare(test, fieldName+" - min icv", exp.minICV, got.MinICV)
	compare(test, fieldName+" - max icv", exp.maxICV, got.MaxICV)
}
