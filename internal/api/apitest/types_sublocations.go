package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Sublocation struct {
	testGeneral
	Unique
	parentLocation        int32
	connectedSublocations []int32
	areas                 []int32
	LocRel
}

func (e Sublocation) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e Sublocation) Compare(test test, got apitypes.Sublocation) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("location", apilookups.E.Locations.Endpoint, e.parentLocation, got.ParentLocation)
	test.compTestStructPtrs("music", e.music, got.Music)
	test.checkResIDsInSlice("connected sublocations", apilookups.E.Sublocations.Endpoint, e.connectedSublocations, got.ConnectedSublocations)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, e.areas, got.Areas)
	e.LocRel.Compare(test, got.LocRel)
}
