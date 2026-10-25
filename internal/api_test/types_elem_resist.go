package apiTest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type testElemResist struct {
	element  int32
	affinity database.ElementalAffinity
}

func compareElemResists(test test, fieldName string, exp testElemResist, got api.ElementalResist) {
	test.t.Helper()

	compIdApiResource(test, fieldName+" - elements", test.cfg.E.Elements.Endpoint, exp.element, got.Element)
	compare(test, fieldName+" - affinities", string(exp.affinity), got.Affinity)
}
