package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Area struct {
	TestGeneral
	NameVer
	DisplayName       string
	ParentLocation    int32
	ParentSublocation int32
	ConnectedAreas    []int32
	LocRel
}

func (e Area) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Area) Compare(test Test, got apitypes.Area) {
	test.T.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("displayName", e.DisplayName, got.DisplayName)
	test.compIDAPIResource("location", apilookups.E.Locations.Endpoint, e.ParentLocation, got.ParentLocation)
	test.compIDAPIResource("sublocation", apilookups.E.Sublocations.Endpoint, e.ParentSublocation, got.ParentSublocation)
	test.checkResIDsInSlice("connected areas", apilookups.E.Areas.Endpoint, e.ConnectedAreas, got.ConnectedAreas)
	e.LocRel.Compare(test, got.LocRel)
}
