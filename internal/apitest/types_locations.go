package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

type expLocation struct {
	testGeneral
	expUnique
	connectedLocations []int32
	sublocations       []int32
	expLocRel
}

func (e expLocation) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareLocations(test test, exp expLocation, got apitypes.Location) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compTestStructPtrs("music", exp.music, got.Music, compareLocMusic)
	test.checkResIDsInSlice("connected locations", apilookups.E.Locations.Endpoint, exp.connectedLocations, got.ConnectedLocations)
	test.checkResIDsInSlice("sublocations", apilookups.E.Sublocations.Endpoint, exp.sublocations, got.Sublocations)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}
