package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

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

func compareAreas(test test, exp expArea, got apitypes.Area) {
	test.t.Helper()
	test.compareExpNameVer(exp.expNameVer, got.ID, got.Name, got.Version)
	test.compare("displayName", exp.displayName, got.DisplayName)
	test.compIDAPIResource("location", apilookups.E.Locations.Endpoint, exp.parentLocation, got.ParentLocation)
	test.compIDAPIResource("sublocation", apilookups.E.Sublocations.Endpoint, exp.parentSublocation, got.ParentSublocation)
	test.checkResIDsInSlice("connected areas", apilookups.E.Areas.Endpoint, exp.connectedAreas, got.ConnectedAreas)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}
