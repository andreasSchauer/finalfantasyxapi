package apiTest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expSublocation struct {
	testGeneral
	expUnique
	parentLocation        int32
	connectedSublocations []int32
	areas                 []int32
	expLocRel
}

func (e expSublocation) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSublocations(test test, exp expSublocation, got api.Sublocation) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "location", test.cfg.E.Locations.Endpoint, exp.parentLocation, got.ParentLocation)
	compTestStructPtrs(test, "music", exp.music, got.Music, compareLocMusic)
	checkResIDsInSlice(test, "connected sublocations", test.cfg.E.Sublocations.Endpoint, exp.connectedSublocations, got.ConnectedSublocations)
	checkResIDsInSlice(test, "areas", test.cfg.E.Areas.Endpoint, exp.areas, got.Areas)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}
