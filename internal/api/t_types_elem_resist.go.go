package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"

type testElemResist struct {
	element  int32
	affinity database.ElementalAffinity
}

func compareElemResists(test test, fieldName string, exp testElemResist, got ElementalResist) {
	test.t.Helper()

	compIdApiResource(test, fieldName+" - elements", test.cfg.e.elements.endpoint, exp.element, got.Element)
	compare(test, fieldName+" - affinities", string(exp.affinity), got.Affinity)
}
