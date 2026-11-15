package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type TestElemResist struct {
	Element  int32
	Affinity database.ElementalAffinity
}

func (e TestElemResist) Compare(test Test, fieldName string, got apitypes.ElementalResist) {
	test.T.Helper()

	test.compIDAPIResource(fieldName+" - elements", apilookups.E.Elements.Endpoint, e.Element, got.Element)
	test.compare(fieldName+" - affinities", string(e.Affinity), got.Affinity)
}
