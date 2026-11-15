package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type TestAgilityParams struct {
	AgilityTier int32
	TickSpeed   int32
	MinICV      *int32
	MaxICV      *int32
}

func (e TestAgilityParams) Compare(test Test, fieldName string, got apitypes.AgilityParams) {
	test.T.Helper()
	test.compIDAPIResource(fieldName+" - agility tier", apilookups.E.AgilityTiers.Endpoint, e.AgilityTier, got.AgilityTier)
	test.compare(fieldName+" - tick speed", e.TickSpeed, got.TickSpeed)
	test.compare(fieldName+" - min icv", e.MinICV, got.MinICV)
	test.compare(fieldName+" - max icv", e.MaxICV, got.MaxICV)
}
