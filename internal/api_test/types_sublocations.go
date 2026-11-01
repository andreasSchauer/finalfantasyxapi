package apitest

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
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("location", test.cfg.E.Locations.Endpoint, exp.parentLocation, got.ParentLocation)
	test.compTestStructPtrs("music", exp.music, got.Music, compareLocMusic)
	test.checkResIDsInSlice("connected sublocations", test.cfg.E.Sublocations.Endpoint, exp.connectedSublocations, got.ConnectedSublocations)
	test.checkResIDsInSlice("areas", test.cfg.E.Areas.Endpoint, exp.areas, got.Areas)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}
