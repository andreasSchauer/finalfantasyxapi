package main


type testElemResist struct {
	element  int32
	affinity int32
}

func compareElemResists(test test, fieldName string, exp testElemResist, got ElementalResist) {
	elemEndpoint := test.cfg.e.elements.endpoint
	affinityEndpoint := test.cfg.e.affinities.endpoint

	compIdApiResource(test, fieldName+" - elements", elemEndpoint, exp.element, got.Element)
	compIdApiResource(test, fieldName+" - affinities", affinityEndpoint, exp.affinity, got.Affinity)
}