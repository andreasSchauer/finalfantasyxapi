package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareLocations(test Test, exp Location, got apitypes.Location) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compTestStructPtrs("music", exp.music, got.Music, compareLocMusic)
	test.checkResIDsInSlice("connected locations", apilookups.E.Locations.Endpoint, exp.connectedLocations, got.ConnectedLocations)
	test.checkResIDsInSlice("sublocations", apilookups.E.Sublocations.Endpoint, exp.sublocations, got.Sublocations)
	compareLocRel(test, exp.LocRel, got.LocRel)
}
