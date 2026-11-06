package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type testElemResist struct {
	element  int32
	affinity database.ElementalAffinity
}

func compareElemResists(test test, fieldName string, exp testElemResist, got apitypes.ElementalResist) {
	test.t.Helper()

	test.compIDAPIResource(fieldName+" - elements", api.E.Elements.Endpoint, exp.element, got.Element)
	test.compare(fieldName+" - affinities", string(exp.affinity), got.Affinity)
}
