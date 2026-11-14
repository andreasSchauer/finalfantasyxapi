package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Location struct {
	testGeneral
	Unique
	connectedLocations []int32
	sublocations       []int32
	LocRel
}

func (e Location) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Location) Compare(test test, got apitypes.Location) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compTestStructPtrs("music", e.music, got.Music)
	test.checkResIDsInSlice("connected locations", apilookups.E.Locations.Endpoint, e.connectedLocations, got.ConnectedLocations)
	test.checkResIDsInSlice("sublocations", apilookups.E.Sublocations.Endpoint, e.sublocations, got.Sublocations)
	e.LocRel.Compare(test, got.LocRel)
}
