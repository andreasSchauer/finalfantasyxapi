package apiTest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareLocations(test test, exp expLocation, got api.Location) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compTestStructPtrs(test, "music", exp.music, got.Music, compareLocMusic)
	checkResIDsInSlice(test, "connected locations", test.cfg.E.Locations.Endpoint, exp.connectedLocations, got.ConnectedLocations)
	checkResIDsInSlice(test, "sublocations", test.cfg.E.Sublocations.Endpoint, exp.sublocations, got.Sublocations)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}
