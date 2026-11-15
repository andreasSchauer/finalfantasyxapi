package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Location struct {
	TestGeneral
	Unique
	ConnectedLocations []int32
	Sublocations       []int32
	LocRel
}

func (e Location) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Location) Compare(test Test, got apitypes.Location) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compTestStructPtrs("music", e.Music, got.Music)
	test.CheckResIDsInSlice("connected locations", apilookups.E.Locations.Endpoint, e.ConnectedLocations, got.ConnectedLocations)
	test.CheckResIDsInSlice("sublocations", apilookups.E.Sublocations.Endpoint, e.Sublocations, got.Sublocations)
	e.LocRel.Compare(test, got.LocRel)
}
