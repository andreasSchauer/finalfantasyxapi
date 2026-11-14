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

func (e Area) Compare(test test, got apitypes.Area) {
	test.t.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("displayName", e.displayName, got.DisplayName)
	test.compIDAPIResource("location", apilookups.E.Locations.Endpoint, e.parentLocation, got.ParentLocation)
	test.compIDAPIResource("sublocation", apilookups.E.Sublocations.Endpoint, e.parentSublocation, got.ParentSublocation)
	test.checkResIDsInSlice("connected areas", apilookups.E.Areas.Endpoint, e.connectedAreas, got.ConnectedAreas)
	e.LocRel.Compare(test, got.LocRel)
}
