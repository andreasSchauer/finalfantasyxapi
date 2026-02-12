package api

type expArea struct {
	testGeneral
	expNameVer
	displayName       string
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
	compare(test, "displayName", exp.displayName, got.DisplayName)
	compIdApiResource(test, "location", test.cfg.e.locations.endpoint, exp.parentLocation, got.ParentLocation)
	compIdApiResource(test, "sublocation", test.cfg.e.sublocations.endpoint, exp.parentSublocation, got.ParentSublocation)
	checkResIDsInSlice(test, "connected areas", test.cfg.e.areas.endpoint, exp.connectedAreas, got.ConnectedAreas)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}
