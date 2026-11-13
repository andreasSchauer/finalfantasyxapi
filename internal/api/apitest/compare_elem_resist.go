package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareElemResists(test Test, fieldName string, exp TestElemResist, got apitypes.ElementalResist) {
	test.t.Helper()

	test.compIDAPIResource(fieldName+" - elements", apilookups.E.Elements.Endpoint, exp.element, got.Element)
	test.compare(fieldName+" - affinities", string(exp.affinity), got.Affinity)
}
