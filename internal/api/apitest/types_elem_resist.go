package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type testElemResist struct {
	element  int32
	affinity database.ElementalAffinity
}

func (e testElemResist) Compare(test test, fieldName string, got apitypes.ElementalResist) {
	test.t.Helper()

	test.compIDAPIResource(fieldName+" - elements", apilookups.E.Elements.Endpoint, e.element, got.Element)
	test.compare(fieldName+" - affinities", string(e.affinity), got.Affinity)
}
