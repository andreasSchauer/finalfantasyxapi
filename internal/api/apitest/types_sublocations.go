package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

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

func compareSublocations(test test, exp expSublocation, got apitypes.Sublocation) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("location", apilookups.E.Locations.Endpoint, exp.parentLocation, got.ParentLocation)
	test.compTestStructPtrs("music", exp.music, got.Music, compareLocMusic)
	test.checkResIDsInSlice("connected sublocations", apilookups.E.Sublocations.Endpoint, exp.connectedSublocations, got.ConnectedSublocations)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, exp.areas, got.Areas)
	compareLocRel(test, exp.expLocRel, got.LocRel)
}
