package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Area struct {
	testGeneral
	NameVer
	displayName       string
	parentLocation    int32
	parentSublocation int32
	connectedAreas    []int32
	LocRel
}

func (e Area) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAreas(test test, exp Area, got apitypes.Area) {
	test.t.Helper()
	test.compareExpNameVer(exp.NameVer, got.ID, got.Name, got.Version)
	test.compare("displayName", exp.displayName, got.DisplayName)
	test.compIDAPIResource("location", apilookups.E.Locations.Endpoint, exp.parentLocation, got.ParentLocation)
	test.compIDAPIResource("sublocation", apilookups.E.Sublocations.Endpoint, exp.parentSublocation, got.ParentSublocation)
	test.checkResIDsInSlice("connected areas", apilookups.E.Areas.Endpoint, exp.connectedAreas, got.ConnectedAreas)
	compareLocRel(test, exp.LocRel, got.LocRel)
}
