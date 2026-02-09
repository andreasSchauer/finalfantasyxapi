package main

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

func compareSublocations(test test, exp expSublocation, got Sublocation) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "location", test.cfg.e.locations.endpoint, exp.parentLocation, got.ParentLocation)
	compTestStructPtrs(test, "music", exp.music, got.Music, compareLocMusic)
	checkResIDsInSlice(test, "connected sublocations", test.cfg.e.sublocations.endpoint, exp.connectedSublocations, got.ConnectedSublocations)
	checkResIDsInSlice(test, "areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}