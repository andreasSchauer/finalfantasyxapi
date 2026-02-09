package main

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

func compareLocations(test test, exp expLocation, got Location) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compTestStructPtrs(test, "music", exp.music, got.Music, compareLocMusic)
	checkResIDsInSlice(test, "connected locations", test.cfg.e.locations.endpoint, exp.connectedLocations, got.ConnectedLocations)
	checkResIDsInSlice(test, "sublocations", test.cfg.e.sublocations.endpoint, exp.sublocations, got.Sublocations)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}