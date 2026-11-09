package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type testElemResist struct {
	element  int32
	affinity database.ElementalAffinity
}

func compareElemResists(test test, fieldName string, exp testElemResist, got apitypes.ElementalResist) {
	test.t.Helper()

	test.compIDAPIResource(fieldName+" - elements", apilookups.E.Elements.Endpoint, exp.element, got.Element)
	test.compare(fieldName+" - affinities", string(exp.affinity), got.Affinity)
}
