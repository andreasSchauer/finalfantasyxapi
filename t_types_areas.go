package main

type expArea struct {
	testGeneral
	expNameVer
	parentLocation    int32
	parentSublocation int32
	connectedAreas    []int32
	expLocRel
}

func (e expArea) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAreas(test test, exp expArea, got Area) {
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compIdApiResource(test, "location", test.cfg.e.locations.endpoint, exp.parentLocation, got.ParentLocation)
	compIdApiResource(test, "sublocation", test.cfg.e.sublocations.endpoint, exp.parentSublocation, got.ParentSublocation)
	checkResIDsInSlice(test, "connected areas", test.cfg.e.areas.endpoint, exp.connectedAreas, got.ConnectedAreas)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}