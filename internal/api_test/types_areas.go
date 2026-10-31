package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareAreas(test test, exp expArea, got api.Area) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "displayName", exp.displayName, got.DisplayName)
	compIDAPIResource(test, "location", test.cfg.E.Locations.Endpoint, exp.parentLocation, got.ParentLocation)
	compIDAPIResource(test, "sublocation", test.cfg.E.Sublocations.Endpoint, exp.parentSublocation, got.ParentSublocation)
	checkResIDsInSlice(test, "connected areas", test.cfg.E.Areas.Endpoint, exp.connectedAreas, got.ConnectedAreas)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}
